package user

import (
	"context"
	"nam_0801/internal/model"
	db "nam_0801/internal/repo/dbmodel"
	"nam_0801/pkg/util/password"
)

func (s *Service) CreateUser(ctx context.Context, req *model.CreateUserRequest) (*model.CreateUserResponse, error) {
	hashPassword, err := password.HashPassword(req.Password)
	if err != nil {
		return nil, err
	}

	newUser := db.CreateUserParams{
		Name:     req.Name,
		Email:    req.Email,
		Password: hashPassword,
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
