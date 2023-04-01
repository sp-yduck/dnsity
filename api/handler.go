package api

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/sp-yduck/dnsity/dns"
)

type DnsRequest struct {
	Type    string `json:"type"`
	Name    string `json:"name"`
	Content string `json:"content,omitempty"`
}

func getRecords(c *gin.Context) {
	records := dns.Records()
	c.JSON(http.StatusOK, gin.H{
		"records": records,
	})
}

func registerRecord(c *gin.Context) {
	var body DnsRequest
	if err := c.Bind(&body); err != nil {
		logger.Errorf("cannot bind request : %v", err)
	}
	logger.Debugf("request body : %v", body)
	switch body.Type {
	case "A":
		dns.RegisterRecord(body.Name, body.Content)
	default:
		logger.Errorf("not supported type")
	}
	c.JSON(http.StatusOK, body)
}
