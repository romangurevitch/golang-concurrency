package main

import (
	"context"
	"fmt"
	"github.com/romangurevitch/golang-concurrency/internal/pkg/logger"
	"github.com/romangurevitch/golang-concurrency/internal/server/rest/account"
)

func init() {
	logger.Init(logger.Default())
}

func main() {
	ctx := context.Background()
	log := logger.WithContext(ctx)

	server := account.NewServer()
	if err := server.Run(fmt.Sprintf(":%s", "8080")); err != nil {
		log.WithError(err).Fatalf("%s server error", "concurrency example server")
	}
}
