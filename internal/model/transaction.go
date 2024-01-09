package model

import (
	db "nam_0801/internal/repo/dbmodel"
)

type Transaction struct {
	ID              int32              `json:"id"`
	AccountID       int32              `json:"account_id"`
	Bank            db.BankName        `json:"bank"`
	Amount          float64            `json:"amount"`
	TransactionType db.TransactionType `json:"transaction_type"`
}

type CreateTransactionRequest struct {
	AccountID       int32              `json:"account_id"`
	Amount          float64            `json:"amount"`
	TransactionType db.TransactionType `json:"transaction_type"`
}

type CreateTransactionResponse struct {
	Transaction
}

type ListTransactionRequest struct {
	AccountID int32 `json:"account_id"`
}

type ListTransactionResponse struct {
	Transactions []Transaction
}

type GetAccountResponse struct {
	Account
}

type ListAccountsResponse struct {
	Accounts []Account
}
