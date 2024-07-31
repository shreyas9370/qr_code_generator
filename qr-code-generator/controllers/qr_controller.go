package controllers

import (
	"net/http"
	"qr-code-generator/services"

	"github.com/gin-gonic/gin"
)

type GenerateRequest struct {
	URL string `json:"url" binding:"required"`
}

func GenerateQRCode(c *gin.Context) {
	var req GenerateRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	qrCodeID, err := services.GenerateAndSaveQRCode(req.URL)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"id": qrCodeID})
}

func RedirectTOURL(c *gin.Context) {
	id := c.Param("id")
	url, err := services.GetURLByQRCodeID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "QR code not found"})
		return
	}
	c.Redirect(http.StatusPermanentRedirect, url)
}
