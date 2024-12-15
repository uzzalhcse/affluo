// internal/dto/auth_response.go
package dto

type AuthResponse struct {
	Token     string `json:"token"`
	Type      string `json:"type"`
	ExpiresAt int64  `json:"expires_at"`
}
