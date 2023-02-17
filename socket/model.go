package socket

// Message is a type for sending message to a particular user
type Message struct {
	UserID string
	Body   []byte
}

// Notification is body of a notification to be sent.
type Notification struct {
	Title    string `json:"title"`
	Content  string `json:"content"`
	StatusID string `json:"status_id"`
}
