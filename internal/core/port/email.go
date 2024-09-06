package port

type EmailItf interface {
	SetSender(sender string)
	SetReciever(to ...string)
	SetSubject(subject string)
	SetBodyHTML(path string, data interface{}) error
	Send() error
}
