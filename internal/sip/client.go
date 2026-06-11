package sip

import (
	"crypto/tls"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/websocket"
)

type Client struct {
	conn *websocket.Conn
}

func Connect(url string) (*Client, error) {

	dialer := websocket.Dialer{
		TLSClientConfig: &tls.Config{
			MinVersion: tls.VersionTLS12,
		},
		HandshakeTimeout: 10 * time.Second,

		Subprotocols: []string{"sip"},
	}

	headers := http.Header{}
	headers.Set("Origin", "https://localhost")

	conn, resp, err := dialer.Dial(url, headers)

	if err != nil {
		if resp != nil {
			log.Printf("HTTP status: %s", resp.Status)
		}
		return nil, err
	}

	log.Printf("Connected")

	if resp != nil {

		log.Printf("Status: %s", resp.Status)

		for k, v := range resp.Header {
			log.Printf("%s: %v", k, v)
		}
	}

	log.Printf(
		"Negotiated subprotocol: %s",
		conn.Subprotocol(),
	)
	return &Client{
		conn: conn,
	}, nil
}

func (c *Client) Close() error {
	return c.conn.Close()
}
func (c *Client) ReadLoop() {

	for {

		msgType, data, err := c.conn.ReadMessage()

		if err != nil {
			log.Printf("read error: %v", err)
			return
		}

		log.Printf(
			"RECV type=%d\n%s\n",
			msgType,
			string(data),
		)
	}
}
func (c *Client) Send(text string) error {

	log.Printf(
		"SEND\n%s\n",
		text,
	)

	return c.conn.WriteMessage(
		websocket.TextMessage,
		[]byte(text),
	)
}

func (c *Client) StartPing() {

	ticker := time.NewTicker(20 * time.Second)

	go func() {

		defer ticker.Stop()

		for range ticker.C {

			err := c.conn.WriteControl(
				websocket.PingMessage,
				[]byte("ping"),
				time.Now().Add(5*time.Second),
			)

			if err != nil {
				log.Printf(
					"ping error: %v",
					err,
				)
				return
			}

			log.Printf("PING")
		}
	}()
}
