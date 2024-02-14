package services

import (
	"blog/internal/modules/user/requests/auth"
	UserResponse "blog/internal/modules/user/responses"
)

type UserServiceInterface interface {
	Create(request auth.RegisterRequest) (UserResponse.User, error)
	CheckUserExist(email string) bool
	HandleUserLogin(request auth.LoginRequest) (UserResponse.User, error)
}
