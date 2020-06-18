package errs

import "fmt"

// Btw is a helper func to create decendent error
func Btw(err error, extra string) error {
	return fmt.Errorf("%s/%w", extra, err)
}

var (
	// ErrInputParams means the input format is invalid
	ErrInputParams error = fmt.Errorf("ErrInputParams")
	// ErrInputEmptyStringParams means the string input is empty
	ErrInputEmptyStringParams error = Btw(ErrInputParams, "ErrInputNullStringParams")
	// ErrInputTooLongStringParams means the string input is too long
	ErrInputTooLongStringParams error = Btw(ErrInputParams, "ErrInputTooLongStringParams")
)
