package response

import ( "net/http"
 	"encoding/json"
 	"fmt"
 	"strings"
	"github.com/go-playground/validator/v10"
 )


type Response struct {
	Status string `json:"status"`
	Error string `json:"error"`
}

const (
	StatusOk = "OK"
	StatusError = "Error"
)

func WriteJSON(w http.ResponseWriter, status int, data interface{}) error {

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)

	return json.NewEncoder(w).Encode(data)
}

func GenerateError(err error) Response {

	return Response{
		Status: StatusError,
		Error: err.Error(),
	}
}

func ValidateError(errs validator.ValidationErrors) Response {

	var errMsgs []string

	for _, err := range errs {
		switch err.ActualTag() {
		case "required":
			errMsgs = append(errMsgs, fmt.Sprintf("field %s is required", err.Field()))
		case "email":
			errMsgs = append(errMsgs, fmt.Sprintf("field %s is not a valid email", err.Field()))
		case "min":
			errMsgs = append(errMsgs, fmt.Sprintf("field %s must be greater than %s", err.Field(), err.Param()))
		case "max":
			errMsgs = append(errMsgs, fmt.Sprintf("field %s must be less than %s", err.Field(), err.Param()))
		default:
			errMsgs = append(errMsgs, fmt.Sprintf("field %s is not valid", err.Field()))
		}
	}

	return Response{
		Status: StatusError,
		Error: strings.Join(errMsgs, ", "),
	}
}