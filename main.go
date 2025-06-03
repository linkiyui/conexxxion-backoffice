package main

import (
	"flag"

	"gitlab.com/conexxxion/conexxxion-backoffice/config"
	server "gitlab.com/conexxxion/conexxxion-backoffice/server"
)

func main() {
	configFilePath := flag.String("config.file", "", "Configuration file path")
	flag.Parse()
	config.Init(*configFilePath)
	server.Start()
}
