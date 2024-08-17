package homelab_web_client

import (
	"fmt"
	"log"
)

func Run() error {
	currentPath := getURLPath()

	switch currentPath {
	case "/":
		// Do something
	case "/install":
		StartInstaller()
	default:
		log.Fatal("Unknown path")
	}

	fmt.Println("Running")

	<-make(chan struct{})

	return nil
}
