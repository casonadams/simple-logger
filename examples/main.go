package main

import logger "github.com/casonadams/simple-logger/v2"

func main() {
	log := logger.NewLogger("test")
	log.Debug("debug message")
	log.Trace("trace message")
	log.Info("Hello World")
	log.Warn("Warn message")
	log.Error("this is an error")
	log.Panic("panic")
}
