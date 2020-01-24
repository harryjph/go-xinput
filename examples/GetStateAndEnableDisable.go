package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	xinput "go-xinput"
	"log"
	"os"
	"strings"
	"time"
)

func main() {
	if xinput.LoadError != nil {
		log.Fatalln(xinput.LoadError)
	}

	go monitorInput()

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
	if err := xinput.EnableInput(); err != nil {
		log.Fatalln(err)
	}
}

func disable() {
	if err := xinput.DisableInput(); err != nil {
		log.Fatalln(err)
	}
}
