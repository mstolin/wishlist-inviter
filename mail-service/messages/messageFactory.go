package messages

type MessageFactory struct {
	DatabaseService string
}

const mailBody = "To: <%s>\nFrom: <%s>\nSubject: %s\n%s"
