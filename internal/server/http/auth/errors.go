package auth

type ErrInvalidGrand struct {
	Description string
	Inner       error
}

func (e ErrInvalidGrand) Error() string {
	return e.Description
}

func (e ErrInvalidGrand) Unwrap() error {
	return e.Inner
}
