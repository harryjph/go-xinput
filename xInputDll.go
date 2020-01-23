package go_xinput

import "syscall"

var (
	dll *syscall.LazyDLL
	LoadError error

	xInputEnable *syscall.LazyProc
	xInputGetAudioDeviceIds *syscall.LazyProc
	xInputGetBatteryInformation *syscall.LazyProc
	xInputGetCapabilities *syscall.LazyProc
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
	xInputGetState = loadProc("XInputGetState")
	xInputSetState = loadProc("XInputSetState")
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
