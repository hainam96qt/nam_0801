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
)

type (
	Endpoint struct {
		transactionSvc TransactionService
	}

	TransactionService interface {
		CreateTransaction(ctx context.Context, req *model.CreateTransactionRequest) (*model.CreateTransactionResponse, error)
	}
)

func InitTransactionHandler(r *chi.Mux, transactionSvc TransactionService) {
	transactionEndpoint := &Endpoint{
		transactionSvc: transactionSvc,
	}
	r.Route("/", func(r chi.Router) {
		r.Post("/transactions", transactionEndpoint.createTransaction)
	})
}

func (e *Endpoint) createTransaction(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var req model.CreateTransactionRequest
	if err := request.DecodeJSON(ctx, r.Body, &req); err != nil {
		log.Printf("read request body error: %s \n", err)
		response.Error(w, error2.NewXError(err.Error(), http.StatusBadRequest))
		return
	}

	res, err := e.transactionSvc.CreateTransaction(ctx, &req)
	if err != nil {
		log.Printf("failed to get list wagers: %s \n", err)
		response.Error(w, err)
		return
	}

	response.JSON(w, http.StatusCreated, res)
}
