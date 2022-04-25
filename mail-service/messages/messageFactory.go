package messages

type MessageFactory struct {
	DatabaseService string
}

const mailBody = "To: <%s>\n" +
	"From: <%s>\n" +
	"Subject: %s\n" +
	"%s"
