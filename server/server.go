package server

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"gitlab.com/conexxxion/conexxxion-backoffice/api_server"
	clog "gitlab.com/conexxxion/conexxxion-backoffice/logger"
	"gitlab.com/conexxxion/conexxxion-backoffice/translations"
)

func Start() {
	sigc := make(chan os.Signal, 1)
	signal.Notify(sigc,
		os.Kill,
		os.Interrupt,
		syscall.SIGHUP,
		syscall.SIGINT,
		syscall.SIGTERM,
		syscall.SIGQUIT)
	go func() {
		s := <-sigc
		fmt.Println("terminating CNXXXN-BACKOFFICE | signal:", s)
		clog.Finish()
		os.Exit(0)
	}()

	clog.Init()
	translations.LoadTranslations()
	go api_server.StartMetricsServer()

	api_server.NewApiServer().Start()
}
