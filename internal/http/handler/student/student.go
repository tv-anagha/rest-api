package student

import "net/http"
import "log/slog"
import "github.com/tv-anagha/rest-api/internal/types"
import "encoding/json"
import "errors"
import "io"
import "github.com/tv-anagha/rest-api/internal/utils/response"
import "github.com/go-playground/validator/v10"


func New() http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {

		slog.Info("creating a student")

		var student types.Student

		err := json.NewDecoder(r.Body).Decode(&student)
		if errors.Is(err, io.EOF){
			response.WriteJSON(w, http.StatusBadRequest, response.GenerateError(err))
			return
		}

		if err != nil {
			response.WriteJSON(w, http.StatusBadRequest, response.GenerateError(err))
			return
		}

		//request body validation

		if err := validator.New().Struct(student); err != nil {

			//type cast
			validationErrors := err.(validator.ValidationErrors)

			response.WriteJSON(w, http.StatusBadRequest, response.ValidateError(validationErrors))
			return	
		}

		response.WriteJSON(w, http.StatusCreated, map[string]string{"success": "ok"})

	}
}