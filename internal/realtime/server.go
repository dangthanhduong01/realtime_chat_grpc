package realtime

import (
	"log"
	"net/http"
	pb "snowApp/gen"

	"github.com/gorilla/websocket"
	"google.golang.org/grpc"
)

type Client struct {
	conn     *websocket.Conn
	userId   string
	sendChan chan []byte
}

type RealtimeServer struct {
	clients    map[string]*Client
	chatClient pb.ChatServiceClient
	upgrader   websocket.Upgrader
}

func NewRealtimeServer(grpcAddr string) *RealtimeServer {
	// Connect to chat service through gRPC
	conn, err := grpc.Dial(grpcAddr, grpc.WithInsecure())
	if err != nil {
		log.Fatal("Failed to connect to gRPC server:", err)
	}

	return &RealtimeServer{
		clients:    make(map[string]*Client),
		chatClient: pb.NewChatServiceClient(conn),
		upgrader: websocket.Upgrader{
			CheckOrigin: func(r *http.Request) bool {
				return true
			},
		},
	}
}

func (s *RealtimeServer) HandleConnection(w http.ResponseWriter, r *http.Request) {
	conn, err := s.upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Printf("Upgrade error: %v", err)
		return
	}

	// Authenticate user (JWT token từ query params)
	userID := r.URL.Query().Get("user_id")
	if userID == "" {
		conn.Close()
		return
	}

	client := &Client{
		conn:     conn,
		userId:   userID,
		sendChan: make(chan []byte, 256),
	}

	s.clients[userID] = client

	go client.writePump()
	go client.readPump(s)
}

func (c *Client) readPump(s *RealtimeServer) {
	defer func() {
		delete(s.clients, c.userId)
		c.conn.Close()
	}()

	for {
		_, message, err := c.conn.ReadMessage()
		if err != nil {
			log.Printf("Read error: %v", err)
			break
		}

		c.sendChan <- message
	}
}

func (c *Client) writePump() {
	defer func() {
		c.conn.Close()
	}()

	for {
		select {
		case message, ok := <-c.sendChan:
			if !ok {
				c.conn.WriteMessage(websocket.CloseMessage, []byte{})
				return
			}

			err := c.conn.WriteMessage(websocket.TextMessage, message)
			if err != nil {
				log.Printf("Write error: %v", err)
				return
			}
		}
	}
}
