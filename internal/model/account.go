package model

import (
	db "nam_0801/internal/repo/dbmodel"
	"time"
)

type Account struct {
	ID        int32       `json:"id"`
	UserID    int32       `json:"user_id"`
	Name      string      `json:"name"`
	Bank      db.BankName `json:"bank"`
	Balance   float64     `json:"balance"`
	CreatedAt time.Time   `json:"created_at"`
}

type CreateAccountRequest struct {
	Name string      `json:"name"`
	Bank db.BankName `json:"bank"`
}

type CreateAccountResponse struct {
	Account
}
