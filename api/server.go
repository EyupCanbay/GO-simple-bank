package api

import (
	db "simple_bank/db/sqlc"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
)

// server serves HTTP request for our banking service
type Server struct {
	store  db.Store
	router *gin.Engine
}

// NewServer creates a new HTTP server and setup routing
func NewServer(store db.Store) *Server {
	server := &Server{store: store}
	router := gin.Default()

	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		v.RegisterValidation("currency", validCurrency)
	}

	router.POST("/accounts", server.createAcount)
	router.GET("/accounts/:account_id", server.getByIdAccount)
	router.GET("/accounts", server.listAccounts)
	router.DELETE("/accounts/:account_id", server.deleteAccount)
	router.PUT("/accounts/:account_id", server.updateAccount)

	router.POST("/transfers", server.createTransfer)

	server.router = router
	return server
}

// start run the HTTP server on a specific adddress
func (server *Server) Start(address string) error {
	return server.router.Run(address)
}

func errorResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}
