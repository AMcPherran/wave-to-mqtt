package main

import (
	"fmt"
	"log"
	"os"
	"time"

	gowave "github.com/AMcPherran/go-wave"
	mqtt "github.com/eclipse/paho.mqtt.golang"
)

var mqClient mqtt.Client

var mqHost string = os.Getenv("MQTT_BROKER")
var mqPort string = os.Getenv("MQTT_PORT")
var mqUser string = os.Getenv("MQTT_USER")
var mqPass string = os.Getenv("MQTT_PASS")

func main() {
	// Always be trying to connect
	for true {
		wave, err := gowave.Connect()
		if err != nil {
			fmt.Println(err)
			log.Println("Did not connect to Wave, will try again in a sec")
			time.Sleep(time.Second)
			continue
		}

		// Start receiving data from Wave
		if err := wave.HandleNotifications(); err != nil {
			log.Fatalf("subscribe failed: %s", err)
		}
		log.Printf("Receiving incoming data from Wave")

		// Get an MQTT client
		mqClient = getMQTTClient(mqHost, mqPort, mqUser, mqPass)

		handleWave(wave)

		//<-wave.BLE.Client.Disconnected()
		// Disconnect the connections
		wave.Disconnect()
		mqClient.Disconnect(5)
		fmt.Println("Disconnected from Wave and MQTT broker")
	}
}

func handleWave(w *gowave.Wave) {
	// Main loop for reading and acting on wave.State
	var lastState gowave.WaveState
	for true {
		select {
		case <-w.BLE.Client.Disconnected():
			return
		default:
			handleButtons(w, &lastState)
			handleMotion(w, &lastState)
			handleBatteryStatus(w, &lastState)
			time.Sleep(500 * time.Microsecond)
		}
	}
}
