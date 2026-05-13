package student

import "net/http"
import "log/slog"
import "github.com/tv-anagha/rest-api/internal/types"
import "encoding/json"
import "errors"
import "io"
import "github.com/tv-anagha/rest-api/internal/utils/response"
import "github.com/go-playground/validator/v10"
import "github.com/tv-anagha/rest-api/internal/storage"


func New(storage storage.Storage) http.HandlerFunc {

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

		lastId, err := storage.CreateStudent(student.Name, student.Email, student.Age)
		if err != nil {
			response.WriteJSON(w, http.StatusInternalServerError, response.GenerateError(err))
			return
		}

		slog.Info("student created successfully", slog.Int("userId", lastId), slog.String("name", student.Name), slog.String("email", student.Email), slog.Int("age", student.Age))

		response.WriteJSON(w, http.StatusCreated, map[string]int{"id": lastId})

	}
}