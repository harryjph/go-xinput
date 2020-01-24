package go_xinput

import (
	"errors"
	"fmt"
)

var FunctionNotAvailable = errors.New("the selected operation is not supported by the XInput library on this system")
var ControllerIndexOutOfRange = errors.New("controller index must be in range 0-3, it is recommended to use the available constants")

type ControllerIndex uint8

const (
	Controller1 ControllerIndex = iota
	Controller2
	Controller3
	Controller4
)

func (index ControllerIndex) String() string {
	switch index {
	case Controller1:
		return "Controller 1"
	case Controller2:
		return "Controller 2"
	case Controller3:
		return "Controller 3"
	case Controller4:
		return "Controller 4"
	default:
		return fmt.Sprintf("Unrecognized Controller Index: %02x", uint8(index))
	}
}
