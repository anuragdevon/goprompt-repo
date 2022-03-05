package unixSignals

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
)

func unixSignals() {
	signalChanel := make(chan os.Signal, 1)
	signal.Notify(signalChanel,
		syscall.SIGINT)

	exit_chan := make(chan int)
	go func() {
		for {
			s := <-signalChanel
			fmt.Println("Signal Received: ", s)
			switch s {
			case syscall.SIGINT:
				fmt.Println("Signal interrupt triggered.")

			default:
				fmt.Println("Unknown signal.")
				exit_chan <- 1
			}
		}
	}()
	exitCode := <-exit_chan
	os.Exit(exitCode)
}
