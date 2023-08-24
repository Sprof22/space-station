package main

import (
	"net/http"
	"time"
)

func main() {
	relay := RelayEndpoint{
		DataChan: make(chan DataPacket),
	}

	http.HandleFunc("/accept", relay.AcceptDataPacket)
	http.HandleFunc("/provide", relay.ProvideDataPacket)

	go http.ListenAndServe(":8080", nil)

	for i := 1; i <= 5; i++ {
		starship := Starship{
			ID: i,
		}
		go starship.GatherAndSendDataPacket("http://localhost:8080")
	}

	for i := 1; i <= 3; i++ {
		station := SpaceStation{
			ID:          i,
			RequestChan: make(chan bool),
			DataChan:    make(chan DataPacket),
		}
		go station.RequestAndProcessDataPacket("http://localhost:8080")
	}

	time.Sleep(30 * time.Second)
}
