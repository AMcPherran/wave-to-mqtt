package main

func RemapFloat(x, in_min, in_max, out_min, out_max float32) float32 {
	return (x-in_min)*(out_max-out_min)/(in_max-in_min) + out_min
}
