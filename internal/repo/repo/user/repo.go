package user

import (
	"context"
	"database/sql"

	db "nam_0801/internal/repo/dbmodel"
)

type Repository struct {
	DatabaseConn *sql.DB
	Query        *db.Queries
}

func NewPostgresRepository(databaseConn *sql.DB) *Repository {
	query := db.New(databaseConn)
	return &Repository{
		Query:        query,
		DatabaseConn: databaseConn,
	}
}

func (r *Repository) CreateUser(ctx context.Context, user db.CreateUserParams) (db.User, error) {
	u, err := r.Query.CreateUser(ctx, user)
	if err != nil {
		return u, err
	}

	return u, nil
}

func (r *Repository) GetUserByEmail(ctx context.Context, email string) (db.User, error) {
	u, err := r.Query.GetUserByEmail(ctx, email)
	if err != nil {
		return u, err
	}

	return u, nil
}
