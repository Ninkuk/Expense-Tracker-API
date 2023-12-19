package main

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"

	"github.com/ninkuk/expense-tracker-api/internal/routes"
)

func main() {
	router := chi.NewRouter()

	// Middleware
	router.Use(middleware.Logger)
	// r.Use(middleware.RequestID)
	// r.Use(middleware.Recoverer)
	// r.Use(middleware.URLFormat)
	// r.Use(render.SetContentType(render.ContentTypeJSON))

	// Mount Routers
	router.Mount("/category", routes.CategoryRouters())
	router.Mount("/user", routes.UserRouters())
	router.Mount("/expense", routes.ExpenseRouters())

	// Create Server at specified port
	server := &http.Server{
		Addr:    ":3000",
		Handler: router,
	}

	err := server.ListenAndServe()

	if err != nil {
		fmt.Println("Failed to listen to server!", err)
	}
}
