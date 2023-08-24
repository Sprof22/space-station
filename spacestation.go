package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type SpaceStation struct {
	ID          int
	RequestChan chan bool
	DataChan    chan DataPacket
}

func (ss *SpaceStation) RequestAndProcessDataPacket(relayURL string) {
	for {
		// Request data packet from the relay endpoint
		resp, err := http.Get(relayURL + "/provide")
		if err != nil {
			fmt.Println("Error requesting data packet:", err)
			continue
		}

		var packet DataPacket
		err = json.NewDecoder(resp.Body).Decode(&packet)
		if err != nil {
			fmt.Println("Error decoding data packet:", err)
			resp.Body.Close()
			continue
		}

		resp.Body.Close()

		// Process the received data packet
		fmt.Printf("Space Station %d received and processed data: %+v\n", ss.ID, packet) // Print the data packet

		// Signal readiness to receive more data
		// ss.RequestChan <- true
	}
}
