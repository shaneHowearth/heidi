package sender

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"time"

	"github.com/shanehowearth/heidi/events/profile/v1"
	"github.com/shanehowearth/heidi/profile/bus"
)

func SendMessages(carrier bus.Carrier, topic string) {

	events := []profile.Updated{{
		PatientID:  1,
		FirstName:  "John",
		LastName:   "Doe",
		IsPregnant: false,
		UpdatedAt:  time.Now().Format(time.RFC3339Nano),
	},
		{
			PatientID:  1,
			FirstName:  "John",
			LastName:   "Doe",
			IsPregnant: true,
			UpdatedAt:  time.Now().Format(time.RFC3339Nano),
		},
	}

	for idx, evt := range events {
		marshaledEvent, err := json.Marshal(evt)
		if err != nil {
			log.Printf("Unable to marshal event %d with error %v", idx, err)
			continue
		}

		if err := carrier.SendMessage(context.Background(), []byte(fmt.Sprintf("Event%02d", idx)), marshaledEvent, topic); err != nil {
			log.Printf("Unable to send event %d with error %v", idx, err)
		}

	}

}
