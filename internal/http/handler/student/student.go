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
import "strconv"

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


func GetById(storage storage.Storage) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {

		idStr := r.PathValue("id")
		slog.Info("getting a student by id", slog.String("id", idStr))

		studentId, err := strconv.Atoi(idStr)
		if err != nil {
			response.WriteJSON(w, http.StatusBadRequest, response.GenerateError(err))
			return
		}

		student, err := storage.GetStudentById(studentId)

		if err != nil {
			response.WriteJSON(w, http.StatusInternalServerError, response.GenerateError(err))
			return
		}
		response.WriteJSON(w, http.StatusOK, student)
	}
}

func GetList(storage storage.Storage) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {
		slog.Info("getting a list of students")

		students, err := storage.GetStudents()
		if err != nil {
			response.WriteJSON(w, http.StatusInternalServerError, response.GenerateError(err))
			return
		}
		response.WriteJSON(w, http.StatusOK, students)
	
	}
}

func UpdateStudent(storage storage.Storage) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {
		slog.Info("updating a student by Id")

		idStr := r.PathValue("id")
		slog.Info("updating a student by id", slog.String("id", idStr))

		studentId, err := strconv.Atoi(idStr)
		if err != nil {
			response.WriteJSON(w, http.StatusBadRequest, response.GenerateError(err))
			return
		}

		var student types.Student

		err = json.NewDecoder(r.Body).Decode(&student)
		if errors.Is(err, io.EOF) {
			response.WriteJSON(w, http.StatusBadRequest, response.GenerateError(err))
			return
		}

		if err != nil {
			response.WriteJSON(w, http.StatusBadRequest, response.GenerateError(err))
			return
		}

		if err := validator.New().Struct(student); err != nil {
			validationErrors := err.(validator.ValidationErrors)
			response.WriteJSON(w, http.StatusBadRequest, response.ValidateError(validationErrors))
			return
		}

		affectedRows, err := storage.UpdateStudent(studentId, student.Name, student.Email, student.Age)
		if err != nil {
			response.WriteJSON(w, http.StatusInternalServerError, response.GenerateError(err))
			return
		}
		response.WriteJSON(w, http.StatusOK, map[string]int{"affectedRows": affectedRows})
	}
}