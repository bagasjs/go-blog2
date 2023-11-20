package service

import "github.com/bagasjs/go-blog/model"

type UserService interface {
	Create(request model.CreateUserRequest) (response model.CreateUserResponse, err error)
	List() (response []model.GetUserResponse, err error)
}
