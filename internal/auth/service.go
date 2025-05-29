package auth

// import (
// 	"context"
// 	"snowApp/internal/db"
// 	"snowApp/internal/model"
// )

// type AuthService interface {
// 	RegisterUser(context.Context, UserRegister) (model.User, error)
// 	LoginUser(context.Context, UserLogin) (model.User, error)
// }

// type UserRegister struct {
// 	ID        string `json:"id"`
// 	Name      string `json:"name"`
// 	AvatarUrl string `json:"avatar_url"`
// 	Bio       string `json:"bio"`
// 	Birthday  string `json:"birthday"`
// }

// type UserLogin struct {
// 	ID       string `json:"id"`
// 	Password string `json:"password"`
// }

// type UserService struct {
// 	queries db.Querier
// }

// func (s *UserService) RegisterUser(ctx context.Context, user model.User) (model.User, error) {
// 	// Check if user already exists
// 	existingUser, err := s.queries.GetUserByID(ctx, user.ID)
// 	if err == nil {
// 		return existingUser, nil
// 	}

// 	// Create new user
// 	newUser, err := s.queries.CreateUser(ctx, db.CreateUserParams{
// 		ID:        user.ID,
// 		Name:      user.Name,
// 		AvatarUrl: user.AvatarUrl,
// 		Bio:       user.Bio,
// 		Birthday:  user.Birthday.Format("2006-01-02"),
// 	})
// 	if err != nil {
// 		return model.User{}, err
// 	}

// 	return newUser, nil
// }
