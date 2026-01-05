package webrtcpeer

import (
	"encoding/json"
	"fmt"

	"github.com/pion/webrtc/v3"
)

func CreatePeer(send func([]byte)) *webrtc.PeerConnection {
	pc, _ := webrtc.NewPeerConnection(webrtc.Configuration{})

	pc.OnConnectionStateChange(func(s webrtc.PeerConnectionState) {
		fmt.Println("Connection State:", s.String())
	})

	pc.OnDataChannel(func(dc *webrtc.DataChannel) {
		dc.OnMessage(func(msg webrtc.DataChannelMessage) {
			fmt.Println("Chat:", string(msg.Data))
			dc.SendText("Server echo: " + string(msg.Data))
		})
	})

	pc.OnICECandidate(func(c *webrtc.ICECandidate) {
		if c != nil {
			data, _ := json.Marshal(c.ToJSON())
			send(data)
		}
	})

	return pc
}
