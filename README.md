# StableBase Go SDK

A Go SDK for interacting with StableBase payment smart contracts on EVM-compatible blockchains. StableBase enables stablecoin payments between merchants, customers, and platforms with automatic fee distribution.

[![Go Reference](https://pkg.go.dev/badge/github.com/Dbriane208/stablebase-go-sdk.svg)](https://pkg.go.dev/github.com/Dbriane208/stablebase-go-sdk)
[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)

---

## 📚 Table of Contents

- [Overview](#overview)
- [Installation](#installation)
- [Quick Start](#quick-start)
- [SDK Architecture](#sdk-architecture)
- [Network Configuration](#network-configuration)
- [API Reference](#api-reference)
  - [Client](#client)
  - [Merchant Module](#merchant-module)
  - [Order Module](#order-module)
  - [Platform Module](#platform-module)
  - [QR Code Module](#qr-code-module)
- [Payment Flow](#payment-flow)
- [Contract Addresses](#contract-addresses)
- [Error Handling](#error-handling)
- [Testing](#testing)
- [Contributing](#contributing)
- [License](#license)

---

## Overview

StableBase is a blockchain-based payment system that enables stablecoin (USDC) payments between three types of users:

| User Type     | Description                   | Primary Actions              |
| ------------- | ----------------------------- | ---------------------------- |
| **Merchants** | Businesses accepting payments | Register once, receive funds |
| **Customers** | Users making payments         | Create orders, pay orders    |
| **Platform**  | System administrator          | Settle orders, collect fees  |

### How It Works

1. **Merchant** registers on the blockchain (one-time setup)
2. **Customer** creates an order for a merchant
3. **Customer** pays the order (funds held in smart contract)
4. **Platform** settles the order → **98%** goes to merchant, **2%** platform fee

---

## Installation

```bash
go get github.com/Dbriane208/stablebase-go-sdk
```

### Requirements

- Go 1.21 or higher
- Access to an EVM-compatible network (Base Sepolia or Polygon Amoy testnet recommended)

---

## Quick Start

### Step 1: Set Up Your Network Configuration

```go
package main

import (
    "math/big"

    "github.com/Dbriane208/stablebase-go-sdk/client"
    "github.com/ethereum/go-ethereum/common"
)

// Configure the network you want to connect to
var baseSepoliaConfig = client.NetworkConfig{
    NetworkName:             "base-sepolia",
    ChainID:                 big.NewInt(84532),
    RPCURL:                  "https://sepolia.base.org",
    USDCAddress:             common.HexToAddress("0x036CbD53842c5426634e7929541eC2318f3dCF7e"),
    PaymentProcessorAddress: common.HexToAddress("0x7c39408AC96a1b9a2722056eDE90b54D2B260380"),
    MerchantRegistryAddress: common.HexToAddress("0x93e93Dfa36C87De32B9118CA5D9BAd1Db892002d"),
    ExplorerURL:             "https://sepolia.basescan.org",
}
```

### Step 2: Create a Client

```go
import (
    "os"
    "log"

    "github.com/Dbriane208/stablebase-go-sdk/client"
)

func main() {
    // Get your private key from environment variable (never hardcode!)
    privateKey := os.Getenv("PRIVATE_KEY")

    // Create a new client
    c, err := client.NewClient(privateKey, baseSepoliaConfig)
    if err != nil {
        log.Fatal("Failed to create client:", err)
    }

    // Client is ready to use!
    log.Println("Connected! Wallet address:", c.WalletAddress.Hex())
}
```

> ⚠️ **Security Warning**: Never hardcode your private key. Always use environment variables or secure key management.

---

## SDK Architecture

```
stablebase-go-sdk/
├── client/          # Core client - blockchain connection
├── merchant/        # Merchant operations (register, update, refund)
├── order/           # Customer order operations (create, pay, cancel)
├── platform/        # Platform admin operations (settle, configure)
├── qrcode/          # QR code generation for payments
├── contracts/       # Auto-generated smart contract bindings
└── utils/           # Validation helpers and utilities
```

### Module Overview

| Module     | Who Uses It | Purpose                    |
| ---------- | ----------- | -------------------------- |
| `client`   | Everyone    | Connect to blockchain      |
| `merchant` | Merchants   | Registration & management  |
| `order`    | Customers   | Create and pay orders      |
| `platform` | Admins      | Settlement & configuration |
| `qrcode`   | Merchants   | Generate payment QR codes  |

---

## Network Configuration

### Supported Networks

| Network      | Chain ID | Status         |
| ------------ | -------- | -------------- |
| Base Sepolia | 84532    | ✅ Testnet     |
| Polygon Amoy | 80002    | ✅ Testnet     |
| Base Mainnet | 8453     | 🔜 Coming Soon |

### Pre-configured Networks

<details>
<summary><b>Base Sepolia Configuration</b></summary>

```go
var BaseSepoliaConfig = client.NetworkConfig{
    NetworkName:             "base-sepolia",
    ChainID:                 big.NewInt(84532),
    RPCURL:                  "https://sepolia.base.org",
    USDCAddress:             common.HexToAddress("0x036CbD53842c5426634e7929541eC2318f3dCF7e"),
    PaymentProcessorAddress: common.HexToAddress("0x7c39408AC96a1b9a2722056eDE90b54D2B260380"),
    MerchantRegistryAddress: common.HexToAddress("0x93e93Dfa36C87De32B9118CA5D9BAd1Db892002d"),
    ExplorerURL:             "https://sepolia.basescan.org",
}
```

</details>

<details>
<summary><b>Polygon Amoy Configuration</b></summary>

```go
var PolygonAmoyConfig = client.NetworkConfig{
    NetworkName:             "polygon-amoy",
    ChainID:                 big.NewInt(80002),
    RPCURL:                  "https://rpc-amoy.polygon.technology",
    USDCAddress:             common.HexToAddress("0x41E94Eb019C0762f9Bfcf9Fb1E58725BfB0e7582"),
    PaymentProcessorAddress: common.HexToAddress("0x3B08Be115E1672cE8A6618D932a97B2Cc251d853"),
    MerchantRegistryAddress: common.HexToAddress("0xE664919f8a195d44c8a137C71cBeb967A71eD3DF"),
    ExplorerURL:             "https://amoy.polygonscan.com",
}
```

</details>

---

## API Reference

### Client

The client is the foundation of the SDK. It manages blockchain connections and provides access to all modules.

#### Creating a Client

```go
import "github.com/Dbriane208/stablebase-go-sdk/client"

c, err := client.NewClient(privateKeyHex, networkConfig)
```

**Parameters:**
| Parameter | Type | Description |
|-----------|------|-------------|
| `privateKeyHex` | `string` | Your wallet's private key (without `0x` prefix) |
| `networkConfig` | `NetworkConfig` | Network configuration struct |

**Returns:** `*Client` - A client instance with all modules available.

**Client Properties:**

```go
type Client struct {
    EthClient     *ethclient.Client  // Raw Ethereum client
    PrivateKey    *ecdsa.PrivateKey  // For signing transactions
    WalletAddress common.Address     // Your wallet address
    ChainID       *big.Int           // Network chain ID

    // Contract instances (ready to use)
    PaymentProcessor *contracts.PaymentProcessor
    MerchantRegistry *contracts.MerchantRegistry
}
```

---

### Merchant Module

The merchant module handles merchant registration and management. Import it like this:

```go
import "github.com/Dbriane208/stablebase-go-sdk/merchant"

// Create merchant instance
m := merchant.New(c) // 'c' is your client
```

#### Register a Merchant

Register a new merchant on the blockchain. This is a **one-time** operation.

```go
event, merchantId, receipt, err := m.RegisterMerchant(
    ctx,
    payoutWallet,  // Where you'll receive payments
    metadataURI,   // URL to your merchant info JSON
)
```

**Parameters:**
| Parameter | Type | Description |
|-----------|------|-------------|
| `ctx` | `context.Context` | Context for cancellation |
| `payoutWallet` | `common.Address` | Wallet address to receive payments |
| `metadataURI` | `string` | URL to merchant metadata JSON |

**Returns:**
| Return | Type | Description |
|--------|------|-------------|
| `event` | `*MerchantRegistryMerchantRegistered` | Event data from blockchain |
| `merchantId` | `[32]byte` | Your unique merchant ID |
| `receipt` | `*types.Receipt` | Transaction receipt |
| `err` | `error` | Error if any |

**Example:**

```go
ctx := context.Background()
payoutWallet := common.HexToAddress("0xYourPayoutWallet")
metadataURI := "https://yoursite.com/merchant-info.json"

event, merchantId, receipt, err := m.RegisterMerchant(ctx, payoutWallet, metadataURI)
if err != nil {
    log.Fatal("Registration failed:", err)
}

fmt.Printf("✅ Registered! Merchant ID: 0x%x\n", merchantId)
fmt.Printf("📜 Transaction: %s\n", receipt.TxHash.Hex())
```

---

#### Update Merchant

Update your merchant details (payout wallet or metadata).

```go
event, receipt, err := m.UpdateMerchant(ctx, merchantId, newPayoutWallet, newMetadataURI)
```

**Parameters:**
| Parameter | Type | Description |
|-----------|------|-------------|
| `ctx` | `context.Context` | Context for cancellation |
| `merchantId` | `[32]byte` | Your merchant ID |
| `newPayoutWallet` | `common.Address` | New payout wallet address |
| `newMetadataURI` | `string` | New metadata URL |

---

#### Get Merchant Info

Retrieve merchant information from the blockchain.

```go
info, err := m.GetMerchantInfo(ctx, merchantId)
```

**Returns:** `*IMerchantRegistryMerchant` containing:

- Payout wallet address
- Metadata URI
- Verification status
- Registration timestamp

---

#### Check Merchant Verification

```go
// Simple boolean check
isVerified, err := m.IsMerchantVerified(ctx, merchantId)

// Get detailed status (0=Pending, 1=Verified, 2=Rejected, 3=Suspended)
status, err := m.GetMerchantVerificationStatus(ctx, merchantId)
```

---

#### Get Merchant Token Balance

Check how much of a specific token the merchant has.

```go
balance, err := m.GetMerchantTokenBalance(ctx, merchantWalletAddress, tokenAddress)
```

---

#### Refund an Order

Merchants can refund paid orders back to the customer.

```go
event, receipt, err := m.RefundOrder(ctx, orderId)
```

---

### Order Module

The order module handles customer payment operations. This is where most of the action happens!

```go
import "github.com/Dbriane208/stablebase-go-sdk/order"

// Create order instance
o := order.New(c) // 'c' is your client
```

#### Complete Payment Flow Example

Here's a full example of how a customer makes a payment:

```go
package main

import (
    "context"
    "log"
    "math/big"

    "github.com/Dbriane208/stablebase-go-sdk/client"
    "github.com/Dbriane208/stablebase-go-sdk/order"
    "github.com/ethereum/go-ethereum/common"
)

func main() {
    // Setup (see Quick Start for full setup)
    c, _ := client.NewClient(privateKey, networkConfig)
    o := order.New(c)
    ctx := context.Background()

    // Payment details
    merchantId := [32]byte{...}  // The merchant you're paying
    tokenAddress := common.HexToAddress("0x036CbD53842c5426634e7929541eC2318f3dCF7e") // USDC
    amount := big.NewInt(5000000) // 5 USDC (6 decimals)

    // Step 1: Approve the PaymentProcessor to spend your USDC
    _, err := o.ApproveToken(ctx, tokenAddress, c.PaymentProcessorAddress, amount)
    if err != nil {
        log.Fatal("Approval failed:", err)
    }
    log.Println("✅ Token approved")

    // Step 2: Create the order
    event, orderId, _, err := o.CreateOrder(ctx, merchantId, tokenAddress, amount, "ipfs://order-metadata")
    if err != nil {
        log.Fatal("Order creation failed:", err)
    }
    log.Printf("✅ Order created! ID: 0x%x\n", orderId)

    // Step 3: Pay the order
    payEvent, _, err := o.PayOrder(ctx, orderId)
    if err != nil {
        log.Fatal("Payment failed:", err)
    }
    log.Println("✅ Payment complete!")
}
```

---

#### Approve Token

Before creating orders, you must approve the PaymentProcessor contract to spend your tokens.

```go
receipt, err := o.ApproveToken(ctx, tokenAddress, spenderAddress, amount)
```

**Parameters:**
| Parameter | Type | Description |
|-----------|------|-------------|
| `tokenAddress` | `common.Address` | Token contract (e.g., USDC) |
| `spenderAddress` | `common.Address` | PaymentProcessor contract address |
| `amount` | `*big.Int` | Amount to approve (in smallest units) |

> 💡 **Tip:** USDC has 6 decimals. To approve 10 USDC: `big.NewInt(10_000_000)`

---

#### Create Order

Create a new payment order for a merchant.

```go
event, orderId, receipt, err := o.CreateOrder(ctx, merchantId, tokenAddress, amount, metadataURI)
```

**Parameters:**
| Parameter | Type | Description |
|-----------|------|-------------|
| `merchantId` | `[32]byte` | The merchant's ID |
| `tokenAddress` | `common.Address` | Payment token (e.g., USDC) |
| `amount` | `*big.Int` | Payment amount |
| `metadataURI` | `string` | Order metadata URL |

**Returns:**
| Return | Type | Description |
|--------|------|-------------|
| `event` | `*PaymentProcessorOrderCreated` | Order creation event |
| `orderId` | `[32]byte` | Unique order ID |
| `receipt` | `*types.Receipt` | Transaction receipt |

---

#### Pay Order

Execute payment for a created order. Funds are transferred to the smart contract.

```go
event, receipt, err := o.PayOrder(ctx, orderId)
```

---

#### Cancel Order

Cancel an unpaid order (only the order creator can cancel).

```go
event, receipt, err := o.CancelOrder(ctx, orderId)
```

---

#### Get Order Information

```go
// Get full order details
orderInfo, err := o.GetOrder(ctx, orderId)

// Get just the status (0=Created, 1=Paid, 2=Settled, 3=Cancelled, 4=Refunded, 5=Expired)
status, err := o.GetOrderStatus(ctx, orderId)

// Check if order has expired
isExpired, err := o.IsOrderExpired(ctx, orderId)

// Get when order was created (Unix timestamp)
createdAt, err := o.GetOrderCreatedAt(ctx, orderId)

// Get time remaining before expiration (in seconds)
remainingTime, err := o.GetOrderRemainingTime(ctx, orderId)
```

---

### Platform Module

The platform module is for **administrators only**. It handles order settlement and system configuration.

```go
import "github.com/Dbriane208/stablebase-go-sdk/platform"

p := platform.New(c)
```

#### Settle Order

Distribute funds from a paid order: 98% to merchant, 2% platform fee.

```go
event, receipt, err := p.SettleOrder(ctx, orderId)
```

**Example:**

```go
event, receipt, err := p.SettleOrder(ctx, orderId)
if err != nil {
    log.Fatal("Settlement failed:", err)
}

fmt.Printf("✅ Order settled!\n")
fmt.Printf("   Merchant received: %s\n", event.MerchantAmount)
fmt.Printf("   Platform fee: %s\n", event.PlatformFee)
```

---

#### Refund Order

Platform can also refund orders back to customers.

```go
event, receipt, err := p.RefundOrder(ctx, orderId)
```

---

#### Emergency Withdraw

Withdraw funds from the contract in emergency situations (admin only).

```go
event, receipt, err := p.EmergencyWithdraw(ctx, tokenAddress, receiverAddress, amount)
```

---

#### Enable/Disable Emergency Withdrawal

```go
event, receipt, err := p.SetEmergencyWithdrawalEnabled(ctx, true)  // Enable
event, receipt, err := p.SetEmergencyWithdrawalEnabled(ctx, false) // Disable
```

---

#### Update Merchant Registry

Point to a new merchant registry contract.

```go
event, receipt, err := p.UpdateMerchantRegistry(ctx, newRegistryAddress)
```

---

#### Set Token Support

Enable or disable support for a token.

```go
// Status: 0=Disabled, 1=Enabled
event, receipt, err := p.SetTokenSupport(ctx, tokenAddress, big.NewInt(1))
```

---

#### Update Order Expiration Time

Change how long orders remain valid before expiring.

```go
// Time in seconds (e.g., 3600 = 1 hour)
event, receipt, err := p.UpdateOrderExpirationTime(ctx, big.NewInt(3600))
```

---

#### Check Balances

```go
// Get platform's token balance
balance, err := p.GetPlatformTokenBalance(ctx, platformWallet, tokenAddress)

// Get the contract's token balance
balance, err := p.GetContractTokenBalance(ctx, tokenAddress)
```

---

### QR Code Module

Generate QR codes for payment information.

```go
import "github.com/Dbriane208/stablebase-go-sdk/qrcode"
```

#### Generate QR Code

```go
data := qrcode.PaymentQRData{
    MerchantID:   "0x1234...",
    Amount:       "5000000",      // 5 USDC
    TokenAddress: "0x036CbD53842c5426634e7929541eC2318f3dCF7e",
    OrderID:      "0xabcd...",
    ChainID:      "84532",
}

// Generate PNG image (256x256 pixels)
pngBytes, err := qrcode.Generate(data, 256)
if err != nil {
    log.Fatal(err)
}

// Save to file
os.WriteFile("payment-qr.png", pngBytes, 0644)
```

#### Parse QR Code Data

```go
jsonString := `{"merchant_id":"0x123...","amount":"5000000",...}`

data, err := qrcode.Parse(jsonString)
if err != nil {
    log.Fatal(err)
}

fmt.Println("Merchant:", data.MerchantID)
fmt.Println("Amount:", data.Amount)
```

---

## Payment Flow

Here's a visual representation of how payments work:

```
┌──────────────┐                              ┌──────────────┐
│   MERCHANT   │                              │   CUSTOMER   │
└──────┬───────┘                              └──────┬───────┘
       │                                             │
       │ 1. RegisterMerchant()                       │
       │    → gets merchantId                        │
       │                                             │
       │ 2. Generate & display QR code               │
       │    (contains merchantId)                    │
       │                                             │
       │                                    3. Scan QR code
       │                                       │
       │                                    4. ApproveToken()
       │                                       │
       │                                    5. CreateOrder()
       │                                       → gets orderId
       │                                       │
       │                                    6. PayOrder()
       │                                       → USDC → Contract
       │                                             │
       │                              ┌──────────────┴──────────────┐
       │                              │          PLATFORM           │
       │                              └──────────────┬──────────────┘
       │                                             │
       │                                    7. SettleOrder()
       │                                       → 98% to Merchant
       │◄──────────────────────────────────────→ 2% to Platform
       │
   Receives funds!
```

### Order States

| Status    | Value | Description                              |
| --------- | ----- | ---------------------------------------- |
| Created   | 0     | Order created, awaiting payment          |
| Paid      | 1     | Customer has paid, awaiting settlement   |
| Settled   | 2     | Funds distributed to merchant & platform |
| Cancelled | 3     | Order cancelled by customer              |
| Refunded  | 4     | Payment refunded to customer             |
| Expired   | 5     | Order expired without payment            |

---

## Contract Addresses

### Base Sepolia (Testnet)

| Contract          | Address                                      |
| ----------------- | -------------------------------------------- |
| Payment Processor | `0x7c39408AC96a1b9a2722056eDE90b54D2B260380` |
| Merchant Registry | `0x93e93Dfa36C87De32B9118CA5D9BAd1Db892002d` |
| USDC              | `0x036CbD53842c5426634e7929541eC2318f3dCF7e` |

### Polygon Amoy (Testnet)

| Contract          | Address                                      |
| ----------------- | -------------------------------------------- |
| Payment Processor | `0x3B08Be115E1672cE8A6618D932a97B2Cc251d853` |
| Merchant Registry | `0xE664919f8a195d44c8a137C71cBeb967A71eD3DF` |
| USDC              | `0x41E94Eb019C0762f9Bfcf9Fb1E58725BfB0e7582` |

---

## Error Handling

The SDK uses custom error types for better error handling:

### ValidationError

Returned when input validation fails.

```go
type ValidationError struct {
    Field   string  // Which field failed
    Message string  // What went wrong
}
```

**Example:**

```go
_, _, _, err := o.CreateOrder(ctx, [32]byte{}, tokenAddress, amount, "")
if err != nil {
    if valErr, ok := err.(*utils.ValidationError); ok {
        fmt.Printf("Validation error in '%s': %s\n", valErr.Field, valErr.Message)
    }
}
```

### TransactionError

Returned when a blockchain transaction fails.

```go
type TransactionError struct {
    TxHash  string  // Transaction hash
    Message string  // Error description
}
```

### Common Error Scenarios

| Error                                  | Cause                   | Solution                      |
| -------------------------------------- | ----------------------- | ----------------------------- |
| "merchant ID cannot be empty"          | Empty merchant ID       | Provide valid merchant ID     |
| "amount must be greater than zero"     | Zero or negative amount | Use positive amount           |
| "payout wallet address cannot be zero" | Invalid address         | Use valid Ethereum address    |
| "transaction failed"                   | On-chain error          | Check transaction on explorer |

---

## Testing

Run the test suite:

```bash
# Run all tests
go test ./...

# Run tests with verbose output
go test -v ./...

# Run tests with coverage
go test -cover ./...

# Run tests for a specific package
go test -v ./merchant/...
go test -v ./order/...
go test -v ./platform/...
```

---

## Contributing

Contributions are welcome! Here's how to get started:

1. **Fork** the repository
2. **Create** your feature branch
   ```bash
   git checkout -b feature/amazing-feature
   ```
3. **Write tests** for your changes
4. **Run tests** to make sure everything passes
   ```bash
   go test ./...
   ```
5. **Commit** your changes
   ```bash
   git commit -m 'Add amazing feature'
   ```
6. **Push** to the branch
   ```bash
   git push origin feature/amazing-feature
   ```
7. **Open** a Pull Request

### Development Guidelines

- Follow Go best practices and conventions
- Add tests for new functionality
- Update documentation for API changes
- Keep commits focused and atomic

---

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

---

## Resources

- [Base Sepolia Explorer](https://sepolia.basescan.org) - View transactions on Base Sepolia
- [Polygon Amoy Explorer](https://amoy.polygonscan.com) - View transactions on Polygon Amoy
- [Go Ethereum Documentation](https://geth.ethereum.org/docs) - Learn about go-ethereum
- [USDC Documentation](https://developers.circle.com/stablecoins/docs) - Learn about USDC

---

## Support

Having issues? Here are some resources:

- 📖 Check this documentation thoroughly
- 🔍 Search [existing issues](https://github.com/Dbriane208/stablebase-go-sdk/issues)
- 🐛 [Report a bug](https://github.com/Dbriane208/stablebase-go-sdk/issues/new)
- 💡 [Request a feature](https://github.com/Dbriane208/stablebase-go-sdk/issues/new)

---

<p align="center">
  Made with ❤️ for the Web3 community
</p>
