package errs

import "testing"

func TestExampleErrorDisplay(t *testing.T) {
	type testCase struct {
		Expect ErrorDisplay
		Params []string
		Code   StatCode
	}

	testCases := []testCase{
		{
			Expect: ErrorDisplay{
				Code:    Ok,
				Message: Ok.String(),
				Detail:  nil,
			},
			Params: nil,
			Code:   Ok,
		},
		{
			Expect: ErrorDisplay{
				Code:    UnknownError,
				Message: UnknownError.String(),
				Detail: map[string]interface{}{
					"info": "hhhh",
				},
			},
			Params: []string{"hhhh"},
			Code:   UnknownError,
		},
	}

	for i, cas := range testCases {
		v := Dis(cas.Code, cas.Params...)
		if v.Code != cas.Expect.Code {
			t.Error("code error in display code at index ", i)
		}
		if v.Message != cas.Expect.Message {
			t.Error("msg error in display code at index ", i)
		}
		if v.Detail != nil && cas.Expect.Detail == nil ||
			v.Detail == nil && cas.Expect.Detail != nil {
			t.Error("detail map error in display code at index", i)
		}
		if v.Detail != nil && cas.Expect.Detail != nil {
			for k1, v1 := range v.Detail {
				if v2, ok := cas.Expect.Detail[k1]; ok {
					if v2 != v1 {
						t.Error("value not equal. \ndetail map error in display code at index", i)
					}
				} else {
					t.Error("key not exist. \ndetail map error in display code at index", i)
				}
			}
		}
	}
}
