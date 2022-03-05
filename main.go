package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"strings"

	"gash/history"
	"gash/prompt"
)

var lineNumber int = 0

const ClearLine = "\n\033[1A\033[K"

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
					lineNumber -= 1
					input := history.ReadGashHistory(lineNumber)
					input = strings.TrimSuffix(input, "\n")
					fmt.Print(ClearLine)
					prompt.Prompt()
					fmt.Print(input)
					prevCommand = input
					os.Stdin.Read(b)

					executionStatus = decisionTree(b, executionStatus, prevCommand)

				} else if string(d) == string(byte(66)) {
					// read latest
					lineNumber += 1
					input := history.ReadGashHistory(lineNumber)
					input = strings.TrimSuffix(input, "\n")
					fmt.Print(ClearLine)
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
				// fmt.Print("\b")
				check(err)
				input = string(b) + input

				history.EditGashHistory(input)
				executionStatus = true

				if err = execInput(input); err != nil {
					fmt.Fprintln(os.Stderr, err)
				}

			} else {
				fmt.Print(string(b))
				input = prevCommand + string(b)
				// input = strings.TrimSuffix(input, "\n")
				// // Enable chacter display on screen
				// exec.Command("stty", "-F", "/dev/tty", "echo").Run()
				// reader := bufio.NewReader(os.Stdin)
				// extra_input, err := reader.ReadString('\n')
				// check(err)

				// input = input + extra_input

				history.EditGashHistory(input)
				executionStatus = true

				if err := execInput(input); err != nil {
					fmt.Fprintln(os.Stderr, err)
				}
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
	// exec.Command("stty", "-F", "/dev/tty", "cbreak", "min", "1").Run()
	var b []byte = make([]byte, 1)
	executionStatus := false
	for {
		// disble chacter display on screen
		// exec.Command("stty", "-F", "/dev/tty", "-echo").Run()
		lineNumber = history.FileLines() + 1
		prevCommand := ""
		prompt.Prompt()
		os.Stdin.Read(b)
		decisionTree(b, executionStatus, prevCommand)
		// unixSignals()
	}
}
