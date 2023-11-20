package repository

import "github.com/bagasjs/go-blog/entity"

type UserRepository interface {
    Insert(user entity.User) error
    FindAll() (users []entity.User, err error)
    DeleteAll() error
}
