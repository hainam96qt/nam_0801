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
	}

	AccountRepository interface {
		GetWagerForUpdate(ctx context.Context, tx *sql.Tx, accountID int32) (*db.Account, error)
	}
)

func NewTransactionService(DatabaseConn *sql.DB, transactionRepo TransactionRepository, accountRepo AccountRepository) *Service {
	return &Service{
		DatabaseConn:    DatabaseConn,
		transactionRepo: transactionRepo,
		accountRepo:     accountRepo,
	}
}
