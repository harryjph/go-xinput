package go_xinput

import (
	"fmt"
	"syscall"
	"unsafe"
)

const (
	xinputBatteryDevtypeGamepad   = 0x00
	xinputBatteryDevtypeHeadset   = 0x01

	xinputBatteryTypeDisconnected = 0x00
	xinputBatteryTypeWired        = 0x01
	xinputBatteryTypeAlkaline     = 0x02
	xinputBatteryTypeNimh         = 0x03
	xinputBatteryTypeUnknown      = 0xFF

	xinputBatteryLevelEmpty       = 0x00
	xinputBatteryLevelLow         = 0x01
	xinputBatteryLevelMedium      = 0x02
	xinputBatteryLevelFull        = 0x03
)

type xinputBatteryInformation struct {
	batteryType uint8
	batteryLevel uint8
}

func (info *xinputBatteryInformation) toBatteryInformation() (*BatteryInformation, error) {
	var batteryType BatteryType
	switch info.batteryType {
	case xinputBatteryTypeDisconnected: batteryType = Disconnected
	case xinputBatteryTypeWired: batteryType = Wired
	case xinputBatteryTypeAlkaline: batteryType = Alkaline
	case xinputBatteryTypeNimh: batteryType = NiMH
	case xinputBatteryTypeUnknown: batteryType = Unknown
	default: return nil, fmt.Errorf("unrecognized battery type: %02x", info.batteryType)
	}

	var batteryLevel BatteryLevel
	switch info.batteryLevel {
	case xinputBatteryLevelEmpty: batteryLevel = Empty
	case xinputBatteryLevelLow: batteryLevel = Low
	case xinputBatteryLevelMedium: batteryLevel = Medium
	case xinputBatteryLevelFull: batteryLevel = Full
	default: return nil, fmt.Errorf("unrecognized battery level: %02x", info.batteryLevel)
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
	var xinputBatteryInformation xinputBatteryInformation
	r, _, _ := xInputGetBatteryInformation.Call(userIndex, devtype, uintptr(unsafe.Pointer(&xinputBatteryInformation)))
	if r == 0 {
		 return xinputBatteryInformation.toBatteryInformation()
	} else {
		return nil, syscall.Errno(r)
	}
}
