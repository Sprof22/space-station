package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"
	"time"
)

type Starship struct {
	ID        int
	DataPacks []DataPacket
}

func (s *Starship) GatherAndSendDataPacket(relayURL string) {
	for {
		packet := DataPacket{Data: rand.Intn(100) + 1}

		payload, _ := json.Marshal(packet)
		_, err := http.Post(relayURL+"/accept", "application/json", bytes.NewReader(payload))
		if err != nil {
			fmt.Println("Error sending data packet:", err)
		}

		time.Sleep(time.Duration(rand.Intn(1)) * time.Second)
	}

}
