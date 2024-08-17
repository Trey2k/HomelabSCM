package main

import (
	"log"

	"homelabscm.com/scm/internal/app/homelab_scm_shell"
)

func main() {
	log.Fatal(homelab_scm_shell.Run())
}