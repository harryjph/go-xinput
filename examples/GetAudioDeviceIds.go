package main

import (
	"encoding/hex"
	"fmt"
	xinput "go-xinput"
	"log"
)

func main() {
	if xinput.LoadError != nil {
		log.Fatalln(xinput.LoadError)
	}

	inputDeviceId, outputDeviceId, err := xinput.GetAudioDeviceIds(xinput.Controller1)
	if err != nil {
		log.Fatalln(err)
	}

	inputDeviceIdHex := hex.EncodeToString(inputDeviceId)
	outputDeviceIdHex := hex.EncodeToString(outputDeviceId)
	fmt.Println("Input Device ID: {}, Output Device ID: {}", inputDeviceIdHex, outputDeviceIdHex)
}
