package account

import (
	"context"
	"nam_0801/internal/model"
	db "nam_0801/internal/repo/dbmodel"
)

func (s *Service) CreateAccount(ctx context.Context, userID int32, req *model.CreateAccountRequest) (*model.CreateAccountResponse, error) {
	newAccount := db.CreateAccountParams{
		Name:   req.Name,
		Bank:   req.Bank,
		UserID: userID,
	}
	a, err := s.accountRepo.CreateAccount(ctx, newAccount)
	if err != nil {
		return nil, err
	}

	return &model.CreateAccountResponse{
		Account: convertAccountDBToAPI(a),
	}, nil
}

func convertAccountDBToAPI(account db.Account) model.Account {
	return model.Account{
		ID:        account.ID,
		UserID:    account.UserID,
		Name:      account.Name,
		Bank:      account.Bank,
		Balance:   account.Balance,
		CreatedAt: account.CreatedAt,
	}
}
