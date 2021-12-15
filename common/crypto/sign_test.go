package crypto

import (
	"encoding/hex"
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestSignAuthHeader(t *testing.T) {
	privateKey := "a845d6234627fd4927c354c7d0cc1a6469678954c03ff98f4af836fff409b2da"
	address := "0x74628399D6BE80f83F418f1DF6c7b7aa18C13db0"
	message := "Authentication:2022-12-01T00:00:00Z"
	authSignature, err := SignAuthHeader(privateKey, address, message)
	require.NoError(t, err)

	t.Log("authSignature", authSignature)

	authSignaturePart := strings.Split(authSignature, "#")
	require.Len(t, authSignaturePart, 3)

	valid, err := IsValidAuthSignature(address, message, authSignaturePart[2])
	require.NoError(t, err)
	require.True(t, valid)
}

func TestIsValidAuthSignature(t *testing.T) {
	authToken := "0x74628399d6be80f83f418f1df6c7b7aa18c13db0#Authentication:2021-11-16T16:09:16.927Z#0x6993693062f54a1e8a543f3b4aef93fbf03bfbdb3350b2a8a1f9fb645528e65f05648c8c69b4af86dac5f45a32931c060e25442b6220ca697b937182b2818f081c"
	auth := strings.Split(authToken, "#")
	valid, err := IsValidAuthSignature(auth[0], auth[1], auth[2])
	require.NoError(t, err)
	t.Log("Token is valid: ", valid)
}

func TestSignOrderID(t *testing.T) {
	// first taker, second maker
	privateKey := "a845d6234627fd4927c354c7d0cc1a6469678954c03ff98f4af836fff409b2da"
	// privateKey := "898348a863bdb21871aae286c289b7b8e570f274130eae1182371af713e508a8"
	orderHash, err := hex.DecodeString("0x02a372c3722050717928aea5fa5d51c45692a1f3acff2463326011934a082004"[2:])
	require.NoError(t, err)

	sign, err := SignEthereumMessage(privateKey, orderHash)
	require.NoError(t, err)

	t.Log("sign", "0x"+hex.EncodeToString(sign))

	address := "0x74628399D6BE80f83F418f1DF6c7b7aa18C13db0"
	// address := "0x9C8c32f402dC69b1811Dd3652A4C26d7849FfE95"
	valid, err := IsValidOrderSignature(address, orderHash, sign)
	require.NoError(t, err)
	require.True(t, valid)
}

func TestIsValidOrderSignature(t *testing.T) {
	// first taker, second maker
	// address := "0x74628399D6BE80f83F418f1DF6c7b7aa18C13db0"
	address := "0x9C8c32f402dC69b1811Dd3652A4C26d7849FfE95"

	orderHash, err := hex.DecodeString("0xb7cb927bf3751e9034446b6b667296ab8dc7e1f17baf6933091368b28ae4b53a"[2:])
	require.NoError(t, err)

	sign, err := hex.DecodeString("7a91bb1453719a6b7266ddc37dea0c1bfae0139046a2e5f25f250131e1399c211518a5955578891b55beb201a9d1b3aa9ca5c0fc41593d56174d07c0f96806e91b")
	require.NoError(t, err)

	valid, err := IsValidOrderSignature(address, orderHash, sign)
	require.NoError(t, err)
	require.True(t, valid)
}
