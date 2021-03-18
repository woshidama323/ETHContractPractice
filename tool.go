package main

import (
	"errors"
	"math/big"
)

func StringToBigInt(reserve string, precision uint64) (*big.Int, error) {

	unit := big.NewInt(0).Exp(big.NewInt(10), big.NewInt(int64(precision)), big.NewInt(0))
	if minBalance, ok := big.NewFloat(0).SetString(reserve); ok {
		minBalance = big.NewFloat(0).Mul(minBalance, big.NewFloat(0).SetInt(unit))

		minBalance.Int(unit)
		return unit, nil
	}
	return nil, errors.New("failed to parse string to big int")
}
