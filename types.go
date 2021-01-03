package main

import "math/big"

//DistributionValue distribution的结果
type DistributionValue struct {
	ReturnAmount *big.Int
	Distribution []*big.Int
}
