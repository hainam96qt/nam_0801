package transaction

import (
	"context"
	"errors"
	"github.com/DATA-DOG/go-sqlmock"
	"nam_0801/internal/model"
	db "nam_0801/internal/repo/dbmodel"
	"testing"
	"time"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"

	mock_transaction "nam_0801/internal/service/transaction/mock"
)

func TestService_CreateTransaction(t *testing.T) {
	t.Run("happy case deposit", func(t *testing.T) {
		ctx := context.Background()
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		transactionRepoMock := mock_transaction.NewMockTransactionRepository(ctrl)
		accountRepoMock := mock_transaction.NewMockAccountRepository(ctrl)

		dbConn, mockDBConn, err := sqlmock.New()
		assert.Nil(t, err, "Error should be nil")
		testTransactionService := &Service{
			DatabaseConn:    dbConn,
			transactionRepo: transactionRepoMock,
			accountRepo:     accountRepoMock,
		}

		req := &model.CreateTransactionRequest{
			AccountID:       111,
			Amount:          100.0,
			TransactionType: db.TransactionTypeDeposit,
		}

		mockDBConn.ExpectBegin()
		mockDBConn.ExpectCommit()
		mockDBConn.ExpectClose()
		accountRepoMock.EXPECT().GetAccountForUpdate(ctx, gomock.Any(), db.GetAccountForUpdateParams{
			ID:     111,
			UserID: 123,
		}).Return(&db.Account{
			ID:        111,
			UserID:    123,
			Name:      "account-1",
			Bank:      "VCB",
			Balance:   2000.0,
			CreatedAt: time.Now(),
		}, nil).Times(1)

		newTrans := db.CreateTransactionParams{
			AccountID:       111,
			Amount:          100.0,
			TransactionType: db.TransactionTypeDeposit,
		}

		var timeCreateTransaction = time.Now().Add(3600)
		transactionRepoMock.EXPECT().CreateTransaction(ctx, gomock.Any(), newTrans).Return(db.Transaction{
			ID:              321,
			Amount:          100.0,
			AccountID:       1,
			TransactionType: "VCB",
			CreatedAt:       timeCreateTransaction,
		}, nil).Times(1)

		accountRepoMock.EXPECT().UpdateAccountBalance(ctx, gomock.Any(), db.UpdateAccountBalanceParams{
			Balance: 2100.0,
			ID:      111,
		}).Times(1)

		result, err := testTransactionService.CreateTransaction(ctx, 123, req)

		assert.Nil(t, err, "Error should be nil")
		assert.Equal(t, result.ID, int32(321))
		assert.Equal(t, result.Amount, 100.0)
		assert.Equal(t, result.AccountID, int32(111))
		assert.Equal(t, result.Bank, db.BankNameVCB)
		assert.Equal(t, result.CreatedAt, timeCreateTransaction)
	})

	t.Run("happy case withdraw", func(t *testing.T) {
		ctx := context.Background()
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		transactionRepoMock := mock_transaction.NewMockTransactionRepository(ctrl)
		accountRepoMock := mock_transaction.NewMockAccountRepository(ctrl)

		dbConn, mockDBConn, err := sqlmock.New()
		assert.Nil(t, err, "Error should be nil")
		testTransactionService := &Service{
			DatabaseConn:    dbConn,
			transactionRepo: transactionRepoMock,
			accountRepo:     accountRepoMock,
		}

		req := &model.CreateTransactionRequest{
			AccountID:       111,
			Amount:          100.0,
			TransactionType: db.TransactionTypeWithdraw,
		}

		mockDBConn.ExpectBegin()
		mockDBConn.ExpectCommit()
		mockDBConn.ExpectClose()
		accountRepoMock.EXPECT().GetAccountForUpdate(ctx, gomock.Any(), db.GetAccountForUpdateParams{
			ID:     111,
			UserID: 123,
		}).Return(&db.Account{
			ID:        111,
			UserID:    123,
			Name:      "account-1",
			Bank:      "VCB",
			Balance:   2000.0,
			CreatedAt: time.Now(),
		}, nil).Times(1)

		newTrans := db.CreateTransactionParams{
			AccountID:       111,
			Amount:          100.0,
			TransactionType: db.TransactionTypeWithdraw,
		}

		var timeCreateTransaction = time.Now().Add(3600)
		transactionRepoMock.EXPECT().CreateTransaction(ctx, gomock.Any(), newTrans).Return(db.Transaction{
			ID:              321,
			Amount:          100.0,
			AccountID:       1,
			TransactionType: "VCB",
			CreatedAt:       timeCreateTransaction,
		}, nil).Times(1)

		accountRepoMock.EXPECT().UpdateAccountBalance(ctx, gomock.Any(), db.UpdateAccountBalanceParams{
			Balance: 1900.0,
			ID:      111,
		}).Times(1)

		result, err := testTransactionService.CreateTransaction(ctx, 123, req)

		assert.Nil(t, err, "Error should be nil")
		assert.Equal(t, result.ID, int32(321))
		assert.Equal(t, result.Amount, 100.0)
		assert.Equal(t, result.AccountID, int32(111))
		assert.Equal(t, result.Bank, db.BankNameVCB)
		assert.Equal(t, result.CreatedAt, timeCreateTransaction)
	})

	t.Run("fail case amount less then 0", func(t *testing.T) {
		ctx := context.Background()
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		transactionRepoMock := mock_transaction.NewMockTransactionRepository(ctrl)
		accountRepoMock := mock_transaction.NewMockAccountRepository(ctrl)

		dbConn, mockDBConn, err := sqlmock.New()
		assert.Nil(t, err, "Error should be nil")
		testTransactionService := &Service{
			DatabaseConn:    dbConn,
			transactionRepo: transactionRepoMock,
			accountRepo:     accountRepoMock,
		}

		req := &model.CreateTransactionRequest{
			AccountID:       111,
			Amount:          -100.0,
			TransactionType: db.TransactionTypeWithdraw,
		}

		mockDBConn.ExpectBegin()
		mockDBConn.ExpectCommit()
		mockDBConn.ExpectRollback()

		_, err = testTransactionService.CreateTransaction(ctx, 123, req)
		assert.NotNil(t, err, "Error should be not nil")
	})

	t.Run("fail case balance not enough", func(t *testing.T) {
		ctx := context.Background()
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		transactionRepoMock := mock_transaction.NewMockTransactionRepository(ctrl)
		accountRepoMock := mock_transaction.NewMockAccountRepository(ctrl)

		dbConn, mockDBConn, err := sqlmock.New()
		assert.Nil(t, err, "Error should be nil")
		testTransactionService := &Service{
			DatabaseConn:    dbConn,
			transactionRepo: transactionRepoMock,
			accountRepo:     accountRepoMock,
		}

		req := &model.CreateTransactionRequest{
			AccountID:       111,
			Amount:          100.0,
			TransactionType: db.TransactionTypeWithdraw,
		}

		mockDBConn.ExpectBegin()
		mockDBConn.ExpectCommit()
		mockDBConn.ExpectRollback()
		accountRepoMock.EXPECT().GetAccountForUpdate(ctx, gomock.Any(), db.GetAccountForUpdateParams{
			ID:     111,
			UserID: 123,
		}).Return(&db.Account{
			ID:        111,
			UserID:    123,
			Name:      "account-1",
			Bank:      "VCB",
			Balance:   10.0,
			CreatedAt: time.Now(),
		}, nil).Times(1)

		_, err = testTransactionService.CreateTransaction(ctx, 123, req)

		assert.NotNil(t, err, "Error should be not nil")
	})

	t.Run("exception error roll back record already insert", func(t *testing.T) {
		ctx := context.Background()
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		transactionRepoMock := mock_transaction.NewMockTransactionRepository(ctrl)
		accountRepoMock := mock_transaction.NewMockAccountRepository(ctrl)

		dbConn, mockDBConn, err := sqlmock.New()
		assert.Nil(t, err, "Error should be nil")
		testTransactionService := &Service{
			DatabaseConn:    dbConn,
			transactionRepo: transactionRepoMock,
			accountRepo:     accountRepoMock,
		}

		req := &model.CreateTransactionRequest{
			AccountID:       111,
			Amount:          100.0,
			TransactionType: db.TransactionTypeWithdraw,
		}

		mockDBConn.ExpectBegin()
		mockDBConn.ExpectCommit()
		mockDBConn.ExpectRollback()
		accountRepoMock.EXPECT().GetAccountForUpdate(ctx, gomock.Any(), db.GetAccountForUpdateParams{
			ID:     111,
			UserID: 123,
		}).Return(&db.Account{
			ID:        111,
			UserID:    123,
			Name:      "account-1",
			Bank:      "VCB",
			Balance:   2000.0,
			CreatedAt: time.Now(),
		}, nil).Times(1)

		newTrans := db.CreateTransactionParams{
			AccountID:       111,
			Amount:          100.0,
			TransactionType: db.TransactionTypeWithdraw,
		}

		var timeCreateTransaction = time.Now().Add(3600)
		transactionRepoMock.EXPECT().CreateTransaction(ctx, gomock.Any(), newTrans).Return(db.Transaction{
			ID:              321,
			Amount:          100.0,
			AccountID:       1,
			TransactionType: "VCB",
			CreatedAt:       timeCreateTransaction,
		}, nil).Times(1)

		accountRepoMock.EXPECT().UpdateAccountBalance(ctx, gomock.Any(), db.UpdateAccountBalanceParams{
			Balance: 1900.0,
			ID:      111,
		}).Return(errors.New("this is fail update")).Times(1)

		_, err = testTransactionService.CreateTransaction(ctx, 123, req)

		assert.NotNil(t, err, "Error should be not nil")
	})

}
