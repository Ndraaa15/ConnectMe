package errx

type Errx struct {
	Code    int
	Message string
	Err     error
}

func New(code int, message string, err error) *Errx {
	return &Errx{
		Code:    code,
		Message: message,
		Err:     err,
	}
}

func (e *Errx) Error() string {
	return e.Message
}
