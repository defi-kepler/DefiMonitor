package monitor

import (
	"crypto/ecdsa"
	"goskeleton/abi/defi/dispatcher"
	"goskeleton/abi/defi/erc20"
	"goskeleton/app/global/variable"
	"goskeleton/app/utils/blockchain"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"go.uber.org/zap"
)

type DispatcherService struct {
	//owner             accounts.Account
	instance                   *dispatcher.Dispatcher
	dispatcherAddress          common.Address
	treasuryAddress            common.Address
	privateKey                 *ecdsa.PrivateKey
	owner                      common.Address
	chainBridgeId              *big.Int
	withdrawalTokenAddress     common.Address
	withdrawalAccountAddress   common.Address
	withdrawalAccountThreshold *big.Int
	withdrawalTokenErc20       *erc20.Erc20
}

func CreateDispatcherService() *DispatcherService {
	dispatcherAddress := common.HexToAddress(variable.ConfigDefiYml.GetString("DispatcherAddress"))
	instance, err := dispatcher.NewDispatcher(dispatcherAddress, blockchain.Client)
	if err != nil {
		panic(err)
	}
	privateKey, err := crypto.HexToECDSA(variable.ConfigDefiYml.GetString("PrivateKey"))
	owner := blockchain.PrivateKeyToAddress(privateKey)
	variable.ZapLog.Info("owner = ", zap.String("owner", owner.Hex()))
	if err != nil {
		panic(err)
	}
	treasuryAddress := common.HexToAddress(variable.ConfigDefiYml.GetString("TreasuryAddress"))
	chainBridgeId := big.NewInt(variable.ConfigDefiYml.GetInt64("ChainBridgeId"))
	var withdrawalAccountThreshold = big.NewInt(variable.ConfigDefiYml.GetInt64("WithdrawalAccountThreshold"))
	withdrawalAccountThreshold = blockchain.ToWei(withdrawalAccountThreshold, 18)
	withdrawalTokenAddress := common.HexToAddress(variable.ConfigDefiYml.GetString("WithdrawalTokenAddress"))
	withdrawalTokenErc20, err := erc20.NewErc20(withdrawalTokenAddress, blockchain.Client)
	if err != nil {
		panic(err)
	}
	withdrawalAccountAddress := common.HexToAddress(variable.ConfigDefiYml.GetString("WithdrawalAccountAddress"))
	return &DispatcherService{
		instance, dispatcherAddress, treasuryAddress, privateKey, owner,
		chainBridgeId,
		withdrawalTokenAddress,
		withdrawalAccountAddress,
		withdrawalAccountThreshold,
		withdrawalTokenErc20,
	}
}

func (service *DispatcherService) TreasuryWithdrawAndDispatch() (tx *types.Transaction, _err error) {
	gas, err := blockchain.BuildEstimateGas(service.owner, dispatcher.DispatcherABI, &service.dispatcherAddress,
		"treasuryWithdrawAndDispatch", service.treasuryAddress)
	if err != nil {
		variable.ZapLog.Error("error = ", zap.Error(err))
		_err = err
		return
	}
	variable.ZapLog.Info("gas = ", zap.Uint64("gas", gas))
	auth, _ := blockchain.BuildTransactor(service.privateKey)
	auth.GasLimit = uint64(gas + 120000) // in units
	return service.instance.TreasuryWithdrawAndDispatch(auth, service.treasuryAddress)
}

func (service *DispatcherService) ChainBridgeToWithdrawalAccount() (tx *types.Transaction, _err error) {
	gas, err := blockchain.BuildEstimateGas(service.owner, dispatcher.DispatcherABI, &service.dispatcherAddress,
		"chainBridgeToWithdrawalAccount", service.chainBridgeId, service.withdrawalTokenAddress, service.withdrawalAccountAddress)
	if err != nil {
		variable.ZapLog.Error("error = ", zap.Error(err))
		_err = err
		return
	}
	variable.ZapLog.Info("gas = ", zap.Uint64("gas", gas))
	auth, _ := blockchain.BuildTransactor(service.privateKey)
	auth.GasLimit = uint64(gas + 120000) // in units
	return service.instance.ChainBridgeToWithdrawalAccount(auth, service.chainBridgeId, service.withdrawalTokenAddress, service.withdrawalAccountAddress)
}

func (service *DispatcherService) CanChainBridgeToWithdrawalAccount() bool {
	if balanceOf, err := service.withdrawalTokenErc20.BalanceOf(nil, service.withdrawalAccountAddress); err != nil {
		return false
	} else if balanceOf.Cmp(service.withdrawalAccountThreshold) < 0 {
		return true
	}
	return false
}
