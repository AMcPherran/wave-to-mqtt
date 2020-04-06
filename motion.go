package main

import (
	"fmt"
	"math"

	gowave "github.com/AMcPherran/go-wave"
)

func handleMotion(w *gowave.Wave, lastState *gowave.WaveState) {
	mb := w.State.Buttons.Middle()

	md := w.State.GetMotionData()
	lastMD := lastState.GetMotionData()
	lastState.SetMotionData(md)

	// When the button is released, print the difference from start pos to finish pos
	if (mb.Action == "LongUp" || mb.Action == "ExtraLongUp") && mb != lastState.Buttons.Middle() {
		processMotion(lastMD)
	}

	if mb != lastState.Buttons.Middle() {
		handleMiddleButton(w, mb)
		lastState.Buttons.Set(mb)
	}
}

func processMotion(m gowave.MotionData) {
	euler := m.Euler
	x := float64(euler.X)
	y := float64(euler.Y)
	z := float64(euler.Z)
	var keyDimension string
	// Determine which dimension had most significant movement
	if math.Abs(z) > 0.8 {
		keyDimension = "z"
	} else if math.Abs(x) > math.Abs(y) {
		keyDimension = "x"
	} else {
		keyDimension = "y"
	}
	if keyDimension == "x" {
		v := fmt.Sprintf("%f", euler.X)
		token := mqClient.Publish("wave/euler/x", 0, false, v)
		token.Wait()
	} else if keyDimension == "y" {
		v := fmt.Sprintf("%f", euler.Y)
		token := mqClient.Publish("wave/euler/y", 0, false, v)
		token.Wait()
	} else if keyDimension == "z" {
		v := fmt.Sprintf("%f", euler.Z)
		token := mqClient.Publish("wave/euler/z", 0, false, v)
		token.Wait()
	}

}
