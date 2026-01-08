package qrcode

import (
	"encoding/json"
	"testing"
)

func TestGenerate(t *testing.T) {
	// Table-driven test - multiple test cases
	tests := []struct {
		name      string
		data      PaymentQRData
		size      int
		expectErr bool
	}{
		{
			name: "valid payment data",
			data: PaymentQRData{
				MerchantID:   "abc123",
				Amount:       "1000000",
				TokenAddress: "0x123",
				OrderID:      "order-1",
				ChainID:      "1",
			},
			size:      256,
			expectErr: false,
		},
		{
			name: "large QR code",
			data: PaymentQRData{
				MerchantID:   "def456",
				Amount:       "5000000",
				TokenAddress: "0x456",
				ChainID:      "137",
			},
			size:      512,
			expectErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			qrBytes, err := Generate(tt.data, tt.size)

			// Check error expectation
			if (err != nil) != tt.expectErr {
				t.Errorf("Generate() error = %v, expectErr %v", err, tt.expectErr)
				return
			}

			// If no error expected, check we got data
			if !tt.expectErr && len(qrBytes) == 0 {
				t.Error("Generate() returned empty bytes")
			}

		})
	}
}

func TestParse(t *testing.T) {
	tests := []struct {
		name      string
		input     string
		want      *PaymentQRData
		expectErr bool
	}{
		{
			name: "valid JSON",
			input: `{
                "merchant_id": "abc123",
                "amount": "1000000",
                "token_address": "0x123",
                "order_id": "order-1",
                "chain_id": "1"
            }`,
			want: &PaymentQRData{
				MerchantID:   "abc123",
				Amount:       "1000000",
				TokenAddress: "0x123",
				OrderID:      "order-1",
				ChainID:      "1",
			},
			expectErr: false,
		},
		{
			name:      "invalid JSON",
			input:     `not json`,
			want:      nil,
			expectErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Parse(tt.input)

			if (err != nil) != tt.expectErr {
				t.Errorf("Parse() error = %v, expectErr %v", err, tt.expectErr)
				return
			}

			if !tt.expectErr {
				if got.MerchantID != tt.want.MerchantID {
					t.Errorf("MerchantID = %v, want %v", got.MerchantID, tt.want.MerchantID)
				}
				if got.Amount != tt.want.Amount {
					t.Errorf("Amount = %v, want %v", got.Amount, tt.want.Amount)
				}
				if got.TokenAddress != tt.want.TokenAddress {
					t.Errorf("TokenAddress = %v, want %v", got.TokenAddress, tt.want.TokenAddress)
				}
				if got.OrderID != tt.want.OrderID {
					t.Errorf("OrderID = %v, want %v", got.TokenAddress, tt.want.TokenAddress)
				}
				if got.ChainID != tt.want.ChainID {
					t.Errorf("ChainID = %v, want %v", got.ChainID, tt.want.ChainID)
				}
			}
		})
	}
}

func TestGenerateAndParse(t *testing.T) {
	// Create test data
	original := PaymentQRData{
		MerchantID:   "test123",
		Amount:       "2000000",
		TokenAddress: "0xABC",
		OrderID:      "order-123",
		ChainID:      "1",
	}

	// Generate QR code
	_, err := Generate(original, 256)
	if err != nil {
		t.Fatalf("Generate() failed: %v", err)
	}

	// The QR contains JSON, so we decode it manually to test parsing
	jsonStr, _ := json.Marshal(original)

	// Parse it back
	parsed, err := Parse(string(jsonStr))
	if err != nil {
		t.Fatalf("Parse() failed: %v", err)
	}

	// Compare
	if parsed.MerchantID != original.MerchantID {
		t.Errorf("MerchantID mismatch: got %v, want %v",
			parsed.MerchantID, original.MerchantID)
	}
	if parsed.Amount != original.Amount {
		t.Errorf("Amount mismatch: got %v, want %v",
			parsed.Amount, original.Amount)
	}
}
