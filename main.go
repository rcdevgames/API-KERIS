package main

import (
	"QAPI/configs"
	"QAPI/core"
	"QAPI/logger"
	"flag"
)

var service = flag.String("service", "rest", "rest | job")
var port = flag.String("port", "3000", "Listen Port")

func main() {
	flag.Parse()
	logger.InitLog()
	configs.ConnectDB()

	if *service == "job" {
		core.StartJob()
	} else {
		core.StartRest(*port)
	}
}
