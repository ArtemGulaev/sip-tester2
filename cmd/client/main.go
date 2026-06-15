package main

import (
	"fmt"
	"log"
	"sip-client/internal/auth"
	"sip-client/internal/sip"
	"time"
)

const (
	ServerURL = "wss://uc.satel.org:5059/"
	Domain    = "uc.satel.org"

	Login    = "10009"
	Password = "changeme9"
	Phone    = "2782"
)

func buildRegister(login string) string {

	return fmt.Sprintf(
		"REGISTER sip:%s SIP/2.0\r\n"+
			"Via: SIP/2.0/WSS localhost;branch=z9hG4bK123456\r\n"+
			"Max-Forwards: 70\r\n"+
			"To: <sip:%s@%s>\r\n"+
			"From: <sip:%s@%s>;tag=12345\r\n"+
			"Call-ID: reg-123456789\r\n"+
			"CSeq: 1 REGISTER\r\n"+
			"Contact: <sip:%s@localhost;transport=ws>\r\n"+
			"User-Agent: Go SIP Client\r\n"+
			"Expires: 3600\r\n"+
			"Content-Length: 0\r\n"+
			"\r\n",
		Domain,
		login,
		Domain,
		login,
		Domain,
		login,
	)
}
func buildAuthorizedRegister(
	login string,
	realm string,
	nonce string,
	response string,
) string {

	return fmt.Sprintf(
		"REGISTER sip:%s SIP/2.0\r\n"+
			"Via: SIP/2.0/WSS localhost;branch=z9hG4bK123456\r\n"+
			"Max-Forwards: 70\r\n"+
			"To: <sip:%s@%s>\r\n"+
			"From: <sip:%s@%s>;tag=12345\r\n"+
			"Call-ID: reg-123456789\r\n"+
			"CSeq: 2 REGISTER\r\n"+
			"Contact: <sip:%s@localhost;transport=ws>\r\n"+
			"Authorization: Digest username=\"%s\", realm=\"%s\", nonce=\"%s\", uri=\"sip:%s\", response=\"%s\"\r\n"+
			"User-Agent: Go SIP Client\r\n"+
			"Expires: 3600\r\n"+
			"Content-Length: 0\r\n"+
			"\r\n",
		Domain,
		login,
		Domain,
		login,
		Domain,
		login,
		login,
		realm,
		nonce,
		Domain,
		response,
	)
}

func main() {

	client, err := sip.Connect(ServerURL)
	if err != nil {
		log.Fatal(err)
	}

	defer client.Close()

	go client.ReadLoop()

	client.StartPing()

	time.Sleep(2 * time.Second)

	log.Println("connection established")

	register := buildRegister(Login)

	err = client.Send(register)
	if err != nil {
		log.Fatal(err)
	}
	time.Sleep(2 * time.Second)

	response := auth.BuildResponse(
		Login,
		Password,
		"SIP-REGISTRAR",
		"24787bba658411f1a958005056856172",
		"REGISTER",
		"sip:uc.satel.org",
	)

	authorizedRegister := buildAuthorizedRegister(
		Login,
		"SIP-REGISTRAR",
		"24787bba658411f1a958005056856172",
		response,
	)

	err = client.Send(authorizedRegister)
	if err != nil {
		log.Fatal(err)
	}

	log.Println("AUTHORIZED REGISTER sent")

	log.Println("Digest response:", response)
	log.Println("REGISTER sent")

	select {}
}
