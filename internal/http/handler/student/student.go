package student

import "net/http"
import "log/slog"
import "github.com/anaghabodhe/Rest-api/internal/types"

func New() http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {

		var student types.Student

		err := json.NewDecoder(r.Body).Decode(&student)
		if errors.Is(err, io.EOF){

		}

		slog.Info("creating a student")


		w.Write([]byte("Hello, World welcome to student-api!"))
	}
}