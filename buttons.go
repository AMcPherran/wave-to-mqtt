package main

import (
	"fmt"

	gowave "github.com/AMcPherran/go-wave"
)

func handleButtons(w *gowave.Wave, lastState *gowave.WaveState) {
	tb := w.State.Buttons.Top()
	bb := w.State.Buttons.Bottom()

	if tb != lastState.Buttons.Top() {
		handleTopButton(w, tb)
		lastState.Buttons.Set(tb)
	}

	if bb != lastState.Buttons.Bottom() {
		handleBottomButton(w, bb)
		lastState.Buttons.Set(bb)
	}

}

func handleMiddleButton(w *gowave.Wave, b gowave.ButtonEvent) {
	// Indicate that button was clicked
	if b.Action == "Down" {
		w.Recenter()
		frame := [][]byte{
			{000, 000, 000, 000, 000, 000, 000, 000, 000},
			{000, 000, 000, 000, 255, 000, 000, 000, 000},
			{000, 000, 000, 255, 255, 255, 000, 000, 000},
			{000, 000, 000, 000, 255, 000, 000, 000, 000},
			{000, 000, 000, 000, 000, 000, 000, 000, 000},
		}
		w.SetDisplay(frame)
		// Send request for current battery status
		w.SendBatteryStatusRequest()
		// Indicate that the Wave is ready to process motion
	} else if b.Action == "Long" || b.Action == "ExtraLong" {
		w.Recenter()
		frame := [][]byte{
			{000, 000, 000, 000, 255, 000, 000, 000, 000},
			{000, 000, 000, 255, 100, 255, 000, 000, 000},
			{000, 000, 255, 100, 255, 100, 255, 000, 000},
			{000, 000, 000, 255, 100, 255, 000, 000, 000},
			{000, 000, 000, 000, 255, 000, 000, 000, 000},
		}
		w.SetDisplay(frame)
	} else {
		if b.Action == "Up" {
			topic := fmt.Sprintf("wave/buttons/%s", b.ID)
			token := mqClient.Publish(topic, 0, false, b.Action)
			token.Wait()
		}
		frame := gowave.BlankDisplayFrame()
		w.SetDisplay(frame)
	}
}

func handleTopButton(w *gowave.Wave, b gowave.ButtonEvent) {
	down := buttonDown(b)
	if down {
		var frame [][]byte
		if b.Action == "Down" {
			frame = [][]byte{
				{000, 000, 000, 000, 000, 000, 255, 000, 000},
				{000, 000, 000, 000, 000, 000, 000, 255, 000},
				{000, 000, 000, 000, 000, 000, 000, 000, 255},
				{000, 000, 000, 000, 000, 000, 000, 255, 000},
				{000, 000, 000, 000, 000, 000, 255, 000, 000},
			}
		} else {
			frame = [][]byte{
				{000, 000, 000, 000, 000, 000, 255, 000, 000},
				{000, 000, 000, 000, 000, 255, 000, 255, 000},
				{000, 000, 000, 000, 000, 000, 255, 000, 255},
				{000, 000, 000, 000, 000, 255, 000, 255, 000},
				{000, 000, 000, 000, 000, 000, 255, 000, 000},
			}
		}
		w.SetDisplay(frame)
		w.Recenter()
	} else {
		frame := gowave.BlankDisplayFrame()
		w.SetDisplay(frame)
		topic := fmt.Sprintf("wave/buttons/%s", b.ID)
		token := mqClient.Publish(topic, 0, false, b.Action)
		token.Wait()
	}
}

func handleBottomButton(w *gowave.Wave, b gowave.ButtonEvent) {
	down := buttonDown(b)
	if down {
		var frame [][]byte
		if b.Action == "Down" {
			frame = [][]byte{
				{000, 000, 255, 000, 000, 000, 000, 000, 000},
				{000, 255, 000, 000, 000, 000, 000, 000, 000},
				{255, 000, 000, 000, 000, 000, 000, 000, 000},
				{000, 255, 000, 000, 000, 000, 000, 000, 000},
				{000, 000, 255, 000, 000, 000, 000, 000, 000},
			}
		} else if b.Action == "Long" {
			frame = [][]byte{
				{000, 000, 255, 000, 000, 000, 000, 000, 000},
				{000, 255, 000, 255, 000, 000, 000, 000, 000},
				{255, 000, 255, 000, 000, 000, 000, 000, 000},
				{000, 255, 000, 255, 000, 000, 000, 000, 000},
				{000, 000, 255, 000, 000, 000, 000, 000, 000},
			}
		} else {
			frame = [][]byte{
				{000, 000, 255, 255, 255, 000, 000, 000, 000},
				{000, 255, 255, 255, 255, 000, 000, 000, 000},
				{255, 255, 255, 255, 255, 000, 000, 000, 000},
				{000, 255, 255, 255, 255, 000, 000, 000, 000},
				{000, 000, 255, 255, 255, 000, 000, 000, 000},
			}
		}
		w.SetDisplay(frame)
		w.Recenter()
	} else {
		frame := gowave.BlankDisplayFrame()
		w.SetDisplay(frame)
		topic := fmt.Sprintf("wave/buttons/%s", b.ID)
		token := mqClient.Publish(topic, 0, false, b.Action)
		token.Wait()
	}
}

func buttonDown(b gowave.ButtonEvent) bool {
	if b.Action == "Down" || b.Action == "Long" || b.Action == "ExtraLong" {
		return true
	} else if b.Action == "Up" || b.Action == "LongUp" || b.Action == "ExtraLongUp" {
		return false
	}
	return false
}
