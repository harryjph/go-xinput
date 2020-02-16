package main

import (
	"encoding/json"
	"fmt"
	xinput "go-xinput"
	"log"
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
			if jsonData, err := json.Marshal(newState); err != nil {
				log.Fatalln(err)
			} else {
				fmt.Printf("%s\r", jsonData)
			}
			oldState = newState
		}
		time.Sleep(50 * time.Millisecond)
	}
}
