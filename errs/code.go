package errs

// StatCode is the request result to front
type StatCode int

const (
	// Ok means everything is on the plan
	Ok StatCode = iota
	// UnknownError means server do not know the error reason for now
	UnknownError
	// RequestError means request is unexpected
	RequestError
	// RequestParseError means data receive from client is parse failed
	RequestParseError

)

func (sc StatCode) String() string {
	return [...]string{
		"Ok",
		"UnknownError",
		"RequestError",
		"RequestParseError",
	}[sc]
}
