package account

import (
	"github.com/gin-gonic/gin"
	"github.com/romangurevitch/golang-concurrency/internal/server/rest"
	"github.com/romangurevitch/golang-concurrency/internal/server/rest/middleware"
)

func NewServer() rest.Server {
	r := gin.New()
	r.Use(middleware.JSONLogMiddleware())
	r.Use(gin.Recovery())

	aServer := &accountServer{
		engine: r,
	}
	aServer.createRoutes()
	return aServer
}

type accountServer struct {
	engine *gin.Engine
}

func (a *accountServer) Run(addr ...string) error {
	return a.engine.Run(addr...)
}

func (a *accountServer) createRoutes() {
	router := a.engine.Group("/api")
	{
		router.GET("/accounts", a.handleGetAccounts)
		router.GET("/health", a.handleHealthCheck)
	}
}
