package main

import (
	"context"
	"encoding/json"
	"firebase.google.com/go"
	"firebase.google.com/go/messaging"
	"github.com/julienschmidt/httprouter"
	"google.golang.org/api/option"
	"log"
	"net/http"
	"time"
)

type State struct {
	DeviceInfo      map[string][]DeviceMetadata
	RideNumber      int
	DriverData      []DriverMetadata
	FirebaseApp     *firebase.App
	MessagingClient *messaging.Client
}

type DriverMetadata struct {
	Wallet   string `json:"wallet"`
	RideData `json:"ride-data"`
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
}

var state State

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
	router.POST("/availableRides", seekRide)
	log.Fatal(http.ListenAndServe("10.177.1.130:8080", router))
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

	err := json.NewDecoder(request.Body).Decode(rideRequest)
	if err != nil {
		log.Println(err.Error())
		writer.WriteHeader(400)
		_, _ = writer.Write([]byte("Invalid request body"))
		return
	}

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
