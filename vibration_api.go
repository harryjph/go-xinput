package go_xinput

func SetVibration(controllerIndex ControllerIndex, lowFrequencyAmount float32, highFrequencyAmount float32) error {
	return setState(controllerIndex, lowFrequencyAmount, highFrequencyAmount)
}
