package repositories

import "blog/api/src/domain"

type IUserRepository interface {
	CreateUser(userDTO *domain.UserDTO) (*domain.User, error)
	FindByUsername(username string) (*domain.User, error)
}
