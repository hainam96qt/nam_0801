package transaction

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

func (r *Repository) CreateTransaction(ctx context.Context, tx *sql.Tx, transaction db.CreateTransactionParams) (db.Transaction, error) {
	var query = r.Query
	if tx != nil {
		query = query.WithTx(tx)
	}
	t, err := query.CreateTransaction(ctx, transaction)
	if err != nil {
		return t, err
	}

	return t, nil
}

const listTransactions = `-- name: ListTransactions :many
SELECT 
	ts.id, 
	ts.amount, 
	ts.account_id, 
	ts.transaction_type, 
	ts.created_at 
FROM transactions ts join accounts ac on ts.account_id = ac.id
WHERE
	ac.user_id = {{.UserID}}
	{{With .AccountID}}
		AND ac.id = {{.}}
	{{End}}
`

func (r *Repository) ListTransactions(ctx context.Context, userID int32, accountID *int32) ([]db.Transaction, error) {
	conditions := map[string]interface{}{
		"AccountID": accountID,
		"UserID":    userID,
	}

	sqlStr, err := template.TemplateSQL(listTransactions, conditions)
	if err != nil {
		return nil, err
	}

	rows, err := r.DatabaseConn.QueryContext(ctx, sqlStr)
	if err != nil {
		return nil, err
	}

	defer rows.Close()
	var items []db.Transaction
	for rows.Next() {
		var i db.Transaction
		if err := rows.Scan(
			&i.ID,
			&i.Amount,
			&i.AccountID,
			&i.TransactionType,
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
