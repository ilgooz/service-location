package main

import (
	"flag"
	"log"

	"github.com/ilgooz/service-location/location"
	"github.com/mesg-foundation/core/x/xsignal"
	mesg "github.com/mesg-foundation/go-service"
)

var (
	maxmindDBPath = flag.String("maxmindDBPath", "", "MaxMind city database path")
)

func main() {
	flag.Parse()

	service, err := mesg.New()
	if err != nil {
		log.Fatal(err)
	}

	ls := location.New(service, *maxmindDBPath)

	// start the location service.
	go func() {
		log.Println("location service has been started")

		if err := ls.Start(); err != nil {
			log.Fatal(err)
		}
	}()

	// wait for interrupt and gracefully shutdown the location service.
	<-xsignal.WaitForInterrupt()

	log.Println("shutting down...")

	if err := ls.Close(); err != nil {
		log.Fatal(err)
	}

	log.Println("shutdown")
}
