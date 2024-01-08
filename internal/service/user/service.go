package user

import (
	"context"
	"database/sql"

	db "nam_0801/internal/repo/dbmodel"
)

type (
	Service struct {
		DatabaseConn *sql.DB

		userRepo UserRepository
	}

	UserRepository interface {
		CreateUser(ctx context.Context, user db.CreateUserParams) (db.User, error)
	}
)

func NewUserService(DatabaseConn *sql.DB, userRepo UserRepository) *Service {
	return &Service{
		DatabaseConn: DatabaseConn,
		userRepo:     userRepo,
	}
}
