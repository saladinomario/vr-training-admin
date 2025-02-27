# VR Training WebSocket API

## Overview
This WebSocket API provides real-time session management for VR training applications. Clients can create, update, and retrieve session information through WebSocket connections.

## Connection
- **Endpoint**: `/ws/sessions`
- **Protocol**: WebSocket

## Message Structure
All messages follow this JSON structure:
```json
{
  "command": string,
  "data": object,
  "response_type": string (optional),
  "message": string (optional)
}
```

## Commands

### 1. Create Session
- **Command**: `create_session`
- **Description**: Create a new VR training session
- **Request Payload**:
```json
{
  "command": "create_session",
  "data": {
    "scenarioId": string,
    "avatarId": string,
    "observerId": string
  }
}
```
- **Success Response**:
  - `response_type`: `"success"`
  - Returns created session details

### 2. Update Session
- **Command**: `update_session`
- **Description**: Update an existing session's status
- **Request Payload**:
```json
{
  "command": "update_session",
  "data": {
    "sessionId": string,
    "status": string,
    "notes": string (optional),
    "score": number (optional)
  }
}
```
- **Allowed Status Values**:
  - `running`
  - `paused`
  - `completed`
  - `failed`

### 3. Get Active Sessions
- **Command**: `get_active_sessions`
- **Description**: Retrieve all currently active sessions
- **Request Payload**:
```json
{
  "command": "get_active_sessions"
}
```
- **Returns**: List of active sessions (running or paused)

### 4. Get Session JSON
- **Command**: `get_session_json`
- **Description**: Retrieve specific session details or all sessions
- **Request Payload**:
```json
{
  "command": "get_session_json",
  "data": {
    "sessionId": string (optional)
  }
}
```
- **Behavior**:
  - If `sessionId` provided: Returns details for specific session
  - If no `sessionId`: Returns all sessions

## Error Handling
- Error responses have:
  - `response_type`: `"error"`
  - `message`: Description of the error

## Example Workflow
1. Connect to WebSocket
2. Create a session
3. Update session status
4. Retrieve active sessions

## Notes
- Ensure proper error handling in your client implementation
- WebSocket connection remains open for multiple commands
- All commands are case-sensitive

## Authentication
*(To be implemented)*
- Future versions may require authentication token

## Potential Improvements
- Add more detailed session metadata
- Implement real-time session tracking
- Add authentication mechanism
