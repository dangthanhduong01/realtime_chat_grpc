package api

// import (
// 	"snowApp/internal/model"
// 	"time"

// 	"github.com/gin-gonic/gin"
// )

// type createUserRequest struct {
// 	ID        string `json:"id" binding:"required"`
// 	Username  string `json:"username" binding:"required"`
// 	Email     string `json:"email" binding:"required,email"`
// 	AvatarUrl string `json:"avatar_url" binding:"required"`
// 	Bio       string `json:"bio" binding:"required"`
// 	Password  string `json:"password" binding:"required"`
// 	Birthday  string `json:"birthday" binding:"required"`
// }

// type userResponse struct {
// 	ID        string `json:"id"`
// 	Username  string `json:"username"`
// 	Email     string `json:"email"`
// 	AvatarUrl string `json:"avatar_url"`
// 	Bio       string `json:"bio"`
// 	Birthday  string `json:"birthday"`
// 	CreatedAt string `json:"created_at"`
// }

// type loginUserRequest struct {
// 	Email    string `json:"email" binding:"required"`
// 	Password string `json:"password" binding:"required"`
// }

// type loginUserResponse struct {
// 	ID        string `json:"id"`
// 	Username  string `json:"username"`
// 	Email     string `json:"email"`
// 	AvatarUrl string `json:"avatar_url"`
// 	Bio       string `json:"bio"`
// 	Birthday  string `json:"birthday"`
// 	CreatedAt string `json:"created_at"`
// }

// func newUserResponse(user model.User) userResponse {
// 	return userResponse{
// 		ID:        user.ID,
// 		Username:  user.Username,
// 		Email:     user.Email,
// 		Birthday:  user.Birthday.Format("2006-01-02"),
// 		CreatedAt: user.CreatedAt.Format("2006-01-02 15:04:05"),
// 	}
// }

// func (server *Server) registerUser(ctx *gin.Context) {
// 	var req createUserRequest
// 	if err := ctx.ShouldBindJSON(&req); err != nil {
// 		ctx.JSON(400, errorResponse(err))
// 		return
// 	}

// 	birthday, err := time.Parse("2006-01-02", req.Birthday)
// 	if err != nil {
// 		ctx.JSON(400, errorResponse(err))
// 		return
// 	}

// 	user := model.User{
// 		// ID:        req.ID,
// 		Username:  req.Username,
// 		Email:     req.Email,
// 		AvatarUrl: req.AvatarUrl,
// 		Bio:       req.Bio,
// 		Birthday:  birthday,
// 	}

// 	newUser, err := server.authService.Register(ctx,
// 		user.Username, user.Email, user.Birthday.Format("2006-01-02"), user.AvatarUrl, user.Bio, req.Password)
// 	if err != nil {
// 		ctx.JSON(500, errorResponse(err))
// 		return
// 	}

// 	ctx.JSON(200, newUserResponse(*newUser))
// }

// func (server *Server) loginUser(ctx *gin.Context) {
// 	var req loginUserRequest
// 	if err := ctx.ShouldBindJSON(&req); err != nil {
// 		ctx.JSON(400, errorResponse(err))
// 		return
// 	}

// 	user, accessToken, refreshToken, err := server.authService.Login(ctx, req.Email, req.Password)
// 	if err != nil {
// 		ctx.JSON(401, errorResponse(err))
// 		return
// 	}

// 	response := loginUserResponse{
// 		ID:        user.ID,
// 		Username:  user.Username,
// 		Email:     user.Email,
// 		AvatarUrl: user.AvatarUrl,
// 		Bio:       user.Bio,
// 		Birthday:  user.Birthday.Format("2006-01-02"),
// 		CreatedAt: user.CreatedAt.Format("2006-01-02 15:04:05"),
// 	}

// 	ctx.JSON(200, gin.H{
// 		"user":          response,
// 		"access_token":  accessToken,
// 		"refresh_token": refreshToken,
// 	})
// }
