package destroy

import (
	"goskeleton/app/core/event_manage"
	"goskeleton/app/global/consts"
	"goskeleton/app/global/variable"
	"os"
	"os/signal"
	"syscall"

	"go.uber.org/zap"
)

func init() {

	go func() {
		c := make(chan os.Signal)
		signal.Notify(c, os.Interrupt, os.Kill, syscall.SIGQUIT, syscall.SIGINT, syscall.SIGTERM) //
		received := <-c                                                                           //
		variable.ZapLog.Warn(consts.ProcessKilled, zap.String("signal", received.String()))
		(event_manage.CreateEventManageFactory()).FuzzyCall(variable.EventDestroyPrefix)
		close(c)
		os.Exit(1)
	}()

}
