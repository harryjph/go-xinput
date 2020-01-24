package go_xinput

const (
	xInputControllerIndex1 = 0
	xInputControllerIndex2 = 1
	xInputControllerIndex3 = 2
	xInputControllerIndex4 = 3
)

func (controllerIndex ControllerIndex) toUserIndex() (uintptr, error) {
	switch controllerIndex {
	case Controller1:
		return xInputControllerIndex1, nil
	case Controller2:
		return xInputControllerIndex2, nil
	case Controller3:
		return xInputControllerIndex3, nil
	case Controller4:
		return xInputControllerIndex4, nil
	default:
		return 0, ControllerIndexOutOfRange
	}
}
