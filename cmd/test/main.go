package main

import (
	"context"
	"fmt"
	"goskeleton/app/global/variable"
	"goskeleton/app/utils/blockchain"
	_ "goskeleton/bootstrap"
	"goskeleton/hdwallet"
	"log"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
)

var (
	logger = variable.ZapLog.Sugar()
)

func main() {
	sign()
	check_sign()
	//init1()
	//fmt.Println("", com.RandNum(8))
}

func init1() {
	mnemonic := variable.ConfigDefiYml.GetString("Mnemonic")

	wallet, err := hdwallet.NewFromMnemonic(mnemonic)
	if err != nil {
		log.Fatal(err)
	}
	path := hdwallet.MustParseDerivationPath("m/44'/60'/0'/0/0")

	account, err := wallet.Derive(path, false)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(account.Address.Hex()) // 0xC2142aB6AFcb127e7af15C87BDC58AC8C503a782

	path = hdwallet.MustParseDerivationPath("m/44'/60'/0'/0/1")

	account, err = wallet.Derive(path, false)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(account.Address.Hex()) // 0xDa446Ef1E3Aa8DaBA66dEB844d883097369915eB

	privateKeyHex, _ := wallet.PrivateKeyHex(account)
	fmt.Println(" privateKeyHex = ", privateKeyHex)
}

func transformEth() {
	mnemonic := variable.ConfigDefiYml.GetString("Mnemonic")

	wallet, err := hdwallet.NewFromMnemonic(mnemonic)
	if err != nil {
		log.Fatal(err)
	}
	path := hdwallet.MustParseDerivationPath("m/44'/60'/0'/0/0")
	account, err := wallet.Derive(path, true)
	if err != nil {
		log.Fatal(err)
	}
	value := big.NewInt(1000000000000000000)
	toAddress := common.HexToAddress("0x0")
	gasLimit := uint64(21000)
	gasPrice := big.NewInt(21000000000)
	var data []byte

	client, err := ethclient.Dial(variable.ConfigDefiYml.GetString("RpcUrl"))
	if err != nil {
		logger.Error(err)
	}
	nonce, _ := client.PendingNonceAt(context.Background(), account.Address)
	tx := types.NewTransaction(nonce, toAddress, value, gasLimit, gasPrice, data)
	signedTx, err := wallet.SignTx(account, tx, nil)
	if err != nil {
		log.Fatal(err)
	}
	if err != nil {
		log.Fatal(err)
	}
	err = client.SendTransaction(context.Background(), signedTx)
	if err != nil {
		logger.Error(err)
	}

	fmt.Printf("tx sent: %s", signedTx.Hash().Hex())
}

func sign() {
	mnemonic := variable.ConfigDefiYml.GetString("Mnemonic")

	wallet, err := hdwallet.NewFromMnemonic(mnemonic)
	if err != nil {
		log.Fatal(err)
	}
	path := hdwallet.MustParseDerivationPath("m/44'/60'/0'/0/0")
	account, err := wallet.Derive(path, true)
	if err != nil {
		log.Fatal(err)
	}
	str := "a"
	str1 := fmt.Sprintf("\u0019Ethereum Signed Message:\n%d%v", len(str), str)
	data := []byte(str1)
	hash := crypto.Keccak256Hash(data)
	fmt.Println("hash: ", hexutil.Encode(hash.Bytes()))

	//hash := common.BytesToHash(data)

	sign1, _ := wallet.SignHash(account, hash.Bytes())

	fmt.Println("sign1: ", hexutil.Encode(sign1))

}

func check_sign() {
	str := "a"
	str1 := fmt.Sprintf("\x19Ethereum Signed Message:\n%d%v", len(str), str)
	data := []byte(str1)
	hash := crypto.Keccak256Hash(data)
	fmt.Println("hash =", common.Bytes2Hex(hash.Bytes()))
	var signatureStr = "0x900f278dd2be8c63696a426e762896871bd8012939e40b12306a8c6fe776ed4b278afcad0524a2d6732dd456c69867ccacf97d22e3e85ff5622d43174721b5a41c"

	signatureStr = signatureStr[:len(signatureStr)-1] + "0"

	signature, _ := hexutil.Decode(signatureStr)
	add := blockchain.GetSigner(signature, hash.Bytes())
	fmt.Println("add =", add)
}

func queryTx() {

	receipt, err := blockchain.QueryTx("0xf6061446bacf05f9ab6287daacfcd98f1d9daac9652d87f8c585a9fd0577e646")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(receipt.Status) // 1
	fmt.Println(receipt.Logs)   // ...
}
