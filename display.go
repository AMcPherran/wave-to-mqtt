package main

import (
	"log"

	gowave "github.com/AMcPherran/go-wave"
)

func handleDisplay(w *gowave.Wave, desiredState *gowave.WaveState) {
	wd := w.State.GetDisplayState()
	ds := desiredState.GetDisplayState()
	if wd.Timestamp < ds.Timestamp {
		err := w.SetDisplay(ds.Frame)
		if err != nil {
			log.Printf("Error sending DisplayFrame Request in response to DesiredState: %s\n", err)
		}
	}
}

func getNotificationDisplayFrame() [][]byte {
	frame := [][]byte{
		{000, 000, 000, 000, 155, 000, 000, 000, 000},
		{000, 000, 000, 000, 000, 000, 000, 000, 000},
		{000, 000, 000, 000, 255, 000, 000, 000, 000},
		{000, 000, 000, 000, 255, 000, 000, 000, 000},
		{000, 000, 000, 000, 230, 000, 000, 000, 000},
	}
	return frame
}

func getRoomChangeDisplayFrame() [][]byte {
	frame := [][]byte{
		{255, 255, 255, 255, 255, 255, 255, 255, 255},
		{255, 000, 000, 000, 000, 000, 000, 000, 255},
		{255, 000, 000, 000, 000, 000, 000, 000, 255},
		{255, 000, 000, 000, 000, 000, 000, 000, 255},
		{255, 255, 255, 255, 255, 255, 255, 255, 255},
	}
	return frame
}
