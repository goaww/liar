package mail

type Sender interface {
	Send(msg Message) error
}
