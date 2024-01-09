package transaction

import (
	"context"
	"database/sql"
	"nam_0801/internal/model"
	db "nam_0801/internal/repo/dbmodel"
	error2 "nam_0801/pkg/error"
	"net/http"
)

func (s *Service) CreateTransaction(ctx context.Context, userID int32, req *model.CreateTransactionRequest) (*model.CreateTransactionResponse, error) {
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
	account, err := s.accountRepo.GetAccountForUpdate(ctx, tx, db.GetAccountForUpdateParams{
		ID:     req.AccountID,
		UserID: userID,
	})
	if err != nil {
		return nil, err
	}

	newAccountBalance := account.Balance
	if req.TransactionType == db.TransactionTypeWithdraw {
		newAccountBalance = account.Balance - req.Amount
	} else {
		newAccountBalance = account.Balance + req.Amount
	}

	if req.Amount < 0 || newAccountBalance < 0 {
		return nil, error2.NewXError("invalid amount", http.StatusBadRequest)
	}

	newTrans := db.CreateTransactionParams{
		AccountID:       req.AccountID,
		Amount:          req.Amount,
		TransactionType: req.TransactionType,
	}

	t, err := s.transactionRepo.CreateTransaction(ctx, tx, newTrans)
	if err != nil {
		return nil, err
	}

	err = s.accountRepo.UpdateAccountBalance(ctx, tx, db.UpdateAccountBalanceParams{
		Balance: newAccountBalance,
		ID:      req.AccountID,
	})
	if err != nil {
		return nil, err
	}

	return &model.CreateTransactionResponse{
		Transaction: convertTransactionDBToAPI(t, *account),
	}, nil
}

func convertTransactionDBToAPI(trans db.Transaction, account db.Account) model.Transaction {
	return model.Transaction{
		ID:              trans.ID,
		AccountID:       account.ID,
		Bank:            account.Bank,
		Amount:          trans.Amount,
		TransactionType: trans.TransactionType,
	}
}
