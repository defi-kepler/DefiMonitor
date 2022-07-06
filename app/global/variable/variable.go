package variable

import (
	"goskeleton/app/global/my_errors"

	"goskeleton/app/utils/yml_config/ymlconfig_interf"
	"log"
	"os"
	"strings"

	"go.uber.org/zap"
)

var (
	BasePath           string
	EventDestroyPrefix = "Destroy_"
	ConfigKeyPrefix    = "Config_"
	DateFormat         = "2006-01-02 15:04:05"

	ZapLog *zap.Logger

	ConfigYml       ymlconfig_interf.YmlConfigInterf
	ConfigGormv2Yml ymlconfig_interf.YmlConfigInterf
	ConfigDefiYml   ymlconfig_interf.YmlConfigInterf

	//websocket
	WebsocketHub              interface{}
	WebsocketHandshakeSuccess = `{"code":200,"msg":"ws","data":""}`
	WebsocketServerPingMsg    = "Server->Ping->Client"
)

func init() {

	if curPath, err := os.Getwd(); err == nil {

		if len(os.Args) > 1 && strings.HasPrefix(os.Args[1], "-test") {
			BasePath = strings.Replace(strings.Replace(curPath, `\test`, "", 1), `/test`, "", 1)

		} else {
			BasePath = curPath
			//BasePath = "/share/project/goproject/src/DefiMonitor"
		}
	} else {
		log.Fatal(my_errors.ErrorsBasePath)
	}
}
