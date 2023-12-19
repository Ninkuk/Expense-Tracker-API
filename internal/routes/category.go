package routes

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/go-chi/chi/v5"
	"github.com/ninkuk/expense-tracker-api/internal/utils"
	"github.com/ninkuk/expense-tracker-api/models"
)

func CategoryRouter() chi.Router {
	router := chi.NewRouter()

	router.Get("/", listCategories)
	router.Get("/{id}", displayCategory)
	router.Get("/{id}/", displayCategory)
	router.Get("/search", searchCategory)

	return router
}

func listCategories(w http.ResponseWriter, r *http.Request) {
	categories := utils.GetCategories()

	bytes, err := json.Marshal(categories)

	if err != nil {
		fmt.Println(err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(bytes)
}

func displayCategory(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	categories := utils.GetCategories().Category

	for i := 0; i < len(categories); i++ {
		if strings.EqualFold(categories[i].ID, id) {
			bytes, err := json.Marshal(categories[i])

			if err != nil {
				fmt.Println(err)
			}

			w.Header().Set("Content-Type", "application/json")
			w.Write(bytes)

			return
		}
	}

	ResourceNotFound(w, r)
}

func searchCategory(w http.ResponseWriter, r *http.Request) {
	queryParams := r.URL.Query()
	searchQuery := strings.ToLower(queryParams.Get("q"))

	if searchQuery == "" {
		ResourceNotFound(w, r)
		return
	}

	categories := utils.GetCategories().Category

	var matchCategories models.Categories

	for i := 0; i < len(categories); i++ {
		if strings.Contains(categories[i].ID, searchQuery) {
			matchCategories.Category = append(matchCategories.Category, categories[i])
		} else if strings.Contains(categories[i].Name, searchQuery) {
			matchCategories.Category = append(matchCategories.Category, categories[i])
		} else if strings.Contains(categories[i].Description, searchQuery) {
			matchCategories.Category = append(matchCategories.Category, categories[i])
		}
	}

	if len(matchCategories.Category) == 0 {
		ResourceNotFound(w, r)
		return
	}

	bytes, err := json.Marshal(matchCategories)

	if err != nil {
		fmt.Println(err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(bytes)
}
