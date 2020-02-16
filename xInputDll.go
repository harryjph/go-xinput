//+build windows

package go_xinput

import (
	"syscall"
)

var (
	dll *syscall.LazyDLL

	xInputEnable *syscall.LazyProc

	// Only available in XInput >= 1.4
	xInputGetAudioDeviceIds *syscall.LazyProc

	xInputGetBatteryInformation *syscall.LazyProc

	xInputGetCapabilities *syscall.LazyProc

	// Only available in XInput < 1.4
	xInputGetDSoundAudioDeviceGuids *syscall.LazyProc

	xInputGetKeystroke *syscall.LazyProc

	xInputGetState *syscall.LazyProc

	xInputSetState *syscall.LazyProc
)

func init() {
	dll, LoadError = loadDll("xinput1_4.dll")
	if LoadError != nil {
		dll, LoadError = loadDll("xinput1_3.dll")
		if LoadError != nil {
			dll, LoadError = loadDll("xinput9_1_0.dll")
			if LoadError != nil {
				return
			}
		}
	}

	xInputEnable = loadProc("XInputEnable")
	xInputGetAudioDeviceIds = loadProc("XInputGetAudioDeviceIds")
	xInputGetBatteryInformation = loadProc("XInputGetBatteryInformation")
	xInputGetCapabilities = loadProc("XInputGetCapabilities")
	xInputGetDSoundAudioDeviceGuids = loadProc("XInputGetDSoundAudioDeviceGuids")
	xInputGetKeystroke = loadProc("XInputGetKeystroke")
	xInputSetState = loadProc("XInputSetState")

	// TODO you're supposed to (in C) call GetProcAddress(HMODULE(hGetProcIDDLL), (LPCSTR)100);... This doesn't seem to work here. Does it only work with xinput 1.3?
	xInputGetState = loadProc(string(100))
	if xInputGetState == nil {
		//log.Println("Failed to load guide-supporting GetState")
		xInputGetState = loadProc("XInputGetState")
	}
}

func loadDll(name string) (*syscall.LazyDLL, error) {
	dll := syscall.NewLazyDLL(name)
	return dll, dll.Load()
}

func loadProc(name string) *syscall.LazyProc {
	proc := dll.NewProc(name)
	if err := proc.Find(); err == nil {
		return proc
	} else {
		return nil
	}
}
