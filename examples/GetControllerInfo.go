package main

import (
	"encoding/json"
	"github.com/harry1453/go-xinput/xinput"
	"log"
	"os"
)

func main() {
	if xinput.LoadError != nil {
		log.Fatalln(xinput.LoadError)
	}

	info, err := xinput.GetControllerInfo(xinput.Controller1)
	if err != nil {
		log.Fatalln(err)
	}

	if err := json.NewEncoder(os.Stdout).Encode(info); err != nil {
		log.Fatalln(err)
	}
}
