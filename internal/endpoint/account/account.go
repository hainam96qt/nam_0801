package account

import (
	"context"
	"github.com/go-chi/chi/v5"
	"log"
	"nam_0801/internal/model"
	error2 "nam_0801/pkg/error"
	"nam_0801/pkg/midleware"
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
		CreateAccount(ctx context.Context, userID int32, req *model.CreateAccountRequest) (*model.CreateAccountResponse, error)
		ListAccount(ctx context.Context, userID int32) (*model.ListAccountsResponse, error)
		GetAccount(ctx context.Context, userID int32, accountID int32) (*model.GetAccountResponse, error)
	}
)

func InitAccountHandler(r *chi.Mux, accountSvc AccountService) {
	accountEndpoint := &Endpoint{
		accountSvc: accountSvc,
	}

	r.Route("/api/users/{user_id}/accounts", func(r chi.Router) {
		r.Use(midleware.Auth.ValidateRoleUser)

		r.Post("/", accountEndpoint.createAccount)
		r.Get("/", accountEndpoint.listAccounts)
		r.Get("/{account_id}", accountEndpoint.getAccounts)
	})
}

func (e Endpoint) createAccount(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	userID, err := strconv.Atoi(chi.URLParam(r, "user_id"))
	if err != nil {
		log.Printf("failed to get query 'page' for list transactions: %s \n", err)
		response.Error(w, error2.NewXError(err.Error(), http.StatusBadRequest))
		return
	}

	userIDCtx := ctx.Value("UserID").(int32)
	if userIDCtx != int32(userID) {
		err := error2.NewXError("user can not access info in other user", http.StatusUnauthorized)
		log.Printf("failed to get list transactions: %s \n", err)
		response.Error(w, error2.NewXError(err.Error(), http.StatusBadRequest))
		return
	}

	var req model.CreateAccountRequest
	if err := request.DecodeJSON(ctx, r.Body, &req); err != nil {
		log.Printf("read request body error: %s \n", err)
		response.Error(w, error2.NewXError(err.Error(), http.StatusBadRequest))
		return
	}

	res, err := e.accountSvc.CreateAccount(ctx, int32(userID), &req)
	if err != nil {
		log.Printf("failed to get list accounts: %s \n", err)
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

	userIDCtx := ctx.Value("UserID").(int32)
	if userIDCtx != int32(userID) {
		err := error2.NewXError("user can not access info in other user", http.StatusUnauthorized)
		log.Printf("failed to get query 'page' for get account: %s \n", err)
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
		log.Printf("failed to get account: %s \n", err)
		response.Error(w, err)
		return
	}

	response.JSON(w, http.StatusOK, res.Account)
}

func (e Endpoint) listAccounts(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	userID, err := strconv.Atoi(chi.URLParam(r, "user_id"))
	if err != nil {
		log.Printf("failed to get query 'page' for list transactions: %s \n", err)
		response.Error(w, error2.NewXError(err.Error(), http.StatusBadRequest))
		return
	}

	userIDCtx := ctx.Value("UserID").(int32)
	if userIDCtx != int32(userID) {
		err := error2.NewXError("user can not access info in other user", http.StatusUnauthorized)
		log.Printf("failed to get list accounts: %s \n", err)
		response.Error(w, error2.NewXError(err.Error(), http.StatusBadRequest))
		return
	}

	res, err := e.accountSvc.ListAccount(ctx, int32(userID))
	if err != nil {
		log.Printf("failed to get list accountss: %s \n", err)
		response.Error(w, err)
		return
	}

	response.JSON(w, http.StatusOK, res.Accounts)
}
