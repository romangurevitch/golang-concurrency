package account

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/romangurevitch/golang-concurrency/internal/pkg/logger"
	"github.com/romangurevitch/golang-concurrency/pkg/api/error"
	"net/http"
)

func (a *accountServer) handleGetAccounts(c *gin.Context) {
	ctx := c.Request.Context()
	email, exists := c.GetQuery("email")
	logger.WithContext(ctx).WithField("email", email).Info()
	if !exists {
		c.JSON(error.NewAPIError(ctx, http.StatusBadRequest, "missing required query param", errors.New("email is a required param")))
		return
	}

	c.JSON(http.StatusOK, "")
}

func (a *accountServer) handleHealthCheck(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"status": "OK",
	})
}
