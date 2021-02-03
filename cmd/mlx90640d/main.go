package main

import (
	"log"

	"golang.org/x/exp/io/i2c"
)

var version = "<not set>"

func main() {
	if err := runMain(); err != nil {
		log.Fatal(err)
	}
}

func runMain() error {
	log.SetFlags(0) // Removes timestamp from logs

	log.Printf("version: %s", version)

	d, err := i2c.Open(&i2c.Devfs{Dev: "/dev/i2c-1"}, 0x33)
	if err != nil {
		return err
	}

	_ = d

	return nil
}
