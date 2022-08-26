package main

import (
	"github.com/XelaMP/getIP.git/cmd/bootstrap"
	"log"
)

func main() {
	if err := bootstrap.Run(); err != nil {
		log.Fatal(err)
	}
}
