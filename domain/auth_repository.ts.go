package domain

type AuthRepository interface {
	Login(body LoginBody) LoginResponse
	VerifyExistingUser(body LoginBody) (bool bool, err error)
	Register(body RegisterBody) RegisterResponse
}
