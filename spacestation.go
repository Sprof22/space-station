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
	resp, err := http.Get(relayURL + "/provide")
	if err != nil {
		fmt.Println("Error requesting data packet:", err)
		return
	}
	defer resp.Body.Close()

	var packet DataPacket
	err = json.NewDecoder(resp.Body).Decode(&packet)
	if err != nil {
		fmt.Println("Error decoding data packet:", err)
		return
	}

	// Process the received data packet
	fmt.Printf("SpaceStation %d received and processed data: %d\n", ss.ID, packet.Data)

	ss.RequestChan <- true
}
