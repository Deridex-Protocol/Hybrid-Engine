package common

import (
	"context"
	"fmt"
	"math/big"
	"sort"
	"strings"
	"testing"
	"time"

	"bitbucket.ideasoft.io/dex/dex-backend/common/contract"
	commoncrypto "bitbucket.ideasoft.io/dex/dex-backend/common/crypto"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	ethcommon "github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/shopspring/decimal"
	"github.com/stretchr/testify/require"
)

func TestGetTradeFlags(t *testing.T) {
	buyFlags, err := GetTradeFlags(true)
	require.NoError(t, err)
	require.Equal(t, buyFlags[31], uint8(TraderFlagIsBuyByte))

	sellFlags, err := GetTradeFlags(false)
	require.NoError(t, err)
	require.Equal(t, sellFlags[31], uint8(TraderFlagIsNullByte))
}

// [Prod kovan bot 1 0x74628399D6BE80f83F418f1DF6c7b7aa18C13db0] 10000125029 true 0 false
// [Prod kovan bot 2 0x9C8c32f402dC69b1811Dd3652A4C26d7849FfE95] 9999873836 true 0 false
// [Prod kovan relayer 0xB43dFDBDF0defdb267C5fBe9A968F51BAF4621eE] 9999999749 true 0 false
// [Local kovan test1 0x7A7050988b93BaAcDF5c6Bac19847aef0998Bd58] 1000147 true 0 false
// [Local kovan test2 0x904B8339Af65FFf3E95CEBcA3f7da1223c67F89B] 999828 true 0 false
// [Local kovan test relayer 0x76EF0ABDB567AC99AF92782093232FD726aaaC1A] 1000000 true 0 false
func TestGetAccountBalance(t *testing.T) {
	client, err := ethclient.Dial("https://kovan.infura.io/v3/10e6ea75b3174d3ca1c8e6cda6de0eac")
	require.NoError(t, err)

	caller, err := contract.NewContract(ethcommon.HexToAddress("0x6E8aDE8AD61bff517a0E946b98D7AD788b05D911"), client)
	require.NoError(t, err)

	accounts := [][]string{
		{"Prod kovan bot 1", "0x74628399D6BE80f83F418f1DF6c7b7aa18C13db0"},
		{"Prod kovan bot 2", "0x9C8c32f402dC69b1811Dd3652A4C26d7849FfE95"},
		{"Prod kovan relayer", "0xB43dFDBDF0defdb267C5fBe9A968F51BAF4621eE"},
		{"Local kovan test1", "0x7A7050988b93BaAcDF5c6Bac19847aef0998Bd58"},
		{"Local kovan test2", "0x904B8339Af65FFf3E95CEBcA3f7da1223c67F89B"},
		{"Local kovan test relayer", "0x76EF0ABDB567AC99AF92782093232FD726aaaC1A"},
	}

	for i := range accounts {
		balance, err := caller.GetAccountBalance(nil, ethcommon.HexToAddress(accounts[i][1]))
		require.NoError(t, err)
		fmt.Println(accounts[i], balance.Margin, balance.MarginIsPositive, balance.Position, balance.PositionIsPositive)
	}
}

func TestTrade(t *testing.T) {
	// local test accounts
	address1PrivateKey := "77a18d8357dacca41b60a951afa9a246639febb5094d7b09f37269ffcb5816bc"
	address2PrivateKey := "76b0fcdf656c0dfd4bbd7b6f199c7166b7ae9be538bf81fec3590c73802fcd7b"
	addressRelayerPrivateKey := "8f43a52c295e88e3e3d4c2e16c2af3cc9f3e596a3205a75985ede44c1dde96aa"

	// prod bot accounts
	// address1PrivateKey = "a845d6234627fd4927c354c7d0cc1a6469678954c03ff98f4af836fff409b2da"
	// address2PrivateKey = "898348a863bdb21871aae286c289b7b8e570f274130eae1182371af713e508a8"
	// addressRelayerPrivateKey = "f02b4ce795e371a5a42d488f301f85dbb8c5af50082c2bf2c0f8ff6a20dcd45f"

	amount := big.NewInt(100) // 1 = 1000000. decimals 6
	price := decimal.New(1000, 18).BigInt()
	makerIsBuy := false

	// kovan smart contract
	addressDxlnPerpetualProxy := "0x6E8aDE8AD61bff517a0E946b98D7AD788b05D911"
	addressDxlnOrders := ethcommon.HexToAddress("0xC050fE2c8f2F7423a4d6E1f3b104341c07E7f532")

	client, err := ethclient.Dial("https://kovan.infura.io/v3/10e6ea75b3174d3ca1c8e6cda6de0eac")
	require.NoError(t, err)

	caller, err := contract.NewContract(ethcommon.HexToAddress(addressDxlnPerpetualProxy), client)
	require.NoError(t, err)

	relayerPrivateKey, err := crypto.HexToECDSA(addressRelayerPrivateKey)
	require.NoError(t, err)
	relayerEthAddress := crypto.PubkeyToAddress(relayerPrivateKey.PublicKey)

	makerPrivateKey, err := crypto.HexToECDSA(address1PrivateKey)
	require.NoError(t, err)
	makerEthAddress := crypto.PubkeyToAddress(makerPrivateKey.PublicKey)

	takerPrivateKey, err := crypto.HexToECDSA(address2PrivateKey)
	require.NoError(t, err)
	takerEthAddress := crypto.PubkeyToAddress(takerPrivateKey.PublicKey)

	addressNonce, err := client.NonceAt(context.Background(), crypto.PubkeyToAddress(relayerPrivateKey.PublicKey), nil)
	require.NoError(t, err)

	opts := &bind.TransactOpts{
		From:  relayerEthAddress,
		Nonce: big.NewInt(int64(addressNonce)),
		Signer: func(address ethcommon.Address, tx *types.Transaction) (*types.Transaction, error) {
			signature, err := crypto.Sign(types.HomesteadSigner{}.Hash(tx).Bytes(), relayerPrivateKey)
			if err != nil {
				return nil, err
			}
			return tx.WithSignature(types.HomesteadSigner{}, signature)
		},
		GasPrice: big.NewInt(2 * 1000000000), // 1 GWEI
		GasLimit: 500000,
	}

	flags, err := GetTradeFlags(makerIsBuy)
	require.NoError(t, err)

	makerTradeData := contract.DxlnOrdersTradeData{
		Order: contract.DxlnOrdersOrder{
			Flags:        flags,
			Amount:       amount,
			LimitPrice:   price,
			TriggerPrice: big.NewInt(0),
			LimitFee:     big.NewInt(0),
			Maker:        makerEthAddress,
			Taker:        relayerEthAddress,
			Expiration:   big.NewInt(0),
		},
		Fill: contract.DxlnOrdersFill{
			Amount:        amount,
			Price:         price,
			Fee:           big.NewInt(0),
			IsNegativeFee: false,
		},
	}

	makerSign, err := commoncrypto.SignOrder(address1PrivateKey, &makerTradeData.Order, 42, addressDxlnOrders)
	require.NoError(t, err)

	makerSign = append(makerSign, SignatureTypeDecimal)

	copy(makerTradeData.Signature.R[:], makerSign[:32])
	copy(makerTradeData.Signature.S[:], makerSign[32:64])
	copy(makerTradeData.Signature.VType[:], makerSign[64:])

	abiEncoder, err := abi.JSON(strings.NewReader(contract.TradeMetaData.ABI))
	require.NoError(t, err)

	tradeMethod, ok := abiEncoder.Methods["tradeExample"]
	require.True(t, ok)

	makerData, err := tradeMethod.Inputs.Pack(makerTradeData)
	require.NoError(t, err)

	flags, err = GetTradeFlags(!makerIsBuy)
	require.NoError(t, err)

	takerTradeData := contract.DxlnOrdersTradeData{
		Order: contract.DxlnOrdersOrder{
			Flags:        flags,
			Amount:       amount,
			LimitPrice:   price,
			TriggerPrice: big.NewInt(0),
			LimitFee:     big.NewInt(0),
			Maker:        takerEthAddress,
			Taker:        relayerEthAddress,
			Expiration:   big.NewInt(0),
		},
		Fill: contract.DxlnOrdersFill{
			Amount:        amount,
			Price:         price,
			Fee:           big.NewInt(0),
			IsNegativeFee: false,
		},
	}

	takerSign, err := commoncrypto.SignOrder(address2PrivateKey, &takerTradeData.Order, 42, addressDxlnOrders)
	require.NoError(t, err)

	takerSign = append(takerSign, SignatureTypeDecimal)

	copy(takerTradeData.Signature.R[:], takerSign[:32])
	copy(takerTradeData.Signature.S[:], takerSign[32:64])
	copy(takerTradeData.Signature.VType[:], takerSign[64:])

	takerData, err := tradeMethod.Inputs.Pack(takerTradeData)
	require.NoError(t, err)

	var accounts = []ethcommon.Address{makerEthAddress, takerEthAddress, relayerEthAddress}
	sort.SliceStable(accounts, func(i, j int) bool {
		return accounts[i].String() < accounts[j].String()
	})

	var makerAddressIndex, takerAddressIndex, relayerAddressIndex *big.Int
	for i := range accounts {
		switch accounts[i].Hex() {
		case makerEthAddress.Hex():
			makerAddressIndex = big.NewInt(int64(i))
		case takerEthAddress.Hex():
			takerAddressIndex = big.NewInt(int64(i))
		case relayerEthAddress.Hex():
			relayerAddressIndex = big.NewInt(int64(i))
		}
	}

	var trades = []contract.DxlnTradeTradeArg{
		{
			TakerIndex: relayerAddressIndex,
			MakerIndex: makerAddressIndex,
			Trader:     addressDxlnOrders,
			Data:       makerData,
		}, {
			TakerIndex: relayerAddressIndex,
			MakerIndex: takerAddressIndex,
			Trader:     addressDxlnOrders,
			Data:       takerData,
		},
	}

	tx, err := caller.Trade(opts, accounts, trades)
	require.NoError(t, err)

	fmt.Println("Transaction hash", tx.Hash())

	for {
		time.Sleep(time.Second)

		ctx := context.Background()
		_, isPending, err := client.TransactionByHash(ctx, tx.Hash())
		if isPending || err == ethereum.NotFound {
			continue
		}
		require.NoError(t, err)

		receipt, err := client.TransactionReceipt(ctx, tx.Hash())
		require.NoError(t, err)

		if receipt.Status == types.ReceiptStatusSuccessful {
			fmt.Println("Transaction is SUCCESSFUL")
		} else {
			fmt.Println("Transaction is FAILED")
		}
		break
	}
}
