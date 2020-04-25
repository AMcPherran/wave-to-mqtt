package main

import (
	"fmt"
	"log"

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
		err := w.SetDisplay(frame)
		if err != nil {
			log.Printf("Error sending DisplayFrame Request: %s\n", err)
		}
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
		err := w.SetDisplay(frame)
		if err != nil {
			log.Printf("Error sending DisplayFrame Request: %s\n", err)
		}
	} else {
		if b.Action == "Up" || b.Action == "Click" {
			topic := fmt.Sprintf("wave/buttons/%s", b.ID)
			token := mqClient.Publish(topic, 0, false, "Up")
			token.Wait()
		}
		frame := gowave.BlankDisplayFrame()
		err := w.SetDisplay(frame)
		if err != nil {
			log.Printf("Error sending DisplayFrame Request: %s\n", err)
		}
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
		} else if b.Action == "Long" {
			frame = [][]byte{
				{000, 000, 000, 000, 000, 000, 255, 000, 000},
				{000, 000, 000, 000, 000, 255, 000, 255, 000},
				{000, 000, 000, 000, 000, 000, 255, 000, 255},
				{000, 000, 000, 000, 000, 255, 000, 255, 000},
				{000, 000, 000, 000, 000, 000, 255, 000, 000},
			}
		} else {
			frame = [][]byte{
				{000, 000, 000, 000, 255, 255, 255, 000, 000},
				{000, 000, 000, 000, 255, 255, 255, 255, 000},
				{000, 000, 000, 000, 255, 255, 255, 255, 255},
				{000, 000, 000, 000, 255, 255, 255, 255, 000},
				{000, 000, 000, 000, 255, 255, 255, 000, 000},
			}
		}
		err := w.SetDisplay(frame)
		if err != nil {
			log.Printf("Error sending DisplayFrame Request: %s\n", err)
		}
		err = w.Recenter()
		if err != nil {
			log.Printf("Error sending Recenter Request: %s\n", err)
		}
	} else {
		frame := gowave.BlankDisplayFrame()
		err := w.SetDisplay(frame)
		if err != nil {
			log.Printf("Error sending DisplayFrame Request: %s\n", err)
		}
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
		err := w.SetDisplay(frame)
		if err != nil {
			log.Printf("Error sending DisplayFrame Request: %s\n", err)
		}
		err = w.Recenter()
		if err != nil {
			log.Printf("Error sending Recenter Request: %s\n", err)
		}
	} else {
		frame := gowave.BlankDisplayFrame()
		err := w.SetDisplay(frame)
		if err != nil {
			log.Printf("Error sending DisplayFrame Request: %s\n", err)
		}
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
