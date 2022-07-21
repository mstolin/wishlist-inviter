package messages

import "fmt"

const mailBody = "To: <%s>\nFrom: <%s>\nSubject: %s\n%s"

type MessageFactory struct {
	senderMail string
}

func NewMessageFactory(senderMail string) (MessageFactory, error) {
	msgFactory := MessageFactory{}
	if senderMail == "" {
		return msgFactory, fmt.Errorf("senderMail can't be empty")
	} else {
		msgFactory.senderMail = senderMail
	}
	return msgFactory, nil
}
