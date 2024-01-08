package model

import (
	db "nam_0801/internal/repo/dbmodel"
	"time"
)

type Transaction struct {
	ID        int32       `json:"id"`
	UserID    int32       `json:"user_id"`
	Name      string      `json:"name"`
	Bank      db.BankName `json:"bank"`
	Balance   float64     `json:"balance"`
	CreatedAt time.Time   `json:"created_at"`
}

type CreateTransactionRequest struct {
	AccountID       int32              `json:"account_id"`
	Amount          float64            `json:"amount"`
	TransactionType db.TransactionType `json:"transaction_type"`
}

type CreateTransactionResponse struct {
	Transaction
}
