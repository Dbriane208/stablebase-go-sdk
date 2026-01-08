package merchant

import (
	"fmt"
	"math/big"

	"github.com/Dbriane208/stablebase-go-sdk/qrcode"
	"github.com/ethereum/go-ethereum/common"
)

// GeneratePaymentQR creates a payment QR code for the merchant
func (m *Merchant) GeneratePaymentQR(merchantID [32]byte, amount *big.Int, tokenAddress common.Address, orderID [32]byte) ([]byte, error) {
	// Prepare the data
	data := qrcode.PaymentQRData{
		MerchantID:   fmt.Sprintf("%x", merchantID),
		Amount:       amount.String(),
		TokenAddress: tokenAddress.Hex(),
		OrderID:      fmt.Sprintf("%x", orderID),
		ChainID:      m.Client.ChainID.String(),
	}

	// Generate the QR code
	return qrcode.Generate(data, 256)
}
