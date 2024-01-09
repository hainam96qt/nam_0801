package authentication

import (
	"context"
	"database/sql"
	"github.com/dgrijalva/jwt-go"
	"nam_0801/internal/model"
	db "nam_0801/internal/repo/dbmodel"
	configs "nam_0801/pkg/config"
	error2 "nam_0801/pkg/error"
	"nam_0801/pkg/util/password"
	"net/http"
	"time"
)

type (
	Service struct {
		cfg          configs.Token
		DatabaseConn *sql.DB

		userRepo UserRepository
	}

	UserRepository interface {
		GetUserByEmail(ctx context.Context, email string) (db.User, error)
	}
)

func NewAuthenticationService(DatabaseConn *sql.DB, cfg configs.Token, userRepo UserRepository) *Service {
	return &Service{
		DatabaseConn: DatabaseConn,
		cfg:          cfg,
		userRepo:     userRepo,
	}
}

func (s *Service) Login(ctx context.Context, req *model.LoginRequest) (*model.LoginResponse2, error) {
	user, err := s.userRepo.GetUserByEmail(ctx, req.Email)
	if err != nil {
		return nil, err
	}

	// validate password
	if !password.CheckPasswordHash(req.Password, user.Password) {
		return nil, error2.NewXError("password incorrect", http.StatusBadRequest)
	}

	token, err := s.generateJWT(user.ID)
	if err != nil {
		return nil, err
	}
	return &model.LoginResponse2{Token: token}, nil
}

func (s *Service) generateJWT(userID int32) (string, error) {
	// Create custom claims
	claims := &model.UserClaims{
		UserID: userID,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(s.cfg.TimeToExpired).Unix(),
			IssuedAt:  time.Now().Unix(),
		},
	}

	// Create a new JWT token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Sign the token with the secret key
	tokenString, err := token.SignedString([]byte(s.cfg.JwtSecretKey))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}
