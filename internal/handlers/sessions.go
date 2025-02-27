// internal/handlers/sessions.go
package handlers

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/gorilla/websocket"
	"github.com/saladinomario/vr-training-admin/internal/models"
	"github.com/saladinomario/vr-training-admin/templates/components/sessions"
	"github.com/saladinomario/vr-training-admin/templates/pages"
)

// WebSocket upgrader with configuration
var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		// Allow all connections for now - can be restricted in production
		return true
	},
}

var SessionStore *models.SessionStore

func init() {
	// Create data directory if it doesn't exist
	dataDir := "./data"
	if err := os.MkdirAll(dataDir, 0755); err != nil {
		log.Printf("Error creating data directory: %v", err)
	}

	// Initialize session store
	sessionFilePath := dataDir + "/sessions.json"
	SessionStore = models.NewSessionStore(sessionFilePath)
}

// StartSessionHandler handles the session creation form submission
func StartSessionHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	if err := r.ParseForm(); err != nil {
		http.Error(w, "Failed to parse form", http.StatusBadRequest)
		return
	}

	// Extract form values
	scenarioID := r.FormValue("scenario_id")
	avatarID := r.FormValue("avatar_id")
	observerID := r.FormValue("observer_id")

	// Validate required fields
	if scenarioID == "" || avatarID == "" || observerID == "" {
		http.Error(w, "Missing required fields", http.StatusBadRequest)
		return
	}

	// Create session
	session, err := SessionStore.Create(scenarioID, avatarID, observerID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Start the session in Unreal Engine
	go startUnrealEngineSession(session.ID)

	// Return success response
	if r.Header.Get("HX-Request") == "true" {
		// Get updated dashboard content
		recentSessions := SessionStore.GetRecent(5)
		// Get the system status
		status := models.SystemStatusStoreInstance.Get()
		component := pages.DashboardContent(recentSessions, status)

		w.Header().Set("Content-Type", "text/html; charset=utf-8")
		if err := component.Render(r.Context(), w); err != nil {
			log.Printf("Error rendering dashboard content: %v", err)
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		}
		return
	}

	// Redirect to dashboard for non-HTMX requests
	http.Redirect(w, r, "/", http.StatusSeeOther)
}

// SessionStatusHandler handles updating session status
func SessionStatusHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	// Extract session ID from URL
	parts := strings.Split(r.URL.Path, "/")
	if len(parts) < 3 {
		http.Error(w, "Invalid URL", http.StatusBadRequest)
		return
	}
	sessionID := parts[2]

	if err := r.ParseForm(); err != nil {
		http.Error(w, "Failed to parse form", http.StatusBadRequest)
		return
	}

	// Extract status from form
	status := r.FormValue("status")
	if status == "" {
		http.Error(w, "Status is required", http.StatusBadRequest)
		return
	}

	// Only allow valid status values
	validStatus := map[string]bool{
		sessions.StatusRunning:   true,
		sessions.StatusPaused:    true,
		sessions.StatusCompleted: true,
	}

	if !validStatus[status] {
		http.Error(w, "Invalid status value", http.StatusBadRequest)
		return
	}

	// Update session status
	err := SessionStore.Update(sessionID, status)
	if err != nil {
		if err == models.ErrSessionNotFound {
			http.NotFound(w, r)
		} else {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		return
	}

	// Update session in Unreal Engine
	go updateUnrealEngineSession(sessionID, status)

	// Return success response
	if r.Header.Get("HX-Request") == "true" {
		// Get updated dashboard content
		recentSessions := SessionStore.GetRecent(5)
		component := pages.RecentActivity(recentSessions)

		w.Header().Set("Content-Type", "text/html; charset=utf-8")
		w.Header().Set("HX-Trigger", "closeModal")
		if err := component.Render(r.Context(), w); err != nil {
			log.Printf("Error rendering recent activity: %v", err)
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		}
		return
	}

	// Redirect to dashboard for non-HTMX requests
	http.Redirect(w, r, "/", http.StatusSeeOther)
}

// SessionFormHandler handles serving the new session form
func SessionFormHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	// Get all scenarios, avatars, and observers for form dropdowns
	allScenarios := ScenarioStore.GetAll()
	allAvatars := AvatarStore.GetAll()
	allObservers := ObserverStore.GetAll()

	// Render form page
	component := pages.SessionNew(allScenarios, allAvatars, allObservers)

	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	if err := component.Render(r.Context(), w); err != nil {
		log.Printf("Error rendering session form: %v", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}
}

// WebSocket message types
const (
	// Command types
	CmdCreateSession     = "create_session"
	CmdUpdateSession     = "update_session"
	CmdGetActiveSessions = "get_active_sessions"
	CmdGetSessionJSON    = "get_session_json"

	// Response types
	ResponseSuccess = "success"
	ResponseError   = "error"
)

// WebSocket message structure
type WebSocketMessage struct {
	Command      string          `json:"command"`
	Data         json.RawMessage `json:"data"`
	ResponseType string          `json:"response_type,omitempty"`
	Message      string          `json:"message,omitempty"`
}

// SessionCreateRequest contains data needed to create a new session
type SessionCreateRequest struct {
	ScenarioID string `json:"scenarioId"`
	AvatarID   string `json:"avatarId"`
	ObserverID string `json:"observerId"`
}

// SessionUpdateRequest contains data needed to update a session
type SessionUpdateRequest struct {
	SessionID string `json:"sessionId"`
	Status    string `json:"status"`
	Notes     string `json:"notes,omitempty"`
	Score     *int   `json:"score,omitempty"`
}

// WebSocketSessionHandler handles WebSocket connections for session operations
func WebSocketSessionHandler(w http.ResponseWriter, r *http.Request) {
	// Upgrade connection to WebSocket
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Printf("Error upgrading to WebSocket: %v", err)
		return
	}
	defer func() {
		// Update system status when connection closes
		models.SystemStatusStoreInstance.UpdateWebSocketConnection(false)
		conn.Close()
		log.Println("WebSocket connection closed")
	}()

	// Update system status when connection is established
	models.SystemStatusStoreInstance.UpdateWebSocketConnection(true)
	log.Println("WebSocket connection established")

	// Process messages
	for {
		// Read message
		_, message, err := conn.ReadMessage()
		if err != nil {
			log.Printf("Error reading WebSocket message: %v", err)
			break
		}

		// Parse message
		var msg WebSocketMessage
		if err := json.Unmarshal(message, &msg); err != nil {
			log.Printf("Error parsing WebSocket message: %v", err)
			sendErrorResponse(conn, "Invalid message format")
			continue
		}

		// Process commands
		switch msg.Command {
		case CmdCreateSession:
			handleCreateSession(conn, msg.Data)
		case CmdUpdateSession:
			handleUpdateSession(conn, msg.Data)
		case CmdGetActiveSessions:
			handleGetActiveSessions(conn)
		case CmdGetSessionJSON:
			handleGetSessionJSON(conn, msg.Data)
		default:
			sendErrorResponse(conn, "Unknown command")
		}
	}
}

// handleCreateSession handles session creation requests via WebSocket
func handleCreateSession(conn *websocket.Conn, data json.RawMessage) {
	// Parse request
	var req SessionCreateRequest
	if err := json.Unmarshal(data, &req); err != nil {
		sendErrorResponse(conn, "Invalid create session request")
		return
	}

	// Validate required fields
	if req.ScenarioID == "" || req.AvatarID == "" || req.ObserverID == "" {
		sendErrorResponse(conn, "Missing required fields")
		return
	}

	// Create session
	session, err := SessionStore.Create(req.ScenarioID, req.AvatarID, req.ObserverID)
	if err != nil {
		sendErrorResponse(conn, "Failed to create session: "+err.Error())
		return
	}

	// Start the session in Unreal Engine
	go startUnrealEngineSession(session.ID)

	// Send success response
	sendResponse(conn, ResponseSuccess, "Session created successfully", session)
}

// handleUpdateSession handles session update requests via WebSocket
func handleUpdateSession(conn *websocket.Conn, data json.RawMessage) {
	// Parse request
	var req SessionUpdateRequest
	if err := json.Unmarshal(data, &req); err != nil {
		sendErrorResponse(conn, "Invalid update session request")
		return
	}

	// Validate required fields
	if req.SessionID == "" || req.Status == "" {
		sendErrorResponse(conn, "Session ID and status are required")
		return
	}

	// Only allow valid status values
	validStatus := map[string]bool{
		sessions.StatusRunning:   true,
		sessions.StatusPaused:    true,
		sessions.StatusCompleted: true,
		sessions.StatusFailed:    true,
	}

	if !validStatus[req.Status] {
		sendErrorResponse(conn, "Invalid status value")
		return
	}

	// Update session status
	err := SessionStore.Update(req.SessionID, req.Status)
	if err != nil {
		if err == models.ErrSessionNotFound {
			sendErrorResponse(conn, "Session not found")
		} else {
			sendErrorResponse(conn, "Failed to update session: "+err.Error())
		}
		return
	}

	// Update session in Unreal Engine
	go updateUnrealEngineSession(req.SessionID, req.Status)

	// Get updated session
	session, err := SessionStore.GetByID(req.SessionID)
	if err != nil {
		sendErrorResponse(conn, "Session updated but could not retrieve: "+err.Error())
		return
	}

	// Send success response
	sendResponse(conn, ResponseSuccess, "Session updated successfully", session)
}

// handleGetActiveSessions handles requests for active sessions via WebSocket
func handleGetActiveSessions(conn *websocket.Conn) {
	// Get all sessions
	allSessions := SessionStore.GetAll()

	// Filter for active sessions
	activeSessions := make([]*sessions.Session, 0)
	for _, session := range allSessions {
		if session.Status == sessions.StatusRunning || session.Status == sessions.StatusPaused {
			activeSessions = append(activeSessions, session)
		}
	}

	// Send response
	sendResponse(conn, ResponseSuccess, "Active sessions retrieved", activeSessions)
}

// handleGetSessionJSON handles requests for full session JSON via WebSocket
func handleGetSessionJSON(conn *websocket.Conn, data json.RawMessage) {
	// Parse request
	var req struct {
		SessionID string `json:"sessionId"`
	}

	if err := json.Unmarshal(data, &req); err != nil || req.SessionID == "" {
		// If no specific session requested, return all sessions
		allSessions := SessionStore.GetAll()
		sendResponse(conn, ResponseSuccess, "All sessions retrieved", allSessions)
		return
	}

	// Get specific session details
	details, err := SessionStore.GetSessionDetails(req.SessionID, ScenarioStore, AvatarStore, ObserverStore)
	if err != nil {
		if err == models.ErrSessionNotFound {
			sendErrorResponse(conn, "Session not found")
		} else {
			sendErrorResponse(conn, "Failed to get session: "+err.Error())
		}
		return
	}

	// Send response
	sendResponse(conn, ResponseSuccess, "Session details retrieved", details)
}

// sendResponse sends a response via WebSocket
func sendResponse(conn *websocket.Conn, responseType, message string, data interface{}) {
	response := WebSocketMessage{
		ResponseType: responseType,
		Message:      message,
	}

	// Convert data to json.RawMessage
	if data != nil {
		jsonData, err := json.Marshal(data)
		if err != nil {
			log.Printf("Error marshaling response data: %v", err)
			sendErrorResponse(conn, "Internal server error")
			return
		}
		response.Data = jsonData
	}

	// Send response
	if err := conn.WriteJSON(response); err != nil {
		log.Printf("Error sending WebSocket response: %v", err)
	}
}

// sendErrorResponse sends an error response via WebSocket
func sendErrorResponse(conn *websocket.Conn, message string) {
	sendResponse(conn, ResponseError, message, nil)
}

// SetupSessionRoutes registers all session-related routes
func SetupSessionRoutes(mux *http.ServeMux) {
	log.Println("Setting up session routes...")

	// Session form
	mux.HandleFunc("/sessions/new", SessionFormHandler)

	// Start session
	mux.HandleFunc("/sessions/start", StartSessionHandler)

	// Update session status
	mux.HandleFunc("/sessions/", func(w http.ResponseWriter, r *http.Request) {
		if strings.HasPrefix(r.URL.Path, "/sessions/") && r.Method == http.MethodPost {
			SessionStatusHandler(w, r)
			return
		}
		http.NotFound(w, r)
	})

	// WebSocket endpoint for session operations
	mux.HandleFunc("/ws/sessions", WebSocketSessionHandler)

	log.Println("Session routes registered successfully")
}

// Unreal Engine integration functions

// startUnrealEngineSession sends a request to start a session in Unreal Engine
func startUnrealEngineSession(sessionID string) {
	// Update status to "running"
	err := SessionStore.Update(sessionID, sessions.StatusRunning)
	if err != nil {
		log.Printf("Error updating session status: %v", err)
		return
	}

	// Create payload for Unreal Engine
	payload, err := SessionStore.CreateURESessionPayload(sessionID, ScenarioStore, AvatarStore, ObserverStore)
	if err != nil {
		log.Printf("Error creating UE payload: %v", err)
		return
	}

	// Send to Unreal Engine
	sendToUnrealEngine(payload)
}

// updateUnrealEngineSession sends a request to update a session in Unreal Engine
func updateUnrealEngineSession(sessionID, status string) {
	// Create payload for Unreal Engine
	payload, err := SessionStore.CreateURESessionPayload(sessionID, ScenarioStore, AvatarStore, ObserverStore)
	if err != nil {
		log.Printf("Error creating UE payload: %v", err)
		return
	}

	// Send to Unreal Engine
	sendToUnrealEngine(payload)
}

// sendToUnrealEngine sends data to the Unreal Engine endpoint
func sendToUnrealEngine(payload []byte) {
	// TODO: Get this from settings
	unrealEndpoint := "http://localhost:8081/api/vr-session"

	// Update system status to show connection attempt
	models.SystemStatusStoreInstance.UpdateUnrealEngineConnection(true)

	// For now, just log the payload
	log.Printf("Would send to Unreal Engine: %s", unrealEndpoint)

	// Create a nicely formatted version of the payload for logging
	var prettyJSON bytes.Buffer
	if err := json.Indent(&prettyJSON, payload, "", "  "); err == nil {
		log.Printf("Payload: %s", prettyJSON.String())
	}

	// Simulate Unreal Engine ready state for development
	models.SystemStatusStoreInstance.UpdateUnrealEngineReady(true)

	// In production, uncomment this code to actually send the request
	/*
		client := &http.Client{
			Timeout: 10 * time.Second,
		}

		req, err := http.NewRequest("POST", unrealEndpoint, bytes.NewBuffer(payload))
		if err != nil {
			log.Printf("Error creating request: %v", err)
			// Update system status to show connection failure
			models.SystemStatusStoreInstance.UpdateUnrealEngineConnection(false)
			models.SystemStatusStoreInstance.UpdateUnrealEngineReady(false)
			return
		}

		req.Header.Set("Content-Type", "application/json")

		resp, err := client.Do(req)
		if err != nil {
			log.Printf("Error sending request to Unreal Engine: %v", err)
			// Update system status to show connection failure
			models.SystemStatusStoreInstance.UpdateUnrealEngineConnection(false)
			models.SystemStatusStoreInstance.UpdateUnrealEngineReady(false)
			return
		}
		defer resp.Body.Close()

		// Update system status based on response
		if resp.StatusCode >= 200 && resp.StatusCode < 300 {
			models.SystemStatusStoreInstance.UpdateUnrealEngineReady(true)
		} else {
			models.SystemStatusStoreInstance.UpdateUnrealEngineReady(false)
		}

		log.Printf("Unreal Engine response status: %s", resp.Status)
	*/
}
