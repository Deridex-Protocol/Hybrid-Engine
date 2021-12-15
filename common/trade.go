package common

import (
	"crypto/rand"
)

func GetTradeFlags(isBuy bool) ([32]byte, error) {
	randomBytes := make([]byte, 31)
	_, err := rand.Read(randomBytes)
	if err != nil {
		return [32]byte{}, err
	}

	var flags [32]byte
	copy(flags[:31], randomBytes)
	if isBuy {
		flags[31] = TraderFlagIsBuyByte
	} else {
		flags[31] = TraderFlagIsNullByte
	}

	return flags, nil
}
