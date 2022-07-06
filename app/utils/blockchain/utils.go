package blockchain

import (
	"context"
	"fmt"
	"goskeleton/abi/defi/erc20"
	"goskeleton/app/global/variable"
	"goskeleton/hdwallet"
	"log"
	"math/big"
	"strings"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts"
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
	Wallet *hdwallet.Wallet
)

func init() {
	rpcUrl := variable.ConfigDefiYml.GetString("RpcUrl")
	var err error
	Client, err = ethclient.Dial(rpcUrl)
	if err != nil {
		variable.ZapLog.Error("RPC 链接异常")
	}
	mnemonic := variable.ConfigDefiYml.GetString("Mnemonic")
	Wallet, err = hdwallet.NewFromMnemonic(mnemonic)
	if err != nil {
		log.Fatal(err)
	}
}

func ERC20Transfer(tokenAdd string, pathStr string, toAddr string, amount *big.Int) (tx *types.Transaction, err error) {
	tokenAddress := common.HexToAddress(tokenAdd)
	instance, _ := erc20.NewErc20(tokenAddress, Client)
	toAddress := common.HexToAddress(toAddr)
	path := hdwallet.MustParseDerivationPath(pathStr)
	account, err := Wallet.Derive(path, false)
	if err != nil {
		variable.ZapLog.Sugar().Error(err)
	}
	nonce, err := Client.PendingNonceAt(context.Background(), account.Address)
	if err != nil {
		variable.ZapLog.Sugar().Error(err)
	}
	gasPrice, err := Client.SuggestGasPrice(context.Background())
	if err != nil {
		variable.ZapLog.Sugar().Error(err)
	}
	privateKey, err := Wallet.PrivateKey(account)
	if err != nil {
		variable.ZapLog.Sugar().Error(err)
	}
	chainID, _ := Client.ChainID(context.Background())
	auth, _ := bind.NewKeyedTransactorWithChainID(privateKey, chainID)
	auth.Nonce = big.NewInt(int64(nonce))
	auth.Value = big.NewInt(0)     // in wei
	auth.GasLimit = uint64(300000) // in units
	auth.GasPrice = gasPrice
	tx, err = instance.Transfer(auth, toAddress, amount)
	if err != nil {
		variable.ZapLog.Sugar().Error(err)
	}
	variable.ZapLog.Sugar().Info(tx.Hash())
	return
}

func ERC20TransferBySymbol(symbol string, pathStr string, toAddr string, amount *big.Int) (tx *types.Transaction, err error) {
	if strings.ToUpper(symbol) == "BUSD" {
		busd := variable.ConfigDefiYml.GetString("BUSD")
		tx, err = ERC20Transfer(busd, pathStr, toAddr, amount)
	} else if strings.ToLower(symbol) == "usdt" {
		USDT := variable.ConfigDefiYml.GetString("USDT")
		tx, err = ERC20Transfer(USDT, pathStr, toAddr, amount)
	} else {
		panic("symbol error")
	}
	return
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

func BuildEstimateGas(owner accounts.Account, abiStr string, contractAddress *common.Address, method string, args ...interface{}) (uint64, error) {
	abi, _ := abi.JSON(strings.NewReader(abiStr))
	var data []byte
	if len(args) > 0 {
		data, _ = abi.Pack(method, args...)
	} else {
		data, _ = abi.Pack(method)
	}
	ethRPCParams := ethereum.CallMsg{
		From:  owner.Address,
		To:    contractAddress,
		Value: big.NewInt(0),
		Data:  data,
	}
	return Client.EstimateGas(context.Background(), ethRPCParams)
}

func BuildTransactor(owner accounts.Account) (opts *bind.TransactOpts, _err error) {
	nonce, err := Client.PendingNonceAt(context.Background(), owner.Address)
	if err != nil {
		_err = err
		return
	}
	gasPrice, err := Client.SuggestGasPrice(context.Background())
	if err != nil {
		_err = err
		return
	}
	privateKey, err := Wallet.PrivateKey(owner)
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
