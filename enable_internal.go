package go_xinput

import "syscall"

func setEnabled(enabled bool) error {
	if xInputEnable == nil {
		return FunctionNotAvailable
	}

	var enabledInt uintptr
	if enabled {
		enabledInt = 1
	} else {
		enabledInt = 0
	}
	r, _, _ := xInputEnable.Call(enabledInt)
	if r == 0 {
		return nil
	} else {
		return syscall.Errno(r)
	}
}
