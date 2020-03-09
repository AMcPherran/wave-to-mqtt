package main

import (
	"encoding/json"

	gowave "github.com/AMcPherran/go-wave"
)

func handleBatteryStatus(w *gowave.Wave, lastState *gowave.WaveState) {
	bs := w.State.GetBatteryStatus()
	if bs != lastState.GetBatteryStatus() {
		lastState.SetBatteryStatus(bs)
		publishBatteryStatus(bs)
	}
}

func publishBatteryStatus(bs gowave.BatteryStatus) {
	topic := "wave/battery"
	p, _ := json.Marshal(bs)
	token := mqClient.Publish(topic, 0, false, p)
	token.Wait()
}
