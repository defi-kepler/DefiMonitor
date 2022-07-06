package bootstrap

import (
	_ "goskeleton/app/core/destroy" //
	"goskeleton/app/global/variable"
	"goskeleton/app/service/sys_log_hook"

	"goskeleton/app/utils/yml_config"
	"goskeleton/app/utils/zap_factory"
)

func checkRequiredFolders() {
}

func init() {

	checkRequiredFolders()

	variable.ConfigYml = yml_config.CreateYamlFactory()
	variable.ConfigYml.ConfigFileChangeListen()

	variable.ConfigDefiYml = variable.ConfigYml.Clone("defi")
	variable.ConfigDefiYml.ConfigFileChangeListen()

	variable.ZapLog = zap_factory.CreateZapFactory(sys_log_hook.ZapLogHandler)

}
