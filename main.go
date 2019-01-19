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
	RideNumber      string
	DriverData      []DriverMetadata
	FirebaseApp     *firebase.App
	MessagingClient *messaging.Client
}

type DriverMetadata struct {
	Wallet string `json:"wallet"`
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

var state State

func main() {
	state := State{
		DeviceInfo:  map[string][]DeviceMetadata{},
		RideNumber:  "",
		DriverData:  []DriverMetadata{},
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
		log.Println("айайай драйв")
	}
	_ = append(state.DriverData, *driverMetadata)
}

func seekRide(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {

}

func registerDevice(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	metadata := new(DeviceMetadata)

	err := json.NewDecoder(r.Body).Decode(metadata)
	if err != nil {
		log.Println("ойойой")
	}
	_ = append(state.DeviceInfo[metadata.TokenId], *metadata)

	//state.FirebaseApp.

}
