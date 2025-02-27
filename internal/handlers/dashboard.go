// internal/handlers/dashboard.go
package handlers

import (
	"log"
	"net/http"

	"github.com/saladinomario/vr-training-admin/internal/models"
	"github.com/saladinomario/vr-training-admin/templates/pages"
)

func DashboardHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}

	component := pages.Dashboard()

	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	if err := component.Render(r.Context(), w); err != nil {
		log.Printf("Error rendering dashboard: %v", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}
}

func DashboardContentHandler(w http.ResponseWriter, r *http.Request) {
	// Get recent sessions
	recentSessions := SessionStore.GetRecent(5)

	// Get the system status
	status := models.SystemStatusStoreInstance.Get()

	//Render component
	component := pages.DashboardContent(recentSessions, status)

	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	if err := component.Render(r.Context(), w); err != nil {
		log.Printf("Error rendering dashboard content: %v", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}
}
