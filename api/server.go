package api

import (
	db "github/heimaolst/simplebank/db/sqlc"

	"github.com/gin-gonic/gin"
)

type Server struct {
	store  db.Store
	router *gin.Engine
}

func NewServer(store db.Store) *Server {
	server := &Server{store: store}
	router := gin.Default()

	router.POST("/accounts", server.createAccount)
	router.GET("/accounts/:id", server.getAccount)
	router.GET("/accounts", server.listAccounts)
	router.DELETE("/accounts/:id", server.deleteAccount)
	router.PUT("/accounts", server.updateAccountInfo)
	server.router = router

	return server
}

func (server *Server) Start(address string) error {
	return server.router.Run(address)

}
func errResponse(error error) gin.H {
	return gin.H{"error": error.Error()}
}
