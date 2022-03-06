package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"strings"
	"time"

	"gash/globals"
	"gash/history"
	"gash/prompt"
	"gash/unixSignals"
)

func check(e error) {
	if e != nil {
		// panic(e)
		fmt.Println(e)
	}
}

func decisionTree(b []byte, executionStatus bool, prevCommand string) bool {
	if !executionStatus {
		var c []byte = make([]byte, 1)
		var d []byte = make([]byte, 1)

		if string(b) == string(byte(27)) {
			os.Stdin.Read(c)
			os.Stdin.Read(d)

			if string(c) == string(byte(91)) {
				if string(d) == string(byte(65)) {
					// read history
					globals.LineNumber -= 1
					input := history.ReadGashHistory(globals.LineNumber)
					input = strings.TrimSuffix(input, "\n")
					fmt.Print(globals.ClearLine)
					prompt.Prompt()
					fmt.Print(input)
					prevCommand = input
					os.Stdin.Read(b)

					executionStatus = decisionTree(b, executionStatus, prevCommand)

				} else if string(d) == string(byte(66)) {
					// read latest
					globals.LineNumber += 1
					input := history.ReadGashHistory(globals.LineNumber)
					input = strings.TrimSuffix(input, "\n")
					fmt.Print(globals.ClearLine)
					prompt.Prompt()
					fmt.Print(input)
					prevCommand = input
					os.Stdin.Read(b)

					executionStatus = decisionTree(b, executionStatus, prevCommand)
				}
			}
		} else {
			input := ""
			if prevCommand == "" {
				fmt.Print(string(b))

				// Enable chacter display on screen
				exec.Command("stty", "-F", "/dev/tty", "echo").Run()
				reader := bufio.NewReader(os.Stdin)
				input, err := reader.ReadString('\n')

				check(err)
				input = string(b) + input

				history.EditGashHistory(input)
				executionStatus = true

				t1 := time.Now()

				if err := execInput(input); err != nil {
					fmt.Fprintln(os.Stderr, err)
				}

				t2 := time.Now()
				diff := t2.Sub(t1)
				fmt.Println("time elapsed: ", diff)

			} else {
				fmt.Print(string(b))
				input = prevCommand + string(b)
				input = strings.TrimSuffix(input, "\n")
				// Enable chacter display on screen
				exec.Command("stty", "-F", "/dev/tty", "echo").Run()
				reader := bufio.NewReader(os.Stdin)
				extra_input, err := reader.ReadString('\n')
				check(err)

				input = input + extra_input

				history.EditGashHistory(input)
				executionStatus = true

				t1 := time.Now()

				if err := execInput(input); err != nil {
					fmt.Fprintln(os.Stderr, err)
				}

				t2 := t1.Add(time.Second * 341)
				diff := t2.Sub(t1)
				fmt.Println("time elapsed: ", diff)
			}

		}
	}
	return executionStatus
}

func execInput(input string) error {

	input = strings.TrimSuffix(input, "\n")

	args := strings.Split(input, " ")

	switch args[0] {
	case "cd":
		// 'cd' to home dir with empty path not yet supported.
		if len(args) < 2 || args[1] == "" {
			dir := "/home/" + "anurag"
			return os.Chdir(dir)
		}
		// Change the directory and return the error
		return os.Chdir(args[1])
	case "exit":
		os.Exit(0)
	}
	cmd := exec.Command(args[0], args[1:]...)

	// Set the correct output device
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	return cmd.Run()
}

func main() {
	// disable input buffering
	// TODO: check validity of this
	exec.Command("stty", "-F", "/dev/tty", "cbreak", "min", "1").Run()
	var b []byte = make([]byte, 1)
	executionStatus := false
	for {
		// disble chacter display on screen
		exec.Command("stty", "-F", "/dev/tty", "-echo").Run()
		globals.LineNumber = history.FileLines() + 1
		prevCommand := ""
		prompt.Prompt()

		// Invoke Unix Signals Handler
		go unixSignals.SingHandler()

		os.Stdin.Read(b)
		decisionTree(b, executionStatus, prevCommand)
	}
}
