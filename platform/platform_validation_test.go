package platform

import (
	"context"
	"math/big"
	"testing"

	"github.com/Dbriane208/stablebase-go-sdk/client"
	"github.com/ethereum/go-ethereum/common"
	"github.com/stretchr/testify/assert"
)

// TestPlatformMethodsValidation tests that validation is properly called in production methods
func TestSettleOrderValidation(t *testing.T) {
	// Create a minimal platform instance (won't actually call blockchain)
	p := &Platform{
		Client: &client.Client{},
	}
	ctx := context.Background()

	t.Run("invalid inputs caught before blockchain call", func(t *testing.T) {
		// Test with empty order ID
		_, _, err := p.SettleOrder(ctx, [32]byte{})
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "order")
	})
}

func TestRefundOrderValidation(t *testing.T) {
	p := &Platform{
		Client: &client.Client{},
	}
	ctx := context.Background()

	t.Run("invalid inputs caught before blockchain call", func(t *testing.T) {
		// Test with empty order ID
		_, _, err := p.RefundOrder(ctx, [32]byte{})
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "order")
	})
}

func TestEmergencyWithdrawValidation(t *testing.T) {
	p := &Platform{
		Client: &client.Client{},
	}
	ctx := context.Background()

	t.Run("invalid inputs caught before blockchain call", func(t *testing.T) {
		// Test with zero token address
		_, _, err := p.EmergencyWithdraw(ctx, common.Address{}, common.HexToAddress("0x123"), big.NewInt(100))
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "token")

		// Test with zero receiver address
		_, _, err = p.EmergencyWithdraw(ctx, common.HexToAddress("0x123"), common.Address{}, big.NewInt(100))
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "receiver")

		// Test with nil amount
		_, _, err = p.EmergencyWithdraw(ctx, common.HexToAddress("0x123"), common.HexToAddress("0x456"), nil)
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "amount")

		// Test with zero amount
		_, _, err = p.EmergencyWithdraw(ctx, common.HexToAddress("0x123"), common.HexToAddress("0x456"), big.NewInt(0))
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "amount")

		// Test with negative amount
		_, _, err = p.EmergencyWithdraw(ctx, common.HexToAddress("0x123"), common.HexToAddress("0x456"), big.NewInt(-1))
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "amount")
	})
}

func TestUpdateMerchantRegistryValidation(t *testing.T) {
	p := &Platform{
		Client: &client.Client{},
	}
	ctx := context.Background()

	t.Run("invalid inputs caught before blockchain call", func(t *testing.T) {
		// Test with zero address
		_, _, err := p.UpdateMerchantRegistry(ctx, common.Address{})
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "address")
	})
}

func TestSetTokenSupportValidation(t *testing.T) {
	p := &Platform{
		Client: &client.Client{},
	}
	ctx := context.Background()

	t.Run("invalid inputs caught before blockchain call", func(t *testing.T) {
		// Test with zero token address
		_, _, err := p.SetTokenSupport(ctx, common.Address{}, big.NewInt(1))
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "token")

		// Test with nil status
		_, _, err = p.SetTokenSupport(ctx, common.HexToAddress("0x123"), nil)
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "status")

		// Test with invalid status (not 0 or 1)
		_, _, err = p.SetTokenSupport(ctx, common.HexToAddress("0x123"), big.NewInt(2))
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "status")

		// Test with negative status
		_, _, err = p.SetTokenSupport(ctx, common.HexToAddress("0x123"), big.NewInt(-1))
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "status")
	})
}

func TestGetPlatformTokenBalanceValidation(t *testing.T) {
	p := &Platform{
		Client: &client.Client{},
	}
	ctx := context.Background()

	t.Run("invalid inputs caught before blockchain call", func(t *testing.T) {
		// Test with zero platform wallet address
		_, err := p.GetPlatformTokenBalance(ctx, common.Address{}, common.HexToAddress("0x123"))
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "platform")

		// Test with zero token address
		_, err = p.GetPlatformTokenBalance(ctx, common.HexToAddress("0x123"), common.Address{})
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "token")
	})
}

func TestGetContractTokenBalanceValidation(t *testing.T) {
	p := &Platform{
		Client: &client.Client{},
	}
	ctx := context.Background()

	t.Run("invalid inputs caught before blockchain call", func(t *testing.T) {
		// Test with zero token address
		_, err := p.GetContractTokenBalance(ctx, common.Address{})
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "address")
	})
}

