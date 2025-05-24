package api

import (
	db "simple_bank/db/sqlc"

	"github.com/gin-gonic/gin"
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

	router.POST("/accounts", server.createAcount)
	router.GET("/accounts/:account_id", server.getByIdAccount)
	router.GET("/accounts", server.listAccounts)
	router.DELETE("/accounts/:account_id", server.deleteAccount)
	router.PUT("/accounts/:account_id", server.updateAccount)

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
