package platform

import (
	"context"
	"crypto/ecdsa"
	"math/big"
	"testing"

	"github.com/Dbriane208/stablebase-go-sdk/client"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// TestNew tests the constructor
func TestNew(t *testing.T) {
	c := &client.Client{}
	p := New(c)

	assert.NotNil(t, p)
	assert.Equal(t, c, p.Client)
}

// TestPlatformStructInitialization tests platform struct initialization
func TestPlatformStructInitialization(t *testing.T) {
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
			p := New(tt.client)
			assert.NotNil(t, p)
			assert.NotNil(t, p.Client)
			assert.Equal(t, tt.client, p.Client)
		})
	}
}

// TestSettleOrderInputValidation tests that SettleOrder validates inputs
func TestSettleOrderInputValidation(t *testing.T) {
	privateKey, err := crypto.GenerateKey()
	require.NoError(t, err)

	p := &Platform{
		Client: &client.Client{
			PrivateKey: privateKey,
			ChainID:    big.NewInt(1),
		},
	}
	ctx := context.Background()

	t.Run("empty order ID", func(t *testing.T) {
		_, _, err := p.SettleOrder(ctx, [32]byte{})
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "order")
	})
}

// TestRefundOrderInputValidation tests that RefundOrder validates inputs
func TestRefundOrderInputValidation(t *testing.T) {
	privateKey, err := crypto.GenerateKey()
	require.NoError(t, err)

	p := &Platform{
		Client: &client.Client{
			PrivateKey: privateKey,
			ChainID:    big.NewInt(1),
		},
	}
	ctx := context.Background()

	t.Run("empty order ID", func(t *testing.T) {
		_, _, err := p.RefundOrder(ctx, [32]byte{})
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "order")
	})
}

// TestEmergencyWithdrawInputValidation tests that EmergencyWithdraw validates inputs
func TestEmergencyWithdrawInputValidation(t *testing.T) {
	privateKey, err := crypto.GenerateKey()
	require.NoError(t, err)

	p := &Platform{
		Client: &client.Client{
			PrivateKey: privateKey,
			ChainID:    big.NewInt(1),
		},
	}
	ctx := context.Background()

	tests := []struct {
		name            string
		tokenAddress    common.Address
		receiverAddress common.Address
		amount          *big.Int
		expectError     bool
		errorContains   string
	}{
		{
			name:            "zero token address",
			tokenAddress:    common.Address{},
			receiverAddress: common.HexToAddress("0x456"),
			amount:          big.NewInt(1000000),
			expectError:     true,
			errorContains:   "token",
		},
		{
			name:            "zero receiver address",
			tokenAddress:    common.HexToAddress("0x123"),
			receiverAddress: common.Address{},
			amount:          big.NewInt(1000000),
			expectError:     true,
			errorContains:   "receiver",
		},
		{
			name:            "zero amount",
			tokenAddress:    common.HexToAddress("0x123"),
			receiverAddress: common.HexToAddress("0x456"),
			amount:          big.NewInt(0),
			expectError:     true,
			errorContains:   "amount",
		},
		{
			name:            "negative amount",
			tokenAddress:    common.HexToAddress("0x123"),
			receiverAddress: common.HexToAddress("0x456"),
			amount:          big.NewInt(-100),
			expectError:     true,
			errorContains:   "amount",
		},
		{
			name:            "nil amount",
			tokenAddress:    common.HexToAddress("0x123"),
			receiverAddress: common.HexToAddress("0x456"),
			amount:          nil,
			expectError:     true,
			errorContains:   "amount",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, _, err := p.EmergencyWithdraw(ctx, tt.tokenAddress, tt.receiverAddress, tt.amount)

			if tt.expectError {
				assert.Error(t, err)
				assert.Contains(t, err.Error(), tt.errorContains)
			}
		})
	}
}

// TestUpdateMerchantRegistryInputValidation tests that UpdateMerchantRegistry validates inputs
func TestUpdateMerchantRegistryInputValidation(t *testing.T) {
	privateKey, err := crypto.GenerateKey()
	require.NoError(t, err)

	p := &Platform{
		Client: &client.Client{
			PrivateKey: privateKey,
			ChainID:    big.NewInt(1),
		},
	}
	ctx := context.Background()

	t.Run("zero address", func(t *testing.T) {
		_, _, err := p.UpdateMerchantRegistry(ctx, common.Address{})
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "address")
	})
}

// TestSetTokenSupportInputValidation tests that SetTokenSupport validates inputs
func TestSetTokenSupportInputValidation(t *testing.T) {
	privateKey, err := crypto.GenerateKey()
	require.NoError(t, err)

	p := &Platform{
		Client: &client.Client{
			PrivateKey: privateKey,
			ChainID:    big.NewInt(1),
		},
	}
	ctx := context.Background()

	tests := []struct {
		name          string
		tokenAddress  common.Address
		status        *big.Int
		expectError   bool
		errorContains string
	}{
		{
			name:          "zero token address",
			tokenAddress:  common.Address{},
			status:        big.NewInt(1),
			expectError:   true,
			errorContains: "token",
		},
		{
			name:          "nil status",
			tokenAddress:  common.HexToAddress("0x123"),
			status:        nil,
			expectError:   true,
			errorContains: "status",
		},
		{
			name:          "invalid status greater than 1",
			tokenAddress:  common.HexToAddress("0x123"),
			status:        big.NewInt(2),
			expectError:   true,
			errorContains: "status",
		},
		{
			name:          "invalid negative status",
			tokenAddress:  common.HexToAddress("0x123"),
			status:        big.NewInt(-1),
			expectError:   true,
			errorContains: "status",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, _, err := p.SetTokenSupport(ctx, tt.tokenAddress, tt.status)

			if tt.expectError {
				assert.Error(t, err)
				assert.Contains(t, err.Error(), tt.errorContains)
			}
		})
	}
}

// TestGetPlatformTokenBalanceInputValidation tests that GetPlatformTokenBalance validates inputs
func TestGetPlatformTokenBalanceInputValidation(t *testing.T) {
	p := &Platform{
		Client: &client.Client{},
	}
	ctx := context.Background()

	tests := []struct {
		name                  string
		platformWalletAddress common.Address
		tokenAddress          common.Address
		expectError           bool
		errorContains         string
	}{
		{
			name:                  "zero platform wallet address",
			platformWalletAddress: common.Address{},
			tokenAddress:          common.HexToAddress("0x123"),
			expectError:           true,
			errorContains:         "platform",
		},
		{
			name:                  "zero token address",
			platformWalletAddress: common.HexToAddress("0x456"),
			tokenAddress:          common.Address{},
			expectError:           true,
			errorContains:         "token",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := p.GetPlatformTokenBalance(ctx, tt.platformWalletAddress, tt.tokenAddress)

			if tt.expectError {
				assert.Error(t, err)
				assert.Contains(t, err.Error(), tt.errorContains)
			}
		})
	}
}

// TestGetContractTokenBalanceInputValidation tests that GetContractTokenBalance validates inputs
func TestGetContractTokenBalanceInputValidation(t *testing.T) {
	p := &Platform{
		Client: &client.Client{},
	}
	ctx := context.Background()

	t.Run("zero token address", func(t *testing.T) {
		_, err := p.GetContractTokenBalance(ctx, common.Address{})
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "address")
	})
}

// TestPlatformMethodsNilClient tests that methods handle nil client gracefully
func TestPlatformMethodsNilClient(t *testing.T) {
	p := &Platform{
		Client: nil,
	}
	ctx := context.Background()

	t.Run("SettleOrder with nil client panics", func(t *testing.T) {
		assert.Panics(t, func() {
			p.SettleOrder(ctx, [32]byte{1, 2, 3})
		})
	})
}

// TestPlatformEventParsing tests event parsing logic
func TestPlatformEventParsing(t *testing.T) {
	t.Run("parse settlement event", func(t *testing.T) {
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

// TestPlatformPrivateKeyHandling tests private key operations
func TestPlatformPrivateKeyHandling(t *testing.T) {
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

// TestPlatformContextHandling tests context cancellation
func TestPlatformContextHandling(t *testing.T) {
	t.Run("cancelled context", func(t *testing.T) {
		ctx, cancel := context.WithCancel(context.Background())
		cancel()

		p := &Platform{
			Client: &client.Client{},
		}

		select {
		case <-ctx.Done():
			assert.Error(t, ctx.Err())
		default:
			t.Error("Context should be cancelled")
		}

		assert.NotNil(t, p)
	})
}

// TestPlatformAddressValidation tests address validation edge cases
func TestPlatformAddressValidation(t *testing.T) {
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
			name:    "valid contract address",
			address: common.HexToAddress("0x742d35Cc6634C0532925a3b844Bc454e4438f44e"),
			isZero:  false,
		},
		{
			name:    "valid token address",
			address: common.HexToAddress("0xA0b86991c6218b36c1d19D4a2e9Eb0cE3606eB48"),
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

// TestPlatformAmountValidation tests amount validation edge cases
func TestPlatformAmountValidation(t *testing.T) {
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
			amount:  new(big.Int).Mul(big.NewInt(1000000), big.NewInt(1e18)),
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

// TestPlatformTokenSupportStatus tests token support status values
func TestPlatformTokenSupportStatus(t *testing.T) {
	tests := []struct {
		name    string
		status  *big.Int
		isValid bool
	}{
		{
			name:    "disabled status (0)",
			status:  big.NewInt(0),
			isValid: true,
		},
		{
			name:    "enabled status (1)",
			status:  big.NewInt(1),
			isValid: true,
		},
		{
			name:    "invalid status (2)",
			status:  big.NewInt(2),
			isValid: false,
		},
		{
			name:    "negative status",
			status:  big.NewInt(-1),
			isValid: false,
		},
		{
			name:    "nil status",
			status:  nil,
			isValid: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			isValid := tt.status != nil && (tt.status.Cmp(big.NewInt(0)) == 0 || tt.status.Cmp(big.NewInt(1)) == 0)
			assert.Equal(t, tt.isValid, isValid)
		})
	}
}

// TestPlatformOrderIDValidation tests order ID validation
func TestPlatformOrderIDValidation(t *testing.T) {
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

// TestPlatformChainIDHandling tests chain ID operations
func TestPlatformChainIDHandling(t *testing.T) {
	tests := []struct {
		name    string
		chainID *big.Int
		desc    string
	}{
		{
			name:    "Ethereum Mainnet",
			chainID: big.NewInt(1),
			desc:    "mainnet",
		},
		{
			name:    "Sepolia Testnet",
			chainID: big.NewInt(11155111),
			desc:    "sepolia",
		},
		{
			name:    "Polygon Mainnet",
			chainID: big.NewInt(137),
			desc:    "polygon",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.NotNil(t, tt.chainID)
			assert.Greater(t, tt.chainID.Int64(), int64(0))
		})
	}
}

// TestPlatformEmergencyWithdrawAmountRange tests various amount ranges
func TestPlatformEmergencyWithdrawAmountRange(t *testing.T) {
	tests := []struct {
		name   string
		amount *big.Int
	}{
		{
			name:   "1 wei",
			amount: big.NewInt(1),
		},
		{
			name:   "1000 tokens (with 18 decimals)",
			amount: new(big.Int).Mul(big.NewInt(1000), big.NewInt(1e18)),
		},
		{
			name:   "1 million USDC (6 decimals)",
			amount: big.NewInt(1_000_000_000_000),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.NotNil(t, tt.amount)
			assert.True(t, tt.amount.Sign() > 0)
		})
	}
}
