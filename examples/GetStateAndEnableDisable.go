package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	xInput "go-xinput"
	"log"
	"os"
	"strings"
	"time"
)

func main() {
	if xInput.LoadError != nil {
		log.Fatalln(xInput.LoadError)
	}

	go monitorInput()

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

func monitorInput() {
	for {
		reader := bufio.NewReader(os.Stdin)
		fmt.Println("Enter command (on/off)")
		text, _ := reader.ReadString('\n')
		switch strings.TrimSpace(text) {
		case "on":
			enable()
		case "off":
			disable()
		}
	}
}

func enable() {
	if err := xInput.EnableInput(); err != nil {
		log.Fatalln(err)
	}
}

func disable() {
	if err := xInput.DisableInput(); err != nil {
		log.Fatalln(err)
	}
}
