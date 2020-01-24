package go_xinput

// https://docs.microsoft.com/en-us/windows/win32/xinput/xinput-and-controller-subtypes
type ControllerType uint8

const (
	UnknownType ControllerType = iota
	Gamepad
	Wheel
	ArcadeStick
	FlightStick
	DancePad
	Guitar
	GuitarAlternative
	GuitarBass
	DrumKit
	ArcadePad
)

type ControllerFeatures struct {
	VoiceSupported     bool
	VibrationSupported bool
	Wireless           bool
	// Reports whether the device supports plug-in modules. Although XInput can
	// report whether the device supports plug-in modules, it does not currently support them.
	// For example, the text input device (keyboard) does not work through XInput.
	PluginModulesSupported bool
	HasNavigationButtons   bool
}

type ControllerInfo struct {
	Subtype  ControllerType
	Features ControllerFeatures
	// TODO figure out what these 2 do... do they just report current state?
	State     ControllerState
	Vibration ControllerVibration
}

func GetControllerInfo(controllerIndex ControllerIndex) (*ControllerInfo, error) {
	return getControllerInfo(controllerIndex)
}
