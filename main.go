package main

import (
	core "MailingReports/MailingReports/Resources/Utilities/Core"
	"fmt"
)

func main() {
	sender := core.New()
	m := core.NewMessage("Test", "Body message.")
	m.To = []string{"to@gmail.com"}
	m.CC = []string{"copy1@gmail.com", "copy2@gmail.com"}
	m.BCC = []string{"bc@gmail.com"}
	m.AttachFile("/path/to/file")
	fmt.Println(sender.Send(m))
}
