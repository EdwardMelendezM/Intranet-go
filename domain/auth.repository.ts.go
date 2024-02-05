package domain

type AuthRepository interface {
	Login(body LoginBody) LoginResponse
	Register(body RegisterBody) RegisterResponse
}
