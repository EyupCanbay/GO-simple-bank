package api

import (
	"database/sql"
	"fmt"
	"net/http"
	db "simple_bank/db/sqlc"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgtype"
)

func convertToPgInt8(value int64) pgtype.Int8 {
	return pgtype.Int8{
		Int64: value,
		Valid: true,
	}
}

type transferRequest struct {
	FromAccountID int64  `json:"from_account_id" binding:"required,min=1"`
	ToAccountID   int64  `json:"to_account_id" binding:"required,min=1"`
	Amount        int64  `json:"amount" binding:"required,gt=0"`
	Currency      string `json:"currency" binding:"required,currency"`
}

func (server *Server) createTransfer(ctx *gin.Context) {
	var req transferRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	if !server.validAccount(ctx, req.FromAccountID, req.Currency, req.Amount) {
		ctx.JSON(http.StatusBadRequest, errorResponse(fmt.Errorf("dont match currency")))
		return
	}
	if !server.validAccount(ctx, req.ToAccountID, req.Currency, req.Amount) {
		ctx.JSON(http.StatusBadRequest, errorResponse(fmt.Errorf("dont match currency")))
		return
	}
	if !server.IsEnughBalance(ctx, req.FromAccountID, req.Amount) {
		ctx.JSON(http.StatusBadRequest, errorResponse(fmt.Errorf("must be enugh balance")))
		return
	}

	arg := db.TransferTxParams{
		FromAccountID: convertToPgInt8(req.FromAccountID),
		ToAccountID:   convertToPgInt8(req.ToAccountID),
		Amount:        req.Amount,
	}

	result, err := server.store.TransferTx(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, result)
}

func (server *Server) validAccount(ctx *gin.Context, accountId int64, currency string, amount int64) bool {
	account, err := server.store.GetAccount(ctx, accountId)
	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, errorResponse(err))
			return false
		}

		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return false
	}

	if account.Currency != currency {
		err := fmt.Errorf("account[%d]courrrency mismatch: %s vs %s", account.ID, account.Currency, currency)
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return false
	}

	return true
}

func (server *Server) IsEnughBalance(ctx *gin.Context, fromAccountId int64, amount int64) bool {
	from_account, err := server.store.GetAccount(ctx, fromAccountId)
	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, errorResponse(err))
			return false
		}

		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return false
	}

	if from_account.Balance < amount {
		return false
	}

	return true

}
