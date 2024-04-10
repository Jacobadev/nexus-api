package model

// Session model
type Session struct {
	SessionID string `json:"session_id" redis:"session_id"`
	ID        int    `json:"user_id" redis:"user_id"`
}
