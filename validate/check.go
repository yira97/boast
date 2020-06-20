package validate

const (
	// MinLen is a bitmask for checker
	MinLen int64 = 1 << iota
	// MaxLen is a bitmask for checker
	MaxLen
)

// Checker is for check struct tidy
type Checker interface {
	Check(v interface{}) error
	SetMode(m int64)
}
