package client

import (
	"crypto/ecdsa"
	"math/big"

	"github.com/Dbriane208/stablebase-go-sdk/contracts"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
)

type Client struct {
	EthClient     *ethclient.Client // Connection to blockchain
	PrivateKey    *ecdsa.PrivateKey // For signing transactions
	WalletAddress common.Address    // Derived from private key
	ChainID       *big.Int          // Network chain ID

	// Contract instances
	PaymentProcessor *contracts.PaymentProcessor
	MerchantRegistry *contracts.MerchantRegistry

	// Contract addresses
	PaymentProcessorAddress common.Address
	MerchantRegistryAddress common.Address
}

type NetworkConfig struct {
	NetworkName string
	ChainID     *big.Int
	RPCURL      string
	USDCAddress common.Address
	ExplorerURL string

	// Contract address
	PaymentProcessorAddress common.Address
	MerchantRegistryAddress common.Address
}

func NewClient(privateKeyHex string, network NetworkConfig) (*Client, error) {
	// Connect to the blockchain using network.RPCURL
	ethClient, err := ethclient.Dial(network.RPCURL)
	if err != nil {
		return nil, err
	}

	// Parse the private key from hex string
	privateKey, err := crypto.HexToECDSA(privateKeyHex)
	if err != nil {
		return nil, err
	}

	// Derive wallet address from private key
	walletAddress := crypto.PubkeyToAddress(privateKey.PublicKey)

	// Create contract instances
	paymentProcessor, err := contracts.NewPaymentProcessor(network.PaymentProcessorAddress, ethClient)
	if err != nil {
		return nil, err
	}
	merchantRegistry, err := contracts.NewMerchantRegistry(network.MerchantRegistryAddress, ethClient)
	if err != nil {
		return nil, err
	}

	// Return the Client struct with all fields filled
	return &Client{
		EthClient:               ethClient,
		PrivateKey:              privateKey,
		WalletAddress:           walletAddress,
		ChainID:                 network.ChainID,
		PaymentProcessor:        paymentProcessor,
		MerchantRegistry:        merchantRegistry,
		PaymentProcessorAddress: network.PaymentProcessorAddress,
		MerchantRegistryAddress: network.MerchantRegistryAddress,
	}, nil
}
