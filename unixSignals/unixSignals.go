package unixSignals

// package unixSignals

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"gash/globals"
	"gash/prompt"
)

func SingHandler() {
	signalChanel := make(chan os.Signal, 1)
	signal.Notify(signalChanel,
		syscall.SIGINT)

	for {
		s := <-signalChanel

		switch s {
		case syscall.SIGINT:
			fmt.Println()
			fmt.Print(globals.ClearLine)
			prompt.Prompt()
			continue

			// default:
			// 	fmt.Println("Unknown signal.")
			// 	fmt.Println()
			// 	fmt.Print(globals.ClearLine)
			// 	prompt.Prompt()
			// 	continue

		}
	}
}
