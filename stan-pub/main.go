package main

import (
	"flag"
	"log"
	"os"

	"github.com/nats-io/stan.go"
)


func main() {

	var model string
	flag.StringVar(&model, "model-path", "stan-pub/model.json", "path to model json")
	flag.Parse()

	file, err := os.ReadFile(model)
	if err != nil {
		  log.Fatal(err)
	}
	sc, err := stan.Connect("test-cluster", "pub")
	if err != nil {
		log.Fatalf("Can't connect: %v.\n", err)
	}
	defer sc.Close()

	subj, msg := "foo", []byte(file)


	err = sc.Publish(subj, msg)
	if err != nil {
		log.Fatalf("Error during publish: %v\n", err)
	}
	log.Printf("Published [%s] : '%s'\n", subj, msg)
}
