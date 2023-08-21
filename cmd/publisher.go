package main

import (
	"github.com/pamallika/WBL0v2/internal/service/publisher"
	"log"
)

const (
	NATSStreamingURL = "localhost:4223"
	clusterID        = "test-cluster"
	clientID         = "test-publisher"
	channel          = "test-channel"
)

func main() {
	nc := publisher.CreateSTAN()
	err := nc.Connect(clusterID, clientID, NATSStreamingURL)
	defer nc.Close()
	if err != nil {
		log.Fatalf("Error connecting to nats : %s", err)
	}
	_ = nc.PublishFromCLI(channel)
}
