package main

import (
	"go.uber.org/zap"

	"github.com/sp-yduck/dnsity/api"
	"github.com/sp-yduck/dnsity/dns"
	"github.com/sp-yduck/dnsity/utils"
)

var logger *zap.SugaredLogger

func init() {
	logger = utils.NewLogger("main")
}

func main() {

	logger.Info("main function")

	// dns server
	if err := dns.Run(); err != nil {
		logger.Errorf("failed to start dns server : %v", err)
	}

	// api server
	if err := api.Run(); err != nil {
		logger.Errorf("failed to start api server : %v", err)
	}
}
