package validate

import "github.com/yrfg/boast/errs"

var _ Checker = (*StringCheck)(nil)

// StringCheck is a validator for string
type StringCheck struct {
	Mode   int64
	MinLen int
	MaxLen int
}

// SetMode is the impl for Checker
func (sc *StringCheck) SetMode(mode int64) {
	sc.Mode = mode
}

// Check is the impl for Checker
func (sc *StringCheck) Check(v interface{}) error {
	s, ok := v.(string)
	if !ok {
		return errs.ErrInvalidType
	}

	if sc.Mode&MinLen == 1 && len(s) < sc.MinLen {
		return errs.ErrStringTooShort
	}

	if sc.Mode&MaxLen == 1 && len(s) > sc.MaxLen {
		return errs.ErrStringTooLong
	}

	return nil
}
