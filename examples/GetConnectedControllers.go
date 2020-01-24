package main

import (
	"fmt"
	xinput "go-xinput"
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
