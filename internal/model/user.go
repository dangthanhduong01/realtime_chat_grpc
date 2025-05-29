package model

import (
	"time"

	"golang.org/x/crypto/bcrypt"
)

type User struct {
	ID           string    `bson:"_id" json:"id"`
	Username     string    `bson:"username" json:"username"`
	Email        string    `bson:"email" json:"email"`
	AvatarUrl    string    `bson:"avatar_url" json:"avatar_url"`
	Bio          string    `bson:"bio" json:"bio"`
	PasswordHash string    `bson:"password_hash" json:"-"`
	Birthday     time.Time `bson:"birthday" json:"birthday"`
	CreatedAt    time.Time `bson:"created_at" json:"created_at"`
	UpdatedAt    time.Time `bson:"updated_at" json:"updated_at"`
}

func (u *User) SetPassword(password string) error {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	u.PasswordHash = string(hash)
	return nil
}

func (u *User) CheckPassword(password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(u.PasswordHash), []byte(password))
	return err == nil
}
