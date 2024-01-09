package transaction

import (
	"context"
	"database/sql"
	db "nam_0801/internal/repo/dbmodel"
)

type (
	Service struct {
		DatabaseConn *sql.DB

		transactionRepo TransactionRepository
		accountRepo     AccountRepository
	}

	TransactionRepository interface {
		CreateTransaction(ctx context.Context, tx *sql.Tx, transaction db.CreateTransactionParams) (db.Transaction, error)
		ListTransactions(ctx context.Context, userID int32, accountID *int32) ([]db.Transaction, error)
	}

	AccountRepository interface {
		GetAccountForUpdate(ctx context.Context, tx *sql.Tx, arg db.GetAccountForUpdateParams) (*db.Account, error)
		UpdateAccountBalance(ctx context.Context, tx *sql.Tx, arg db.UpdateAccountBalanceParams) error
		ListAccounts(ctx context.Context, userID int32, accountIDs []int32) ([]db.Account, error)
	}
)

func NewTransactionService(DatabaseConn *sql.DB, transactionRepo TransactionRepository, accountRepo AccountRepository) *Service {
	return &Service{
		DatabaseConn:    DatabaseConn,
		transactionRepo: transactionRepo,
		accountRepo:     accountRepo,
	}
}
