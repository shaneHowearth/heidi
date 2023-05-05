package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLocalMain(t *testing.T) {
	testcases := map[string]struct {
		errorString     string
		busAddr         string
		setBusAddr      bool // need this so we can set empty string for busAddr
		profileTopic    string
		setProfileTopic bool
	}{
		"no bus address": {
			errorString: "BUS_ADDRESS has not been set",
		},
		"bus address isn't addr:port": {
			setBusAddr:  true,
			errorString: "BUS_ADDRESS must be in the form LOCATION:PORT",
		},
		"bad bus address port": {
			setBusAddr:  true,
			busAddr:     "bad:port",
			errorString: "BUS_ADDRESS port is invalid",
		},
		"no profile topic": {
			setBusAddr:  true,
			busAddr:     "test:600",
			errorString: "PROFILE_TOPIC has not been set",
		},
	}
	for name, tc := range testcases {
		t.Run(name, func(t *testing.T) {
			if tc.setBusAddr {
				t.Setenv("BUS_ADDRESS", tc.busAddr)
			}
			if tc.setProfileTopic {
				t.Setenv("PROFILE_TOPIC", tc.profileTopic)
			}
			assert.PanicsWithValue(t, tc.errorString, main, "main did not panic")
		})
	}

}
