# StableBase Go SDK

A Go SDK for interacting with StableBase payment smart contracts on EVM-compatible blockchains. StableBase enables stablecoin payments between merchants, customers, and platforms with automatic fee distribution.

[![Go Reference](https://pkg.go.dev/badge/github.com/Dbriane208/stablebase-go-sdk.svg)](https://pkg.go.dev/github.com/Dbriane208/stablebase-go-sdk)
[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)

## Overview

StableBase enables three types of users:

- **Merchants** - Businesses accepting stablecoin payments (minimal interaction)
- **Customers** - Users making payments (main actors)
- **Platform** - Settles orders and collects fees

### Key Concept: Customer-Initiated Payments

The SDK is designed around customer-initiated payments where:

- Customers create orders
- Customers pay orders
- Platform settles (distributes funds: 98% to merchant, 2% platform fee)

## Installation

```bash
go get github.com/Dbriane208/stablebase-go-sdk
```

## Quick Start

### Merchant Registration (One-time Setup)

```go
package main

import (
    "fmt"
    "os"

    sdk "github.com/Dbriane208/stablebase-go-sdk"
    "github.com/Dbriane208/stablebase-go-sdk/qrcode"
)

func main() {
    // Initialize SDK with merchant's private key
    client, _ := sdk.NewClient(sdk.Config{
        Network:    "base-sepolia",
        PrivateKey: os.Getenv("MERCHANT_PRIVATE_KEY"),
    })

    // Register as a merchant (one-time)
    merchantId, _ := client.Merchant.Register(
        payoutWalletAddress,  // where funds will be sent
        "https://merchant.com/info.json",
    )
    fmt.Println("Your Merchant ID:", merchantId)

    // Generate static QR code for your store
    qrPNG, _ := qrcode.GenerateMerchantQR(merchantId, "Coffee Shop", qrcode.Options{
        Size: 256,
    })
    // Display/print this QR code at your store
}
```

## Payment Flow

```
┌──────────────┐                              ┌──────────────┐
│   MERCHANT   │                              │   CUSTOMER   │
└──────┬───────┘                              └──────┬───────┘
       │                                             │
       │ 1. Register merchant (one-time)             │
       │    → gets merchantId                        │
       │                                             │
       │ 2. Display QR code                          │
       │    (contains merchantId)                    │
       │                                             │
       │                                    3. Scan QR code
       │                                       │
       │                                    4. Enter amount
       │                                       │
       │                                    5. approve(USDC)
       │                                       │
       │                                    6. createOrder(merchantId, amount)
       │                                       → gets orderId
       │                                       │
       │                                    7. payOrder(orderId)
       │                                       → USDC sent to contract
       │                                             │
       │                              ┌──────────────┴──────────────┐
       │                              │          PLATFORM           │
       │                              └──────────────┬──────────────┘
       │                                             │
       │                                    8. settleOrder(orderId)
       │                                       → 98% to Merchant
       │◄──────────────────────────────────────→ 2% to Platform
       │
   Receives funds!
```

## SDK Structure

```
stablebase-go-sdk/
├── client/          # Main SDK client
├── merchant/        # Merchant registration & QR code generation
├── customer/        # Customer operations (approve, create order, pay)
├── platform/        # Platform operations (settlement)
├── qrcode/          # QR code generation & parsing
└── contracts/       # Smart contract bindings
```

## Supported Networks

| Network      | Chain ID | Status         |
| ------------ | -------- | -------------- |
| Base Sepolia | 84532    | ✅ Testnet     |
| Polygon Amoy | 80002    | ✅ Testnet     |
| Base Mainnet | 8453     | 🔜 Coming Soon |

## QR Code Format

Merchant QR codes contain the following information:

```json
{
  "v": "1",
  "type": "stablebase_merchant",
  "merchantId": "0xabcd...1234",
  "merchantName": "Coffee Shop",
  "network": "base-sepolia",
  "chainId": 84532,
  "paymentProcessor": "0x7c39408AC96a1b9a2722056eDE90b54D2B260380",
  "supportedTokens": ["USDC"]
}
```

## Events

The SDK supports listening to contract events:

| Event            | Description                        |
| ---------------- | ---------------------------------- |
| `OrderCreated`   | When a customer creates an order   |
| `OrderPaid`      | When a customer pays an order      |
| `OrderSettled`   | When the platform settles an order |
| `OrderCancelled` | When an order is cancelled         |
| `OrderExpired`   | When an order expires              |

```go
client.Events.Subscribe(sdk.EventOrderPaid, func(event OrderPaidEvent) {
    fmt.Printf("Order %s paid by %s\n", event.OrderID, event.Payer)
})
```

## Contract Addresses

### Base Sepolia (Testnet)

- **Payment Processor**: `0x7c39408AC96a1b9a2722056eDE90b54D2B260380`
- **Merchant Registry**: `0x93e93Dfa36C87De32B9118CA5D9BAd1Db892002d`
- **USDC**: `0x036CbD53842c5426634e7929541eC2318f3dCF7e`

### Polygon Amoy (Testnet)

- **Payment Processor**: `0x3B08Be115E1672cE8A6618D932a97B2Cc251d853`
- **Merchant Registry**: `0xE664919f8a195d44c8a137C71cBeb967A71eD3DF`
- **USDC**: `0x41E94Eb019C0762f9Bfcf9Fb1E58725BfB0e7582`

## Development

### Prerequisites

- Go 1.21 or higher
- Access to an EVM-compatible network (testnet recommended for development)

### Testing

```bash
# Run tests
go test ./...

# Run tests with coverage
go test -cover ./...
```

## Contributing

Contributions are welcome! Please feel free to submit a Pull Request.

1. Fork the repository
2. Create your feature branch (`git checkout -b feature/amazing-feature`)
3. Commit your changes (`git commit -m 'Add some amazing feature'`)
4. Push to the branch (`git push origin feature/amazing-feature`)
5. Open a Pull Request

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

## Resources

- [Base Sepolia Explorer](https://sepolia.basescan.org)
- [Polygon Amoy Explorer](https://amoy.polygonscan.com)
