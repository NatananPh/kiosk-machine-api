package service

import (
	"github.com/NatananPh/kiosk-machine-api/pkg/user/model"
	"github.com/NatananPh/kiosk-machine-api/pkg/user/repository"
)

type UserServiceImpl struct {
	userRepository repository.UserRepository
}

func NewUserService(userRepository repository.UserRepository) UserService {
	return &UserServiceImpl{
		userRepository: userRepository,
	}
}

func (us *UserServiceImpl) GetUsers() ([]*model.User, error) {
	users, err := us.userRepository.GetUsers()
	if err != nil {
		return nil, err
	}
	var modelUsers []*model.User
	for _, user := range users {
		modelUsers = append(modelUsers, &model.User{
			ID:    user.ID,
			Username:  user.Username,
			Password: user.Password,
			RoleID: user.RoleID,
		})
	}
	return modelUsers, nil
}