package monitor

import (
	"goskeleton/app/global/variable"
	"goskeleton/app/service/monitor"
	"goskeleton/app/utils/blockchain"
	"math/big"
	"time"

	"github.com/robfig/cron"
	"github.com/spf13/cobra"
	"go.uber.org/zap"
)

type TxHash struct {
	tx   string
	time time.Time
}

const (
	ACTION_WITHDRAW uint8 = 0
	ACTION_DISPATCH uint8 = 1
	ACTION_EXECUTE  uint8 = 100
	TIME_DAY        int64 = 86400
)

var (
	LogAction       string
	Date            string
	BUSDTreasuryAdd string            = variable.ConfigDefiYml.GetString("BUSDTreasury")
	lastHash        map[uint8]*TxHash = make(map[uint8]*TxHash)
	logger                            = variable.ZapLog.Sugar()
)

var Monitor = &cobra.Command{
	Use:     "m",
	Aliases: []string{"monitor"}, //
	Short:   "",
	Long: `Call method:
	1. Enter the project root directory (ginkeleton).
	2. Execute go run cmd/cli/main go demo_ Simple -h / / you can view the user guide
	3. Execute go run cmd/cli/main go demo_ Simple -a create / / execute the corresponding command through the action action action
	`,

	Run: func(cmd *cobra.Command, args []string) {
		start(LogAction, Date)
	},
}

func init() {

}

func start(actionName, Date string) {
	crontab := cron.New()

	spec := "*/10 * * * * ?"
	crontab.AddFunc(spec, func() { WithdrawFn(BUSDTreasuryAdd) })
	crontab.AddFunc(spec, DispatchFn)
	crontab.AddFunc(spec, ExecuteFn)
	crontab.Start()
	select {}
}

//From recharge contract withdrawal to dispatching contract
func WithdrawFn(treasuryAdd string) {
	logger.Info("WithdrawFn")

	var lastTx *TxHash = lastHash[ACTION_WITHDRAW]
	if CanCall(lastTx, TIME_DAY) {
		treasuryService := monitor.CreateTreasuryService(treasuryAdd)
		tx, _ := treasuryService.Withdraw()
		if tx != nil {
			variable.ZapLog.Info("---", zap.String("tx", tx.Hash().Hex()))
			lastHash[ACTION_WITHDRAW] = &TxHash{tx.Hash().Hex(), time.Now()}
		}
	}
}

func DispatchFn() {
	logger.Info("Dispatch ....")
	var lastTx *TxHash = lastHash[ACTION_DISPATCH]
	if CanCall(lastTx, TIME_DAY) {
		dispatcherService := monitor.CreateDispatcherService()
		tx, _ := dispatcherService.Dispatch()
		if tx != nil {
			variable.ZapLog.Info("---", zap.String("tx", tx.Hash().Hex()))
			lastHash[ACTION_DISPATCH] = &TxHash{tx.Hash().Hex(), time.Now()}
		}
	}
}

func ExecuteFn() {

	const LENGTH uint8 = 2
	var i uint8 = 0
	dispatcherService := monitor.CreateDispatcherService()
	for ; i < LENGTH; i++ {
		var lastTx *TxHash = lastHash[ACTION_EXECUTE+i]
		if CanCall(lastTx, TIME_DAY) {
			logger.Info("Execute ", zap.Uint("i", uint(i)))
			if !dispatcherService.CanExecute(big.NewInt(int64(i))) {
				continue
			}
			tx, _ := dispatcherService.Execute(big.NewInt(int64(i)))
			if tx != nil {
				variable.ZapLog.Info("---", zap.String("tx", tx.Hash().Hex()))
				lastHash[ACTION_EXECUTE+i] = &TxHash{tx.Hash().Hex(), time.Now()}
			}
		}
	}

}

func CanCall(tx *TxHash, interval int64) bool {
	if tx == nil {
		return true
	}
	now := time.Now().Unix()
	lastTime := tx.time.Unix()
	if now-lastTime > interval {
		return true
	}
	receipt, err := blockchain.QueryTx(tx.tx)
	if err != nil || receipt == nil {
		return false
	}
	return receipt.Status != 1
}
