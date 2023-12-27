package routes

import (
	"net/http"

	"github.com/ninkuk/expense-tracker-api/internal/utils"
)

// Serves the API documentation as HTML (default) or txt file
func ServeAPIDoc(w http.ResponseWriter, r *http.Request) {
	if r.Header.Get("Accept") == "text/plain" {
		w.Header().Set("Content-Type", "text/plain")
		w.Write(utils.GetDocTXT())
	} else {
		w.Header().Set("Content-Type", "text/html")
		w.Write(utils.GetDocHTML())
	}
}

func ResourceNotFound(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotFound)
	w.Write([]byte("The resource you're looking for doesn't exist."))
}
