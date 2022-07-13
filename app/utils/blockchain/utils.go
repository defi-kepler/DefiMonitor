package blockchain

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"goskeleton/app/global/variable"
	"math/big"
	"strings"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/shopspring/decimal"
)

var (
	Client *ethclient.Client
)

func init() {
	rpcUrl := variable.ConfigDefiYml.GetString("RpcUrl")
	var err error
	Client, err = ethclient.Dial(rpcUrl)
	if err != nil {
		variable.ZapLog.Error("RPC 链接异常")
	}

}

func GetSigner(signature []byte, data []byte) string {
	sigPublicKeyECDSA, _ := crypto.SigToPub(data, signature)
	return crypto.PubkeyToAddress(*sigPublicKeyECDSA).Hex()
}

func GetSignerByHex(signatureStr string, data []byte) string {
	signature, _ := hexutil.Decode(signatureStr)
	fmt.Println(GetSigner(signature, data))
	return strings.ToLower(GetSigner(signature, data))
}

func QueryTx(tx string) (receipt *types.Receipt, err error) {
	receipt, err = Client.TransactionReceipt(context.Background(), common.HexToHash(tx))
	return
}

func ToWei(iamount interface{}, decimals int) *big.Int {
	amount := decimal.NewFromFloat(0)
	switch v := iamount.(type) {
	case string:
		amount, _ = decimal.NewFromString(v)
	case float64:
		amount = decimal.NewFromFloat(v)
	case int64:
		amount = decimal.NewFromFloat(float64(v))
	case decimal.Decimal:
		amount = v
	case *decimal.Decimal:
		amount = *v
	}
	mul := decimal.NewFromFloat(float64(10)).Pow(decimal.NewFromFloat(float64(decimals)))
	result := amount.Mul(mul)
	wei := new(big.Int)
	wei.SetString(result.String(), 10)
	return wei
}

func BuildEstimateGas(from common.Address, abiStr string, contractAddress *common.Address, method string, args ...interface{}) (uint64, error) {
	abi, _ := abi.JSON(strings.NewReader(abiStr))
	var data []byte
	if len(args) > 0 {
		data, _ = abi.Pack(method, args...)
	} else {
		data, _ = abi.Pack(method)
	}
	ethRPCParams := ethereum.CallMsg{
		From:  from,
		To:    contractAddress,
		Value: big.NewInt(0),
		Data:  data,
	}
	return Client.EstimateGas(context.Background(), ethRPCParams)
}

func PrivateKeyToAddress(privateKey *ecdsa.PrivateKey) common.Address {
	return crypto.PubkeyToAddress(privateKey.PublicKey)
}

func BuildTransactor(privateKey *ecdsa.PrivateKey) (opts *bind.TransactOpts, _err error) {

	nonce, err := Client.PendingNonceAt(context.Background(), PrivateKeyToAddress(privateKey))
	if err != nil {
		_err = err
		return
	}
	gasPrice, err := Client.SuggestGasPrice(context.Background())
	if err != nil {
		_err = err
		return
	}
	chainID, _ := Client.ChainID(context.Background())
	auth, _ := bind.NewKeyedTransactorWithChainID(privateKey, chainID)
	auth.Nonce = big.NewInt(int64(nonce))
	auth.Value = big.NewInt(0)     // in wei
	auth.GasLimit = uint64(350000) // in units
	auth.GasPrice = gasPrice
	opts = auth
	return
}
