package main

import (
	"github.com/alifnh/bjb-auction-backend/internal/config"
	"github.com/alifnh/bjb-auction-backend/internal/httpserver"
	"github.com/alifnh/bjb-auction-backend/internal/pkg/logger"
)

func main() {
	cfg := config.InitConfig()

	logger.SetLogrusLogger(cfg)
	httpserver.StartGinHttpServer(cfg)
}
