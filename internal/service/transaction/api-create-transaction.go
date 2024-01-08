package transaction

import (
	"context"
	"database/sql"
	"nam_0801/internal/model"
	db "nam_0801/internal/repo/dbmodel"
	error2 "nam_0801/pkg/error"
	"net/http"
)

func (s *Service) CreateTransaction(ctx context.Context, req *model.CreateTransactionRequest) (*model.CreateTransactionResponse, error) {
	tx, err := s.DatabaseConn.BeginTx(ctx, &sql.TxOptions{})
	if err != nil {
		return nil, err
	}
	defer func(tx *sql.Tx) {
		if r := recover(); r != nil {
			err, _ = r.(error)
			err := tx.Rollback()
			if err != nil {
				return
			}
			return
		}
		err := tx.Commit()
		if err != nil {
			return
		}
	}(tx)


	// get account
	account, err := s.accountRepo.GetWagerForUpdate(ctx, tx, req.AccountID)
	if err != nil {
		return nil, err
	}

	newAccountBalance := account.Balance
	if req.TransactionType == db.TransactionTypeWithdraw {
		newAccountBalance = account.Balance - req.Amount
 	} else {
		newAccountBalance = account.Balance + req.Amount
	}

	if newAccountBalance < 0 {
		return nil, error2.NewXError("invalid amount", http.StatusBadRequest)
	}

	newTrans := db.CreateTransactionParams{
		AccountID:       req.AccountID,
		Amount:          req.Amount,
		TransactionType: req.TransactionType,
	}

	a, err := s.transactionRepo.CreateTransaction(ctx, tx, newTrans)
	if err != nil {
		return nil, err
	}

	return &model.CreateTransactionResponse{
		Transaction: convertTransactionDBToAPI(a),
	}, nil
}

func convertTransactionDBToAPI(trans db.Transaction) model.Transaction {
	return model.Transaction{
		ID:        trans.ID,
		UserID:    trans.,
		Name:      trans.Name,
		Bank:      trans.Bank,
		Balance:   trans.Balance,
		CreatedAt: account.CreatedAt,
	}
}
