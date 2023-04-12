package users

import (
	"time"

	database "github.com/anirudhmpai/database/sqlc"
)

// ? user represents data about a users data.
type User struct {
	ID    string `json:"id"`
	Name  string `json:"name"`
	Age   int64  `json:"age"`
	Email string `json:"email"`
}

// ? SignInInput struct
type SignInInput struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

// ? UserResponse struct
type UserResponse struct {
	ID        string    `json:"id,omitempty"`
	Name      string    `json:"name,omitempty"`
	Email     string    `json:"email,omitempty"`
	Role      string    `json:"role,omitempty"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func FilteredResponse(user database.User) UserResponse {
	return UserResponse{
		ID:        user.ID.String(),
		Email:     user.Email,
		Name:      user.Name,
		Role:      user.Role,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}
}
