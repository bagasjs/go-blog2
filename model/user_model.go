package model

type CreateUserRequest struct {
    Email                string
	Name                 string
	Password             string
	PasswordConfirmation string
}

type CreateUserResponse struct {
	Id    int
	Name  string
    Email string
}

type GetUserResponse struct {
	Id    int
	Name  string
	Email string
}
