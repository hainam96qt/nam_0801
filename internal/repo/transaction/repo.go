package transaction

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

func (r *Repository) CreateTransaction(ctx context.Context, trans db.CreateTransactionParams) (db.Transaction, error) {
	t, err := r.Query.CreateTransaction(ctx, trans)
	if err != nil {
		return t, err
	}

	return t, nil
}
