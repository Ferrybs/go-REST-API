package repositories

import (
	"blog/api/src/domain"
	"blog/api/src/utils"

	"github.com/google/uuid"
)

type UserRepositoryMemory struct {
	users []*domain.User
}

func NewUserRepositoryMemory() *UserRepositoryMemory {
	return &UserRepositoryMemory{}
}

func (r *UserRepositoryMemory) CreateUser(userDTO *domain.UserDTO) (*domain.User, error) {
	hashedPassword, err := utils.HashPassword(userDTO.Password)
	if err != nil {
		return nil, err
	}
	user := &domain.User{
		UserID:   domain.UserID(uuid.New().String()),
		Username: userDTO.Username,
		Password: hashedPassword,
	}

	r.users = append(r.users, user)
	return user, nil
}

func (r *UserRepositoryMemory) FindByUsername(username string) (*domain.User, error) {
	for _, user := range r.users {
		if user.Username == username {
			return user, nil
		}
	}
	return nil, nil
}
