package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type RelayEndpoint struct {
	DataChan chan DataPacket
}

func (re *RelayEndpoint) AcceptDataPacket(w http.ResponseWriter, r *http.Request) {
	var packet DataPacket
	err := json.NewDecoder(r.Body).Decode(&packet)
	if err != nil {
		http.Error(w, "Failed to decode data packet", http.StatusBadRequest)
		return
	}

	fmt.Printf("Received data packet: %+v\n", packet)

	re.DataChan <- packet

	w.WriteHeader(http.StatusOK)
}

func (re *RelayEndpoint) ProvideDataPacket(w http.ResponseWriter, r *http.Request) {
	packet := <-re.DataChan

	// Print the data packet before sending it instead of console log
	fmt.Printf("Providing data packet: %+v\n", packet)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(packet)
}
