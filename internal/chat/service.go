package chat

// import (
// 	"context"
// 	pb "snowApp/gen"
// 	mongo "snowApp/internal/repository/mongo"
// 	redis "snowApp/internal/repository/redis"
// 	"time"

// 	"go.mongodb.org/mongo-driver/bson"
// 	// "go.mongodb.org/mongo-driver/mongo"

// 	// "go.mongodb.org/mongo-driver/mongo/options"
// 	"google.golang.org/grpc/codes"
// 	"google.golang.org/grpc/status"

// 	"google.golang.org/protobuf/types/known/timestamppb"
// )

// type ChatServer struct {
// 	pb.UnimplementedChatServiceServer
// 	mongoClient *mongo.MongoClient
// 	redisClient *redis.RedisClient
// }

// func NewChatServer(mongoRepo *mongo.MongoClient, redisRepo *redis.RedisClient) *ChatServer {
// 	// mongoClient, err := mongo.Connect(context.Background(), options.Client().ApplyURI(mongoURI))
// 	// if err != nil {
// 	// 	return nil, err
// 	// }

// 	// redisClient := redis.NewClient(&redis.Options{
// 	// 	Addr:     "localhost:6379",
// 	// 	Password: "",
// 	// 	DB:       0,
// 	// })

// 	return &ChatServer{
// 		mongoClient: mongoRepo,
// 		redisClient: redisRepo,
// 	}
// }

// func (s *ChatServer) SendMessage(ctx context.Context, req *pb.SendMessageRequest) (*pb.SendMessageResponse, error) {
// 	collection := s.mongoClient.Database("chatdb").Collection("messages")

// 	message := pb.Message{
// 		SenderId:       req.SenderId,
// 		ConversationId: req.ConversationId,
// 		Content:        req.Content,
// 		SentAt:         timestamppb.Now(),
// 	}

// 	insertResult, err := collection.InsertOne(ctx, message)
// 	if err != nil {
// 		return nil, status.Errorf(codes.Internal, "Failed to insert message: %v", err)
// 	}

// 	err = s.redisClient.Publish(ctx, req.ConversationId, insertResult.InsertedID).Err()
// 	if err != nil {
// 		return nil, status.Errorf(codes.Internal, "Failed to publish message to Redis: %v", err)
// 	}

// 	return &pb.SendMessageResponse{
// 		MessageId: insertResult.InsertedID.(string),
// 		SentAt:    message.SentAt,
// 	}, nil
// }

// func (s *ChatServer) CreateConversation(ctx context.Context, req *pb.CreateConversationRequest) (*pb.Conversation, error) {

// 	// Validate the request
// 	if len(req.ParticipantIds) < 2 {
// 		return nil, status.Errorf(codes.InvalidArgument, "At least two participants are required to create a conversation")
// 	}

// 	conv := pb.Conversation{
// 		Id:             generateConversationId(),
// 		ParticipantIds: req.ParticipantIds,
// 		CreatedAt:      timestamppb.Now(),
// 		UpdatedAt:      timestamppb.Now(),
// 	}

// 	switch req.ConversationType.(type) {
// 	case *pb.CreateConversationRequest_Private:
// 		conv.Type = pb.Conversation_PRIVATE
// 	case *pb.CreateConversationRequest_Group:
// 		group := req.GetGroup()
// 		conv.Type = pb.Conversation_GROUP
// 		conv.Title = group.Title
// 		conv.AvatarUrl = *group.AvatarUrl
// 	}

// 	if err := s.saveConversation(ctx, &conv); err != nil {
// 		return nil, status.Errorf(codes.Internal, "Failed to save conversation: %v", err)
// 	}

// 	return &conv, nil
// }

// func (s *ChatServer) GetConversations(ctx context.Context, req *pb.GetConversationsRequest) (*pb.GetConversationsResponse, error) {
// 	collection := s.mongoClient.Database("chatdb").Collection("conversations")

// 	// Find conversations for the user
// 	cursor, err := collection.Find(ctx, bson.M{"participant_ids": req.UserId})
// 	if err != nil {
// 		return nil, status.Errorf(codes.Internal, "Failed to fetch conversations: %v", err)
// 	}
// 	defer cursor.Close(ctx)

// 	var conversations []*pb.Conversation
// 	for cursor.Next(ctx) {
// 		var conv pb.Conversation
// 		if err := cursor.Decode(&conv); err != nil {
// 			return nil, status.Errorf(codes.Internal, "Failed to decode conversation: %v", err)
// 		}
// 		conversations = append(conversations, &conv)
// 	}

// 	return &pb.GetConversationsResponse{Conversations: conversations}, nil
// }

// func (s *ChatServer) UpdateConversation(ctx context.Context, req *pb.UpdateConversationRequest) (*pb.Conversation, error) {
// 	collection := s.mongoClient.Database("chatdb").Collection("conversations")

// 	conv, err := s.getConversation(ctx, req.ConversationId)
// 	if err != nil {
// 		return nil, status.Errorf(codes.NotFound, "Conversation not found: %v", err)
// 	}

// 	switch update := req.Update.(type) {
// 	case *pb.UpdateConversationRequest_Title:
// 		conv.Title = update.Title
// 	case *pb.UpdateConversationRequest_AvatarUrl:
// 		conv.AvatarUrl = update.AvatarUrl
// 		// case *pb.UpdateConversationRequest_AddParticipants:
// 		// 	conv.ParticipantIds = append(conv.ParticipantIds, update.AddParticipants...)
// 		// case *pb.UpdateConversationRequest_RemoveParticipants:
// 		// 	for _, id := range update.RemoveParticipants {
// 		// 		for i, participantId := range conv.ParticipantIds {
// 		// 			if participantId == id {
// 		// 				conv.ParticipantIds = append(conv.ParticipantIds[:i], conv.ParticipantIds[i+1:]...)
// 		// 				break
// 		// 			}
// 		// 		}
// 		// 	}
// 	}

// 	conv.UpdatedAt = timestamppb.Now()

// 	if _, err := s.updateConversation(ctx, update); err != nil {
// 		return nil, status.Errorf(codes.Internal, "Failed to update conversation: %v", err)
// 	}
// 	// Update the conversation
// 	update := bson.M{
// 		"$set": bson.M{
// 			"title":      req.Title,
// 			"avatar_url": req.AvatarUrl,
// 			"updated_at": time.Now(),
// 		},
// 	}

// 	result, err := collection.UpdateOne(ctx, bson.M{"id": req.ConversationId}, update)
// 	if err != nil {
// 		return nil, status.Errorf(codes.Internal, "Failed to update conversation: %v", err)
// 	}
// 	if result.ModifiedCount == 0 {
// 		return nil, status.Errorf(codes.NotFound, "Conversation not found")
// 	}

// 	return &pb.Conversation{Id: req.ConversationId}, nil
// }

// func generateConversationId() string {
// 	// This is a placeholder. In a real application, you would generate a unique ID.
// 	return "conv_" + time.Now().Format("20060102150405")
// }

// func (s *ChatServer) saveConversation(ctx context.Context, conv *pb.Conversation) error {
// 	collection := s.mongoClient.Database("chatdb").Collection("conversations")

// 	_, err := collection.InsertOne(ctx, conv)
// 	if err != nil {
// 		return status.Errorf(codes.Internal, "Failed to save conversation: %v", err)
// 	}

// 	return nil
// }

// func (s *ChatServer) updateConversation(ctx context.Context, req *pb.UpdateConversationRequest) (*pb.Conversation, error) {
// 	collection := s.mongoClient.Database("chatdb").Collection("conversations")

// 	// Update the conversation
// 	update := bson.M{
// 		"$set": bson.M{
// 			"title":      req.Title,
// 			"avatar_url": req.AvatarUrl,
// 			"updated_at": time.Now(),
// 		},
// 	}

// 	result, err := collection.UpdateOne(ctx, bson.M{"id": req.ConversationId}, update)
// 	if err != nil {
// 		return nil, status.Errorf(codes.Internal, "Failed to update conversation: %v", err)
// 	}
// 	if result.ModifiedCount == 0 {
// 		return nil, status.Errorf(codes.NotFound, "Conversation not found")
// 	}

// 	return &pb.Conversation{Id: req.ConversationId}, nil
// }

// func (s *ChatServer) getConversation(ctx context.Context, conversationId string) (*pb.Conversation, error) {
// 	collection := s.mongoClient.Database("chatdb").Collection("conversations")

// 	var conv pb.Conversation
// 	err := collection.FindOne(ctx, bson.M{"id": conversationId}).Decode(&conv)
// 	if err != nil {
// 		if err == mongo.ErrNoDocuments {
// 			return nil, status.Errorf(codes.NotFound, "Conversation not found")
// 		}
// 		return nil, status.Errorf(codes.Internal, "Failed to fetch conversation: %v", err)
// 	}

// 	return &conv, nil
// }

// // ========== Phân quyền trạng thái người dùng ==========
// func (s *ChatServer) UpdateUserStatus(ctx context.Context, req *pb.UpdateUserStatusRequest) (*pb.UpdateUserStatusResponse, error) {
// 	// Validate
// 	if req.UserId == "" {
// 		return nil, status.Error(codes.InvalidArgument, "user_id is required")
// 	}

// 	collection := s.mongoClient.Database("chatdb").Collection("users")

// 	// Update the user's status
// 	update := bson.M{
// 		"$set": bson.M{
// 			"status":     req.Status,
// 			"updated_at": time.Now(),
// 		},
// 	}

// 	result, err := collection.UpdateOne(ctx, bson.M{"id": req.UserId}, update)
// 	if err != nil {
// 		return nil, status.Errorf(codes.Internal, "Failed to update user status: %v", err)
// 	}
// 	if result.ModifiedCount == 0 {
// 		return nil, status.Errorf(codes.NotFound, "User not found")
// 	}

// 	return &pb.UpdateUserStatusResponse{Success: true}, nil
// }
// func (s *ChatServer) GetUserStatus(ctx context.Context, req *pb.GetUserStatusRequest) (*pb.UserStatus, error) {
// 	if len(req.UserIds) == 0 {
// 		return nil, status.Error(codes.InvalidArgument, "at least one user_id required")
// 	}

// 	statuses, err := s.statusRepo.BatchGetUserStatus(ctx, req.UserIds)
// 	if err != nil {
// 		return nil, status.Errorf(codes.Internal, "failed to get statuses: %v", err)
// 	}

// 	// Trả về status của user đầu tiên (nếu query 1 user)
// 	// Hoặc có thể định nghĩa lại response thành BatchUserStatusResponse
// 	if len(statuses) > 0 {
// 		return statuses[0], nil
// 	}

// 	return nil, status.Error(codes.NotFound, "user status not found")
// }

// // ========== Phần gọi thoại/video ==========

// // func (s *ChatServer) StartCall(ctx context.Context, req *pb.StartCallRequest) (*pb.CallInfo, error) {
// // 	// Kiểm tra quyền
// // 	if !s.canStartCall(ctx, req.CallerId, req.ConversationId) {
// // 		return nil, status.Error(codes.PermissionDenied, "cannot start call in this conversation")
// // 	}

// // 	// Tạo cuộc gọi mới
// // 	call := &pb.CallInfo{
// // 		CallId:         generateCallID(),
// // 		Status:         pb.CallInfo_RINGING,
// // 		ParticipantIds: s.getConversationParticipants(ctx, req.ConversationId),
// // 		StartedAt:      timestamppb.Now(),
// // 	}

// // 	// Lưu thông tin cuộc gọi
// // 	if err := s.callRepo.CreateCall(ctx, call); err != nil {
// // 		return nil, status.Errorf(codes.Internal, "failed to create call: %v", err)
// // 	}

// // 	// Gửi event đến các participants
// // 	s.notifyCallStarted(call, req.CallerId)

// // 	// Trả về ICE servers (STUN/TURN)
// // 	call.IceServers = s.getICEServers()

// // 	return call, nil
// // }

// // func (s *ChatServer) HandleCallSignal(stream pb.ChatService_HandleCallSignalServer) error {
// // 	// Xử lý stream tín hiệu call
// // 	for {
// // 		signal, err := stream.Recv()
// // 		if err != nil {
// // 			return status.Errorf(codes.Internal, "failed to receive signal: %v", err)
// // 		}

// // 		// Validate call và participant
// // 		if !s.validateCallParticipant(signal.CallId, signal.SenderId) {
// // 			return status.Error(codes.PermissionDenied, "invalid participant")
// // 		}

// // 		// Chuyển tiếp tín hiệu đến các participant khác
// // 		if err := s.routeCallSignal(signal); err != nil {
// // 			return status.Errorf(codes.Internal, "failed to route signal: %v", err)
// // 		}
// // 	}
// // }
