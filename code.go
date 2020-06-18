package errs

// StatCode is the request result to front
type StatCode int

const (
	// Ok means everything is on the plan
	Ok StatCode = iota + 100
	// UnknownError means server do not know the error reason for now
	UnknownError
	// RequestParseError means data receive from client is parse failed
	RequestParseError
)

func (sc StatCode) String() string {
	return [...]string{
		"ok",
		"unknown error",
		"request parse error",
	}[sc]
}
