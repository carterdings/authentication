package errs

import "fmt"

const (
	ErrTypeFramework = 1
	ErrTypeBusiness  = 2
)

const (
	ErrCodeSuccess = 0
	ErrCodeUnknown = 9999

	ErrMsgSuccess = "success"
)

var (
	ErrTypeMap = map[int]string{
		ErrTypeFramework: "framework",
		ErrTypeBusiness:  "business",
	}
)

// Err error wrapper
type Err struct {
	Type int
	Code int
	Msg  string
}

// Error error
func (e *Err) Error() string {
	return fmt.Sprintf("type: %s, code: %d, msg: %s", ErrTypeMap[e.Type], e.Code, e.Msg)
}

// NewFrameworkErr new framework error
func NewFrameworkErr(code int, msg string) error {
	return &Err{ErrTypeFramework, code, msg}
}

// New new error
func New(code int, msg string) error {
	return &Err{ErrTypeBusiness, code, msg}
}

// Newf new formatted error
func Newf(code int, format string, args ...interface{}) error {
	return &Err{ErrTypeBusiness, code, fmt.Sprintf(format, args...)}
}

// ErrCode get error code
func ErrCode(err error) int {
	if err == nil {
		return ErrCodeSuccess
	}
	e, ok := err.(*Err)
	if !ok {
		return ErrCodeUnknown
	}
	if e == nil {
		return ErrCodeSuccess
	}
	return e.Code
}

// ErrMsg get error msg
func ErrMsg(err error) string {
	if err == nil {
		return ErrMsgSuccess
	}
	e, ok := err.(*Err)
	if !ok {
		return err.Error()
	}
	if e == nil {
		return ErrMsgSuccess
	}
	return e.Msg
}
