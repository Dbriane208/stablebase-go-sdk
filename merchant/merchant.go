package merchant

import (
	"context"
	"fmt"
	"math/big"

	"github.com/Dbriane208/stablebase-go-sdk/client"
	"github.com/Dbriane208/stablebase-go-sdk/contracts"
	"github.com/Dbriane208/stablebase-go-sdk/utils"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
)

// Merchant handles merchant operations
type Merchant struct {
	Client *client.Client
}

// New creates a new Merchant instance
func New(c *client.Client) *Merchant {
	return &Merchant{Client: c}
}

// RegisterMerchant registers a new merchant on the blockchain
// Returns the merchantId and transaction receipt after confirmation
func (m *Merchant) RegisterMerchant(ctx context.Context, payoutWallet common.Address, metadataURI string) (*contracts.MerchantRegistryMerchantRegistered, [32]byte, *types.Receipt, error) {
	// Validate inputs
	if err := utils.ValidateRegisterMerchantInputs(payoutWallet, metadataURI); err != nil {
		return nil, [32]byte{}, nil, err
	}

	// Create transaction options
	auth, err := bind.NewKeyedTransactorWithChainID(m.Client.PrivateKey, m.Client.ChainID)
	if err != nil {
		return nil, [32]byte{}, nil, err
	}

	// Call the contract
	tx, err := m.Client.MerchantRegistry.RegisterMerchant(auth, payoutWallet, metadataURI)
	if err != nil {
		return nil, [32]byte{}, nil, err
	}

	// Wait for the transaction to be mined
	receipt, err := bind.WaitMined(ctx, m.Client.EthClient, tx)
	if err != nil {
		return nil, [32]byte{}, nil, err
	}

	// Check if transaction succeeded
	if receipt.Status == 0 {
		return nil, [32]byte{}, receipt, fmt.Errorf("transaction failed: %s", tx.Hash().Hex())
	}

	// Parse the MerchantRegistered event to get the merchantId
	for _, log := range receipt.Logs {
		event, err := m.Client.MerchantRegistry.ParseMerchantRegistered(*log)
		if err == nil {
			// Found the merchantId in the event!
			return event, event.MerchantId, receipt, nil
		}
	}

	// If we couldn't find the event, return error
	return nil, [32]byte{}, receipt, fmt.Errorf("merchant registered but event not found in logs")
}

// UpdateMerchant updates a merchant on the blockchain
func (m *Merchant) UpdateMerchant(ctx context.Context, merchantId [32]byte, payoutWallet common.Address, metadataURI string) (*contracts.MerchantRegistryMerchantUpdated, *types.Receipt, error) {
	// Validate inputs
	if err := utils.ValidateUpdateMerchantInputs(merchantId, payoutWallet, metadataURI); err != nil {
		return nil, nil, err
	}

	// Create transaction options
	auth, err := bind.NewKeyedTransactorWithChainID(m.Client.PrivateKey, m.Client.ChainID)
	if err != nil {
		return nil, nil, err
	}

	// Call the contract
	tx, err := m.Client.MerchantRegistry.UpdateMerchant(auth, merchantId, payoutWallet, metadataURI)
	if err != nil {
		return nil, nil, err
	}

	// Wait for the transaction to be mined
	receipt, err := bind.WaitMined(ctx, m.Client.EthClient, tx)
	if err != nil {
		return nil, nil, err
	}

	// Check if transaction succeeded
	if receipt.Status == 0 {
		return nil, receipt, fmt.Errorf("transaction failed: %s", tx.Hash().Hex())
	}

	// Parse the event
	for _, log := range receipt.Logs {
		event, err := m.Client.MerchantRegistry.ParseMerchantUpdated(*log)
		if err == nil {
			return event, receipt, nil
		}
	}

	// If we couldn't find the event, return error
	return nil, receipt, fmt.Errorf("merchant updated but event not found in logs")
}

// RefundOrder allows platform and merchant to return order back to payer
func (m *Merchant) RefundOrder(ctx context.Context, orderId [32]byte) (*contracts.PaymentProcessorOrderRefunded, *types.Receipt, error) {
	// Validate inputs
	if err := utils.ValidateOrderId(orderId); err != nil {
		return nil, nil, err
	}

	// Create transaction options
	auth, err := bind.NewKeyedTransactorWithChainID(m.Client.PrivateKey, m.Client.ChainID)
	if err != nil {
		return nil, nil, err
	}

	// Call the contract
	tx, err := m.Client.PaymentProcessor.RefundOrder(auth, orderId)
	if err != nil {
		return nil, nil, err
	}

	// Wait for the transaction to be mined
	receipt, err := bind.WaitMined(ctx, m.Client.EthClient, tx)
	if err != nil {
		return nil, nil, err
	}

	// Check if transaction succeeded
	if receipt.Status == 0 {
		return nil, receipt, fmt.Errorf("refund transaction failed: %s", tx.Hash().Hex())
	}

	// Parse the RefundOrder event
	for _, log := range receipt.Logs {
		event, err := m.Client.PaymentProcessor.ParseOrderRefunded(*log)
		if err == nil {
			return event, receipt, nil
		}
	}

	// If no event found, still return success but with warning
	return nil, receipt, fmt.Errorf("refund succeeded but event not found in logs")
}

// GetMerchantInfo retrieves merchant information from the blockchain
func (m *Merchant) GetMerchantInfo(ctx context.Context, merchantId [32]byte) (*contracts.IMerchantRegistryMerchant, error) {
	// Validate inputs
	if err := utils.ValidateMerchantId(merchantId); err != nil {
		return nil, err
	}

	// Create call options (read-only, no transaction)
	opts := &bind.CallOpts{Context: ctx}

	// Call the contract to get merchant info
	merchantInfo, err := m.Client.MerchantRegistry.GetMerchantInfo(opts, merchantId)
	if err != nil {
		return nil, err
	}

	return &merchantInfo, nil
}

// IsMerchantVerified checks whether the merchant is verified from the blockchain
func (m *Merchant) IsMerchantVerified(ctx context.Context, merchantId [32]byte) (bool, error) {
	// Validate inputs
	if err := utils.ValidateMerchantId(merchantId); err != nil {
		return false, err
	}

	// Create call options
	opts := &bind.CallOpts{Context: ctx}

	// Call the contract to check whether merchant is verified
	verification, err := m.Client.MerchantRegistry.IsMerchantVerified(opts, merchantId)
	if err != nil {
		return false, err
	}

	return verification, nil
}

// GetMerchantVerificationStatus retrieves merchant verification status from the blockchain
func (m *Merchant) GetMerchantVerificationStatus(ctx context.Context, merchantId [32]byte) (uint8, error) {
	// Validate inputs
	if err := utils.ValidateMerchantId(merchantId); err != nil {
		return 0, err
	}

	// Create call options
	opts := &bind.CallOpts{Context: ctx}

	// Call the contract to check merchant verification status
	// verification status: 0=Pending, 1=Verified, 2=Rejected, 3=Suspended
	status, err := m.Client.MerchantRegistry.GetMerchantVerificationStatus(opts, merchantId)
	if err != nil {
		return 0, err
	}

	return status, nil
}

// GetMerchantCreatedAt returns the timestamp when the merchant was registered
func (m *Merchant) GetMerchantCreatedAt(ctx context.Context, merchantId [32]byte) (*big.Int, error) {
	// Validate inputs
	if err := utils.ValidateMerchantId(merchantId); err != nil {
		return nil, err
	}

	// Create call options
	opts := &bind.CallOpts{Context: ctx}

	// Call the contract to get merchant creation timestamp
	createdAt, err := m.Client.MerchantRegistry.GetMerchantCreatedAt(opts, merchantId)
	if err != nil {
		return nil, err
	}

	return createdAt, nil
}

// GetMerchantTokenBalance gets the merchant's token balance for a specific token
func (m *Merchant) GetMerchantTokenBalance(ctx context.Context, merchantWalletAddress common.Address, tokenAddress common.Address) (*big.Int, error) {
	// Validate inputs
	if err := utils.ValidateMerchatTokenBalanceInputs(merchantWalletAddress, tokenAddress); err != nil {
		return nil, err
	}

	// Create call options
	opts := &bind.CallOpts{Context: ctx}

	// Create ERC20 token instance
	token, err := contracts.NewERC20(tokenAddress, m.Client.EthClient)
	if err != nil {
		return nil, fmt.Errorf("failed to instantiate ERC20: %w", err)
	}

	// Call the token contract to get balance
	balance, err := token.BalanceOf(opts, merchantWalletAddress)
	if err != nil {
		return nil, err
	}

	return balance, nil
}
