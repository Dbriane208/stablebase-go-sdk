package qrcode

import (
	"encoding/json"
	"github.com/skip2/go-qrcode"
)

// PaymentQRData
type PaymentQRData struct {
	MerchantID   string `json:"merchant_id"`
	Amount       string `json:"amount"`
	TokenAddress string `json:"token_address"`
	OrderID      string `json:"order_id"`
	ChainID      string `json:"chain_id"`
}

// Generate creates a QR code image from payment data
// Returns the QR code as PNG
func Generate(data PaymentQRData, size int) ([]byte, error) {
	// Convert struct to JSON string
	jsonData, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}

	// Generate QR code using the library
	qr, err := qrcode.New(string(jsonData), qrcode.Medium)
	if err != nil {
		return nil, err
	}

	// return as PNG bytes
	return qr.PNG(size)
}

// Parse decodes a QR code string into PaymentData
func Parse(qrString string) (*PaymentQRData, error) {
	var data PaymentQRData

	err := json.Unmarshal([]byte(qrString), &data)
	if err != nil {
		return nil, err
	}

	return &data, nil
}
