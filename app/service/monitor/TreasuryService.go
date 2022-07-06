package monitor

import (
	"goskeleton/abi/defi/erc20"
	"goskeleton/abi/defi/treasury"
	"goskeleton/app/global/variable"
	"goskeleton/app/utils/blockchain"
	"goskeleton/hdwallet"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"go.uber.org/zap"
)

type TreasuryService struct {
	owner           accounts.Account
	instance        *treasury.Treasury
	treasuryAddress common.Address
}

func CreateTreasuryService(treasuryAddressStr string) *TreasuryService {
	treasuryAddress := common.HexToAddress(variable.ConfigDefiYml.GetString("BUSDTreasury"))
	instance, err := treasury.NewTreasury(treasuryAddress, blockchain.Client)

	if err != nil {
		panic(err)
	}

	path := hdwallet.MustParseDerivationPath(variable.ConfigDefiYml.GetString("TreasuryOwnerPath"))
	owner, err := blockchain.Wallet.Derive(path, false)
	variable.ZapLog.Info("owner = ", zap.String("owner", owner.Address.Hex()))
	if err != nil {
		panic(err)
	}
	return &TreasuryService{owner, instance, treasuryAddress}
}

func (s *TreasuryService) Withdraw() (tx *types.Transaction, _err error) {
	balance, err := s.TreasuryBalanceOf()
	if err != nil || balance.Uint64() == 0 {
		_err = err
		return
	}
	variable.ZapLog.Info("amount,", zap.Uint64("amount", balance.Uint64()))

	dispatcher := common.HexToAddress(variable.ConfigDefiYml.GetString("Dispatcher"))
	gas, err := blockchain.BuildEstimateGas(s.owner, treasury.TreasuryABI, &s.treasuryAddress, "withdraw", dispatcher)
	if err != nil {
		variable.ZapLog.Error("error = ", zap.Error(err))
		_err = err
		return
	}
	variable.ZapLog.Info("gas = ", zap.Uint64("gas", gas))
	auth, _ := blockchain.BuildTransactor(s.owner)
	auth.GasLimit = uint64(gas + 120000) // in units
	return s.instance.Withdraw(auth, dispatcher)
}

func (s *TreasuryService) TreasuryBalanceOf() (*big.Int, error) {
	tokenAddres, err := s.instance.Token(nil)
	if err != nil {
		return big.NewInt(0), err
	}
	tokenInstance, err := erc20.NewErc20(tokenAddres, blockchain.Client)
	if err != nil {
		return big.NewInt(0), err
	}
	return tokenInstance.BalanceOf(nil, s.treasuryAddress)
}
