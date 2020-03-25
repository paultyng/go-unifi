package main

import "github.com/google/uuid"

type message struct {
	CMD       string      `json:"cmd"`      // "sdp_exchange"
	DeviceID  string      `json:"deviceId"` // TODO: also as device_id?
	RequestID string      `json:"requestId"`
	Payload   interface{} `json:"payload"`
}

type sdpOfferPayload struct {
	Password string `json:"password"`
	SDPOffer string `json:"sdpOffer"`
	STUNURI  string `json:"stunUri"`
	TURNURI  string `json:"turnUri"`
	TTL      int    `json:"ttl"`
	Type     string `json:"type"` // "OFFER"
	Username string `json:"username"`
}

type sdpAnswerPayload struct {
	SDP      string `json:"sdpAnswer"`
	Type     string `json:"type"` // "ANSWER"
	WebRTCID string `json:"webRtcId"`
}

func newMessage(cmd, deviceID string, payload interface{}) *message {
	rID, err := uuid.NewRandom()
	if err != nil {
		panic(err)
	}
	return &message{
		CMD:       cmd,
		DeviceID:  deviceID,
		RequestID: rID.String(),
		Payload:   payload,
	}
}
