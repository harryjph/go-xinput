//+build !windows

package xinput

func setEnabled(enabled bool) error {
	return UnsupportedOS
}
