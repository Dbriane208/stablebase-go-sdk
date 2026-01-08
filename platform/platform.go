package platform

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

type Platform struct {
	Client *client.Client
}

func New(c *client.Client) *Platform {
	return &Platform{Client: c}
}

// SettleOrder distributes the order between the merchant and the platform
func (p *Platform) SettleOrder(ctx context.Context, orderId [32]byte) (*contracts.PaymentProcessorOrderSettled, *types.Receipt, error) {
	//Validate inputs
	if err := utils.ValidateOrderId(orderId); err != nil {
		return nil, nil, err
	}

	// Create transaction options
	auth, err := bind.NewKeyedTransactorWithChainID(p.Client.PrivateKey, p.Client.ChainID)
	if err != nil {
		return nil, nil, err
	}

	// Call the contract
	tx, err := p.Client.PaymentProcessor.SettleOrder(auth, orderId)
	if err != nil {
		return nil, nil, err
	}

	// Wait for the transaction to be mined
	receipt, err := bind.WaitMined(ctx, p.Client.EthClient, tx)
	if err != nil {
		return nil, nil, err
	}

	// Check if transaction succeeded
	if receipt.Status == 0 {
		return nil, receipt, fmt.Errorf("settlement transaction failed: %s", tx.Hash().Hex())
	}

	// Parse the SettleOrder event
	for _, log := range receipt.Logs {
		event, err := p.Client.PaymentProcessor.ParseOrderSettled(*log)
		if err == nil {
			return event, receipt, nil
		}
	}

	// If no event found, still return success but with warning
	return nil, receipt, fmt.Errorf("settlement succeeded but event not found in logs")
}

// RefundOrder allows platform and merchant to return order back to payer
func (p *Platform) RefundOrder(ctx context.Context, orderId [32]byte) (*contracts.PaymentProcessorOrderRefunded, *types.Receipt, error) {
	//Validate inputs
	if err := utils.ValidateOrderId(orderId); err != nil {
		return nil, nil, err
	}

	// Create transaction options
	auth, err := bind.NewKeyedTransactorWithChainID(p.Client.PrivateKey, p.Client.ChainID)
	if err != nil {
		return nil, nil, err
	}

	// Call the contract
	tx, err := p.Client.PaymentProcessor.RefundOrder(auth, orderId)
	if err != nil {
		return nil, nil, err
	}

	// Wait for the transaction to be mined
	receipt, err := bind.WaitMined(ctx, p.Client.EthClient, tx)
	if err != nil {
		return nil, nil, err
	}

	// Check if transaction succeeded
	if receipt.Status == 0 {
		return nil, receipt, fmt.Errorf("refund transaction failed: %s", tx.Hash().Hex())
	}

	// Parse the RefundOrder event
	for _, log := range receipt.Logs {
		event, err := p.Client.PaymentProcessor.ParseOrderRefunded(*log)
		if err == nil {
			return event, receipt, nil
		}
	}

	// If no event found, still return success but with warning
	return nil, receipt, fmt.Errorf("refund succeeded but event not found in logs")
}

// EmergencyWithdraw allows platform admin to withdraw funds from the contract in emergency situations
func (p *Platform) EmergencyWithdraw(ctx context.Context, tokenAddress common.Address, receiverAddress common.Address, amount *big.Int) (*contracts.PaymentProcessorEmergencyWithdrawalSuccess, *types.Receipt, error) {
	// Validate inputs
	if err := utils.ValidateEmergencyWithdrawInputs(tokenAddress, receiverAddress, amount); err != nil {
		return nil, nil, err
	}

	// Create transaction opts
	auth, err := bind.NewKeyedTransactorWithChainID(p.Client.PrivateKey, p.Client.ChainID)
	if err != nil {
		return nil, nil, err
	}

	// Call the contract's emergency withdraw function
	tx, err := p.Client.PaymentProcessor.EmergencyWithdraw(auth, tokenAddress, receiverAddress, amount)
	if err != nil {
		return nil, nil, err
	}

	// Wait for the transaction to be mined
	receipt, err := bind.WaitMined(ctx, p.Client.EthClient, tx)
	if err != nil {
		return nil, nil, err
	}

	// Check if transaction succeeded
	if receipt.Status == 0 {
		return nil, receipt, fmt.Errorf("emergency withdrawal transaction failed: %s", tx.Hash().Hex())
	}

	// Parse the EmergencyWithdrawalSuccess event
	for _, log := range receipt.Logs {
		event, err := p.Client.PaymentProcessor.ParseEmergencyWithdrawalSuccess(*log)
		if err == nil {
			return event, receipt, nil
		}
	}

	// If no event found, still return success but with warning
	return nil, receipt, fmt.Errorf("emergency withdrawal succeeded but event not found in logs")
}

// SetEmergencyWithdrawalEnabled Enables/disable emergency withdrawal functionality
func (p *Platform) SetEmergencyWithdrawalEnabled(ctx context.Context, enable bool) (*contracts.PaymentProcessorEmergencyWithdrawalEnabledUpdated, *types.Receipt, error) {
	// Create transaction options
	auth, err := bind.NewKeyedTransactorWithChainID(p.Client.PrivateKey, p.Client.ChainID)
	if err != nil {
		return nil, nil, err
	}

	// Call the contract
	tx, err := p.Client.PaymentProcessor.SetEmergencyWithdrawalEnabled(auth, enable)
	if err != nil {
		return nil, nil, err
	}

	// Wait for the transaction to be mined
	receipt, err := bind.WaitMined(ctx, p.Client.EthClient, tx)
	if err != nil {
		return nil, nil, err
	}

	// Check if transaction succeed
	if receipt.Status == 0 {
		return nil, receipt, fmt.Errorf("set emergency withdrawal enabled transaction failed: %s", tx.Hash().Hex())
	}

	// Parse the EmergencyWithdrawal event
	for _, log := range receipt.Logs {
		event, err := p.Client.PaymentProcessor.ParseEmergencyWithdrawalEnabledUpdated(*log)
		if err == nil {
			return event, receipt, nil
		}
	}

	// if no event found, still return success but with warning
	return nil, receipt, fmt.Errorf("emergency withdrawal enabled but event not found in logs")
}

// UpdateMerchantRegistry updates the address of new deployed contract
func (p *Platform) UpdateMerchantRegistry(ctx context.Context, newRegistryAddress common.Address) (*contracts.PaymentProcessorMerchantRegistryUpdated, *types.Receipt, error) {
	// Validate inputs
	if err := utils.ValidateUpdateMerchantRegistryInput(newRegistryAddress); err != nil {
		return nil, nil, err
	}

	// Create transaction options
	auth, err := bind.NewKeyedTransactorWithChainID(p.Client.PrivateKey, p.Client.ChainID)
	if err != nil {
		return nil, nil, err
	}

	// Call the contract
	tx, err := p.Client.PaymentProcessor.UpdateMerchantRegistry(auth, newRegistryAddress)
	if err != nil {
		return nil, nil, err
	}

	// Wait for the transaction to be mined
	receipt, err := bind.WaitMined(ctx, p.Client.EthClient, tx)
	if err != nil {
		return nil, nil, err
	}

	// Check if transaction succeeded
	if receipt.Status == 0 {
		return nil, receipt, fmt.Errorf("update merchant registry transaction failed: %s", tx.Hash().Hex())
	}

	// Parse the UpdateMerchantRegistry event
	for _, log := range receipt.Logs {
		event, err := p.Client.PaymentProcessor.ParseMerchantRegistryUpdated(*log)
		if err == nil {
			return event, receipt, nil
		}
	}

	// if no event found, still return success but with warning
	return nil, receipt, fmt.Errorf("update merchant registry succeeded but event not found in logs")
}

// SetTokenSupport updates the status of a token (enable/disable support)
func (p *Platform) SetTokenSupport(ctx context.Context, tokenAddress common.Address, status *big.Int) (*contracts.PaymentProcessorTokenSupportUpdated, *types.Receipt, error) {
	// Validate Input
	if err := utils.ValidateSetTokenSupportInput(tokenAddress, status); err != nil {
		return nil, nil, err
	}

	// Create transaction options
	auth, err := bind.NewKeyedTransactorWithChainID(p.Client.PrivateKey, p.Client.ChainID)
	if err != nil {
		return nil, nil, err
	}

	// Call the contract
	tx, err := p.Client.PaymentProcessor.SetTokenSupport(auth, tokenAddress, status)
	if err != nil {
		return nil, nil, err
	}

	// Wait for the transaction to be mined
	receipt, err := bind.WaitMined(ctx, p.Client.EthClient, tx)
	if err != nil {
		return nil, nil, err
	}

	// Check if transaction succeeded
	if receipt.Status == 0 {
		return nil, receipt, fmt.Errorf("set token support transaction failed: %s", tx.Hash().Hex())
	}

	// Parse the TokenSupportUpdated event
	for _, log := range receipt.Logs {
		event, err := p.Client.PaymentProcessor.ParseTokenSupportUpdated(*log)
		if err == nil {
			return event, receipt, nil
		}
	}

	// if no event found, still return success but with warning
	return nil, receipt, fmt.Errorf("set token support succeeded but event not found in logs")
}

// UpdateProtocolAddress updates the protocol address for the platform
func (p *Platform) UpdateProtocolAddress(ctx context.Context, what [32]byte, newAddress common.Address) (*contracts.PaymentProcessorProtocolAddressUpdated, *types.Receipt, error) {
	// Create transaction options
	auth, err := bind.NewKeyedTransactorWithChainID(p.Client.PrivateKey, p.Client.ChainID)
	if err != nil {
		return nil, nil, err
	}

	// Call the contract
	tx, err := p.Client.PaymentProcessor.UpdateProtocolAddress(auth, what, newAddress)
	if err != nil {
		return nil, nil, err
	}

	// Wait for the transaction to be mined
	receipt, err := bind.WaitMined(ctx, p.Client.EthClient, tx)
	if err != nil {
		return nil, nil, err
	}

	// Check if transaction succeeded
	if receipt.Status == 0 {
		return nil, receipt, fmt.Errorf("update protocol address transaction failed: %s", tx.Hash().Hex())
	}

	// Parse the ProtocolAddressUpdated event
	for _, log := range receipt.Logs {
		event, err := p.Client.PaymentProcessor.ParseProtocolAddressUpdated(*log)
		if err == nil {
			return event, receipt, nil
		}
	}

	// if no event found, still return success but with warning
	return nil, receipt, fmt.Errorf("update protocol address succeeded but event not found in logs")
}

// UpdateOrderExpirationTime updates the protocol address for the platform
func (p *Platform) UpdateOrderExpirationTime(ctx context.Context, newExpirationTime *big.Int) (*contracts.PaymentProcessorOrderExpirationTimeUpdated, *types.Receipt, error) {
	// Create transaction options
	auth, err := bind.NewKeyedTransactorWithChainID(p.Client.PrivateKey, p.Client.ChainID)
	if err != nil {
		return nil, nil, err
	}

	// Call the contract
	tx, err := p.Client.PaymentProcessor.UpdateOrderExpirationTime(auth, newExpirationTime)
	if err != nil {
		return nil, nil, err
	}

	// Wait for the transaction to be mined
	receipt, err := bind.WaitMined(ctx, p.Client.EthClient, tx)
	if err != nil {
		return nil, nil, err
	}

	// Check if transaction succeeded
	if receipt.Status == 0 {
		return nil, receipt, fmt.Errorf("update order expiration time transaction failed: %s", tx.Hash().Hex())
	}

	// Parse the UpdateOrderExpiration event
	for _, log := range receipt.Logs {
		event, err := p.Client.PaymentProcessor.ParseOrderExpirationTimeUpdated(*log)
		if err == nil {
			return event, receipt, nil
		}
	}

	// if no event found, still return success but with warning
	return nil, receipt, fmt.Errorf("update order expiration time update succeeded but event not found in logs")
}

// GetPlatformTokenBalance gets the platform's token balance for a specific token
func (p *Platform) GetPlatformTokenBalance(ctx context.Context, platformWalletAddress common.Address, tokenAddress common.Address) (*big.Int, error) {
	// Validate inputs
	if err := utils.ValidateGetPlatformTokenBalanceInputs(platformWalletAddress, tokenAddress); err != nil {
		return nil, err
	}

	// Create call options
	opts := &bind.CallOpts{Context: ctx}

	// Create ERC20 token instance
	token, err := contracts.NewERC20(tokenAddress, p.Client.EthClient)
	if err != nil {
		return nil, fmt.Errorf("failed to instantiate ERC20: %w", err)
	}

	// Call the token contract to get balance
	balance, err := token.BalanceOf(opts, platformWalletAddress)
	if err != nil {
		return nil, err
	}

	return balance, nil
}

// GetContractTokenBalance gets the PaymentProcessor contract's token balance for a specific token
func (p *Platform) GetContractTokenBalance(ctx context.Context, tokenAddress common.Address) (*big.Int, error) {
	// Validate inputs
	if err := utils.ValidateAddress(tokenAddress, "tokenAddress"); err != nil {
		return nil, err
	}

	// Create call options
	opts := &bind.CallOpts{Context: ctx}

	// Create ERC20 token instance
	token, err := contracts.NewERC20(tokenAddress, p.Client.EthClient)
	if err != nil {
		return nil, fmt.Errorf("failed to instantiate ERC20: %w", err)
	}

	// Get the PaymentProcessor contract address
	contractAddress := p.Client.PaymentProcessorAddress

	// Call the token contract to get contract's balance
	balance, err := token.BalanceOf(opts, contractAddress)
	if err != nil {
		return nil, err
	}

	return balance, nil
}
