package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
	"time"

	"./rides"
	"firebase.google.com/go"
	"firebase.google.com/go/messaging"
	"github.com/julienschmidt/httprouter"
	"google.golang.org/api/option"
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

type ParkingRequest struct {
	ID         string `json:"id"`
	Wallet     string `json:"wallet"`
	PrivateKey string `json:"privateKey"`
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
	From    Coordinates `json:"from"`
	To      Coordinates `json:"to"`
	Message string      `json:"message"`
	Wallet  string      `json:"wallet"`
}

type VerifyRequest struct {
	Wallet       string `json:"wallet"`
	PrivateKey   string `json:"privateKey"`
	DriverWallet string `json:"driverWallet"`
}

type BalanceRequest struct {
	Wallet string `json:"wallet"`
}

var state State
var baseUrl = "https://demo.stax.tlabs.cloud/projects/yetAnotherTeam/contexts/Stax_1/"

const parkingAccount = "0x5a40fE165B188d15a9Fe1411F5Aee170b53bc871"
const IOTAOriginator = "PXKEFBWMND9SVDUHMIXAGHODJWCUZ9XI9FRSZGNJS9TXNZ9XPAZAUMNIFDNWYZTFSMSYSZYXQOYJVURGD"

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
	//s, e := client.Send(ctx, msg)
	//log.Println(s)
	router := httprouter.New()
	router.POST("/registerDevice", registerDevice)
	router.POST("/registerDriver", drive)
	router.POST("/availableRides", seekRide)
	router.POST("/verifyRide", verifyRide)
	router.GET("/balance/:wallet", balance)
	router.POST("/parkCar", parkCar)
	log.Fatal(http.ListenAndServe("10.177.1.130:8080", router))
}

func parkCar(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	parkingRequest := new(ParkingRequest)
	log.Println("Parking car")
	err := json.NewDecoder(request.Body).Decode(parkingRequest)
	if err != nil {
		log.Println(err.Error())
		writer.WriteHeader(400)
		_, _ = writer.Write([]byte("Invalid request body"))

	}

	str := fmt.Sprintf(`{
	"recipient_ref": "%s",
	"amount": 1000
}
`, parkingAccount)
	reader := strings.NewReader(str)
	newRequest, err := http.NewRequest("POST", baseUrl+"payment", reader)
	newRequest.Header.Add("Originator-Ref", parkingRequest.Wallet)
	newRequest.Header.Add("Authorization", "0x"+parkingRequest.PrivateKey)

	log.Println("Paying for parking: " + str)
	client := &http.Client{}
	response, err := client.Do(newRequest)
	if err != nil {
		log.Println("Error trying to conduct payment")
		log.Println(err.Error())
		writer.WriteHeader(500)
		_, _ = writer.Write([]byte("Error trying to conduct payment"))
	}
	bytes, err := ioutil.ReadAll(response.Body)
	log.Println(string(bytes))
	//log.Println("Status code " + string(response.StatusCode))
	log.Println("Storing parking data")
	storeParkingData(parkingRequest)

}

func storeParkingData(parkingRequest *ParkingRequest) {
	storageBody := fmt.Sprintf(`{
			"data": "%s",
			"time": "%s"
		}`, parkingRequest.Wallet, time.Now().String())
	log.Println("store data: " + storageBody)
	bodyReader := strings.NewReader(storageBody)
	request, _ := http.NewRequest("POST", baseUrl+"/storage", bodyReader)
	request.Header.Add("Originator-Ref", IOTAOriginator)
	client := &http.Client{}
	resp, err := client.Do(request)
	log.Println(resp.StatusCode)
	storeResponseBody, _ := ioutil.ReadAll(resp.Body)
	log.Println(string(storeResponseBody))
	if err != nil {
		log.Println(err.Error())
		log.Println("Error storing data")
	}
}

func verifyRide(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	verifyRequest := new(VerifyRequest)
	log.Println("Verifying ride")
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
`, verifyRequest.DriverWallet)
	log.Println("Paying for ride")
	reader := strings.NewReader(str)
	newRequest, err := http.NewRequest("POST", baseUrl+"payment", reader)
	newRequest.Header.Add("Originator-Ref", verifyRequest.Wallet)
	newRequest.Header.Add("Authorization", "0x"+verifyRequest.PrivateKey)

	client := &http.Client{}
	responsePayment, err := client.Do(newRequest)
	bytes, _ := ioutil.ReadAll(responsePayment.Body)
	log.Println(responsePayment.StatusCode)
	log.Println(string(bytes))
	if responsePayment.StatusCode >= 400 {
		writer.WriteHeader(400)
	} else {
		writer.WriteHeader(200)
	}

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
	writer.WriteHeader(200)
}

func seekRide(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	rideRequest := new(RideRequest)

	err := json.NewDecoder(request.Body).Decode(rideRequest)
	if err != nil {
		log.Println(err.Error())
		writer.WriteHeader(400)
		_, _ = writer.Write([]byte("Bad request body"))
	}

	resp, err := http.Get(baseUrl + "account/" + rideRequest.Wallet)
	if err != nil {
		errStr := "Error checking balance"
		log.Println(errStr)
		log.Println(err.Error())
		writer.WriteHeader(500)
		_, _ = writer.Write([]byte(errStr))
		return
	}

	balance := new(BalanceResponse)
	err = json.NewDecoder(resp.Body).Decode(balance)
	if err != nil {
		errStr := "Error checking balance step 2"
		log.Println(errStr)
		log.Println(err.Error())
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

	for _, driver := range state.DriverData {
		first := rides.Distance(driver.RideData.StartPoint.Latitude, driver.RideData.StartPoint.Longitude, rideRequest.From.Latitude, rideRequest.From.Longitude)
		second := rides.Distance(driver.RideData.EndPoint.Latitude, driver.RideData.EndPoint.Longitude, rideRequest.To.Latitude, rideRequest.To.Longitude)
		if first+second < 1000 {
			sendPush(rideRequest.Message)
			//for _, devices := range state.DeviceInfo {
			//	for _, device := range devices {
			//		writer.WriteHeader(200)
			//		_, _ = writer.Write([]byte("Ride request was sent"))
			//	}
			return
			//}
		}
	}
	//return
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

func sendPush(message string) {
	// Obtain a messaging.Client from the App.

	// See documentation on defining a message payload.
	msg := &messaging.Message{
		Notification: &messaging.Notification{
			Title: "Companion found!",
			Body:  "Found person to share a ride with you",
		},
		Data: map[string]string{
			"message": message,
		},
		Topic: "ride",
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

func balance(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	address := params.ByName("wallet")
	resp, err := http.Get(baseUrl + "account/" + address)
	if err != nil {
		errStr := "Error checking balance"
		log.Println(errStr)
		log.Println(err.Error())
		writer.WriteHeader(500)
		_, _ = writer.Write([]byte(errStr))
		return
	}

	balance := new(BalanceResponse)
	err = json.NewDecoder(resp.Body).Decode(balance)
	if err != nil {
		errStr := "Error checking balance step 2"
		log.Printf("%v\n", resp)
		log.Println(errStr)
		log.Println(err.Error())
		writer.WriteHeader(500)
		_, _ = writer.Write([]byte(errStr))
		return
	}
	js, err := json.Marshal(balance)
	if err != nil {
		errStr := "Error checking balance step 3"
		log.Printf("%v\n", resp)
		log.Println(errStr)
		log.Println(err.Error())
		writer.WriteHeader(500)
		_, _ = writer.Write([]byte(errStr))
		return
	}
	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(200)
	writer.Write(js)
	//_, _ = fmt.Fprintf(writer, "%d", balance.Balance)
}
