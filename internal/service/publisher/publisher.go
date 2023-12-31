package publisher

import (
	"fmt"
	"github.com/nats-io/stan.go"
	"log"
	"os"
)

type StanClient struct {
	sc stan.Conn
}

func CreateSTAN() *StanClient {
	st := StanClient{}
	return &st
}

func (sCli *StanClient) Connect(clusterID string, clientID string, URL string) error {
	sc, err := stan.Connect(clusterID, clientID, stan.NatsURL(URL))
	if err != nil {
		return err
	}
	sCli.sc = sc
	return err
}

func (sCli *StanClient) Close() {
	if sCli != nil {
		sCli.sc.Close()
	}
}

func (sCli *StanClient) PublishFromFile(channel string, filepath string) error {
	text, err := os.ReadFile(filepath)
	if err != nil {
		return err
	}
	sCli.sc.Publish(channel, text)
	return err
}

func (sCli *StanClient) PublishFromCLI(channel string) error {
	var filepath string
	var err error
	for {
		var text []byte
		fmt.Print("Enter filepath: ")
		fmt.Fscan(os.Stdin, &filepath)
		if filepath == "exit" {
			return nil
		}
		text, err = os.ReadFile(filepath)
		if err != nil {
			log.Printf("Error reading file : %s", err)
			return err
		}
		err := sCli.sc.Publish(channel, text)
		if err != nil {
			log.Fatalf("Error publish to queue : %s", err)
		}
	}
	return err
}
