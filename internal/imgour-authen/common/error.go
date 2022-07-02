package common

const (
	ErrNotFoundCode = iota
	ErrInternalCode
)

const (
	ErrNotFoundMessage = "Not found"
	ErrInternalMessage = "Internal error"
)

type Error struct {
	Code    int
	Message string
	Extra   error
}

func (m *Error) Error() string {
	return m.Message
}

func ErrNotFound() *Error {
	return &Error{
		Code:    ErrNotFoundCode,
		Message: ErrNotFoundMessage,
		Extra:   nil,
	}
}

func ErrInternal(err error) *Error {
	return &Error{
		Code:    ErrInternalCode,
		Message: ErrInternalMessage,
		Extra:   err,
	}
}
