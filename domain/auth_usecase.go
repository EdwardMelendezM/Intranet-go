package domain

type AuthUseCase interface {
	Login(body LoginBody) LoginResponse
	Register(body RegisterBody) RegisterResponse
}
