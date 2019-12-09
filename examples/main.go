package main

import (
	"fmt"
	"math/rand"
	"time"

	logger "github.com/casonadams/simple-logger"
)

var log logger.Logger

func main() {
	log = *logger.NewLogger("test")
	log.Debug("Debug message")
	log.Trace("Trace message")
	info()
	warn()
	log.Error("Error message")
	log.Fatal("Fatal message")
	log.Panic("Panic message")

	// code never reaches here comment out above to see how fast the logger is

	go info2()
	go warn2()
	go error2()
	go trace2()
	debug2()
}

func trace2() {
	for {
		s := time.Now()
		log.Trace(fmt.Sprintf("\033[%dm%s\033[0m", 96, "trace message"))
		e := time.Now()
		fmt.Printf("%v\n", e.Sub(s))
		n := rand.Intn(1000) // n will be between 0 and 1000
		time.Sleep(time.Duration(n) * time.Microsecond)
	}
}

func info2() {
	for {
		s := time.Now()
		log.Info(fmt.Sprintf("\033[%dm%s\033[0m", 34, "info message"))
		e := time.Now()
		fmt.Printf("%v\n", e.Sub(s))
		n := rand.Intn(1000) // n will be between 0 and 1000
		time.Sleep(time.Duration(n) * time.Microsecond)
	}
}

func error2() {
	for {
		s := time.Now()
		log.Error(fmt.Sprintf("\033[%dm%s\033[0m", 91, "error message"))
		e := time.Now()
		fmt.Printf("%v\n", e.Sub(s))
		n := rand.Intn(1000) // n will be between 0 and 1000
		time.Sleep(time.Duration(n) * time.Microsecond)
	}
}

func warn2() {
	for {
		s := time.Now()
		log.Warn(fmt.Sprintf("\033[%dm%s\033[0m", 93, "warn message"))
		e := time.Now()
		fmt.Printf("%v\n", e.Sub(s))
		n := rand.Intn(1000) // n will be between 0 and 1000
		time.Sleep(time.Duration(n) * time.Microsecond)
	}
}

func debug2() {
	for {
		s := time.Now()
		log.Debug("debug message")
		e := time.Now()
		fmt.Printf("%v\n", e.Sub(s))
		n := rand.Intn(1000) // n will be between 0 and 1000
		time.Sleep(time.Duration(n) * time.Microsecond)
	}
}
