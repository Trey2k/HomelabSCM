package main

import (
	"log"

	"homelabscm.com/scm/internal/app/homelab_scm"
)

func main() {
	log.Fatal(homelab_scm.Run())
}