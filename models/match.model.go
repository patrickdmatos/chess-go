package models

type Match struct {
    ID          string `json:"id"`
    Player1ID  string `json:"player1_id"`
    Player2ID  string `json:"player2_id"`
    Status      string `json:"status"` // e.g., "ongoing", "completed", "casual", "ranked", etc.
    WinnerID    string `json:"winner_id,omitempty"`
}
