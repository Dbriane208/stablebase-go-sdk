package utils

import (
	"math/big"
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/stretchr/testify/assert"
)

func TestValidateRegisterMerchantInputs(t *testing.T) {
	tests := []struct {
		name          string
		payoutWallet  common.Address
		metadataURI   string
		expectValid   bool
		errorContains string
	}{
		{
			name:         "valid inputs",
			payoutWallet: common.HexToAddress("0x742d35Cc6634C0532925a3b844Bc454e4438f44e"),
			metadataURI:  "https://merchant.data/metadata.json",
			expectValid:  true,
		},
		{
			name:          "zero address",
			payoutWallet:  common.Address{},
			metadataURI:   "https://merchant.data/metadata.json",
			expectValid:   false,
			errorContains: "address",
		},
		{
			name:          "empty metadata URI",
			payoutWallet:  common.HexToAddress("0x742d35Cc6634C0532925a3b844Bc454e4438f44e"),
			metadataURI:   "",
			expectValid:   false,
			errorContains: "metadata",
		},
		{
			name:          "whitespace only metadata URI",
			payoutWallet:  common.HexToAddress("0x742d35Cc6634C0532925a3b844Bc454e4438f44e"),
			metadataURI:   "   ",
			expectValid:   false,
			errorContains: "metadata",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := ValidateRegisterMerchantInputs(tt.payoutWallet, tt.metadataURI)

			if tt.expectValid {
				assert.NoError(t, err)
			} else {
				assert.Error(t, err)
				if tt.errorContains != "" {
					assert.Contains(t, err.Error(), tt.errorContains)
				}
			}
		})
	}
}

func TestValidateUpdateMerchantInputs(t *testing.T) {
	validMerchantID := [32]byte{1, 2, 3, 4, 5}
	emptyMerchantID := [32]byte{}

	tests := []struct {
		name          string
		merchantId    [32]byte
		payoutWallet  common.Address
		metadataURI   string
		expectValid   bool
		errorContains string
	}{
		{
			name:         "valid inputs",
			merchantId:   validMerchantID,
			payoutWallet: common.HexToAddress("0x742d35Cc6634C0532925a3b844Bc454e4438f44e"),
			metadataURI:  "https://merchant.data/metadata.json",
			expectValid:  true,
		},
		{
			name:          "empty merchant ID",
			merchantId:    emptyMerchantID,
			payoutWallet:  common.HexToAddress("0x742d35Cc6634C0532925a3b844Bc454e4438f44e"),
			metadataURI:   "https://merchant.data/metadata.json",
			expectValid:   false,
			errorContains: "merchant",
		},
		{
			name:          "zero address",
			merchantId:    validMerchantID,
			payoutWallet:  common.Address{},
			metadataURI:   "https://merchant.data/metadata.json",
			expectValid:   false,
			errorContains: "address",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := ValidateUpdateMerchantInputs(tt.merchantId, tt.payoutWallet, tt.metadataURI)

			if tt.expectValid {
				assert.NoError(t, err)
			} else {
				assert.Error(t, err)
				if tt.errorContains != "" {
					assert.Contains(t, err.Error(), tt.errorContains)
				}
			}
		})
	}
}

func TestValidateCreateOrderInputs(t *testing.T) {
	validMerchantId := [32]byte{1, 2, 3, 4, 5}
	emptyMerchantId := [32]byte{}

	tests := []struct {
		name         string
		merchantId   [32]byte
		tokenAddress common.Address
		amount       *big.Int
		metadataURI  string
		expectValid  bool
	}{
		{
			name:         "valid inputs",
			merchantId:   validMerchantId,
			tokenAddress: common.HexToAddress("0xA0b86991c6218b36c1d19D4a2e9Eb0cE3606eB48"),
			amount:       big.NewInt(1_000_000),
			metadataURI:  "https://order.data/metadata.json",
			expectValid:  true,
		},
		{
			name:         "empty merchant Id",
			merchantId:   emptyMerchantId,
			tokenAddress: common.HexToAddress("0xA0b86991c6218b36c1d19D4a2e9Eb0cE3606eB48"),
			amount:       big.NewInt(1_000_000),
			metadataURI:  "https://order.data/metadata.json",
			expectValid:  false,
		},
		{
			name:         "zero token address",
			merchantId:   validMerchantId,
			tokenAddress: common.Address{},
			amount:       big.NewInt(1_000_000),
			metadataURI:  "https://order.data/metadata.json",
			expectValid:  false,
		},
		{
			name:         "zero amount",
			merchantId:   validMerchantId,
			tokenAddress: common.HexToAddress("0xA0b86991c6218b36c1d19D4a2e9Eb0cE3606eB48"),
			amount:       big.NewInt(0),
			metadataURI:  "https://order.data/metadata.json",
			expectValid:  false,
		},
		{
			name:         "empty metadata uri",
			merchantId:   validMerchantId,
			tokenAddress: common.HexToAddress("0xA0b86991c6218b36c1d19D4a2e9Eb0cE3606eB48"),
			amount:       big.NewInt(1_000_000),
			metadataURI:  "",
			expectValid:  false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := ValidateCreateOrderInputs(tt.merchantId, tt.tokenAddress, tt.amount, tt.metadataURI)

			if tt.expectValid {
				assert.NoError(t, err)
			} else {
				assert.Error(t, err)
			}
		})
	}
}

func TestValidateOrderId(t *testing.T) {
	validOrderId := [32]byte{1, 2, 3, 4, 5}
	emptyOrderId := [32]byte{}

	tests := []struct {
		name        string
		orderId     [32]byte
		expectValid bool
	}{
		{
			name:        "valid order ID",
			orderId:     validOrderId,
			expectValid: true,
		},
		{
			name:        "empty order ID",
			orderId:     emptyOrderId,
			expectValid: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := ValidateOrderId(tt.orderId)

			if tt.expectValid {
				assert.NoError(t, err)
			} else {
				assert.Error(t, err)
			}
		})
	}
}

func TestValidateEmergencyWithdrawInputs(t *testing.T) {
	tests := []struct {
		name            string
		tokenAddress    common.Address
		receiverAddress common.Address
		amount          *big.Int
		expectValid     bool
	}{
		{
			name:            "valid inputs",
			tokenAddress:    common.HexToAddress("0x123"),
			receiverAddress: common.HexToAddress("0x456"),
			amount:          big.NewInt(1_000_000),
			expectValid:     true,
		},
		{
			name:            "empty token address",
			tokenAddress:    common.Address{},
			receiverAddress: common.HexToAddress("0x456"),
			amount:          big.NewInt(1_000_000),
			expectValid:     false,
		},
		{
			name:            "empty receiver address",
			tokenAddress:    common.HexToAddress("0x123"),
			receiverAddress: common.Address{},
			amount:          big.NewInt(1_000_000),
			expectValid:     false,
		},
		{
			name:            "zero amount",
			tokenAddress:    common.HexToAddress("0x123"),
			receiverAddress: common.HexToAddress("0x456"),
			amount:          big.NewInt(0),
			expectValid:     false,
		},
		{
			name:            "negative amount",
			tokenAddress:    common.HexToAddress("0x123"),
			receiverAddress: common.HexToAddress("0x456"),
			amount:          big.NewInt(-100),
			expectValid:     false,
		},
		{
			name:            "nil amount",
			tokenAddress:    common.HexToAddress("0x123"),
			receiverAddress: common.HexToAddress("0x456"),
			amount:          nil,
			expectValid:     false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := ValidateEmergencyWithdrawInputs(tt.tokenAddress, tt.receiverAddress, tt.amount)

			if tt.expectValid {
				assert.NoError(t, err)
			} else {
				assert.Error(t, err)
			}
		})
	}
}

func TestValidateUpdateMerchantRegistryInput(t *testing.T) {
	tests := []struct {
		name                    string
		merchantRegistryAddress common.Address
		expectValid             bool
	}{
		{
			name:                    "valid input",
			merchantRegistryAddress: common.HexToAddress("0x123"),
			expectValid:             true,
		},
		{
			name:                    "invalid input",
			merchantRegistryAddress: common.Address{},
			expectValid:             false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := ValidateUpdateMerchantRegistryInput(tt.merchantRegistryAddress)

			if tt.expectValid {
				assert.NoError(t, err)
			} else {
				assert.Error(t, err)
			}
		})
	}
}

func TestValidateSetTokenSupportInput(t *testing.T) {
	tests := []struct {
		name         string
		tokenAddress common.Address
		status       *big.Int
		expectValid  bool
	}{
		{
			name:         "valid input disabled",
			tokenAddress: common.HexToAddress("0x123"),
			status:       big.NewInt(0),
			expectValid:  true,
		},
		{
			name:         "valid input enabled",
			tokenAddress: common.HexToAddress("0x123"),
			status:       big.NewInt(1),
			expectValid:  true,
		},
		{
			name:         "empty tokenAddress",
			tokenAddress: common.Address{},
			status:       big.NewInt(0),
			expectValid:  false,
		},
		{
			name:         "invalid status",
			tokenAddress: common.HexToAddress("0x123"),
			status:       big.NewInt(5),
			expectValid:  false,
		},
		{
			name:         "nil status",
			tokenAddress: common.HexToAddress("0x123"),
			status:       nil,
			expectValid:  false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := ValidateSetTokenSupportInput(tt.tokenAddress, tt.status)

			if tt.expectValid {
				assert.NoError(t, err)
			} else {
				assert.Error(t, err)
			}
		})
	}
}

func TestValidateAddress(t *testing.T) {
	tests := []struct {
		name         string
		address      common.Address
		field        string
		expectValid  bool
		errorMessage string
	}{
		{
			name:        "valid address",
			address:     common.HexToAddress("0x123"),
			field:       "testAddress",
			expectValid: true,
		},
		{
			name:         "zero address",
			address:      common.Address{},
			field:        "testAddress",
			expectValid:  false,
			errorMessage: "testAddress",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := ValidateAddress(tt.address, tt.field)

			if tt.expectValid {
				assert.NoError(t, err)
			} else {
				assert.Error(t, err)
				assert.Contains(t, err.Error(), tt.errorMessage)
			}
		})
	}
}

func TestValidateGetPlatformTokenBalanceInputs(t *testing.T) {
	tests := []struct {
		name                  string
		platformWalletAddress common.Address
		tokenAddress          common.Address
		expectValid           bool
	}{
		{
			name:                  "valid input",
			platformWalletAddress: common.HexToAddress("0x456"),
			tokenAddress:          common.HexToAddress("0x123"),
			expectValid:           true,
		},
		{
			name:                  "empty tokenAddress",
			platformWalletAddress: common.HexToAddress("0x789"),
			tokenAddress:          common.Address{},
			expectValid:           false,
		},
		{
			name:                  "empty platformWalletAddress",
			platformWalletAddress: common.Address{},
			tokenAddress:          common.HexToAddress("0x789"),
			expectValid:           false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := ValidateGetPlatformTokenBalanceInputs(tt.platformWalletAddress, tt.tokenAddress)

			if tt.expectValid {
				assert.NoError(t, err)
			} else {
				assert.Error(t, err)
			}
		})
	}
}

func TestValidateRefundOrderInputs(t *testing.T) {
	validOrderID := [32]byte{1, 2, 3, 4, 5}
	emptyOrderID := [32]byte{}

	tests := []struct {
		name          string
		orderId       [32]byte
		expectValid   bool
		errorContains string
	}{
		{
			name:        "valid order ID",
			orderId:     validOrderID,
			expectValid: true,
		},
		{
			name:          "empty order ID",
			orderId:       emptyOrderID,
			expectValid:   false,
			errorContains: "order",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := ValidateRefundOrderInputs(tt.orderId)

			if tt.expectValid {
				assert.NoError(t, err)
			} else {
				assert.Error(t, err)
				if tt.errorContains != "" {
					assert.Contains(t, err.Error(), tt.errorContains)
				}
			}
		})
	}
}

func TestCheckTransactionStatus(t *testing.T) {
	tests := []struct {
		name        string
		receipt     *types.Receipt
		expectError bool
	}{
		{
			name: "successful transaction",
			receipt: &types.Receipt{
				Status: types.ReceiptStatusSuccessful,
				TxHash: common.HexToHash("0x123"),
			},
			expectError: false,
		},
		{
			name: "failed transaction",
			receipt: &types.Receipt{
				Status: types.ReceiptStatusFailed,
				TxHash: common.HexToHash("0x456"),
			},
			expectError: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := CheckTransactionStatus(tt.receipt)

			if tt.expectError {
				assert.Error(t, err)
				assert.Contains(t, err.Error(), "transaction failed")
			} else {
				assert.NoError(t, err)
			}
		})
	}
}

func TestValidationError(t *testing.T) {
	tests := []struct {
		name     string
		field    string
		message  string
		expected string
	}{
		{
			name:     "standard validation error",
			field:    "amount",
			message:  "must be greater than zero",
			expected: "amount: must be greater than zero",
		},
		{
			name:     "address validation error",
			field:    "tokenAddress",
			message:  "cannot be zero address",
			expected: "tokenAddress: cannot be zero address",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := &ValidationError{
				Field:   tt.field,
				Message: tt.message,
			}

			assert.Equal(t, tt.expected, err.Error())
		})
	}
}

func TestTransactionError(t *testing.T) {
	tests := []struct {
		name     string
		txHash   string
		message  string
		expected string
	}{
		{
			name:     "transaction failed error",
			txHash:   "0x123abc",
			message:  "out of gas",
			expected: "transaction failed: 0x123abc - out of gas",
		},
		{
			name:     "contract revert error",
			txHash:   "0x456def",
			message:  "execution reverted",
			expected: "transaction failed: 0x456def - execution reverted",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := &TransactionError{
				TxHash:  tt.txHash,
				Message: tt.message,
			}

			assert.Equal(t, tt.expected, err.Error())
		})
	}
}
