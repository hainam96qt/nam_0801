package account

import (
	"context"
	"database/sql"
	db "nam_0801/internal/repo/dbmodel"
)

type (
	Service struct {
		DatabaseConn *sql.DB

		accountRepo AccountRepository
	}

	AccountRepository interface {
		CreateAccount(ctx context.Context, user db.CreateAccountParams) (db.Account, error)
	}
)

func NewAccountService(DatabaseConn *sql.DB, accountRepo AccountRepository) *Service {
	return &Service{
		DatabaseConn: DatabaseConn,
		accountRepo:  accountRepo,
	}
}
