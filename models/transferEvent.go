// models/transfer_event.go
package models

import (
	"database/sql/driver"
	"fmt"
	"math/big"
)

type TransferEvent struct {
	From   string `gorm:"type:varchar(42)"`
	To     string `gorm:"type:varchar(42)"`
	Value  BigInt `gorm:"type:numeric"`
	TxHash string `gorm:"type:varchar(66)"`
}

type BigInt struct {
	*big.Int
}

// Scan implements the sql.Scanner interface
func (b *BigInt) Scan(value interface{}) error {
	switch v := value.(type) {
	case []byte:
		b.Int = new(big.Int)
		b.Int.SetString(string(v), 10)
	case string:
		b.Int = new(big.Int)
		b.Int.SetString(v, 10)
	default:
		return fmt.Errorf("unsupported scan type for BigInt: %T", value)
	}
	return nil
}

// Value implements the driver.Valuer interface
func (b BigInt) Value() (driver.Value, error) {
	return b.String(), nil
}
