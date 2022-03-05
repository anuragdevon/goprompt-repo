package main

// package unixSignals

import (
	"fmt"
	"os"
	"os/exec"
	"os/signal"
	"syscall"
)

func main() {
	signalChanel := make(chan os.Signal, 1)
	signal.Notify(signalChanel,
		syscall.SIGINT)

	exit_chan := make(chan int)

	exec.Command("stty", "-F", "/dev/tty", "cbreak", "min", "1").Run()
	// do not display entered characters on the screen
	exec.Command("stty", "-F", "/dev/tty", "-echo").Run()

	var b []byte = make([]byte, 1)
	os.Stdin.Read(b)
	fmt.Println("Current Byte: ", b, "("+string(b)+")")

	s := <-signalChanel

	switch s {
	case syscall.SIGINT:
		fmt.Println("Signal interrupt triggered.")
		os.Exit(1)

	default:
		fmt.Println("Unknown signal.")
		exit_chan <- 1
	}
	// }
	// }()
	exitCode := <-exit_chan
	os.Exit(exitCode)
}
