package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/anaghabodhe/Rest-api/internal/config"
)

func main() {
	cfg := config.MustLoad()


	//setup router
	router := http.NewServeMux()
	router.HandleFunc("GET /", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello, World welcome to student-api!"))
	})

	//setup server
	server := http.Server{
		Addr:    cfg.Address,
		Handler: router,
	}

	fmt.Println("Server is running on port")


	if err := server.ListenAndServe(); err != nil {
		log.Fatalf("failed to start server: %s", err.Error())
	}

}
