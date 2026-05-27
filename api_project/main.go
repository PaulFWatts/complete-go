package main

import (
	"flag"
	"fmt"
	"net/http"
	"time"

	"github.com/PaulFWatts/complete_go/api_project/internal/app"
)

func main() {
	var port int
	flag.IntVar(&port, "port", 8080, "Port to listen on")
	flag.Parse()

	myapp, err := app.NewApplication()
	if err != nil {
		panic(err)
	}

	http.HandleFunc("/health", HealthCheck)
	server := &http.Server{
		Addr:         fmt.Sprintf(":%d", port),
		IdleTimeout:  time.Minute,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 30 * time.Second,
	}

	myapp.Logger.Printf("Application started successfully on port %d", port)

	err = server.ListenAndServe()
	if err != nil {
		myapp.Logger.Fatal(err)
	}
}

func HealthCheck(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Status is available\n")
}
