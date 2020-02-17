# go-xinput: XInput wrapper for Go

`go-xinput` is a complete XInput wrapper to use XInput controllers in a Golang application. It has a type-safe API, does not require CGo, has no dependencies other than on the Go standard library, and has stubs that produce errors on operating systems other than Windows (so it will compile on other OSes, but will not work).

It exposes all* available function calls for XInput. A comprehensive set of examples is available in the [examples directory](examples).

Use import path `github.com/harry1453/go-xinput/xinput`.

Whilst this library will compile to any target, it only works under Windows.

\* Does not expose `GetKeystroke` but this is planned, and does not expose `GetDSoundAudioDeviceGuids` as this is not included in modern versions of XInput. `GetAudioDeviceIDs` support is currently broken.
  