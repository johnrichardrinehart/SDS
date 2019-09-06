package schemas

import (
	"time"

	"github.com/johnrichardrinehart/go2openapi"
)

type status struct {
	PacketIndex    int       `json:"packetIndex"`
	Path           string    `json:"path"`
	LastReceivedAt time.Time `json:"lastReceivedAt"`
	Ownership      `json:"ownership"`
}

type Ownership struct {
	Source            string `json:"source"`
	HeartbeatInterval int    `json:"heartbeatInterval"`
}

var StreamSchema = map[string]go2openapi.MediaType{
	"application/json": go2openapi.MediaType{
		Example: status{
			PacketIndex:    1050,
			Path:           "/home/12/smartModules/carData/streams/modeCar",
			LastReceivedAt: time.Now(),
			Ownership: Ownership{
				Source:            "deviceID:1234",
				HeartbeatInterval: 900,
			},
		},
	},
}

var StreamsSchema = map[string]go2openapi.MediaType{
	"application/json": go2openapi.MediaType{
		Example: []status{
			status{
				PacketIndex:    1050,
				Path:           "/home/12/smartModules/carData/streams/modeCar",
				LastReceivedAt: time.Now(),
				Ownership: Ownership{
					Source:            "deviceID:1234",
					HeartbeatInterval: 900,
				},
			},
			status{
				PacketIndex:    530,
				Path:           "/home/12/smartModules/carData/streams/acmeCar",
				LastReceivedAt: time.Now(),
				Ownership: Ownership{
					Source:            "deviceID:5678",
					HeartbeatInterval: 600,
				},
			},
		},
	},
}
