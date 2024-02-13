package repositories

import UserModel "blog/internal/modules/user/models"

type UserRepositoryInterface interface {
	Create(user UserModel.User) UserModel.User
	FindByEmail(email string) UserModel.User
}
