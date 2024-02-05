package domain

type LoginBody struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type LoginResponse struct {
	Status       int     `json:"status" binding:"required"`
	Error        *string `json:"error"`
	RefreshToken *string `json:"refresh_token"`
}

type RegisterBody struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
	Birthday string `json:"birthday" binding:"required"`
}

type RegisterResponse struct {
	Status       int     `json:"status" binding:"required"`
	Error        *string `json:"error"`
	RefreshToken *string `json:"refresh_token"`
}
