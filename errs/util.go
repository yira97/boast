package errs

import (
	"encoding/json"
	"net/http"
)

// ErrorDisplay is struct for show detail of an err in package
type ErrorDisplay struct {
	Code StatCode `json:"code"`
	// Message is pair with stat code
	Message string                 `json:"msg"`
	Detail  map[string]interface{} `json:"detail,omitempty"`
}

// Dis is the func to construct ErrorDisplay
//
// type0: Dis(3)
// {
//   "code": 3,
//   "msg": "3 description"
// }
//
// type1: Dis(3, "complex description")
// {
//   "code": 3,
//   "msg": "3 description",
//   "detail": {
//     "info": "complex description"
//   }
// }
//
// type2: Dis(3, "time", "2020/06/18", "result", "some result")
// {
//   "code": 3,
//   "msg": "3 description",
//   "detail": {
//     "time": "2020/06/18",
//     "result": "some result"
//   }
// }

// Dis is a function to create ErrorDisplay
func Dis(code StatCode, details ...string) ErrorDisplay {
	var resp ErrorDisplay
	resp.Code = code
	resp.Message = code.String()
	lenDetail := len(details)
	// type0
	if lenDetail == 0 {
		return resp
	}
	// type1
	resp.Detail = make(map[string]interface{})
	if lenDetail == 1 {
		resp.Detail["info"] = details[0]
		return resp
	}
	// type2
	for i := 0; i+1 < lenDetail; i += 2 {
		resp.Detail[details[i]] = details[i+1]
	}
	return resp
}

// WriteDis is a helper function to write error to http respond directly
func WriteDis(w http.ResponseWriter, httpCode int, serverCode StatCode, details ...string) {
	dis := Dis(serverCode, details...)
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(httpCode)
	if j, err := json.Marshal(dis); err == nil {
		w.Write(j)
	}
}
