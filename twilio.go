package main

import (
	"fmt"
	twilio "github.com/kevinburke/twilio-go"
	// "log"
)

func Send(s *Sms) error {

	// fmt.Println(s)
	client := twilio.NewClient(s.Sid, s.Token, nil)

	// msg, err := client.Messages.SendMessage("+16474901528", "+18196991182", "send via go", nil)
	// msg, err := client.Messages.SendMessage(s.FromNumber, s.ClientNumber, s.Message, nil)
	_, err := client.Messages.SendMessage(s.FromNumber, s.ClientNumber, s.Message, nil)

	// log.Println(msg)

	return err
	// return nil
}

type Sms struct {
	Sid          string
	Token        string
	FromNumber   string
	ClientName   string
	ClientNumber string
	Message      string
}

func (s *Sms) String() string {
	return fmt.Sprintf("sid: %s , token: %s, FromNumber: %s, Client Name: %s, client number: %s, Message: %s",
		s.Sid, s.Token, s.FromNumber, s.ClientName, s.ClientNumber, s.Message)
}
