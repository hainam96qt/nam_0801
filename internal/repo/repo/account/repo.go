package account

import (
	"context"
	"database/sql"
	db "nam_0801/internal/repo/dbmodel"
	"nam_0801/pkg/util/template"
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

func (r *Repository) GetAccountForUpdate(ctx context.Context, tx *sql.Tx, arg db.GetAccountForUpdateParams) (*db.Account, error) {
	var query = r.Query
	if tx != nil {
		query = query.WithTx(tx)
	}
	accountDB, err := query.GetAccountForUpdate(ctx, arg)
	if err != nil {
		return nil, err
	}
	return &accountDB, nil
}

func (r *Repository) GetAccount(ctx context.Context, arg db.GetAccountParams) (*db.Account, error) {
	accountDB, err := r.Query.GetAccount(ctx, arg)
	if err != nil {
		return nil, err
	}
	return &accountDB, nil
}

func (r *Repository) UpdateAccountBalance(ctx context.Context, tx *sql.Tx, arg db.UpdateAccountBalanceParams) error {
	var query = r.Query
	if tx != nil {
		query = query.WithTx(tx)
	}
	err := query.UpdateAccountBalance(ctx, arg)
	if err != nil {
		return err
	}
	return nil
}

const listAccounts = `-- name: ListAccounts :many
SELECT id, user_id, name, bank, balance, created_at 
FROM accounts 
WHERE 
	user_id = {{.UserID}}
	{{With .AccountIDs}} 
		AND id IN ({{ join . "," }})
	{{END}}
`

func (r *Repository) ListAccounts(ctx context.Context, userID int32, accountIDs []int32) ([]db.Account, error) {
	conditions := map[string]interface{}{
		"UserID":     userID,
		"AccountIDs": accountIDs,
	}

	sqlStr, err := template.TemplateSQL(listAccounts, conditions)
	if err != nil {
		return nil, err
	}

	rows, err := r.DatabaseConn.QueryContext(ctx, sqlStr)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []db.Account
	for rows.Next() {
		var i db.Account
		if err := rows.Scan(
			&i.ID,
			&i.UserID,
			&i.Name,
			&i.Bank,
			&i.Balance,
			&i.CreatedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}
