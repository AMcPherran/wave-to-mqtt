package main

import (
	"fmt"

	mqtt "github.com/eclipse/paho.mqtt.golang"
)

var knt int
var f mqtt.MessageHandler = func(client mqtt.Client, msg mqtt.Message) {
	fmt.Printf("MSG: %s\n", msg.Payload())
	text := fmt.Sprintf("this is result msg #%d!", knt)
	knt++
	token := client.Publish("nn/result", 0, false, text)
	token.Wait()
}

func getMQTTClient(host, port, username, password string) mqtt.Client {
	brokerAddr := fmt.Sprintf("tcp://%s:%s", host, port)
	opts := mqtt.NewClientOptions().AddBroker(brokerAddr)
	opts.SetClientID("wave-master")
	opts.SetDefaultPublishHandler(f)
	opts.SetUsername(username)
	opts.SetPassword(password)
	client := mqtt.NewClient(opts)
	if token := client.Connect(); token.Wait() && token.Error() != nil {
		panic(token.Error())
	} else {
		fmt.Printf("Connected to MQTT broker\n")
	}
	return client
}

type ButtonPayload struct {
	Action string `json:"action"`
	RSSI   int    `json:"rssi"`
}
