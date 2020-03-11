package main

import (
	"fmt"
	"log"
	"time"

	gowave "github.com/AMcPherran/go-wave"
	mqtt "github.com/eclipse/paho.mqtt.golang"
)

var displaySetMsgHandler mqtt.MessageHandler = func(client mqtt.Client, msg mqtt.Message) {
	s := fmt.Sprintf("%s", msg.Payload())
	if s == "notify" {
		frame := getNotificationDisplayFrame()
		serverState.SetDisplayState(gowave.DisplayState{
			Frame:     frame,
			Timestamp: time.Now().Unix(),
		})
	}
	if s == "roomChange" {
		frame := getRoomChangeDisplayFrame()
		serverState.SetDisplayState(gowave.DisplayState{
			Frame:     frame,
			Timestamp: time.Now().Unix(),
		})
	}
}

func getMQTTClient(host, port, username, password string) mqtt.Client {
	brokerAddr := fmt.Sprintf("tcp://%s:%s", host, port)
	opts := mqtt.NewClientOptions().AddBroker(brokerAddr)
	opts.SetClientID("wave-master")
	opts.SetDefaultPublishHandler(displaySetMsgHandler)
	opts.SetUsername(username)
	opts.SetPassword(password)
	client := mqtt.NewClient(opts)
	if token := client.Connect(); token.Wait() && token.Error() != nil {
		panic(token.Error())
	} else {
		log.Printf("Connected to MQTT broker \"%s\"\n", host)
	}
	// Subscribe to display payload topic
	client.Subscribe("wave/display/set", 0, displaySetMsgHandler)
	return client
}

type ButtonPayload struct {
	Action string `json:"action"`
	RSSI   int    `json:"rssi"`
}
