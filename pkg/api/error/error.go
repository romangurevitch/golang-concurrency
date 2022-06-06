package error

import (
	"context"
	"errors"
	"fmt"
	"github.com/romangurevitch/golang-concurrency/internal/pkg/logger"
)

var (
	ErrInternalServerError = errors.New("internal server error")
)

func NewAPIError(ctx context.Context, status int, msg string, err error) (int, *APIError) {
	logger.WithContext(ctx).WithError(err).WithField("status", status).Error(msg)
	return status, &APIError{
		Message: msg,
	}
}

type APIError struct {
	Message string `json:"message"`
}

func (e APIError) Error() string {
	return fmt.Sprintf("api error: %v", e.Message)
}
