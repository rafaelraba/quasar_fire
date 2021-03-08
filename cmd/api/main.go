package main

import (
	"github.com/rafaelraba/quasar_fire/cmd/api/bootstrap"
	"log"
)

func main () {
	if err := bootstrap.Run(); err != nil {
		log.Fatal(err)
	}
}