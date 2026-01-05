package order

import (
	"context"
	"fmt"
	"math/big"

	"github.com/Dbriane208/stablebase-go-sdk/client"
	"github.com/Dbriane208/stablebase-go-sdk/contracts"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
)

// Order handles customer order operations
type Order struct {
	Client *client.Client
}

// New creates a new Customer instance
func New(c *client.Client) *Order {
	return &Order{Client: c}
}

// Create the Approval function to allow the PaymentProcessor to spend our usdc
func (o *Order) ApproveToken(ctx context.Context, tokenAddress common.Address, spenderAddress common.Address, amount *big.Int) (*types.Receipt, error) {
	// Create transaction options
	auth, err := bind.NewKeyedTransactorWithChainID(o.Client.PrivateKey, o.Client.ChainID)
	if err != nil {
		return nil, err
	}

	// Create ERC20 contract instance for the token
	token, err := contracts.NewERC20(tokenAddress, o.Client.EthClient)
	if err != nil {
		return nil, fmt.Errorf("failed to instantiate ERC20: %w", err)
	}

	// Call approve on the token contract
	tx, err := token.Approve(auth, spenderAddress, amount)
	if err != nil {
		return nil, err
	}

	// Wait for mining
	receipt, err := bind.WaitMined(ctx, o.Client.EthClient, tx)
	if err != nil {
		return nil, err
	}

	// Check status
	if receipt.Status == 0 {
		return receipt, fmt.Errorf("approval transaction failed: %s", tx.Hash().Hex())
	}

	return receipt, nil
}

// CreateOrder creates a new payment order initiated by the customer
func (o *Order) CreateOrder(ctx context.Context, merchantID [32]byte, tokenAddress common.Address, amount *big.Int, metadataURI string) (*contracts.PaymentProcessorOrderCreated, [32]byte, *types.Receipt, error) {
	// Create transaction options
	auth, err := bind.NewKeyedTransactorWithChainID(o.Client.PrivateKey, o.Client.ChainID)
	if err != nil {
		return nil, [32]byte{}, nil, err
	}

	// Call the contract
	tx, err := o.Client.PaymentProcessor.CreateOrder(auth, merchantID, tokenAddress, amount, metadataURI)
	if err != nil {
		return nil, [32]byte{}, nil, err
	}

	// Wait for the transaction to be mined
	receipt, err := bind.WaitMined(ctx, o.Client.EthClient, tx)
	if err != nil {
		return nil, [32]byte{}, nil, err
	}

	// Check if transaction succeeded
	if receipt.Status == 0 {
		return nil, [32]byte{}, receipt, fmt.Errorf("transaction failed: %s", tx.Hash().Hex())
	}

	// Get the orderId by querying the event
	for _, log := range receipt.Logs {
		event, err := o.Client.PaymentProcessor.ParseOrderCreated(*log)
		if err == nil {
			return event, event.OrderId, receipt, nil
		}
	}

	// If no OrderId found return error
	return nil, [32]byte{}, receipt, fmt.Errorf("order created but event not found in logs")
}

// PayOrder executes payment for an order
func (o *Order) PayOrder(ctx context.Context, orderId [32]byte) (*contracts.PaymentProcessorOrderPaid, *types.Receipt, error) {
	// Create transaction options
	auth, err := bind.NewKeyedTransactorWithChainID(o.Client.PrivateKey, o.Client.ChainID)
	if err != nil {
		return nil, nil, err
	}

	// Call the contract
	tx, err := o.Client.PaymentProcessor.PayOrder(auth, orderId)
	if err != nil {
		return nil, nil, err
	}

	// Wait for the transaction to be mined
	receipt, err := bind.WaitMined(ctx, o.Client.EthClient, tx)
	if err != nil {
		return nil, nil, err
	}

	// Check if transaction succeeded
	if receipt.Status == 0 {
		return nil, receipt, fmt.Errorf("payment transaction failed: %s", tx.Hash().Hex())
	}

	// Parse the OrderPaid event
	for _, log := range receipt.Logs {
		event, err := o.Client.PaymentProcessor.ParseOrderPaid(*log)
		if err == nil {
			// Return the event with useful data
			return event, receipt, nil
		}
	}

	// If no event found, still return success but with warning
	return nil, receipt, fmt.Errorf("payment succeeded but event not found in logs")
}

// Cancel Order allows payer to cancel an order
func (o *Order) CancelOrder(ctx context.Context, orderId [32]byte) (*contracts.PaymentProcessorOrderCancelled, *types.Receipt, error) {
	// Create transaction options
	auth, err := bind.NewKeyedTransactorWithChainID(o.Client.PrivateKey, o.Client.ChainID)
	if err != nil {
		return nil, nil, err
	}

	// Call the contract
	tx, err := o.Client.PaymentProcessor.CancelOrder(auth, orderId)
	if err != nil {
		return nil, nil, err
	}

	// Wait for the transaction to be mined
	receipt, err := bind.WaitMined(ctx, o.Client.EthClient, tx)
	if err != nil {
		return nil, nil, err
	}

	// Check if transaction succeeded
	if receipt.Status == 0 {
		return nil, receipt, fmt.Errorf("cancellation transaction failed: %s", tx.Hash().Hex())
	}

	// Parse the CancelOrder event
	for _, log := range receipt.Logs {
		event, err := o.Client.PaymentProcessor.ParseOrderCancelled(*log)
		if err == nil {
			return event, receipt, nil
		}
	}

	// If no event found, still return success but with warning
	return nil, receipt, fmt.Errorf("cancellation succeeded but event not found in logs")
}

// GetOrder retrieves Order information from the blockchain
func (o *Order) GetOrder(ctx context.Context, orderId [32]byte) (*contracts.IPaymentProcessorOrder, error) {
	// Create call options (read-only, no transaction)
	opts := &bind.CallOpts{Context: ctx}

	// Call the contract to get order info
	orderInfo, err := o.Client.PaymentProcessor.GetOrder(opts, orderId)
	if err != nil {
		return nil, err
	}

	return &orderInfo, nil
}

// GetOrderStatus gets the status of the order
func (o *Order) GetOrderStatus(ctx context.Context, orderId [32]byte) (uint8, error) {
	// Create call options
	opts := &bind.CallOpts{Context: ctx}

	// Call the contract
	status, err := o.Client.PaymentProcessor.GetOrderStatus(opts, orderId)
	if err != nil {
		return 0, err
	}

	return status, nil
}

// IsOrderExpired checks if an order has expired
func (o *Order) IsOrderExpired(ctx context.Context, orderId [32]byte) (bool, error) {
	// Create call options
	opts := &bind.CallOpts{Context: ctx}

	// Call the contract
	isOrderExpired, err := o.Client.PaymentProcessor.IsOrderExpired(opts, orderId)
	if err != nil {
		return false, err
	}

	return isOrderExpired, nil
}

// GetOrderCreatedAt check when the order was created
func (o *Order) GetOrderCreatedAt(ctx context.Context, orderId [32]byte) (*big.Int, error) {
	// Create call options
	opts := &bind.CallOpts{Context: ctx}

	// Call the contract to get creation time
	createdAt, err := o.Client.PaymentProcessor.GetOrderCreatedAt(opts, orderId)
	if err != nil {
		return nil, err
	}

	return createdAt, nil
}

// GetOrderRemainingTime gets remaining time before order expires
func (o *Order) GetOrderRemainingTime(ctx context.Context, orderId [32]byte) (*big.Int, error) {
	// Create call options
	opts := &bind.CallOpts{Context: ctx}

	// Call the contract to get remaining time
	remainingTime, err := o.Client.PaymentProcessor.GetOrderRemainingTime(opts, orderId)
	if err != nil {
		return nil, err
	}

	return remainingTime, nil
}
