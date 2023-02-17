package socket

// Hub maintains the set of active clients and broadcasts messages to the
// clients.
type Hub struct {
	// Registered clients.
	clients map[string]*Client

	// Inbound messages from the clients.
	broadcast chan Message

	// Register requests from the clients.
	register chan *Client

	// Unregister requests from clients.
	unregister chan *Client
}

func NewHub() *Hub {
	return &Hub{
		broadcast:  make(chan Message),
		register:   make(chan *Client),
		unregister: make(chan *Client),
		clients:    make(map[string]*Client),
	}
}

// Run is a function to handle messages that are comming to hub
func (h *Hub) Run() {
	for {
		select {
		case client := <-h.register:
			h.clients[client.ClientID] = client
		case client := <-h.unregister:
			if _, ok := h.clients[client.ClientID]; ok {
				delete(h.clients, client.ClientID)
				close(client.Send)
			}
		case message := <-h.broadcast:
			if client, ok := h.clients[message.UserID]; ok {
				select {
				case client.Send <- message.Body:
				default:
					close(client.Send)
					delete(h.clients, client.ClientID)
				}
			}
		}
	}
}
