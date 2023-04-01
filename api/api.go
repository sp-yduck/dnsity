package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sp-yduck/dnsity/utils"
	"go.uber.org/zap"
)

var logger *zap.SugaredLogger

func init() {
	logger = utils.NewLogger("api")
}

func Run() error {
	logger.Infof("api module")

	r := gin.Default()

	r.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"health": "ok",
		})
	})

	api := r.Group("/api/v1")
	api.GET("/dns", getRecords)
	api.PUT("/dns", registerRecord)

	if err := r.Run("127.0.0.1:8080"); err != nil {
		return err
	}
	return nil
}
