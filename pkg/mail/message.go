package mail

type Message struct {
	Subject     string
	SenderEmail string
	SenderName  string
	Receiver    []string
	Content     string
}
