package render

import (
	"net/http"
)

// R is a default struct for json responses.

//Example:
//
//``` go
//	func MyHandler(w http.ResponseWriter, r *http.Request) {
//		// some code ...
//		// ...
//		res := render.R{
//			Code: http.StatusOk,
//			Message: "User created",
//		}
//		res.Render(w)
//		return
//	}
//```
// Usage of predefined response:
//``` go
//	func MyHandler(w http.ResponseWriter, r *http.Request) {
//		// some code ...
//		// ...
//		render.ResultBadRequest.SetError("Invalid email").Render(w)
//		return
//	}
//```
type R struct {
	Code    int         `json:"errcode"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
	Error   interface{} `json:"errmsg,omitempty"`
}

// SetError adds error details to response.
func (r *R) SetError(val interface{}) *R {
	nr := *r

	switch val.(type) {
	case nil:
		break
	case error:
		nr.Error = val.(error).Error()
	case string:
		nr.Error = val
	case R:
		nr.Error = val.(R).Error
	case *R:
		nr.Error = val.(*R).Error
	default:
		nr.Error = val
	}

	return &nr
}

// SetData sets response data.
func (r *R) SetData(val interface{}) *R {
	nr := *r
	nr.Data = val
	return &nr
}

// Render writes current response as WriteJSON to the `http.ResponseWriter`.
func (r *R) Render(w http.ResponseWriter) {
	WriteJSON(w, r.Code, r)
}
