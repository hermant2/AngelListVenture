package apperror

import "net/http"

type Code int

const (
	General     Code = 18930
	InputZero   Code = 18931
	NoInvestors Code = 18932
)

type Standard struct {
	Status      int    `json:"status"`
	Code        Code   `json:"code"`
	Description string `json:"description"`
}

func (err Standard) Error() string {
	return err.Description
}

func Unprocessable(code Code) Standard {
	return Standard{Status: http.StatusUnprocessableEntity, Code: code, Description: "Unprocessable Entity"}
}

func BadRequest(code Code) Standard {
	return Standard{Status: http.StatusBadRequest, Code: code, Description: "Bad Request"}
}

func NotFound(code Code) Standard {
	return Standard{Status: http.StatusNotFound, Code: code, Description: "Not Found"}
}

func InternalServerError(code Code) Standard {
	return Standard{Status: http.StatusInternalServerError, Code: code, Description: "Internal Server Error"}
}
