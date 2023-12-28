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
	router.Use(middleware.CleanPath)
	router.Use(middleware.Heartbeat("/ping"))
	router.Use(middleware.Recoverer)
	router.Use(middleware.RedirectSlashes)

	// Mount Routers
	router.Get("/", routes.ServeAPIDoc)
	router.Mount("/category", routes.CategoryRouter())
	router.Mount("/user", routes.UserRouter())
	router.Mount("/expense", routes.ExpenseRouter())

	// Handle 404
	router.NotFound(routes.ResourceNotFound)

	// Create Server at specified port
	port := 8080
	server := &http.Server{
		Addr:    fmt.Sprintf(":%d", port),
		Handler: router,
	}

	// Start the server
	fmt.Printf("Server is starting on port %d...\n", port)
	err := server.ListenAndServe()

	if err != nil {
		fmt.Println("Failed to listen to server!", err)
	}
}
