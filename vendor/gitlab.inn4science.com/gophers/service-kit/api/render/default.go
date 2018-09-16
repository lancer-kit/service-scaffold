package render

import "net/http"

var (
	// ResultServerError predefined response for `http.StatusInternalServerError`.
	ResultServerError = &R{
		Code:    http.StatusInternalServerError,
		Message: "Request Failed",
	}

	// ResultBadRequest predefined response for `http.StatusBadRequest`.
	ResultBadRequest = &R{
		Code:    http.StatusBadRequest,
		Message: "Bad Request",
	}

	// ResultSuccess predefined response for `http.StatusOK`.
	ResultSuccess = &R{
		Code:    http.StatusOK,
		Message: "Ok",
	}

	// ResultNotFound predefined response for `http.StatusNotFound`.
	ResultNotFound = &R{
		Code:    http.StatusNotFound,
		Message: "Not Found",
	}

	// ResultUnauthorized predefined response for `http.StatusUnauthorized`.
	ResultUnauthorized = &R{
		Code:    http.StatusUnauthorized,
		Message: "Action Unauthorized",
	}

	// ResultForbidden predefined response for `http.StatusForbidden`.
	ResultForbidden = &R{
		Code:    http.StatusForbidden,
		Message: "Forbidden",
	}

	// ResultConflict predefined response for `http.StatusConflict`.
	ResultConflict = &R{
		Code:    http.StatusConflict,
		Message: "Already created",
	}
)
