// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.25.0
// source: query.sql

package db

import (
	"context"
)

const createAccount = `-- name: CreateAccount :one
INSERT INTO accounts (
    name,
    bank
) VALUES (
    $1, $2
 ) RETURNING id, user_id, name, bank, balance, created_at
`

type CreateAccountParams struct {
	Name string   `json:"name"`
	Bank BankName `json:"bank"`
}

func (q *Queries) CreateAccount(ctx context.Context, arg CreateAccountParams) (Account, error) {
	row := q.db.QueryRowContext(ctx, createAccount, arg.Name, arg.Bank)
	var i Account
	err := row.Scan(
		&i.ID,
		&i.UserID,
		&i.Name,
		&i.Bank,
		&i.Balance,
		&i.CreatedAt,
	)
	return i, err
}

const createTransaction = `-- name: CreateTransaction :one
INSERT INTO transactions (
    account_id,
    amount,
    transaction_type
) VALUES (
    $1, $2, $3
) RETURNING id, amount, account_id, transaction_type, created_at
`

type CreateTransactionParams struct {
	AccountID       int32           `json:"account_id"`
	Amount          float64         `json:"amount"`
	TransactionType TransactionType `json:"transaction_type"`
}

func (q *Queries) CreateTransaction(ctx context.Context, arg CreateTransactionParams) (Transaction, error) {
	row := q.db.QueryRowContext(ctx, createTransaction, arg.AccountID, arg.Amount, arg.TransactionType)
	var i Transaction
	err := row.Scan(
		&i.ID,
		&i.Amount,
		&i.AccountID,
		&i.TransactionType,
		&i.CreatedAt,
	)
	return i, err
}

const createUser = `-- name: CreateUser :one
INSERT INTO users (
    name,
    email,
    password
) VALUES (
    $1, $2, $3
) RETURNING id, name, email, password, created_at
`

type CreateUserParams struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (q *Queries) CreateUser(ctx context.Context, arg CreateUserParams) (User, error) {
	row := q.db.QueryRowContext(ctx, createUser, arg.Name, arg.Email, arg.Password)
	var i User
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.Email,
		&i.Password,
		&i.CreatedAt,
	)
	return i, err
}

const getAccountForUpdate = `-- name: GetAccountForUpdate :one
SELECT id, user_id, name, bank, balance, created_at FROM accounts WHERE id = $1 FOR UPDATE
`

func (q *Queries) GetAccountForUpdate(ctx context.Context, id int32) (Account, error) {
	row := q.db.QueryRowContext(ctx, getAccountForUpdate, id)
	var i Account
	err := row.Scan(
		&i.ID,
		&i.UserID,
		&i.Name,
		&i.Bank,
		&i.Balance,
		&i.CreatedAt,
	)
	return i, err
}

const updateAccountBalance = `-- name: UpdateAccountBalance :exec
UPDATE accounts SET balance = $1 where id = $2
`

type UpdateAccountBalanceParams struct {
	Balance float64 `json:"balance"`
	ID      int32   `json:"id"`
}

func (q *Queries) UpdateAccountBalance(ctx context.Context, arg UpdateAccountBalanceParams) error {
	_, err := q.db.ExecContext(ctx, updateAccountBalance, arg.Balance, arg.ID)
	return err
}
