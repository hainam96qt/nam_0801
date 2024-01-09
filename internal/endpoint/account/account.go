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
	"strconv"
)

type (
	Endpoint struct {
		accountSvc AccountService
	}

	AccountService interface {
		CreateAccount(ctx context.Context, req *model.CreateAccountRequest) (*model.CreateAccountResponse, error)
		ListAccount(ctx context.Context, userID int32) (*model.ListAccountsResponse, error)
		GetAccount(ctx context.Context, userID int32, accountID int32) (*model.GetAccountResponse, error)
	}
)

func InitAccountHandler(r *chi.Mux, accountSvc AccountService) {
	accountEndpoint := &Endpoint{
		accountSvc: accountSvc,
	}
	r.Route("/users/{user_id}", func(r chi.Router) {
		r.Post("/accounts", accountEndpoint.createAccount)
		r.Get("/accounts", accountEndpoint.listAccounts)
		r.Get("/accounts/{account_id}", accountEndpoint.getAccounts)
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

func (e Endpoint) getAccounts(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	userID, err := strconv.Atoi(chi.URLParam(r, "user_id"))
	if err != nil {
		log.Printf("failed to get query 'page' for list transactions: %s \n", err)
		response.Error(w, error2.NewXError(err.Error(), http.StatusBadRequest))
		return
	}

	accountID, err := strconv.Atoi(chi.URLParam(r, "account_id"))
	if err != nil {
		log.Printf("failed to get query 'page' for get account: %s \n", err)
		response.Error(w, error2.NewXError(err.Error(), http.StatusBadRequest))
		return
	}

	res, err := e.accountSvc.GetAccount(ctx, int32(userID), int32(accountID))
	if err != nil {
		log.Printf("failed to get list wagers: %s \n", err)
		response.Error(w, err)
		return
	}

	response.JSON(w, http.StatusCreated, res.Account)
}

func (e Endpoint) listAccounts(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	userID, err := strconv.Atoi(chi.URLParam(r, "user_id"))
	if err != nil {
		log.Printf("failed to get query 'page' for list transactions: %s \n", err)
		response.Error(w, error2.NewXError(err.Error(), http.StatusBadRequest))
		return
	}

	res, err := e.accountSvc.ListAccount(ctx, int32(userID))
	if err != nil {
		log.Printf("failed to get list wagers: %s \n", err)
		response.Error(w, err)
		return
	}

	response.JSON(w, http.StatusCreated, res.Accounts)
}
