package mail

type Message struct {
	Subject     string
	SenderName  string
	SenderEmail string
	Receiver    []string
	Content     string
}
