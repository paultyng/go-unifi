package main

import (
	"bytes"
	"context"
	"encoding/binary"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"
	"time"

	awscreds "github.com/aws/aws-sdk-go/aws/credentials"
	awssigner "github.com/aws/aws-sdk-go/aws/signer/v4"
	mqtt "github.com/eclipse/paho.mqtt.golang"
	webrtc "github.com/pion/webrtc/v2"
)

const deviceID = "<removed>"

func init() {
	mqtt.DEBUG = log.New(os.Stdout, "[debug] ", 0)
	mqtt.ERROR = log.New(os.Stdout, "[error] ", 0)
	mqtt.WARN = log.New(os.Stdout, "[warn] ", 0)
	mqtt.CRITICAL = log.New(os.Stdout, "[critical] ", 0)
}

func main() {
	err := run(context.Background(), os.Args)
	if err != nil {
		log.Fatal(err)
	}
}

type cloudAccessConfig struct {
	Region    string `json:"region"`
	AccountID string `json:"accountId"`
	IOT       struct {
		Host                string `json:"host"`
		SitesListTopic      string `json:"sitesListTopic"`
		UserDataTopic       string `json:"userDataTopic"`
		PingTopic           string `json:"pingTopic"`
		DeviceRole          string `json:"deviceRole"`
		CredentialsProvider string `json:"credentialsProvider"`
	} `json:"iot"`
	APIGateway struct {
		URL string `json:"url"`
	} `json:"apiGateway"`
	APIGatewayUI struct {
		URL string `json:"url"`
	} `json:"apiGatewayUI"`
}

type credentials struct {
	AccessKeyID     string `json:"accessKeyId"`
	SecretKey       string `json:"secretKey"`
	TURNCredentials struct {
		URIs     []string    `json:"uris"`
		Password string      `json:"password"`
		Username string      `json:"username"`
		TTL      json.Number `json:"ttl"`
	} `json:"turnCredentials"`
	IdentityID   string `json:"identityId"`
	SessionToken string `json:"sessionToken"`
	Expiration   int64  `json:"expiration"`
	Region       string `json:"region"`
}

func findURI(uris []string, scheme string, transport string) string {
	for _, u := range uris {
		p, err := url.Parse(u)
		if err != nil {
			return ""
		}
		if p.Scheme != scheme {
			continue
		}
		if p.Query().Get("transport") != transport {
			continue
		}
		return u
	}
	return ""
}

func cloudDo(method, relativeURL string, reqBody interface{}, respBody interface{}, customizations ...func(req *http.Request, body io.ReadSeeker) error) error {
	baseURL, _ := url.Parse("https://cloudaccess.svc.ubnt.com")

	var (
		reqReader io.ReadSeeker
		err       error
		reqBytes  []byte
	)
	if reqBody != nil {

		reqBytes, err = json.Marshal(reqBody)
		if err != nil {
			return err
		}
		reqReader = bytes.NewReader(reqBytes)
	}

	reqURL, err := url.Parse(relativeURL)
	if err != nil {
		return err
	}

	url := baseURL.ResolveReference(reqURL)

	log.Printf("%s %s", method, url)

	req, err := http.NewRequest(method, url.String(), reqReader)
	if err != nil {
		return err
	}

	req.Header.Set("User-Agent", "go-unifi/0.1")

	for _, c := range customizations {
		err = c(req, reqReader)
		if err != nil {
			return err
		}
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusNotFound {
		return fmt.Errorf("not found")
	}

	if resp.StatusCode != 200 {
		// fmt.Printf("Request Body:\n%s\n", string(reqBytes))
		// errBody := struct {
		// 	Meta meta `json:"meta"`
		// }{}
		// err = json.NewDecoder(resp.Body).Decode(&errBody)
		// return fmt.Errorf("%s %s (%s) for %s %s", errBody.Meta.RC, errBody.Meta.Message, resp.Status, method, url.String())
		return fmt.Errorf("%s", resp.Status)
	}

	if respBody == nil || resp.ContentLength == 0 {
		return nil
	}

	respBytes, _ := ioutil.ReadAll(resp.Body)
	fmt.Printf("Response Body:\n%s\n", string(respBytes))

	err = json.NewDecoder(bytes.NewBuffer(respBytes)).Decode(respBody)
	if err != nil {
		return err
	}

	return nil
}

func run(ctx context.Context, args []string) error {
	log.Printf("creating credentials")
	creds := &credentials{}
	err := cloudDo("POST", "/create-credentials", &struct {
		WithTurn bool `json:"withTurn"`
	}{true}, creds, func(req *http.Request, _ io.ReadSeeker) error {
		// set unifi auth cookie
		req.AddCookie(&http.Cookie{
			Name:  "UBIC_AUTH",
			Value: "<removed>",
		})
		return nil
	})
	if err != nil {
		return err
	}

	var cac cloudAccessConfig
	{
		conf := struct {
			Obj cloudAccessConfig `json:"unifiCloudAccess"`
		}{}
		log.Printf("getting cloud access config")
		err = cloudDo("GET", "https://config.ubnt.com/cloudAccessConfig.json", nil, &conf)
		if err != nil {
			return err
		}
		cac = conf.Obj
	}

	log.Printf("cloud access config retrieved (host: %s)", cac.IOT.Host)

	awsCredentials := awscreds.NewStaticCredentials(creds.AccessKeyID, creds.SecretKey, creds.SessionToken)
	signer := awssigner.NewSigner(awsCredentials)

	stunURI := findURI(creds.TURNCredentials.URIs, "stun", "udp")
	turnURI := findURI(creds.TURNCredentials.URIs, "turn", "udp")

	config := webrtc.Configuration{
		ICETransportPolicy:   webrtc.ICETransportPolicyAll,
		BundlePolicy:         webrtc.BundlePolicyBalanced,
		RTCPMuxPolicy:        webrtc.RTCPMuxPolicyRequire,
		ICECandidatePoolSize: 5,
		SDPSemantics:         webrtc.SDPSemanticsUnifiedPlan,
		ICEServers: []webrtc.ICEServer{
			{
				Username:       creds.TURNCredentials.Username,
				CredentialType: webrtc.ICECredentialTypePassword,
				Credential:     creds.TURNCredentials.Password,
				URLs:           []string{stunURI, turnURI},
			},
		},
	}

	var pc *webrtc.PeerConnection
	{
		se := webrtc.SettingEngine{}
		se.SetTrickle(true)
		wrtcAPI := webrtc.NewAPI(webrtc.WithSettingEngine(se))
		log.Printf("new peer connection")
		pc, err = wrtcAPI.NewPeerConnection(config)
		if err != nil {
			return err
		}
	}

	offer2, err := pc.CreateOffer(nil)
	if err != nil {
		return err
	}
	log.Println(offer2.SDP)
	os.Exit(1)

	pc.OnICEConnectionStateChange(func(cs webrtc.ICEConnectionState) {
		fmt.Println(cs.String())
	})

	pc.OnDataChannel(func(d *webrtc.DataChannel) {
		d.OnOpen(func() {
			fmt.Println("generic open")
		})

		d.OnMessage(func(msg webrtc.DataChannelMessage) {
			fmt.Printf("gen: %s", string(msg.Data))
		})
	})

	var dcID uint16 = 1
	dc, err := pc.CreateDataChannel("update_default", &webrtc.DataChannelInit{
		ID: &dcID,
	})
	if err != nil {
		return err
	}

	// Set the handler for ICE connection state
	// This will notify you when the peer has connected/disconnected
	pc.OnICEConnectionStateChange(func(connectionState webrtc.ICEConnectionState) {
		fmt.Printf("ICE Connection State has changed: %s\n", connectionState.String())
	})

	// Register channel opening handling
	dc.OnOpen(func() {
		fmt.Printf("Data channel '%s'-'%d' open. Random messages will now be sent to any connected DataChannels every 5 seconds\n", dc.Label(), dc.ID())

		// for range time.NewTicker(5 * time.Second).C {
		// 	message := signal.RandSeq(15)
		// 	fmt.Printf("Sending '%s'\n", message)

		// 	// Send the message as text
		// 	sendErr := dataChannel.SendText(message)
		// 	if sendErr != nil {
		// 		panic(sendErr)
		// 	}
		// }
	})

	// Register text message handling
	dc.OnMessage(func(msg webrtc.DataChannelMessage) {
		fmt.Printf("Message from DataChannel '%s': '%s'\n", dc.Label(), string(msg.Data))
	})

	log.Print("create mqtt websocket URL")

	var mqttClient mqtt.Client
	{
		mqttURL, err := signAWSMQTTWSURL(cac.IOT.Host, cac.Region, creds.AccessKeyID, creds.SecretKey, creds.SessionToken)
		if err != nil {
			return err
		}

		mqttOpts := mqtt.NewClientOptions().AddBroker(mqttURL)
		// mqttOpts.SetCleanSession(true)
		// mqttOpts.SetAutoReconnect(true)
		mqttOpts.SetDefaultPublishHandler(func(client mqtt.Client, msg mqtt.Message) {
			log.Printf("[mqtt:%s] %s", msg.Topic(), string(msg.Payload()))
		})
		mqttOpts.SetClientID(creds.IdentityID)
		mqttOpts.SetMaxReconnectInterval(1 * time.Second)
		mqttClient = mqtt.NewClient(mqttOpts)

	}

	log.Print("connect to mqtt")
	if token := mqttClient.Connect(); token.Wait() && token.Error() != nil {
		return token.Error()
	}

	mqttClientChan := make(chan []byte)

	topic := fmt.Sprintf("client/%s/%s", creds.IdentityID, deviceID)
	if token := mqttClient.Subscribe(topic, 0, func(client mqtt.Client, msg mqtt.Message) {
		log.Printf("[mqtt:%s] %s", msg.Topic(), string(msg.Payload()))
		mqttClientChan <- msg.Payload()
	}); token.Wait() && token.Error() != nil {
		return token.Error()
	}

	log.Print("mqtt connected")

	log.Print("create offer")

	offer, err := pc.CreateOffer(nil)
	if err != nil {
		return err
	}

	err = pc.SetLocalDescription(offer)
	if err != nil {
		return err
	}

	log.Printf("sdp exchange")
	ttl, err := creds.TURNCredentials.TTL.Int64()
	if err != nil {
		return err
	}
	msg := newMessage("sdp_exchange", deviceID, &sdpOfferPayload{
		SDPOffer: offer.SDP,
		STUNURI:  stunURI,
		TURNURI:  turnURI,
		Username: creds.TURNCredentials.Username,
		Password: creds.TURNCredentials.Password,
		Type:     "OFFER",
		TTL:      int(ttl),
	})

	err = cloudDo("POST", "/post-message", msg, nil, func(req *http.Request, body io.ReadSeeker) error {
		_, err = signer.Sign(req, body, "execute-api", creds.Region, time.Now())
		return err
	})
	if err != nil {
		return err
	}

	sdpRemoteBytes := <-mqttClientChan
	var sdpRemote struct {
		SDP        string `json:"sdp"`
		SDPVersion string `json:"sdpVersion"`
		WebRTCID   int    `json:"webRtcId"`
		RequestID  string `json:"requestId"`
		DeviceID   string `json:"deviceId"`
		IdentityID string `json:"identityId"`
	}
	err = json.Unmarshal(sdpRemoteBytes, &sdpRemote)
	if err != nil {
		return err
	}

	err = pc.SetRemoteDescription(webrtc.SessionDescription{
		Type: webrtc.SDPTypeAnswer,
		SDP:  sdpRemote.SDP,
	})
	if err != nil {
		return err
	}

	/*

		{
			path:
			queryString:
		method:
		contentType:
		"Accept-Encoding": "gzip",
		length: 0,

		}

	*/

	apiMsgBuffer := bytes.NewBuffer([]byte{})
	{
		apiMsg := struct {
			Path           string `json:"path"`
			QueryString    string `json:"queryString"`
			Method         string `json:"method"`
			ContentType    string `json:"contentType"`
			Length         int    `json:"length"`
			AcceptEncoding string `json:"Accept-Encoding"`
		}{
			Path:           "/api/s/default/rest/networkconf",
			Method:         "GET",
			ContentType:    "application/json",
			Length:         0,
			AcceptEncoding: "gzip",
		}
		apiMsgBytes, err := json.Marshal(apiMsg)
		if err != nil {
			return err
		}

		// encode total length in 4 bytes (12+body len)
		err = binary.Write(apiMsgBuffer, binary.LittleEndian, uint32(12+len(apiMsgBytes)))
		if err != nil {
			return err
		}

		// encode message ID in 8 bytes
		err = binary.Write(apiMsgBuffer, binary.LittleEndian, uint64(0))
		if err != nil {
			return err
		}

		// encode body len in 4 bytes
		err = binary.Write(apiMsgBuffer, binary.LittleEndian, uint32(len(apiMsgBytes)))
		if err != nil {
			return err
		}

		_, err = apiMsgBuffer.Write(apiMsgBytes)
		if err != nil {
			return err
		}

	}

	// dcID = 3
	// apiDC, err := pc.CreateDataChannel("api", sdf)
	// if err != nil {
	// 	return err
	// }

	for dc.ReadyState() != webrtc.DataChannelStateOpen {
		log.Println("not yet open...")
		time.Sleep(5 * time.Second)
	}

	// err = dc.Send(apiMsgBuffer.Bytes())
	// if err != nil {
	// 	return err
	// }

	log.Printf("waiting...")

	select {
	case <-ctx.Done():
		return ctx.Err()
	}
}
