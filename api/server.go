package api

import (
	"fmt"
	db "simple_bank/db/sqlc"
	"simple_bank/token"
	"simple_bank/utils"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
)

// server serves HTTP request for our banking service
type Server struct {
	config     utils.Config
	store      db.Store
	tokenMaker token.Maker
	router     *gin.Engine
}

// NewServer creates a new HTTP server and setup routing
func NewServer(config utils.Config, store db.Store) (*Server, error) {
	tokenMaker, err := token.NewPasetoMaker(config.TokenSymetricKey)
	if err != nil {
		return nil, fmt.Errorf("Can not create token")
	}
	server := &Server{
		config:     config,
		store:      store,
		tokenMaker: tokenMaker,
	}

	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		v.RegisterValidation("currency", validCurrency)
	}

	server.setUpRouter()
	return server, nil
}

func (server *Server) setUpRouter() {

	router := gin.Default()

	router.POST("/users", server.createUser)
	router.POST("/users/login", server.loginUser)

	router.POST("/accounts", server.createAcount)
	router.GET("/accounts/:account_id", server.getByIdAccount)
	router.GET("/accounts", server.listAccounts)
	router.DELETE("/accounts/:account_id", server.deleteAccount)
	router.PUT("/accounts/:account_id", server.updateAccount)

	router.POST("/transfers", server.createTransfer)

	server.router = router

}

// start run the HTTP server on a specific adddress
func (server *Server) Start(address string) error {
	return server.router.Run(address)
}

func errorResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}
