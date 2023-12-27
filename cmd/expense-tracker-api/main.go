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

	// TODO: Middleware
	router.Use(middleware.Logger)
	// r.Use(middleware.RequestID)
	// r.Use(middleware.Recoverer)
	// r.Use(middleware.URLFormat)
	// r.Use(render.SetContentType(render.ContentTypeJSON))

	// Mount Routers
	router.Get("/", routes.ServeAPIDoc)
	router.Mount("/category", routes.CategoryRouter())
	router.Mount("/user", routes.UserRouter())
	router.Mount("/expense", routes.ExpenseRouter())

	// Handle Errors
	router.NotFound(routes.ResourceNotFound)
	// TODO: router.MethodNotAllowed(routes.X)

	// Create Server at specified port
	port := 8080
	server := &http.Server{
		Addr:    fmt.Sprintf(":%d", port),
		Handler: router,
	}

	fmt.Printf("Server is starting on port %d...\n", port)
	err := server.ListenAndServe()

	if err != nil {
		fmt.Println("Failed to listen to server!", err)
	}
}
