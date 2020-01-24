package go_xinput

type Joystick struct {
	X float32
	Y float32
}

type Buttons struct {
	A             bool
	B             bool
	X             bool
	Y             bool
	DpadUp        bool
	DpadDown      bool
	DpadLeft      bool
	DpadRight     bool
	LeftShoulder  bool
	RightShoulder bool
	LeftJoystick  bool
	RightJoystick bool
	Start         bool
	Back          bool
}

type ControllerState struct {
	Buttons       Buttons
	LeftJoystick  Joystick
	RightJoystick Joystick
	LeftTrigger   float32
	RightTrigger  float32
}

func GetControllerState(controllerIndex ControllerIndex) (*ControllerState, error) {
	return getState(controllerIndex)
}
