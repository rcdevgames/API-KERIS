package main

import (
	"QAPI/configs"
	"QAPI/core"
	"QAPI/library"
	"flag"
)

var service = flag.String("service", "rest", "rest | job")
var port = flag.String("port", "3000", "Listen Port")

func main() {
	flag.Parse()
	library.InitLog()
	configs.ConnectDB()

	if *service == "job" {
		core.StartJob()
	} else {
		core.StartRest(*port)
	}
}
