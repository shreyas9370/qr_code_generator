package services

import (
	"errors"
	"qr-code-generator/models"

	"github.com/skip2/go-qrcode"
)

func GenerateAndSaveQRCode(url string) (int64, error) {
	qrCode, err := qrcode.New(url, qrcode.Medium)
	if err != nil {
		return 0, err
	}
	qrCodeData, err := qrCode.PNG(-1)
	qrCodeModel := models.QRCode{
		URL:    url,
		QRCode: qrCodeData,
	}
	id, err := qrCodeModel.Save()
	if err != nil {
		return 0, err
	}
	return id, nil
}

func GetURLByQRCodeID(id string) (string, error) {
	qrCodeModel := models.QRCode{}
	if err := qrCodeModel.FindByID(id); err != nil {
		return "", err
	}
	if qrCodeModel.URL == "" {
		return "", errors.New("URL Not Found")
	}
	return qrCodeModel.URL, nil
}
