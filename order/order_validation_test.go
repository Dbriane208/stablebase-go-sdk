package order

import (
	"context"
	"math/big"
	"testing"

	"github.com/Dbriane208/stablebase-go-sdk/client"
	"github.com/ethereum/go-ethereum/common"
	"github.com/stretchr/testify/assert"
)

func TestCreateOrderValidation(t *testing.T) {
	// Order instance
	o := &Order{
		Client: &client.Client{},
	}
	ctx := context.Background()

	t.Run("invalid inputs caught before blockchain call", func(t *testing.T) {
		merchantId := [32]byte{1, 2, 3, 4, 5}
		emptyMerchantId := [32]byte{}

		// Test with zero address
		_, _, _, err := o.CreateOrder(ctx, merchantId, common.Address{}, big.NewInt(1_000_000), "https://order/metadata.json")
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "tokenAddress")

		// Test with empty metadata
		_, _, _, err = o.CreateOrder(ctx, merchantId, common.HexToAddress("0x123"), big.NewInt(1_000_000), "")
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "metadata")

		// Test with empty merchantId
		_, _, _, err = o.CreateOrder(ctx, emptyMerchantId, common.HexToAddress("0x123"), big.NewInt(1_000_000), "https://order/metadata.json")
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "merchantId")

		// Test with zero amount
		_, _, _, err = o.CreateOrder(ctx, merchantId, common.HexToAddress("0x123"), big.NewInt(0), "https://order/metadata.json")
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "amount")
	})
}

func TestPayOrderValidation(t *testing.T) {
	// order instance
	o := &Order{
		Client: &client.Client{},
	}

	ctx := context.Background()

	t.Run("invalid inputs caught before blockchain call", func(t *testing.T) {
		emptyOrderId := [32]byte{}

		// Test with invalid details
		_, _, err := o.PayOrder(ctx, emptyOrderId)
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "orderId")
	})
}

func TestCancelOrderValidation(t *testing.T) {
	// order instance
	o := &Order{
		Client: &client.Client{},
	}

	ctx := context.Background()

	t.Run("invalid inputs caught before blockchain call", func(t *testing.T) {
		emptyOrderId := [32]byte{}

		// Test with invalid details
		_, _, err := o.CancelOrder(ctx, emptyOrderId)
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "orderId")
	})
}

func TestGetOrderValidation(t *testing.T) {
	// order instance
	o := &Order{
		Client: &client.Client{},
	}

	ctx := context.Background()

	t.Run("invalid inputs caught before blockchain call", func(t *testing.T) {
		emptyOrderId := [32]byte{}

		// Test with invalid details
		_, err := o.GetOrder(ctx, emptyOrderId)
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "orderId")
	})
}

func TestGetOrderStatusValidation(t *testing.T) {
	// order instance
	o := &Order{
		Client: &client.Client{},
	}

	ctx := context.Background()

	t.Run("invalid inputs caught before blockchain call", func(t *testing.T) {
		emptyOrderId := [32]byte{}

		// Test with invalid details
		_, err := o.GetOrderStatus(ctx, emptyOrderId)
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "orderId")
	})
}

func TestIsOrderExpiredValidation(t *testing.T) {
	// order instance
	o := &Order{
		Client: &client.Client{},
	}

	ctx := context.Background()

	t.Run("invalid inputs caught before blockchain call", func(t *testing.T) {
		emptyOrderId := [32]byte{}

		// Test with invalid details
		_, err := o.IsOrderExpired(ctx, emptyOrderId)
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "orderId")
	})
}

func TestGetOrderCreatedAtValidation(t *testing.T) {
	// order instance
	o := &Order{
		Client: &client.Client{},
	}

	ctx := context.Background()

	t.Run("invalid inputs caught before blockchain call", func(t *testing.T) {
		emptyOrderId := [32]byte{}

		// Test with invalid details
		_, err := o.GetOrderCreatedAt(ctx, emptyOrderId)
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "orderId")
	})
}

func TestGetOrderRemainingTimeValidation(t *testing.T) {
	// order instance
	o := &Order{
		Client: &client.Client{},
	}

	ctx := context.Background()

	t.Run("invalid inputs caught before blockchain call", func(t *testing.T) {
		emptyOrderId := [32]byte{}

		// Test with invalid details
		_, err := o.GetOrderRemainingTime(ctx, emptyOrderId)
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "orderId")
	})
}
