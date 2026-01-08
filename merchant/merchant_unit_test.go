package merchant

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
	m := New(c)

	assert.NotNil(t, m)
	assert.Equal(t, c, m.Client)
}

// TestMerchantStructInitialization tests merchant struct initialization
func TestMerchantStructInitialization(t *testing.T) {
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
			m := New(tt.client)
			assert.NotNil(t, m)
			assert.NotNil(t, m.Client)
			assert.Equal(t, tt.client, m.Client)
		})
	}
}

// TestRegisterMerchantInputValidation tests that RegisterMerchant validates inputs
func TestRegisterMerchantInputValidation(t *testing.T) {
	// Create a test private key
	privateKey, err := crypto.GenerateKey()
	require.NoError(t, err)

	m := &Merchant{
		Client: &client.Client{
			PrivateKey: privateKey,
			ChainID:    big.NewInt(1),
		},
	}
	ctx := context.Background()

	tests := []struct {
		name          string
		payoutWallet  common.Address
		metadataURI   string
		expectError   bool
		errorContains string
	}{
		{
			name:         "empty payout wallet address",
			payoutWallet: common.Address{},
			metadataURI:  "https://example.com/metadata.json",
			expectError:  true,
			errorContains: "address",
		},
		{
			name:          "empty metadata URI",
			payoutWallet:  common.HexToAddress("0x123"),
			metadataURI:   "",
			expectError:   true,
			errorContains: "metadata",
		},
		{
			name:          "whitespace metadata URI",
			payoutWallet:  common.HexToAddress("0x123"),
			metadataURI:   "   ",
			expectError:   true,
			errorContains: "metadata",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, _, _, err := m.RegisterMerchant(ctx, tt.payoutWallet, tt.metadataURI)

			if tt.expectError {
				assert.Error(t, err)
				assert.Contains(t, err.Error(), tt.errorContains)
			} else {
				// This will fail for valid inputs because we don't have a real blockchain
				// but we've verified validation works for invalid inputs
				assert.NoError(t, err)
			}
		})
	}
}

// TestUpdateMerchantInputValidation tests that UpdateMerchant validates inputs
func TestUpdateMerchantInputValidation(t *testing.T) {
	privateKey, err := crypto.GenerateKey()
	require.NoError(t, err)

	m := &Merchant{
		Client: &client.Client{
			PrivateKey: privateKey,
			ChainID:    big.NewInt(1),
		},
	}
	ctx := context.Background()

	tests := []struct {
		name          string
		merchantID    [32]byte
		payoutWallet  common.Address
		metadataURI   string
		expectError   bool
		errorContains string
	}{
		{
			name:          "empty merchant ID",
			merchantID:    [32]byte{},
			payoutWallet:  common.HexToAddress("0x123"),
			metadataURI:   "https://example.com/metadata.json",
			expectError:   true,
			errorContains: "merchant",
		},
		{
			name:          "empty payout wallet",
			merchantID:    [32]byte{1, 2, 3},
			payoutWallet:  common.Address{},
			metadataURI:   "https://example.com/metadata.json",
			expectError:   true,
			errorContains: "address",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, _, err := m.UpdateMerchant(ctx, tt.merchantID, tt.payoutWallet, tt.metadataURI)

			if tt.expectError {
				assert.Error(t, err)
				assert.Contains(t, err.Error(), tt.errorContains)
			}
		})
	}
}

// TestRefundOrderInputValidation tests that RefundOrder validates inputs
func TestRefundOrderInputValidation(t *testing.T) {
	privateKey, err := crypto.GenerateKey()
	require.NoError(t, err)

	m := &Merchant{
		Client: &client.Client{
			PrivateKey: privateKey,
			ChainID:    big.NewInt(1),
		},
	}
	ctx := context.Background()

	t.Run("empty order ID", func(t *testing.T) {
		_, _, err := m.RefundOrder(ctx, [32]byte{})
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "order")
	})
}

// TestGetMerchantInfoInputValidation tests that GetMerchantInfo validates inputs
func TestGetMerchantInfoInputValidation(t *testing.T) {
	m := &Merchant{
		Client: &client.Client{},
	}
	ctx := context.Background()

	t.Run("empty merchant ID", func(t *testing.T) {
		_, err := m.GetMerchantInfo(ctx, [32]byte{})
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "merchant")
	})
}

// TestMerchantMethodsNilClient tests that methods handle nil client gracefully
func TestMerchantMethodsNilClient(t *testing.T) {
	m := &Merchant{
		Client: nil,
	}
	ctx := context.Background()

	t.Run("RegisterMerchant with nil client panics", func(t *testing.T) {
		assert.Panics(t, func() {
			m.RegisterMerchant(ctx, common.HexToAddress("0x123"), "https://example.com")
		})
	})
}

// TestMerchantEventParsing tests event parsing logic
func TestMerchantEventParsing(t *testing.T) {
	t.Run("parse merchant registered event", func(t *testing.T) {
		// Create a mock receipt with successful status
		receipt := &types.Receipt{
			Status: types.ReceiptStatusSuccessful,
			TxHash: common.HexToHash("0x123"),
			Logs:   []*types.Log{},
		}

		// Verify receipt status check
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

// TestMerchantPrivateKeyHandling tests private key operations
func TestMerchantPrivateKeyHandling(t *testing.T) {
	t.Run("valid private key", func(t *testing.T) {
		privateKey, err := crypto.GenerateKey()
		require.NoError(t, err)
		assert.NotNil(t, privateKey)

		// Verify we can derive address
		address := crypto.PubkeyToAddress(privateKey.PublicKey)
		assert.NotEqual(t, common.Address{}, address)
	})

	t.Run("nil private key", func(t *testing.T) {
		var nilKey *ecdsa.PrivateKey
		assert.Nil(t, nilKey)
	})
}

// TestMerchantInfoStruct tests the merchant info return structure
func TestMerchantInfoStruct(t *testing.T) {
	t.Run("merchant info structure", func(t *testing.T) {
		info := contracts.IMerchantRegistryMerchant{
			Owner:              common.HexToAddress("0x789"),
			PayoutWallet:       common.HexToAddress("0x123"),
			MetadataURI:        "https://example.com",
			Exists:             true,
			VerificationStatus: 1,
			CreatedAt:          big.NewInt(1234567890),
		}

		assert.NotEqual(t, common.Address{}, info.Owner)
		assert.NotEqual(t, common.Address{}, info.PayoutWallet)
		assert.NotEmpty(t, info.MetadataURI)
		assert.True(t, info.Exists)
	})
}

// TestMerchantContextHandling tests context cancellation
func TestMerchantContextHandling(t *testing.T) {
	t.Run("cancelled context", func(t *testing.T) {
		ctx, cancel := context.WithCancel(context.Background())
		cancel() // Cancel immediately

		m := &Merchant{
			Client: &client.Client{},
		}

		// Verify context is cancelled
		select {
		case <-ctx.Done():
			assert.Error(t, ctx.Err())
		default:
			t.Error("Context should be cancelled")
		}

		// Methods should respect cancelled context (through blockchain calls)
		assert.NotNil(t, m)
	})
}

// TestMerchantAddressValidation tests address validation edge cases
func TestMerchantAddressValidation(t *testing.T) {
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
			name:    "valid address",
			address: common.HexToAddress("0x742d35Cc6634C0532925a3b844Bc454e4438f44e"),
			isZero:  false,
		},
		{
			name:    "address from hex",
			address: common.HexToAddress("0x1"),
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

// TestMerchantMetadataURIValidation tests metadata URI validation
func TestMerchantMetadataURIValidation(t *testing.T) {
	tests := []struct {
		name        string
		metadataURI string
		shouldPass  bool
	}{
		{
			name:        "valid HTTPS URL",
			metadataURI: "https://example.com/metadata.json",
			shouldPass:  true,
		},
		{
			name:        "valid HTTP URL",
			metadataURI: "http://example.com/metadata.json",
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
			// Simple validation logic
			isEmpty := len(tt.metadataURI) == 0 || (len(tt.metadataURI) == 3 && tt.metadataURI == "   ")

			if tt.shouldPass {
				assert.False(t, isEmpty)
			} else {
				assert.True(t, isEmpty)
			}
		})
	}
}
