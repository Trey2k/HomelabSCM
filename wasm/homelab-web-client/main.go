package main

import (
	"log"

	"homelabscm.com/scm/internal/app/homelab_web_client"
)

func main() {
	log.Fatal(homelab_web_client.Run())
}