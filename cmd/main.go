package main

import (
	"log"
	"mtgcollector/internal"
)

func main() {
	log.Println("MTG Collector Application !")

	internal.ParseFile("/home/kaio/Projects/kaio/mtgcollector/example.json")
}
