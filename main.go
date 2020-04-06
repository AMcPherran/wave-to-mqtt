package main

import (
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

var serverState gowave.WaveState

func main() {
	// Always be trying to connect
	for true {
		wave, err := gowave.Connect()
		if err != nil {
			log.Printf("Did not connect to Wave, will try again in a few seconds. Reason: %s \n", err)
			time.Sleep(time.Second)
			continue
		}

		// Start receiving data from Wave
		if err := wave.HandleNotifications(); err != nil {
			log.Fatalf("Subscribe failed: %s", err)
		}
		log.Printf("Receiving incoming data from Wave")

		// Get an MQTT client
		mqClient = getMQTTClient(mqHost, mqPort, mqUser, mqPass)

		// Send request for initial battery status
		wave.SendBatteryStatusRequest()

		handleWave(wave)

		//<-wave.BLE.Client.Disconnected()
		// Disconnect the connections
		err = wave.Disconnect()
		if err != nil {
			log.Printf("Error disconnecting from Wave: %s", err)
		}
		mqClient.Disconnect(5)
		log.Println("Disconnected from Wave and MQTT broker")
	}
}

func handleWave(w *gowave.Wave) {
	// Main loop for reading and acting on wave.State
	for true {
		select {
		case <-w.BLE.Client.Disconnected():
			return
		default:
			handleButtons(w, &serverState)
			handleMotion(w, &serverState)
			handleBatteryStatus(w, &serverState)
			//handleDisplay(w, &serverState)
			time.Sleep(500 * time.Microsecond)
		}
	}
}
