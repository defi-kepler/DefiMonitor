package monitor

import (
	"goskeleton/abi/defi/dispatcher"
	"goskeleton/abi/defi/erc20"
	"goskeleton/app/global/variable"
	"goskeleton/app/utils/blockchain"
	"goskeleton/hdwallet"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"go.uber.org/zap"
)

type DispatcherService struct {
	owner             accounts.Account
	instance          *dispatcher.Dispatcher
	dispatcherAddress common.Address
}

func CreateDispatcherService() *DispatcherService {
	dispatcherAddress := common.HexToAddress(variable.ConfigDefiYml.GetString("Dispatcher"))
	instance, err := dispatcher.NewDispatcher(dispatcherAddress, blockchain.Client)
	if err != nil {
		panic(err)
	}
	path := hdwallet.MustParseDerivationPath(variable.ConfigDefiYml.GetString("DispatcherOwnerPath"))
	owner, err := blockchain.Wallet.Derive(path, false)
	variable.ZapLog.Info("owner = ", zap.String("owner", owner.Address.Hex()))
	if err != nil {
		panic(err)
	}
	return &DispatcherService{owner, instance, dispatcherAddress}
}

func (service *DispatcherService) Dispatch() (tx *types.Transaction, _err error) {
	gas, err := blockchain.BuildEstimateGas(service.owner, dispatcher.DispatcherABI, &service.dispatcherAddress, "dispatch")
	if err != nil {
		variable.ZapLog.Error("error = ", zap.Error(err))
		_err = err
		return
	}
	variable.ZapLog.Info("gas = ", zap.Uint64("gas", gas))
	auth, _ := blockchain.BuildTransactor(service.owner)
	auth.GasLimit = uint64(gas + 120000) // in units
	return service.instance.Dispatch(auth)
}

func (service *DispatcherService) Execute(pid *big.Int) (tx *types.Transaction, _err error) {
	gas, err := blockchain.BuildEstimateGas(service.owner, dispatcher.DispatcherABI, &service.dispatcherAddress, "execute", pid)
	if err != nil {
		variable.ZapLog.Error("error = ", zap.Error(err))
		_err = err
		return
	}
	variable.ZapLog.Info("gas = ", zap.Uint64("gas", gas))
	auth, _ := blockchain.BuildTransactor(service.owner)
	auth.GasLimit = uint64(gas + 120000) // in units
	return service.instance.Execute(auth, pid)
}

func (service *DispatcherService) Strategys(pid *big.Int) (struct {
	Strategy common.Address
	Point    *big.Int
}, error) {
	return service.instance.Strategys(nil, pid)
}

func (service *DispatcherService) CanExecute(pid *big.Int) bool {
	token0, _ := service.instance.Token0(nil)
	token1, _ := service.instance.Token1(nil)
	tokenInstance0, _ := erc20.NewErc20(token0, blockchain.Client)
	tokenInstance1, _ := erc20.NewErc20(token1, blockchain.Client)

	balance0, _ := tokenInstance0.BalanceOf(nil, service.dispatcherAddress)
	balance1, _ := tokenInstance1.BalanceOf(nil, service.dispatcherAddress)
	if (balance0 == nil || balance0.Uint64() == 0) && (balance1 == nil || balance1.Uint64() == 0) {
		return false
	}
	strategy, _ := service.instance.Strategys(nil, pid)
	allowance0, _ := tokenInstance0.Allowance(nil, service.dispatcherAddress, strategy.Strategy)
	allowance1, _ := tokenInstance1.Allowance(nil, service.dispatcherAddress, strategy.Strategy)
	if (allowance0 == nil || allowance0.Uint64() == 0) && (allowance1 == nil || allowance1.Uint64() == 0) {
		return false
	}
	if strategy.Point.Uint64() == 0 {
		return false
	}
	return true
}
