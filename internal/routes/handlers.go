package routes

import (
	"fmt"
	"net/http"
)

func ShowDoc(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Hello World!")
	// TODO: serve docs as a static html?
}

func ResourceNotFound(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotFound)
	w.Write([]byte("The resource you're looking for doesn't exist."))
}
