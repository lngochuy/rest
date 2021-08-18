package resources

import "time"

type UserResponse struct {
	ID          uint      `json:"id"`
	Email       string    `json:"email"`
	DateOfBirth time.Time `json:"date_of_birth"`
}

type CreateUserRequest struct {
	Email       string `json:"email" binding:"required"`
	DateOfBirth myTime `json:"date_of_birth" binding:"required"`
}

type UpdateUserRequest struct {
	Email       string `json:"email" binding:"required"`
	DateOfBirth myTime `json:"date_of_birth" binding:"required"`
}
