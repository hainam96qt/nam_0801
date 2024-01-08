package account

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
		accountSvc AccountService
	}

	AccountService interface {
		CreateAccount(ctx context.Context, req *model.CreateAccountRequest) (*model.CreateAccountResponse, error)
	}
)

func InitAccountHandler(r *chi.Mux, accountSvc AccountService) {
	accountEndpoint := &Endpoint{
		accountSvc: accountSvc,
	}
	r.Route("/", func(r chi.Router) {
		r.Post("/accounts", accountEndpoint.createAccount)
	})
}

func (e Endpoint) createAccount(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var req model.CreateAccountRequest
	if err := request.DecodeJSON(ctx, r.Body, &req); err != nil {
		log.Printf("read request body error: %s \n", err)
		response.Error(w, error2.NewXError(err.Error(), http.StatusBadRequest))
		return
	}

	res, err := e.accountSvc.CreateAccount(ctx, &req)
	if err != nil {
		log.Printf("failed to get list wagers: %s \n", err)
		response.Error(w, err)
		return
	}

	response.JSON(w, http.StatusCreated, res)
}
