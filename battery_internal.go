package go_xinput

import (
	"fmt"
	"syscall"
	"unsafe"
)

const (
	xInputBatteryDevtypeGamepad = 0b00
	xInputBatteryDevtypeHeadset = 0x01

	xInputBatteryTypeDisconnected = 0x00
	xInputBatteryTypeWired        = 0x01
	xInputBatteryTypeAlkaline     = 0x02
	xInputBatteryTypeNimh         = 0x03
	xInputBatteryTypeUnknown      = 0xFF

	xInputBatteryLevelEmpty  = 0x00
	xInputBatteryLevelLow    = 0x01
	xInputBatteryLevelMedium = 0x02
	xInputBatteryLevelFull   = 0x03
)

type xInputBatteryInformation struct {
	batteryType  uint8
	batteryLevel uint8
}

func (info *xInputBatteryInformation) toBatteryInformation() (*BatteryInformation, error) {
	var batteryType BatteryType
	switch info.batteryType {
	case xInputBatteryTypeDisconnected:
		batteryType = Disconnected
	case xInputBatteryTypeWired:
		batteryType = Wired
	case xInputBatteryTypeAlkaline:
		batteryType = Alkaline
	case xInputBatteryTypeNimh:
		batteryType = NiMH
	case xInputBatteryTypeUnknown:
		batteryType = Unknown
	default:
		return nil, fmt.Errorf("unrecognized battery type: %02x", info.batteryType)
	}

	var batteryLevel BatteryLevel
	switch info.batteryLevel {
	case xInputBatteryLevelEmpty:
		batteryLevel = Empty
	case xInputBatteryLevelLow:
		batteryLevel = Low
	case xInputBatteryLevelMedium:
		batteryLevel = Medium
	case xInputBatteryLevelFull:
		batteryLevel = Full
	default:
		return nil, fmt.Errorf("unrecognized battery level: %02x", info.batteryLevel)
	}

	return &BatteryInformation{
		BatteryType:  batteryType,
		BatteryLevel: batteryLevel,
	}, nil
}

func getBatteryInformation(controllerIndex ControllerIndex, devtype uintptr) (*BatteryInformation, error) {
	if xInputGetBatteryInformation == nil {
		return nil, FunctionNotAvailable
	}
	userIndex, err := controllerIndex.toUserIndex()
	if err != nil {
		return nil, err
	}
	var xInputBatteryInformation xInputBatteryInformation
	r, _, _ := xInputGetBatteryInformation.Call(userIndex, devtype, uintptr(unsafe.Pointer(&xInputBatteryInformation)))
	if r == 0 {
		return xInputBatteryInformation.toBatteryInformation()
	} else {
		return nil, syscall.Errno(r)
	}
}
