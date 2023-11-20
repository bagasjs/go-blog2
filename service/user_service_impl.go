package service

import (
	"github.com/bagasjs/go-blog/entity"
	"github.com/bagasjs/go-blog/model"
	"github.com/bagasjs/go-blog/repository"
)

type userServiceImpl struct {
	UserRepository repository.UserRepository
}

func (service *userServiceImpl) Create(request model.CreateUserRequest) (response model.CreateUserResponse, err error) {
    user := entity.User{ Name: request.Name, Password: request.Password, Email: request.Email }
    err = service.UserRepository.Insert(user)
    if err != nil {
        return response, err
    }
    response.Name = user.Name
    response.Email = user.Email
    response.Id = 0
    return response, nil
}

func (service *userServiceImpl) List() (responses []model.GetUserResponse, err error) {
	users, err := service.UserRepository.FindAll()
	if err != nil {
		return nil, err
	}

	for _, user := range users {
		responses = append(responses, model.GetUserResponse{
			Id:    user.Id,
			Name:  user.Name,
			Email: user.Email,
		})
	}

    return responses, nil
}

func NewUserService(userRepository repository.UserRepository) UserService {
	return &userServiceImpl{
		UserRepository: userRepository,
	}
}
