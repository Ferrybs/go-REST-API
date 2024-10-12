package usecase

import (
	"blog/api/src/adapters/repositories"
	"blog/api/src/domain"
	"blog/api/src/infrastructure/auth"
	"blog/api/src/utils"
	"errors"
)

type UserUsecase struct {
	userRepository repositories.IUserRepository
}

func NewUserUsecase(userRepository repositories.IUserRepository) *UserUsecase {
	return &UserUsecase{userRepository: userRepository}
}

func (u *UserUsecase) CreateUser(userDTO *domain.UserDTO) (*domain.UserLoginView, error) {
	searched_user, err := u.userRepository.FindByUsername(userDTO.Username)
	if err != nil {
		return nil, err
	}
	if searched_user != nil {
		return nil, errors.New("user already exists")
	}
	user, err := u.userRepository.CreateUser(userDTO)
	if err != nil {
		return nil, err
	}
	token, err := auth.GenerateToken(user.UserID)
	if err != nil {
		return nil, err
	}

	userView := &domain.UserView{
		UserID:   user.UserID,
		Username: user.Username,
	}
	return &domain.UserLoginView{
		AccessToken: token,
		User:        *userView,
	}, nil
}

func (u *UserUsecase) FindByUsername(username string) (*domain.UserView, error) {
	user, err := u.userRepository.FindByUsername(username)
	if err != nil {
		return nil, err
	}
	if user == nil {
		return nil, errors.New("user not found")
	}
	return &domain.UserView{
		UserID:   user.UserID,
		Username: user.Username,
	}, nil
}

func (u *UserUsecase) GetUserAccessToken(userLoginDTO *domain.UserLoginDTO) (domain.UserLoginView, error) {
	searched_user, err := u.userRepository.FindByUsername(userLoginDTO.Username)
	var userLoginView domain.UserLoginView
	if err != nil {
		return userLoginView, err
	}
	if searched_user == nil {
		return userLoginView, errors.New("user not found")
	}
	if !utils.CheckPasswordHash(userLoginDTO.Password, searched_user.Password) {
		return userLoginView, errors.New("invalid password")
	}
	token, err := auth.GenerateToken(searched_user.UserID)
	if err != nil {
		return userLoginView, err
	}
	userView := &domain.UserView{
		UserID:   searched_user.UserID,
		Username: searched_user.Username,
	}
	userLoginView = domain.UserLoginView{
		AccessToken: token,
		User:        *userView,
	}
	return userLoginView, nil
}
