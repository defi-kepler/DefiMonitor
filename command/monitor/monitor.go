package monitor

import (
	"goskeleton/app/global/variable"
	"goskeleton/app/service/monitor"
	"goskeleton/app/utils/blockchain"
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
	ACTION_WITHDRAW      uint8 = 0
	ACTION_DISPATCH      uint8 = 1
	ACTION_EXECUTE       uint8 = 100
	TIME_DAY             int64 = 86400
	ACTION_DISPATCH_TIME int64 = 21600 //6h
	ACTION_WITHDRAW_TIME int64 = 300   // 5m
)

var (
	LogAction string
	Date      string
	lastHash  map[uint8]*TxHash = make(map[uint8]*TxHash)
	logger                      = variable.ZapLog.Sugar()
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
	dispatcherService := monitor.CreateDispatcherService()
	//Check whether there is a balance in the cash out contract every minute, and call the contract if it reaches the threshold
	crontab.AddFunc("0 */1 * * * ?", func() { TreasuryWithdrawAndDispatchFn(dispatcherService) })
	//Transfer from the cash out contract to the dispatch contract and distribute it every 6 hours
	crontab.AddFunc("0 */1 * * * ?", func() { DispatchFn(dispatcherService) })
	crontab.Start()
	select {}
}

//From recharge contract withdrawal to dispatching contract
func TreasuryWithdrawAndDispatchFn(dispatcherService *monitor.DispatcherService) {
	logger.Info("TreasuryWithdrawAndDispatchFn")

	var lastTx *TxHash = lastHash[ACTION_WITHDRAW]
	if CanCall(lastTx, ACTION_WITHDRAW_TIME) {
		if dispatcherService.CanChainBridgeToWithdrawalAccount() {
			tx, _ := dispatcherService.ChainBridgeToWithdrawalAccount()
			if tx != nil {
				variable.ZapLog.Info("---", zap.String("tx", tx.Hash().Hex()))
				lastHash[ACTION_WITHDRAW] = &TxHash{tx.Hash().Hex(), time.Now()}
			}
		}

	}
}

func DispatchFn(dispatcherService *monitor.DispatcherService) {
	logger.Info("Dispatch ....")
	var lastTx *TxHash = lastHash[ACTION_DISPATCH]
	if CanCall(lastTx, ACTION_DISPATCH_TIME) {
		tx, _ := dispatcherService.TreasuryWithdrawAndDispatch()
		if tx != nil {
			variable.ZapLog.Info("---", zap.String("tx", tx.Hash().Hex()))
			lastHash[ACTION_DISPATCH] = &TxHash{tx.Hash().Hex(), time.Now()}
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
