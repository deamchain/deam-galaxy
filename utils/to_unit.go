package utils

import "math/big"

// ToUnit number of DEAM to Wei
func ToUnit(deam uint64) *big.Int {
	return new(big.Int).Mul(new(big.Int).SetUint64(deam), big.NewInt(1e18))
}
