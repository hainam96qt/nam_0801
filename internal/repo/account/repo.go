package account

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

func (r *Repository) CreateAccount(ctx context.Context, account db.CreateAccountParams) (db.Account, error) {
	a, err := r.Query.CreateAccount(ctx, account)
	if err != nil {
		return a, err
	}

	return a, nil
}

func (r *Repository) GetWagerForUpdate(ctx context.Context, tx *sql.Tx, accountID int32) (*db.Account, error) {
	var query = r.Query
	if tx != nil {
		query = query.WithTx(tx)
	}
	wagerDB, err := query.GetAccountForUpdate(ctx, accountID)
	if err != nil {
		return nil, err
	}
	return &wagerDB, nil
}

func (r *Repository) Update(ctx context.Context, tx *sql.Tx, accountID int32) (*db.Account, error) {
	var query = r.Query
	if tx != nil {
		query = query.WithTx(tx)
	}
	wagerDB, err := query.GetAccountForUpdate(ctx, accountID)
	if err != nil {
		return nil, err
	}
	return &wagerDB, nil
}
