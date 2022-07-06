package numeric

import (
	"database/sql/driver"
	"errors"
	"fmt"
	"math/big"
	"strings"
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
	if n.Int == nil {
		return nil, nil
	}
	return n.Int.String(), nil
}

func (n Numeric) MarshalJSON() ([]byte, error) {
	return []byte(n.Int.String()), nil
}

func (n *Numeric) UnmarshalJSON(p []byte) error {
	numericString := strings.Trim(string(p), "\"'")
	if numericString == "null" {
		return nil
	}
	var ok bool
	n.Int, ok = new(big.Int).SetString(numericString, 10)
	if !ok {
		return fmt.Errorf("not a valid big integer: %s", p)
	}
	return nil
}

type NumericList struct {
	List []*big.Int
}

func (NumericList) GormDataType() string {
	return "numeric[]"
}

func (nl *NumericList) Scan(value interface{}) error {
	for _, s := range value.([]string) {
		i, ok := new(big.Int).SetString(s, 10)
		if !ok {
			return errors.New(fmt.Sprint("Failed to unmarshal value:", s))
		}
		nl.List = append(nl.List, i)
	}

	return nil
}

func (nl NumericList) Value() (driver.Value, error) {
	var v []string
	for _, i := range nl.List {
		v = append(v, i.String())
	}
	return v, nil
}
