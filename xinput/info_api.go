package xinput

import "fmt"

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

func (controllerType ControllerType) String() string {
	switch controllerType {
	case UnknownType:
		return "Unknown"
	case Gamepad:
		return "Gamepad"
	case Wheel:
		return "Wheel"
	case ArcadeStick:
		return "Arcade Stick"
	case FlightStick:
		return "Flight Stick"
	case DancePad:
		return "Dance Pad"
	case Guitar:
		return "Guitar"
	case GuitarAlternative:
		return "Alternative Guitar"
	case GuitarBass:
		return "Bass Guitar"
	case DrumKit:
		return "Drum Kit"
	case ArcadePad:
		return "Arcade Pad"
	default:
		return fmt.Sprintf("Unrecognized Controller Type: %02x", uint8(controllerType))
	}
}

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
