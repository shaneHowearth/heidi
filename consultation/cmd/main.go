package main

import (
	"log"
	"os"
	"strconv"
	"strings"

	receiver "github.com/shanehowearth/heidi/consultation"
	"github.com/shanehowearth/heidi/consultation/bus/kafka"
)

func main() {
	busAddr, ok := os.LookupEnv("BUS_ADDRESS")
	if !ok {
		panic("BUS_ADDRESS has not been set")
	}

	// BUS_ADDRESS must be LOCATION:PORT for this instance.
	// We cannot check if this is an IP, or a DNS name.
	addressInfo := strings.Split(busAddr, ":")
	if len(addressInfo) != 2 {
		panic("BUS_ADDRESS must be in the form LOCATION:PORT")
	}

	// Check that the port value is possible.
	port, err := strconv.Atoi(addressInfo[1])
	if err != nil || port < 0 || port >= 65536 {
		panic("BUS_ADDRESS port is invalid")
	}

	profileTopic, ok := os.LookupEnv("PROFILE_TOPIC")
	if !ok {
		panic("PROFILE_TOPIC has not been set")
	}

	productConsumer, err := kafka.NewConsumer(busAddr, profileTopic)
	if err != nil {
		log.Print("Unable to create new bus")
		panic(err.Error())
	}

	receiver.HandleMessages(productConsumer)
}
