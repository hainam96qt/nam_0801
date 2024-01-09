package account

import (
	"context"
	"nam_0801/internal/model"
	db "nam_0801/internal/repo/dbmodel"
)

func (s *Service) GetAccount(ctx context.Context, userID int32, accountID int32) (*model.GetAccountResponse, error) {
	params := db.GetAccountParams{
		ID:     accountID,
		UserID: userID,
	}
	account, err := s.accountRepo.GetAccount(ctx, params)
	if err != nil {
		return nil, err
	}

	return &model.GetAccountResponse{
		Account: convertAccountDBToAPI(*account),
	}, nil
}
