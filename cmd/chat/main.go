package main

import (
	"log"
	"net"
	"os"
	"os/signal"
	"snowApp/api"
	"snowApp/internal/controller"
	"snowApp/internal/repository/mongo"
	"snowApp/internal/repository/redis"
	"snowApp/internal/service"
	"snowApp/internal/utils"
	"snowApp/pkg/jwt"
	"snowApp/pkg/logger"
	"syscall"

	"google.golang.org/grpc"

	pb "snowApp/gen"
)

func main() {
	// Initialize MongoDB and Redis repositories
	mongoRepo, err := mongo.New("mongodb://localhost:27017", "chatdb")
	if err != nil {
		log.Fatal("Failed to connect to MongoDB:", err)
	}
	defer mongoRepo.Close()
	redisRepo := redis.New("redis://localhost:6379")
	defer redisRepo.Close()
	// Initialize the chat service
	tokenManager := jwt.NewManager("your-secret-key")
	messageRepository := mongo.NewMessageRepository(mongoRepo.Client, "chatdb")
	conversationRepository := mongo.NewConversationRepository(mongoRepo.Client, "chatdb")
	eventRepository := redis.NewEventRepository(redisRepo.Client)
	userRepository := mongo.NewUserRepository(mongoRepo.Client, "chatdb")
	logger := logger.NewLogger()

	chatService := service.NewChatService(
		messageRepository,
		conversationRepository,
		eventRepository,
		*logger,
	)
	userService := service.NewAuthService(userRepository,
		tokenManager,
		3600,  // Token expiry in seconds
		86400, // Refresh token expiry in seconds
	)

	chatController := controller.NewChatController(chatService, userService, logger)

	config, err := utils.LoadConfig(".")
	if err != nil {
		log.Fatal("Cannot load config:", err)
	}
	chatServer, err := api.NewServer(config, chatController)
	if err != nil {
		log.Fatal("Failed to create chat server:", err)
	}
	grpcServer := grpc.NewServer(
		grpc.MaxRecvMsgSize(1024*1024*10), // 10 MB
		grpc.MaxSendMsgSize(1024*1024*10), // 10 MB
	)

	// Register the chat service with the gRPC server
	pb.RegisterChatServiceServer(grpcServer, chatServer)

	// Start the gRPC server
	go func() {
		listener, err := net.Listen("tcp", ":50051")
		if err != nil {
			log.Fatal("Failed to listen on port 50051:", err)
		}
		log.Println("Chat server is running on port 50051...")
		if err := grpcServer.Serve(listener); err != nil {
			log.Fatal("Failed to serve gRPC server:", err)
		}
	}()

	// Gracefully shutdown the server on interrupt signal
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGINT, syscall.SIGTERM)
	<-stop
	log.Println("Shutting down the server...")
	grpcServer.GracefulStop()
	log.Println("Server stopped.")

	// listener, err := net.Listen("tcp", ":50051")
	// if err != nil {
	// 	log.Fatal("Failed to listen on port 50051:", err)
	// }
	// log.Println("Chat server is running on port 50051...")
	// if err := grpcServer.Serve(listener); err != nil {
	// 	log.Fatal("Failed to serve gRPC server:", err)
	// }
}
