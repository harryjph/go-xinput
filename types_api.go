package go_xinput

// TODO this might have to be a uint
type ControllerIndex uint8
const (
	Controller1 ControllerIndex = iota
	Controller2
	Controller3
	Controller4
)

type Vector2 struct {
	X float32
	Y float32
}

type ControllerState struct {

}
