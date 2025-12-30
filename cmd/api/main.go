package main

import (
	"financialcontrol/internal/v1/api"
	"fmt"
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"
)

func main() {
	api := api.NewApi(chi.NewMux())

	api.RegisterRoutes()

	port := os.Getenv("GOBID_APP_PORT")
	if port == "" {
		port = "3080"
	}

	fmt.Printf("Starting Server on port :%s\n", port)
	if err := http.ListenAndServe(":"+port, api.Router); err != nil {
		panic(err)
	}
}
