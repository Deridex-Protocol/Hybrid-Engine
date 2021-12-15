package crypto

import (
	"bytes"
	"encoding/hex"
	"fmt"
	"math/big"
	"strings"

	"bitbucket.ideasoft.io/dex/dex-backend/common/contract"
	"github.com/ethereum/go-ethereum/accounts"
	"github.com/ethereum/go-ethereum/accounts/abi"
	ethcommon "github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
)

func SignEthereumMessage(privateKeyHex string, message []byte) ([]byte, error) {
	messageHash := accounts.TextHash(message)
	privateKey, err := crypto.HexToECDSA(privateKeyHex)
	if err != nil {
		return nil, err
	}
	sign, err := crypto.Sign(messageHash, privateKey)
	if err != nil {
		return nil, err
	}
	sign[64] = sign[64] + 27
	return sign, nil
}

func SignAuthHeader(privateKeyHex, address, message string) (string, error) {
	sign, err := SignEthereumMessage(privateKeyHex, []byte(message))
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("%s#%s#%s", address, message, "0x"+hex.EncodeToString(sign)), nil
}

func IsValidAuthSignature(address string, message string, signature string) (bool, error) {
	messageHash := accounts.TextHash([]byte(message))

	if strings.HasPrefix(signature, "0x") || strings.HasPrefix(signature, "0X") {
		signature = signature[2:]
	}
	signatureBytesHex, err := hex.DecodeString(signature)
	if err != nil {
		return false, err
	}

	if signatureBytesHex[64] >= 27 {
		signatureBytesHex[64] -= 27
	}

	publicKey, err := crypto.SigToPub(messageHash, signatureBytesHex)
	if err != nil {
		return false, err
	}
	if publicKey == nil {
		return false, nil
	}
	return bytes.Equal(ethcommon.HexToAddress(address).Bytes(), crypto.PubkeyToAddress(*publicKey).Bytes()), nil
}

func SignOrder(privateKeyHex string, order *contract.DxlnOrdersOrder, chanID int64, contractAddress ethcommon.Address) ([]byte, error) {
	orderHash, err := GetOrderHash(order, chanID, contractAddress)
	if err != nil {
		return nil, err
	}
	return SignEthereumMessage(privateKeyHex, orderHash)
}

func GetOrderHash(order *contract.DxlnOrdersOrder, chanID int64, contractAddress ethcommon.Address) ([]byte, error) {
	bytesType, err := abi.NewType("bytes32", "", nil)
	if err != nil {
		return nil, err
	}
	uint256Type, err := abi.NewType("uint256", "", nil)
	if err != nil {
		return nil, err
	}
	addressType, err := abi.NewType("address", "", nil)
	if err != nil {
		return nil, err
	}

	domainHashArguments := abi.Arguments{
		{Type: bytesType},
		{Type: bytesType},
		{Type: bytesType},
		{Type: uint256Type},
		{Type: addressType},
	}

	var domainSeparatorSchema, domainName, domainVersion [32]byte
	copy(domainSeparatorSchema[:], crypto.Keccak256([]byte(`EIP712Domain(string name,string version,uint256 chainId,address verifyingContract)`)))
	copy(domainName[:], crypto.Keccak256([]byte("DexOrders")))
	copy(domainVersion[:], crypto.Keccak256([]byte("1.0")))

	domainHashData, err := domainHashArguments.Pack(
		domainSeparatorSchema,
		domainName,
		domainVersion,
		big.NewInt(chanID),
		contractAddress,
	)
	if err != nil {
		return nil, err
	}

	structHashArguments := abi.Arguments{
		{Type: bytesType},
		{Type: uint256Type},
		{Type: uint256Type},
		{Type: uint256Type},
		{Type: uint256Type},
		{Type: addressType},
		{Type: addressType},
		{Type: uint256Type},
	}
	structHashData, err := structHashArguments.Pack(
		order.Flags,
		order.Amount,
		order.LimitPrice,
		order.TriggerPrice,
		order.LimitFee,
		order.Maker,
		order.Taker,
		order.Expiration,
	)
	if err != nil {
		return nil, err
	}

	return crypto.Keccak256(
		[]byte{'\x19', '\x01'},
		crypto.Keccak256(domainHashData),
		crypto.Keccak256(
			crypto.Keccak256([]byte(`Order(bytes32 flags,uint256 amount,uint256 limitPrice,uint256 triggerPrice,uint256 limitFee,address maker,address taker,uint256 expiration)`)),
			structHashData,
		),
	), nil
}

func IsValidOrderSignature(address string, orderHash, signature []byte) (bool, error) {
	orderHash = accounts.TextHash(orderHash)
	if signature[64] >= 27 {
		signature[64] -= 27
	}
	publicKey, err := crypto.SigToPub(orderHash, signature)
	if err != nil {
		return false, err
	}
	if publicKey == nil {
		return false, nil
	}
	return bytes.Equal(ethcommon.HexToAddress(address).Bytes(), crypto.PubkeyToAddress(*publicKey).Bytes()), nil
}
