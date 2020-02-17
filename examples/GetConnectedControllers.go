package main

import (
	"fmt"
	"github.com/harry1453/go-xinput/xinput"
	"log"
)

func main() {
	if xinput.LoadError != nil {
		log.Fatalln(xinput.LoadError)
	}

	connectedControllers := xinput.GetConnectedControllers()
	if len(connectedControllers) == 0 {
		fmt.Println("No controllers connected.")
	} else {
		fmt.Println("Connected controllers: ", connectedControllers)
	}
}
