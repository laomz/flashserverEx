package common

import (
	"os"
	"os/signal"
	"syscall"
)

var ServiceConfigNote = `start failed,need config file.`

func ServiceWait(moduleName string) {
	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT)
	<-c
	os.Exit(0)
}
