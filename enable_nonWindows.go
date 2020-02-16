//+build !windows

package go_xinput

func setEnabled(enabled bool) error {
	return UnsupportedOS
}
