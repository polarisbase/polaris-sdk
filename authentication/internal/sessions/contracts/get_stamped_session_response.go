package contracts

type GetStampedSessionResponse struct {
	SessionID string `json:"session_id"`
	Stamp     string `json:"stamp"`
}
