package nats

import (
	"log"
	"github.com/nats-io/stan.go"
)


const (
	ClusterID 	string = "test-cluster"
	NatsUrl 	string = "nats://nats-streaming-container:4222"

)


func NewServer(ClientID string) (stan.Conn, error) {
	sc, err := stan.Connect(ClusterID, ClientID, stan.NatsURL(NatsUrl))

	if err != nil {
		log.Fatal("Error establishing connection to NATS:", err)
		return nil, err
	}
	log.Print("Connection to the NATS is established")
	return sc, nil
}