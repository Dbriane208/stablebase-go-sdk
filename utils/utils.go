package utils

import (
	"fmt"
	"math/big"
	"strings"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
)

// ValidationError represents an input validation error
type ValidationError struct {
	Field   string
	Message string
}

func (e *ValidationError) Error() string {
	return fmt.Sprintf("%s: %s", e.Field, e.Message)
}

// TransactionError represents a blockchain transaction error
type TransactionError struct {
	TxHash  string
	Message string
}

func (e *TransactionError) Error() string {
	return fmt.Sprintf("transaction failed: %s - %s", e.TxHash, e.Message)
}

// ValidateCreateOrderInputs validates inputs for order creation
func ValidateCreateOrderInputs(merchantId [32]byte, tokenAddress common.Address, amount *big.Int, metadataURI string) error {
	if merchantId == [32]byte{} {
		return &ValidationError{
			Field:   "merchantId",
			Message: "merchant ID cannot be empty",
		}
	}

	if tokenAddress == (common.Address{}) {
		return &ValidationError{
			Field:   "tokenAddress",
			Message: "tokenAddress cannot be zero",
		}
	}

	if amount == nil || amount.Sign() <= 0 {
		return &ValidationError{
			Field:   "amount",
			Message: "amount must be greater than zero",
		}
	}

	trimmed := strings.TrimSpace(metadataURI)
	if len(trimmed) == 0 {
		return &ValidationError{
			Field:   "metadataURI",
			Message: "metadata URI cannot be empty",
		}
	}

	return nil
}

// ValidateOrderId
func ValidateOrderId(orderId [32]byte) error {
	if orderId == [32]byte{} {
		return &ValidationError{
			Field:   "orderId",
			Message: "order ID cannot be empty",
		}
	}

	return nil
}

// ValidateRegisterMerchantInputs validates inputs for merchant registration
func ValidateRegisterMerchantInputs(payoutWallet common.Address, metadataURI string) error {
	if payoutWallet == (common.Address{}) {
		return &ValidationError{
			Field:   "payoutWallet",
			Message: "payout wallet address cannot be zero",
		}
	}

	trimmed := strings.TrimSpace(metadataURI)
	if len(trimmed) == 0 {
		return &ValidationError{
			Field:   "metadataURI",
			Message: "metadata URI cannot be empty",
		}
	}

	return nil
}

// ValidateUpdateMerchantInputs validates inputs for merchant update
func ValidateUpdateMerchantInputs(merchantId [32]byte, payoutWallet common.Address, metadataURI string) error {
	if merchantId == [32]byte{} {
		return &ValidationError{
			Field:   "merchantId",
			Message: "merchant ID cannot be empty",
		}
	}

	if payoutWallet == (common.Address{}) {
		return &ValidationError{
			Field:   "payoutWallet",
			Message: "payout wallet address cannot be zero",
		}
	}

	trimmed := strings.TrimSpace(metadataURI)
	if len(trimmed) == 0 {
		return &ValidationError{
			Field:   "metadataURI",
			Message: "metadata URI cannot be empty",
		}
	}

	return nil
}

// ValidateMerchatTokenBalanceInputs
func ValidateMerchatTokenBalanceInputs(merchantWalletAddress common.Address, tokenAddress common.Address) error {
	if merchantWalletAddress == (common.Address{}) {
		return &ValidationError{
			Field:   "merchant wallet address",
			Message: "merchant wallet address cannot be zero",
		}
	}

	if tokenAddress == (common.Address{}) {
		return &ValidationError{
			Field:   "token address",
			Message: "token address cannot be zero",
		}
	}

	return nil
}

// CheckTransactionStatus checks if a transaction was successful
func CheckTransactionStatus(receipt *types.Receipt) error {
	if receipt.Status == types.ReceiptStatusFailed {
		return &TransactionError{
			TxHash:  receipt.TxHash.Hex(),
			Message: "transaction failed",
		}
	}
	return nil
}

// ValidateMerchantId checks if a merchant ID is valid (not empty)
func ValidateMerchantId(merchantId [32]byte) error {
	if merchantId == [32]byte{} {
		return &ValidationError{
			Field:   "merchantId",
			Message: "merchant ID cannot be empty",
		}
	}
	return nil
}

// ValidateAddress checks if an Ethereum address is valid (not zero)
func ValidateAddress(address common.Address, fieldName string) error {
	if address == (common.Address{}) {
		return &ValidationError{
			Field:   fieldName,
			Message: fmt.Sprintf("%s cannot be zero address", fieldName),
		}
	}
	return nil
}

// ValidateRefundOrderInputs validates inputs for order refund
func ValidateRefundOrderInputs(orderId [32]byte) error {
	if orderId == [32]byte{} {
		return &ValidationError{
			Field:   "orderId",
			Message: "order ID cannot be empty",
		}
	}

	return nil
}

// ValidateEmergencyWithdrawInputs validates input for emergency withdrawal
func ValidateEmergencyWithdrawInputs(tokenAddress common.Address, receiverAddress common.Address, amount *big.Int) error {
	if tokenAddress == (common.Address{}) {
		return &ValidationError{
			Field:   "tokenAddress",
			Message: "token address cannot be empty",
		}
	}

	if receiverAddress == (common.Address{}) {
		return &ValidationError{
			Field:   "receiverAddress",
			Message: "receiver address cannot be empty",
		}
	}

	if amount == nil || amount.Sign() <= 0 {
		return &ValidationError{
			Field:   "amount",
			Message: "amount must be greater than zero",
		}
	}

	return nil
}

// ValidateUpdateMerchantRegistry validates the registry address
func ValidateUpdateMerchantRegistryInput(newRegistryAddress common.Address) error {
	if newRegistryAddress == (common.Address{}) {
		return &ValidationError{
			Field:   "newRegistryAddress",
			Message: "new registry address should not be empty",
		}
	}

	return nil
}

// ValidateSetTokenSupportInput validates whether a token is supported
func ValidateSetTokenSupportInput(tokenAddress common.Address, status *big.Int) error {
	if tokenAddress == (common.Address{}) {
		return &ValidationError{
			Field:   "tokenAddress",
			Message: "token address cannot be empty",
		}
	}

	if status == nil {
		return &ValidationError{
			Field:   "status",
			Message: "status cannot be nil",
		}
	}

	if status.Cmp(big.NewInt(0)) != 0 && status.Cmp(big.NewInt(1)) != 0 {
		return &ValidationError{
			Field:   "status",
			Message: "unsupported token status",
		}
	}

	return nil
}

// ValidateGetPlatfromTokenBalanceInputs
func ValidateGetPlatformTokenBalanceInputs(platformWalletAddress common.Address, tokenAddress common.Address) error {
	if platformWalletAddress == (common.Address{}) {
		return &ValidationError{
			Field:   "platform wallet address",
			Message: "platform wallet address cannot be zero",
		}
	}

	if tokenAddress == (common.Address{}) {
		return &ValidationError{
			Field:   "token address",
			Message: "token address cannot be zero",
		}
	}

	return nil
}
