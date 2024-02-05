package domain

type LoginBody struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type LoginResponse struct {
	Status       int
	Error        *string
	RefreshToken *string
}

type RegisterBody struct {
	Username string
	Password string
	Birthday string
}

type RegisterResponse struct {
	Status       int
	Error        *string
	RefreshToken *string
}
