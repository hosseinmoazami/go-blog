package services

import (
	UserModel "blog/internal/modules/user/models"
	"blog/internal/modules/user/requests/auth"

	UserRepository "blog/internal/modules/user/repositories"
	UserResponse "blog/internal/modules/user/responses"
	"errors"

	"golang.org/x/crypto/bcrypt"
)

type UserService struct {
	userRepository UserRepository.UserRepositoryInterface
}

func New() *UserService {
	return &UserService{
		userRepository: UserRepository.New(),
	}
}

func (userService *UserService) Create(request auth.RegisterRequest) (UserResponse.User, error) {
	var response UserResponse.User
	var user UserModel.User
	hashedPass, err := bcrypt.GenerateFromPassword([]byte(request.Password), 12)
	if err != nil {
		return response, errors.New("error hashing the password")
	}

	user.Name = request.Name
	user.Email = request.Email
	user.Password = string(hashedPass)

	newUser := userService.userRepository.Create(user)

	if newUser.ID == 0 {
		return response, errors.New("error on creating the user")
	}
	return UserResponse.ToUser(newUser), nil
}

func (userService *UserService) CheckUserExist(email string) bool {
	user := userService.userRepository.FindByEmail(email)
	return user.ID != 0
}

func (userService *UserService) HandleUserLogin(request auth.LoginRequest) (UserResponse.User, error) {
	var response UserResponse.User

	user := userService.userRepository.FindByEmail(request.Email)

	if user.ID == 0 {
		return response, errors.New("invalid credentials")
	}

	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(request.Password))
	if err != nil {
		return response, errors.New("invalid credentials")
	}

	return UserResponse.ToUser(user), nil
}
