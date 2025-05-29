package api

import (
	"snowApp/internal/controller"
	"snowApp/internal/utils"

	pb "snowApp/gen"

	"github.com/gin-gonic/gin"
)

type Server struct {
	pb.UnimplementedChatServiceServer
	config     utils.Config
	controller *controller.ChatController
	router     *gin.Engine
}

func NewServer(config utils.Config, chatController *controller.ChatController) (*Server, error) {
	server := &Server{
		config:     config,
		controller: chatController,
	}

	router := gin.Default()
	server.router = router

	server.setupRouter()

	return server, nil
}

func (server *Server) setupRouter() {
	router := gin.Default()
	// router.GET("users", server.getUsers)
	// router.POST("/user/register", server.createUser)
	// router.POST("/user/login", server.getUser)
	// router.GET("/user/:id", server.getUser)

	server.router = router
}

func (server *Server) Start(address string) error {
	return server.router.Run(address)
}

func errorResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}
