package main

import (
	"log"

	"github.com/gobuffalo/pop"
)

func test() {
	_, err := pop.Connect("development")
	if err != nil {
		log.Panic(err)
	}
}
