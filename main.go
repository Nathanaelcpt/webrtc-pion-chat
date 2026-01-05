package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/pion/webrtc/v3"

	"webrtc-pion-chat/signaling"
	webrtcpeer "webrtc-pion-chat/webrtc"
)

func main() {
	http.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
		var pc = webrtcpeer.CreatePeer(func(msg []byte) {})

		signaling.HandleWS(w, r, func(msg []byte) {
			var sdp map[string]interface{}
			json.Unmarshal(msg, &sdp)

			if sdp["type"] == "offer" {
				b, _ := json.Marshal(sdp)
				var offer = webrtc.SessionDescription{}
				json.Unmarshal(b, &offer)
				pc.SetRemoteDescription(offer)

				answer, _ := pc.CreateAnswer(nil)
				pc.SetLocalDescription(answer)

				res, _ := json.Marshal(answer)
				fmt.Println(string(res))
			}
		})
	})

	http.ListenAndServe(":8080", nil)
}
