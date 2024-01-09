package transaction

import (
	"context"
	"nam_0801/internal/model"
	db "nam_0801/internal/repo/dbmodel"
)

func (s *Service) ListTransactions(ctx context.Context, userID int32, accountID int32) (*model.ListTransactionResponse, error) {
	var accountIDPtr *int32
	if accountID > 0 {
		accountIDPtr = &accountID
	}

	transactions, err := s.transactionRepo.ListTransactions(ctx, userID, accountIDPtr)
	if err != nil {
		return nil, err
	}

	var accountIDs []int32
	for _, v := range transactions {
		accountIDs = append(accountIDs, v.AccountID)
	}

	accounts, err := s.accountRepo.ListAccounts(ctx, userID, accountIDs)
	if err != nil {
		return nil, err
	}

	return &model.ListTransactionResponse{
		Transactions: convertTransactionsDBToAPI(transactions, accounts),
	}, nil
}

func convertTransactionsDBToAPI(trans []db.Transaction, accounts []db.Account) []model.Transaction {
	var mapAccountByID = make(map[int32]db.Account)
	for k, v := range accounts {
		mapAccountByID[v.ID] = accounts[k]
	}

	result := []model.Transaction{}
	for _, v := range trans {
		result = append(result, convertTransactionDBToAPI(v, mapAccountByID[v.AccountID]))
	}
	return result
}
