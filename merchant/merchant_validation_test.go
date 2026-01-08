package merchant

import (
	"context"
	"testing"

	"github.com/Dbriane208/stablebase-go-sdk/client"
	"github.com/ethereum/go-ethereum/common"
	"github.com/stretchr/testify/assert"
)

// TestMerchantMethodsValidation tests that validation is properly called in production methods
func TestRegisterMerchantValidation(t *testing.T) {
	// Create a minimal merchant instance (won't actually call blockchain)
	m := &Merchant{
		Client: &client.Client{},
	}
	ctx := context.Background()

	t.Run("invalid inputs caught before blockchain call", func(t *testing.T) {
		// Test with zero address
		_, _, _, err := m.RegisterMerchant(ctx, common.Address{}, "https://metadata.json")
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "address")

		// Test with empty metadata
		_, _, _, err = m.RegisterMerchant(ctx, common.HexToAddress("0x123"), "")
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "metadata")
	})
}

func TestUpdateMerchantValidation(t *testing.T) {
	m := &Merchant{
		Client: &client.Client{},
	}
	ctx := context.Background()

	t.Run("invalid inputs caught before blockchain call", func(t *testing.T) {
		// Test with empty merchant ID
		_, _, err := m.UpdateMerchant(ctx, [32]byte{}, common.HexToAddress("0x123"), "https://metadata.json")
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "merchant")

		// Test with zero address
		validID := [32]byte{1, 2, 3}
		_, _, err = m.UpdateMerchant(ctx, validID, common.Address{}, "https://metadata.json")
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "address")
	})
}

func TestRefundOrderValidation(t *testing.T) {
	m := &Merchant{
		Client: &client.Client{},
	}
	ctx := context.Background()

	t.Run("invalid inputs caught before blockchain call", func(t *testing.T) {
		// Test with empty order ID
		_, _, err := m.RefundOrder(ctx, [32]byte{})
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "order")
	})
}

func TestGetMerchantInfoValidation(t *testing.T) {
	m := &Merchant{
		Client: &client.Client{},
	}
	ctx := context.Background()

	t.Run("invalid inputs caught before blockchain call", func(t *testing.T) {
		// Test with empty merchant ID
		_, err := m.GetMerchantInfo(ctx, [32]byte{})
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "merchant")
	})
}
