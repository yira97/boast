package validate

import (
	"testing"

	"github.com/yrfg/boast/errs"
)

func TestStringCheck(t *testing.T) {
	for i, cas := range []struct {
		s      string
		expect error
		sc     StringCheck
	}{
		{
			s:      "hhhhh",
			expect: nil,
			sc: StringCheck{
				MinLen: 5,
				Mode:   MinLen,
			},
		},
		{
			s:      "hhhh",
			expect: errs.ErrStringTooShort,
			sc: StringCheck{
				MinLen: 5,
				Mode:   MinLen,
			},
		},
	} {
		if got := cas.sc.Check(cas.s); got != cas.expect {
			t.Error("string check failed at index", i)
		}
	}

}
