package mail

import "fmt"

func GenerateMessage(from, to, subject, content string) string {
	message := `To: <%s>
From: <%s>
Subject: %s

%s
`
	return fmt.Sprintf(message, from, to, subject, content)
}
