package account

import (
	"context"
	"nam_0801/internal/model"
	db "nam_0801/internal/repo/dbmodel"
)

func (s *Service) ListAccount(ctx context.Context, userID int32) (*model.ListAccountsResponse, error) {
	accounts, err := s.accountRepo.ListAccounts(ctx, userID, nil)
	if err != nil {
		return nil, err
	}

	return &model.ListAccountsResponse{
		Accounts: convertAccountsDBToAPI(accounts),
	}, nil
}

func convertAccountsDBToAPI(accounts []db.Account) []model.Account {
	var result []model.Account
	for _, v := range accounts {
		result = append(result, convertAccountDBToAPI(v))
	}
	return result
}
