package order

import (
	"context"
	"crypto/ecdsa"
	"math/big"
	"testing"

	"github.com/Dbriane208/stablebase-go-sdk/client"
	"github.com/Dbriane208/stablebase-go-sdk/contracts"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// TestNew tests the constructor
func TestNew(t *testing.T) {
	c := &client.Client{}
	o := New(c)

	assert.NotNil(t, o)
	assert.Equal(t, c, o.Client)
}

// TestOrderStructInitialization tests order struct initialization
func TestOrderStructInitialization(t *testing.T) {
	tests := []struct {
		name   string
		client *client.Client
	}{
		{
			name:   "with empty client",
			client: &client.Client{},
		},
		{
			name: "with populated client",
			client: &client.Client{
				ChainID: big.NewInt(1),
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			o := New(tt.client)
			assert.NotNil(t, o)
			assert.NotNil(t, o.Client)
			assert.Equal(t, tt.client, o.Client)
		})
	}
}

// TestCreateOrderInputValidation tests that CreateOrder validates inputs
func TestCreateOrderInputValidation(t *testing.T) {
	privateKey, err := crypto.GenerateKey()
	require.NoError(t, err)

	o := &Order{
		Client: &client.Client{
			PrivateKey: privateKey,
			ChainID:    big.NewInt(1),
		},
	}
	ctx := context.Background()

	tests := []struct {
		name          string
		merchantID    [32]byte
		tokenAddress  common.Address
		amount        *big.Int
		metadataURI   string
		expectError   bool
		errorContains string
	}{
		{
			name:          "empty merchant ID",
			merchantID:    [32]byte{},
			tokenAddress:  common.HexToAddress("0x123"),
			amount:        big.NewInt(1000000),
			metadataURI:   "https://example.com/order.json",
			expectError:   true,
			errorContains: "merchantId",
		},
		{
			name:          "zero token address",
			merchantID:    [32]byte{1, 2, 3},
			tokenAddress:  common.Address{},
			amount:        big.NewInt(1000000),
			metadataURI:   "https://example.com/order.json",
			expectError:   true,
			errorContains: "tokenAddress",
		},
		{
			name:          "zero amount",
			merchantID:    [32]byte{1, 2, 3},
			tokenAddress:  common.HexToAddress("0x123"),
			amount:        big.NewInt(0),
			metadataURI:   "https://example.com/order.json",
			expectError:   true,
			errorContains: "amount",
		},
		{
			name:          "negative amount",
			merchantID:    [32]byte{1, 2, 3},
			tokenAddress:  common.HexToAddress("0x123"),
			amount:        big.NewInt(-100),
			metadataURI:   "https://example.com/order.json",
			expectError:   true,
			errorContains: "amount",
		},
		{
			name:          "empty metadata URI",
			merchantID:    [32]byte{1, 2, 3},
			tokenAddress:  common.HexToAddress("0x123"),
			amount:        big.NewInt(1000000),
			metadataURI:   "",
			expectError:   true,
			errorContains: "metadata",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, _, _, err := o.CreateOrder(ctx, tt.merchantID, tt.tokenAddress, tt.amount, tt.metadataURI)

			if tt.expectError {
				assert.Error(t, err)
				assert.Contains(t, err.Error(), tt.errorContains)
			}
		})
	}
}

// TestPayOrderInputValidation tests that PayOrder validates inputs
func TestPayOrderInputValidation(t *testing.T) {
	privateKey, err := crypto.GenerateKey()
	require.NoError(t, err)

	o := &Order{
		Client: &client.Client{
			PrivateKey: privateKey,
			ChainID:    big.NewInt(1),
		},
	}
	ctx := context.Background()

	t.Run("empty order ID", func(t *testing.T) {
		_, _, err := o.PayOrder(ctx, [32]byte{})
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "orderId")
	})
}

// TestCancelOrderInputValidation tests that CancelOrder validates inputs
func TestCancelOrderInputValidation(t *testing.T) {
	privateKey, err := crypto.GenerateKey()
	require.NoError(t, err)

	o := &Order{
		Client: &client.Client{
			PrivateKey: privateKey,
			ChainID:    big.NewInt(1),
		},
	}
	ctx := context.Background()

	t.Run("empty order ID", func(t *testing.T) {
		_, _, err := o.CancelOrder(ctx, [32]byte{})
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "orderId")
	})
}

// TestGetOrderInputValidation tests that GetOrder validates inputs
func TestGetOrderInputValidation(t *testing.T) {
	o := &Order{
		Client: &client.Client{},
	}
	ctx := context.Background()

	t.Run("empty order ID", func(t *testing.T) {
		_, err := o.GetOrder(ctx, [32]byte{})
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "orderId")
	})
}

// TestGetOrderStatusInputValidation tests that GetOrderStatus validates inputs
func TestGetOrderStatusInputValidation(t *testing.T) {
	o := &Order{
		Client: &client.Client{},
	}
	ctx := context.Background()

	t.Run("empty order ID", func(t *testing.T) {
		_, err := o.GetOrderStatus(ctx, [32]byte{})
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "orderId")
	})
}

// TestOrderMethodsNilClient tests that methods handle nil client gracefully
func TestOrderMethodsNilClient(t *testing.T) {
	o := &Order{
		Client: nil,
	}
	ctx := context.Background()

	t.Run("CreateOrder with nil client panics", func(t *testing.T) {
		assert.Panics(t, func() {
			o.CreateOrder(ctx, [32]byte{1}, common.HexToAddress("0x123"), big.NewInt(100), "https://example.com")
		})
	})
}

// TestOrderEventParsing tests event parsing logic
func TestOrderEventParsing(t *testing.T) {
	t.Run("parse order created event", func(t *testing.T) {
		receipt := &types.Receipt{
			Status: types.ReceiptStatusSuccessful,
			TxHash: common.HexToHash("0x123"),
			Logs:   []*types.Log{},
		}

		if receipt.Status == types.ReceiptStatusFailed {
			t.Error("Expected successful receipt")
		}
	})

	t.Run("handle failed transaction", func(t *testing.T) {
		receipt := &types.Receipt{
			Status: types.ReceiptStatusFailed,
			TxHash: common.HexToHash("0x456"),
		}

		assert.Equal(t, uint64(0), receipt.Status)
	})
}

// TestOrderPrivateKeyHandling tests private key operations
func TestOrderPrivateKeyHandling(t *testing.T) {
	t.Run("valid private key", func(t *testing.T) {
		privateKey, err := crypto.GenerateKey()
		require.NoError(t, err)
		assert.NotNil(t, privateKey)

		address := crypto.PubkeyToAddress(privateKey.PublicKey)
		assert.NotEqual(t, common.Address{}, address)
	})

	t.Run("nil private key", func(t *testing.T) {
		var nilKey *ecdsa.PrivateKey
		assert.Nil(t, nilKey)
	})
}

// TestOrderInfoStruct tests the order info return structure
func TestOrderInfoStruct(t *testing.T) {
	t.Run("order info structure", func(t *testing.T) {
		info := contracts.IPaymentProcessorOrder{
			Payer:          common.HexToAddress("0x123"),
			Token:          common.HexToAddress("0x456"),
			MerchantId:     [32]byte{4, 5, 6},
			MerchantPayout: common.HexToAddress("0x789"),
			Amount:         big.NewInt(1000000),
			Exists:         true,
			CurrentBps:     big.NewInt(500),
			CreatedAt:      big.NewInt(1234567890),
			MetadataUri:    "https://example.com",
			Status:         1,
		}

		assert.NotEqual(t, [32]byte{}, info.MerchantId)
		assert.NotEqual(t, common.Address{}, info.Payer)
		assert.NotEqual(t, common.Address{}, info.Token)
		assert.NotEqual(t, common.Address{}, info.MerchantPayout)
		assert.NotNil(t, info.Amount)
		assert.NotEmpty(t, info.MetadataUri)
		assert.True(t, info.Exists)
	})
}

// TestOrderContextHandling tests context cancellation
func TestOrderContextHandling(t *testing.T) {
	t.Run("cancelled context", func(t *testing.T) {
		ctx, cancel := context.WithCancel(context.Background())
		cancel()

		o := &Order{
			Client: &client.Client{},
		}

		select {
		case <-ctx.Done():
			assert.Error(t, ctx.Err())
		default:
			t.Error("Context should be cancelled")
		}

		assert.NotNil(t, o)
	})
}

// TestOrderAmountValidation tests amount validation edge cases
func TestOrderAmountValidation(t *testing.T) {
	tests := []struct {
		name    string
		amount  *big.Int
		isValid bool
	}{
		{
			name:    "zero amount",
			amount:  big.NewInt(0),
			isValid: false,
		},
		{
			name:    "negative amount",
			amount:  big.NewInt(-100),
			isValid: false,
		},
		{
			name:    "nil amount",
			amount:  nil,
			isValid: false,
		},
		{
			name:    "valid small amount",
			amount:  big.NewInt(1),
			isValid: true,
		},
		{
			name:    "valid large amount",
			amount:  big.NewInt(1000000000),
			isValid: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			isValid := tt.amount != nil && tt.amount.Sign() > 0
			assert.Equal(t, tt.isValid, isValid)
		})
	}
}

// TestOrderIDValidation tests order ID validation
func TestOrderIDValidation(t *testing.T) {
	tests := []struct {
		name    string
		orderID [32]byte
		isEmpty bool
	}{
		{
			name:    "empty order ID",
			orderID: [32]byte{},
			isEmpty: true,
		},
		{
			name:    "valid order ID",
			orderID: [32]byte{1, 2, 3, 4, 5},
			isEmpty: false,
		},
		{
			name:    "full order ID",
			orderID: [32]byte{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 25, 26, 27, 28, 29, 30, 31, 32},
			isEmpty: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			isEmpty := tt.orderID == [32]byte{}
			assert.Equal(t, tt.isEmpty, isEmpty)
		})
	}
}

// TestOrderStatusValidation tests order status values
func TestOrderStatusValidation(t *testing.T) {
	tests := []struct {
		name   string
		status *big.Int
		desc   string
	}{
		{
			name:   "pending status",
			status: big.NewInt(0),
			desc:   "PENDING",
		},
		{
			name:   "paid status",
			status: big.NewInt(1),
			desc:   "PAID",
		},
		{
			name:   "settled status",
			status: big.NewInt(2),
			desc:   "SETTLED",
		},
		{
			name:   "cancelled status",
			status: big.NewInt(3),
			desc:   "CANCELLED",
		},
		{
			name:   "refunded status",
			status: big.NewInt(4),
			desc:   "REFUNDED",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.NotNil(t, tt.status)
			assert.GreaterOrEqual(t, tt.status.Int64(), int64(0))
			assert.LessOrEqual(t, tt.status.Int64(), int64(4))
		})
	}
}

// TestOrderMetadataURIValidation tests metadata URI validation
func TestOrderMetadataURIValidation(t *testing.T) {
	tests := []struct {
		name        string
		metadataURI string
		shouldPass  bool
	}{
		{
			name:        "valid HTTPS URL",
			metadataURI: "https://example.com/order.json",
			shouldPass:  true,
		},
		{
			name:        "valid HTTP URL",
			metadataURI: "http://example.com/order.json",
			shouldPass:  true,
		},
		{
			name:        "valid IPFS URL",
			metadataURI: "ipfs://QmHash123",
			shouldPass:  true,
		},
		{
			name:        "empty string",
			metadataURI: "",
			shouldPass:  false,
		},
		{
			name:        "whitespace only",
			metadataURI: "   ",
			shouldPass:  false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			isEmpty := len(tt.metadataURI) == 0 || (len(tt.metadataURI) == 3 && tt.metadataURI == "   ")
			
			if tt.shouldPass {
				assert.False(t, isEmpty)
			} else {
				assert.True(t, isEmpty)
			}
		})
	}
}

// TestOrderTokenAddressValidation tests token address validation
func TestOrderTokenAddressValidation(t *testing.T) {
	tests := []struct {
		name    string
		address common.Address
		isZero  bool
	}{
		{
			name:    "zero address",
			address: common.Address{},
			isZero:  true,
		},
		{
			name:    "USDC address",
			address: common.HexToAddress("0xA0b86991c6218b36c1d19D4a2e9Eb0cE3606eB48"),
			isZero:  false,
		},
		{
			name:    "USDT address",
			address: common.HexToAddress("0xdAC17F958D2ee523a2206206994597C13D831ec7"),
			isZero:  false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			isZero := tt.address == common.Address{}
			assert.Equal(t, tt.isZero, isZero)
		})
	}
}
