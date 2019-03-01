package modules

import (
	"gopkg.in/alecthomas/kingpin.v2"
	"fmt"
	"os"
	"testProject/core"
	"testProject/modules"
	"testProject/core/server"
	"github.com/sirupsen/logrus"
)

const VERSION = "0.1"

var BUILD = "debug"

var (
	app        = kingpin.New("little coin trader", fmt.Sprintf("Version: %s-%s", VERSION, BUILD))
	configFile = app.Flag("config-file", "config file path(default: ./app.yaml)").Short('c').Default("./app.yaml").String()
	apiServer  = app.Command("server", "start server").Default()
)
func main() {
	command := kingpin.MustParse(app.Parse(os.Args[1:]))
	core.InitConfig(*configFile)
	core.LoadCore()
	modules.LoadModules()
	switch command {
	case apiServer.FullCommand():
		logrus.Info("server start!!!")
		err := server.Start()
		if err != nil {
			logrus.Panic(err)
		}
	default:
		logrus.Error("wrong command!!!")
	}
	logrus.Info("Application exit!!!")
}
