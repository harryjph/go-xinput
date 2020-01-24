package go_xinput

import "errors"

var FunctionNotAvailable = errors.New("the selected operation is not supported by the XInput library on this system")
var ControllerIndexOutOfRange = errors.New("controller index must be in range 0-3, it is recommended to use the available constants")

type ControllerIndex uint8

const (
	Controller1 ControllerIndex = iota
	Controller2
	Controller3
	Controller4
)
