package user

import (
	"context"
	"nam_0801/internal/model"
	db "nam_0801/internal/repo/dbmodel"
)

func (s *Service) CreateUser(ctx context.Context, req *model.CreateUserRequest) (*model.CreateUserResponse, error) {
	newUser := db.CreateUserParams{
		Name:     req.Name,
		Email:    req.Email,
		Password: req.Password, // TODO hash
	}
	u, err := s.userRepo.CreateUser(ctx, newUser)
	if err != nil {
		return nil, err
	}

	return &model.CreateUserResponse{
		ID:    u.ID,
		Token: "123456", // TODO
	}, nil
}
