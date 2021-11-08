package numeric

import (
	"database/sql/driver"
	"errors"
	"fmt"
	"math/big"
)

type Numeric struct {
	Int *big.Int
}

func (Numeric) GormDataType() string {
	return "numeric"
}

func (n *Numeric) Scan(value interface{}) error {
	var ok bool

	n.Int, ok = new(big.Int).SetString(value.(string), 10)
	if !ok {
		return errors.New(fmt.Sprint("Failed to unmarshal value:", value))
	}
	return nil
}

func (n Numeric) Value() (driver.Value, error) {
	return n.Int.String(), nil
}
