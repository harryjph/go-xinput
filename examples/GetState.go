package main

import (
	"encoding/json"
	xInput "go-xInput"
	"log"
	"os"
	"time"
)

func main() {
	if xInput.LoadError != nil {
		log.Fatalln(xInput.LoadError)
	}

	var oldState *xInput.ControllerState
	for {
		newState, err := xInput.GetControllerState(xInput.Controller1)
		if err != nil {
			log.Fatalln(err)
		}
		if oldState == nil || *newState != *oldState {
			if err := json.NewEncoder(os.Stdout).Encode(newState); err != nil {
				log.Fatalln(err)
			}
			oldState = newState
		}
		time.Sleep(50 * time.Millisecond)
	}
}
