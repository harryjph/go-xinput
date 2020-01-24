package go_xinput

func DisableInput() error {
	return setEnabled(false)
}

func EnableInput() error {
	return setEnabled(true)
}
