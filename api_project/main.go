package main

import (
	"flag"
	"fmt"
	"net/http"
	"time"

	"github.com/PaulFWatts/complete_go/api_project/internal/app"
	"github.com/PaulFWatts/complete_go/api_project/internal/routes"
)

func main() {
	var port int
	flag.IntVar(&port, "port", 8080, "Port to listen on")
	flag.Parse()

	myapp, err := app.NewApplication()
	if err != nil {
		panic(err)
	}

	r := routes.SetupRoutes(myapp)
	server := &http.Server{
		Addr:         fmt.Sprintf(":%d", port),
		Handler:      r,
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
