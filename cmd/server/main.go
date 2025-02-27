package main

//Version 1.0

import (
	"log"
	"net/http"
	"os"

	"github.com/saladinomario/vr-training-admin/internal/handlers"
	"github.com/saladinomario/vr-training-admin/internal/models"
)

// Track registered routes
var registeredRoutes []string

func setupRoutes() *http.ServeMux {
	log.Println("Setting up all application routes...")

	mux := http.NewServeMux()

	// Register dashboard handler
	log.Println("Registering dashboard route")
	mux.HandleFunc("/", handlers.DashboardHandler)
	mux.HandleFunc("/dashboard-content", handlers.DashboardContentHandler)

	// Register scenario routes
	log.Println("Setting up scenario routes")
	handlers.SetupScenarioRoutes(mux)

	// Register avatar routes
	log.Println("Setting up avatar routes")
	handlers.SetupAvatarRoutes(mux)

	// Register observer routes
	log.Println("Setting up observer routes")
	handlers.SetupObserverRoutes(mux)

	// Register settings routes
	log.Println("Setting up settings routes")
	handlers.SetupSettingsRoutes(mux)

	// Register session routes (new)
	log.Println("Setting up session routes")
	handlers.SetupSessionRoutes(mux)

	// Serve static files
	log.Println("Setting up static file server")
	fs := http.FileServer(http.Dir("static"))
	mux.Handle("/static/", http.StripPrefix("/static/", fs))

	log.Println("All routes registered successfully")
	return mux
}

// Print dynamically registered routes
func printRegisteredRoutes() {
	log.Println("=== Registered Routes ===")
	for _, route := range registeredRoutes {
		log.Println(route)
	}
	log.Println("==========================")
}

func main() {
	// Initialize System Status store
	unrealEngineUrl := os.Getenv("UNREAL_ENGINE_URL")
	if unrealEngineUrl == "" {
		unrealEngineUrl = "http://localhost:8081"
	}
	models.SystemStatusStoreInstance = models.NewSystemStatusStore(unrealEngineUrl)

	mux := setupRoutes()
	printRegisteredRoutes()

	// Get port from environment variable
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080" // Default to 8080 if PORT not set
	}

	log.Printf("Server starting on port %s", port)
	log.Printf("Visit http://localhost:%s to view the application", port)

	if err := http.ListenAndServe(":"+port, mux); err != nil {
		log.Fatal(err)
	}
}
