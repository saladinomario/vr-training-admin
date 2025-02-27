// internal/models/system.go
package models

import "sync"

// SystemStatus holds the current status of the system components.
type SystemStatus struct {
	UnrealEngineConnected bool   `json:"unrealEngineConnected"`
	WebSocketConnected    bool   `json:"webSocketConnected"`
	UnrealEngineReady     bool   `json:"unrealEngineReady"`
	UnrealEngineUrl       string `json:"unrealEngineUrl"`
}

// SystemStatusStore is a thread-safe store for the system status.
type SystemStatusStore struct {
	status SystemStatus
	mu     sync.RWMutex
}

// NewSystemStatusStore creates a new SystemStatusStore.
func NewSystemStatusStore(unrealEngineUrl string) *SystemStatusStore {
	return &SystemStatusStore{
		status: SystemStatus{
			UnrealEngineConnected: false,
			WebSocketConnected:    false,
			UnrealEngineReady:     false,
			UnrealEngineUrl:       unrealEngineUrl,
		},
	}
}

// Get returns a copy of the current system status.
func (s *SystemStatusStore) Get() SystemStatus {
	s.mu.RLock()
	defer s.mu.RUnlock()
	return s.status
}

// Update updates the system status.
func (s *SystemStatusStore) Update(newStatus SystemStatus) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.status = newStatus
}

// UpdateUnrealEngineConnection updates the UnrealEngineConnected field.
func (s *SystemStatusStore) UpdateUnrealEngineConnection(connected bool) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.status.UnrealEngineConnected = connected
}

// UpdateWebSocketConnection updates the WebSocketConnected field.
func (s *SystemStatusStore) UpdateWebSocketConnection(connected bool) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.status.WebSocketConnected = connected
}

// UpdateUnrealEngineReady updates the UnrealEngineReady field.
func (s *SystemStatusStore) UpdateUnrealEngineReady(ready bool) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.status.UnrealEngineReady = ready
}

// Global variable to store the system status
var SystemStatusStoreInstance *SystemStatusStore
