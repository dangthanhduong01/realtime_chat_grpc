package db

type CreateUserParams struct {
	ID        string `json:"id"`
	Name      string `json:"name"`
	AvatarUrl string `json:"avatar_url"`
	Bio       string `json:"bio"`
	Birthday  string `json:"birthday"`
}

// const createUserSQL = `
// 	INSERT INTO users (id, name, avatar_url, bio, birthday)
// 	VALUES ($1, $2, $3, $4, $5)
// 	RETURNING id, name, avatar_url, bio, birthday, created_at
// `

// func (q *Queries) CreateUser(ctx context.Context, arg CreateUserParams) (model.User, error) {
// 	var user model.User
// 	err := q.db.QueryRowContext(ctx, createUserSQL,
// 		arg.ID,
// 		arg.Name,
// 		arg.AvatarUrl,
// 		arg.Bio,
// 		arg.Birthday,
// 	).Scan(
// 		&user.ID,
// 		&user.Name,
// 		&user.AvatarUrl,
// 		&user.Bio,
// 		&user.Birthday,
// 		&user.CreatedAt,
// 	)
// 	return user, err
// }

// const getUserByID = `SELECT id, name, avatar_url, bio, birthday, created_at
// FROM users WHERE id = $1 LIMIT 1`

// func (q *Queries) GetUserByID(ctx context.Context, id string) (model.User, error) {
// 	var user model.User
// 	err := q.db.QueryRowContext(ctx, getUserByID, id).Scan(
// 		&user.ID,
// 		&user.Name,
// 		&user.AvatarUrl,
// 		&user.Bio,
// 		&user.Birthday,
// 		&user.CreatedAt,
// 	)
// 	return user, err
// }
