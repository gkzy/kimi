package conn

import (
	"github.com/gkzy/gow"
	"github.com/gkzy/gow/lib/config"
	"github.com/gkzy/gow/lib/logy"
	"os"
)

const (
	ServerName = "login-service"
)

// InitLog init logy
func InitLog() {
	runMode := config.DefaultString("run_mode", "dev")
	if runMode == gow.ProdMode || runMode == gow.TestMode {
		logy.SetOutput(
			logy.MultiWriter(
				logy.NewWriter(os.Stdout),
				logy.NewFileWriter(logy.FileWriterOptions{
					Dir:           "./logs",
					Prefix:        "web",
					StorageType:   logy.StorageTypeDay,
					StorageMaxDay: 7,
				}),
			),
			ServerName,
		)
	} else {
		logy.SetOutput(
			logy.WithColor(logy.NewWriter(os.Stdout)),
			ServerName,
		)
	}

	logy.Infof("-------------------------------------------")
	logy.Infof("Start %s ......", ServerName)
	logy.Infof("-------------------------------------------")
}
