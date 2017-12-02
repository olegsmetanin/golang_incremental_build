package main

import (
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"os"
)

// PrimitiveHandler of service
func PrimitiveHandler(w http.ResponseWriter, r *http.Request) {
	// Response
	w.Write([]byte("This is api 21"))
}

func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}

func main() {
	r := mux.NewRouter()
	// Base path for api
	basePath := getEnv("URL_BASE_PATH", "/")
	// Routes consist of a path and a handler function.
	r.HandleFunc(basePath, PrimitiveHandler)

	// Bind to a port and pass our router in
	log.Fatal(http.ListenAndServe(":8080", r))
}
