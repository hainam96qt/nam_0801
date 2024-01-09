package transaction

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
		transactionSvc TransactionService
	}

	TransactionService interface {
		CreateTransaction(ctx context.Context, userID int32, req *model.CreateTransactionRequest) (*model.CreateTransactionResponse, error)
		ListTransactions(ctx context.Context, userID int32, accountID int32) (*model.ListTransactionResponse, error)
	}
)

func InitTransactionHandler(r *chi.Mux, transactionSvc TransactionService) {
	transactionEndpoint := &Endpoint{
		transactionSvc: transactionSvc,
	}
	r.Route("/api/users/{user_id}", func(r chi.Router) {
		r.Get("/transactions", transactionEndpoint.ListTransactions)
		r.Post("/transactions", transactionEndpoint.createTransaction)

	})
}

func (e *Endpoint) ListTransactions(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	userID, err := strconv.Atoi(chi.URLParam(r, "user_id"))
	if err != nil {
		log.Printf("failed to get query 'page' for list transactions: %s \n", err)
		response.Error(w, error2.NewXError(err.Error(), http.StatusBadRequest))
		return
	}

	var accountID int
	accountIDStr := r.URL.Query().Get("accountID")
	if accountIDStr != "" {
		accountID, err = strconv.Atoi(accountIDStr)
		if err != nil {
			log.Printf("failed to get query 'account' for list transactions: %s \n", err)
			response.Error(w, error2.NewXError(err.Error(), http.StatusBadRequest))
			return
		}
	}

	res, err := e.transactionSvc.ListTransactions(ctx, int32(userID), int32(accountID))
	if err != nil {
		log.Printf("failed to get list wagers: %s \n", err)
		response.Error(w, err)
		return
	}

	response.JSON(w, http.StatusCreated, res.Transactions)
}

func (e *Endpoint) createTransaction(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	userID, err := strconv.Atoi(chi.URLParam(r, "user_id"))
	if err != nil {
		log.Printf("failed to get query 'page' for create transaction: %s \n", err)
		response.Error(w, error2.NewXError(err.Error(), http.StatusBadRequest))
		return
	}

	var req model.CreateTransactionRequest
	if err := request.DecodeJSON(ctx, r.Body, &req); err != nil {
		log.Printf("read request body error: %s \n", err)
		response.Error(w, error2.NewXError(err.Error(), http.StatusBadRequest))
		return
	}

	res, err := e.transactionSvc.CreateTransaction(ctx, int32(userID), &req)
	if err != nil {
		log.Printf("failed to get list wagers: %s \n", err)
		response.Error(w, err)
		return
	}

	response.JSON(w, http.StatusCreated, res)
}
