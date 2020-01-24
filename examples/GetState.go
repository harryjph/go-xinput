package main

import (
	"encoding/json"
	xinput "go-xinput"
	"log"
	"os"
	"time"
)

func main() {
	if xinput.LoadError != nil {
		log.Fatalln(xinput.LoadError)
	}

	var oldState *xinput.ControllerState
	for {
		newState, err := xinput.GetControllerState(xinput.Controller1)
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
