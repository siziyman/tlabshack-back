package main

import (
	"context"
	"encoding/json"
	"firebase.google.com/go"
	"firebase.google.com/go/messaging"
	"fmt"
	"github.com/julienschmidt/httprouter"
	"google.golang.org/api/option"
	"log"
	"net/http"
	"strings"
	"time"
	"tlabshack/rides"
)

type BalanceResponse struct {
	Balance int64 `json:"balance"`
}

type State struct {
	DeviceInfo      map[string][]DeviceMetadata
	RideNumber      int
	DriverData      []DriverMetadata
	FirebaseApp     *firebase.App
	MessagingClient *messaging.Client
}

type DriverMetadata struct {
	Wallet   string   `json:"wallet"`
	RideData RideData `json:"rideData"`
}

type DeviceMetadata struct {
	TokenId       string `json:"tokenId"`
	FirebaseToken string `json:"firebaseToken"`
	Wallet        string `json:"wallet"`
}

type Coordinates struct {
	Longitude float64 `json:"longitude"`
	Latitude  float64 `json:"latitude"`
}

type RideData struct {
	StartPoint Coordinates `json:"from"`
	EndPoint   Coordinates `json:"to"`
}

type RideRequest struct {
	RideData RideData `json:"rideData"`
	Message  string   `json:"message"`
	Wallet   string   `json:"wallet"`
}

type VerifyRequest struct {
	Wallet       string `json:"wallet"`
	PrivateKey   string `json:"privateKey"`
	DriverWallet string `json:"driverWallet"`
}

var state State
var baseUrl = "https://demo.stax.tlabs.cloud/projects/yetAnotherTeam/contexts/Stax_1/"

func main() {

	state = State{
		DeviceInfo:  make(map[string][]DeviceMetadata),
		RideNumber:  0,
		DriverData:  make([]DriverMetadata, 0),
		FirebaseApp: nil,
	}
	ctx := context.Background()
	config := &firebase.Config{
		AuthOverride:     nil,
		DatabaseURL:      "",
		ProjectID:        "rideshare-7e771",
		ServiceAccountID: "",
		StorageBucket:    "rideshare-7e771.appspot.com",
	}
	var e error
	opt := option.WithCredentialsFile("rideshare.json")
	state.FirebaseApp, e = firebase.NewApp(ctx, config, opt)
	if e != nil {
		log.Println(e.Error())
	}
	client, _ := state.FirebaseApp.Messaging(context.Background())
	state.MessagingClient = client

	oneHour := time.Duration(1) * time.Hour
	msg := &messaging.Message{
		Android: &messaging.AndroidConfig{
			TTL:      &oneHour,
			Priority: "normal",
			Notification: &messaging.AndroidNotification{
				Title: "HELLO RIDE",
				Body:  "some ride",
				Icon:  "",
				Color: "#f45342",
			},
		},
		Topic: "ride",
	}

	androidConfig := new(messaging.AndroidConfig)
	androidConfig.Notification = new(messaging.AndroidNotification)
	//androidConfig.Notification.Body = "HELLO"
	msg.Android = androidConfig
	//msg.Data = map[string]string{"hell": "hello"}
	s, e := client.Send(ctx, msg)
	log.Println(s)
	router := httprouter.New()
	router.POST("/registerDevice", registerDevice)
	router.POST("/registerDriver", drive)
	router.GET("/availableRides", seekRide)
	router.POST("/verifyRide", verifyRide)
	log.Fatal(http.ListenAndServe("10.177.1.130:8080", router))
}

func verifyRide(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	verifyRequest := new(VerifyRequest)

	err := json.NewDecoder(request.Body).Decode(verifyRequest)
	if err != nil {
		log.Println(err.Error())
		writer.WriteHeader(400)
		_, _ = writer.Write([]byte("Invalid request body"))
		return
	}
	str := fmt.Sprintf(`{
	"recipient_ref": "%s",
	"amount": 1000
}
`)
	reader := strings.NewReader(str)
	newRequest, err := http.NewRequest("POST", baseUrl+"payment", reader)
	newRequest.Header.Add("Originator-Ref", verifyRequest.Wallet)
	newRequest.Header.Add("Authorization", verifyRequest.PrivateKey)

}

func drive(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	driverMetadata := new(DriverMetadata)
	err := json.NewDecoder(request.Body).Decode(driverMetadata)
	if err != nil {
		log.Println(err.Error())
		writer.WriteHeader(400)
		_, _ = writer.Write([]byte("Invalid request body"))
		return
	}

	state.DriverData = append(state.DriverData, *driverMetadata)
}

func seekRide(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	rideRequest := new(RideRequest)

	resp, err := http.Get(baseUrl + "account/" + rideRequest.Wallet)
	if err != nil {
		errStr := "Error checking balance"
		log.Println(errStr)
		writer.WriteHeader(500)
		_, _ = writer.Write([]byte(errStr))
		return
	}

	balance := new(BalanceResponse)
	err = json.NewDecoder(resp.Body).Decode(balance)
	if err != nil {
		errStr := "Error checking balance"
		log.Println(errStr)
		writer.WriteHeader(500)
		_, _ = writer.Write([]byte(errStr))
		return
	}
	if balance.Balance < 1100 {
		errStr := "Not enough balance to pay for the ride"
		log.Println(errStr)
		writer.WriteHeader(402)
		_, _ = writer.Write([]byte(errStr))
		return
	}

	err = json.NewDecoder(request.Body).Decode(rideRequest)
	if err != nil {
		log.Println(err.Error())
		writer.WriteHeader(400)
		_, _ = writer.Write([]byte("Invalid request body"))
		return
	}

	for _, driver := range state.DriverData {
		first := rides.Distance(driver.RideData.StartPoint.Latitude, driver.RideData.StartPoint.Longitude, rideRequest.RideData.StartPoint.Latitude, rideRequest.RideData.StartPoint.Longitude)
		second := rides.Distance(driver.RideData.EndPoint.Latitude, driver.RideData.EndPoint.Longitude, rideRequest.RideData.EndPoint.Latitude, rideRequest.RideData.EndPoint.Longitude)
		if first+second < 1000 {
			for k, devices := range state.DeviceInfo {
				if k == driver.Wallet {
					for _, device := range devices {
						sendPush(device.TokenId, rideRequest.Message)
						writer.WriteHeader(200)
						_, _ = writer.Write([]byte("Ride request was sent"))
						return
					}
				}

			}
		}
	}
	writer.WriteHeader(404)
	_, _ = writer.Write([]byte("Drivers not found"))
}

func registerDevice(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	metadata := new(DeviceMetadata)

	err := json.NewDecoder(r.Body).Decode(metadata)
	if err != nil {
		log.Println(err.Error())
		w.WriteHeader(400)
		_, _ = w.Write([]byte("Invalid request body"))
		return
	}
	if state.DeviceInfo[metadata.TokenId] == nil {
		state.DeviceInfo[metadata.TokenId] = make([]DeviceMetadata, 0)
	}
	state.DeviceInfo[metadata.TokenId] = append(state.DeviceInfo[metadata.TokenId], *metadata)

	_, _ = w.Write([]byte("Successfully registered"))
	//state.FirebaseApp.

}

func sendPush(registrationToken, message string) {
	// Obtain a messaging.Client from the App.

	// See documentation on defining a message payload.
	msg := &messaging.Message{
		Data: map[string]string{
			"message": message,
		},
		Token: registrationToken,
	}

	// Send a message to the device corresponding to the provided
	// registration token.
	response, err := state.MessagingClient.Send(context.Background(), msg)
	if err != nil {
		log.Fatalln(err)
		return
	}
	log.Println(response)
}
