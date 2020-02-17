package xinput

type ControllerVibration struct {
	LowFrequencyLevel  float32
	HighFrequencyLevel float32
}

func SetVibration(controllerIndex ControllerIndex, vibration ControllerVibration) error {
	return setState(controllerIndex, vibration.HighFrequencyLevel, vibration.LowFrequencyLevel)
}
