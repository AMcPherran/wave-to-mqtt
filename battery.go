package main

import (
	"encoding/json"
	"log"
	"time"

	gowave "github.com/AMcPherran/go-wave"
)

func handleBatteryStatus(w *gowave.Wave, lastState *gowave.WaveState) {
	bs := w.State.GetBatteryStatus()
	if bs != lastState.GetBatteryStatus() {
		lastState.SetBatteryStatus(bs)
		publishBatteryStatus(bs)
		log.Printf("Publishing updated battery status: %f \n", bs.Percentage)
	}
	// Update battery status every 5 minutes
	if (time.Now().Unix() - bs.Timestamp) > (60 * 5) {
		// Send request for updated battery status
		err := w.SendBatteryStatusRequest()
		if err != nil {
			log.Printf("Error sending BatteryStatus Request: %s \n", err)
		}
	}
}

func publishBatteryStatus(bs gowave.BatteryStatus) {
	topic := "wave/battery"
	p, _ := json.Marshal(bs)
	token := mqClient.Publish(topic, 0, false, p)
	token.Wait()
}
