package model

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type LoginResponse2 struct {
	Token string `json:"token"`
}
