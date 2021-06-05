package countries

import (
	core "MailingReports/MailingReports/Resources/Utilities/Core"
	"fmt"
)

func Netherlands() {
	sender := core.New()
	m := core.NewMessage("Test", "Body message.")
	m.To = []string{"to@gmail.com"}
	m.CC = []string{"copy1@gmail.com", "copy2@gmail.com"}
	m.BCC = []string{"bc@gmail.com"}
	m.AttachFile("./Reports/Austria/Austria.xlsx")
	fmt.Println(sender.Send(m))
}
