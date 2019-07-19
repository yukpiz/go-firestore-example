package domain

type Message struct {
	SenderID string `json:"sender_id"`
	Message  string `json:"message"`
	SentAt   int64  `json:"sent_at"`
}
