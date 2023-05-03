package receiver

import (
	"fmt"

	"github.com/shanehowearth/heidi/consultation/bus"
)

func HandleMessages(carrier bus.Carrier) {
	for {
		// Read one message at a time from the carrier
		tmp, err := carrier.Read(1)
		if err != nil {
			fmt.Println("ERROR reading", err)
			break
		}

		// TODO: Actually do something with the message.
		fmt.Println(string(tmp[0]))
	}
}
