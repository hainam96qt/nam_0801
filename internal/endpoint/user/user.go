package user

import (
	"context"
	"github.com/go-chi/chi/v5"
	"log"
	"nam_0801/internal/model"
	error2 "nam_0801/pkg/error"
	"nam_0801/pkg/util/request"
	"nam_0801/pkg/util/response"
	"net/http"
)

type (
	Endpoint struct {
		userSvc UserService
	}

	UserService interface {
		CreateUser(ctx context.Context, req *model.CreateUserRequest) (*model.CreateUserResponse, error)
	}
)

func InitUserHandler(r *chi.Mux, userSvc UserService) {
	userEndpoint := &Endpoint{
		userSvc: userSvc,
	}
	r.Route("/", func(r chi.Router) {
		r.Post("/users", userEndpoint.createUser)
	})
}

func (e *Endpoint) createUser(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var req model.CreateUserRequest
	if err := request.DecodeJSON(ctx, r.Body, &req); err != nil {
		log.Printf("read request body error: %s \n", err)
		response.Error(w, error2.NewXError(err.Error(), http.StatusBadRequest))
		return
	}

	res, err := e.userSvc.CreateUser(ctx, &req)
	if err != nil {
		log.Printf("failed to create user: %s \n", err)
		response.Error(w, err)
		return
	}

	response.JSON(w, http.StatusCreated, res)
}
