package go_xinput

const (
	xinputIndex1 = 0
	xinputIndex2 = 1
	xinputIndex3 = 2
	xinputIndex4 = 3
)

func (controllerIndex ControllerIndex) toUserIndex() (uintptr, error) {
	switch controllerIndex {
	case Controller1: return xinputIndex1, nil
	case Controller2: return xinputIndex2, nil
	case Controller3: return xinputIndex3, nil
	case Controller4: return xinputIndex4, nil
	default: return 0, ControllerIndexOutOfRange
	}
}
