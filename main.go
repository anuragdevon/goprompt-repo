package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"os/exec"
	"strings"
)

var LINE_NUMBER int = 0
var HIST_FILE string = "gash_history.log"

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func editGashHistory(input string) {
	f, err := os.OpenFile(HIST_FILE, os.O_APPEND|os.O_CREATE|os.O_WRONLY, os.ModePerm)
	check(err)
	defer f.Close()
	_, errW := f.WriteString(input)
	check(errW)
}

func readGashHistory(lineNumber int) string {
	f, err := os.OpenFile(HIST_FILE, os.O_RDONLY, os.ModePerm)
	check(err)
	defer f.Close()

	rd := bufio.NewReader(f)
	i := 0
	for line, err := rd.ReadString('\n'); err != io.EOF; line, err = rd.ReadString('\n') {
		i += 1
		if lineNumber == i {
			return line
		}
	}
	return ""
}

func total_lines() int {
	f, err := os.OpenFile(HIST_FILE, os.O_RDONLY, os.ModePerm)
	check(err)
	defer f.Close()
	count := 0

	rd := bufio.NewReader(f)
	for _, err := rd.ReadString('\n'); err != io.EOF; _, err = rd.ReadString('\n') {
		count += 1
	}
	return count
}

func decisionTree(b []byte) {
	var c []byte = make([]byte, 1)
	var d []byte = make([]byte, 1)

	if string(b) == string(byte(27)) {
		os.Stdin.Read(c)
		os.Stdin.Read(d)

		if string(c) == string(byte(91)) {
			if string(d) == string(byte(65)) {
				LINE_NUMBER -= 1
				// read history
				input := readGashHistory(LINE_NUMBER)
				input = strings.TrimSuffix(input, "\n")
				fmt.Print("\n\033[1A\033[K")
				promt()
				fmt.Print(input)
				os.Stdin.Read(b)
				// recursive
				decisionTree(b)
				reader := bufio.NewReader(os.Stdin)
				addtional_input, err := reader.ReadString('\n')
				check(err)
				fmt.Println()
				input += addtional_input
				if err = execInput(input); err != nil {
					fmt.Fprintln(os.Stderr, err)
				}

			} else if string(d) == string(byte(66)) {
				// read latest
				readGashHistory(LINE_NUMBER)
			}
		}
	} else {
		fmt.Print(string(b))

		// Enable chacter display on screen
		exec.Command("stty", "-F", "/dev/tty", "echo").Run()
		reader := bufio.NewReader(os.Stdin)
		input, err := reader.ReadString('\n')
		check(err)
		input = string(b) + input
		editGashHistory(input)

		if err = execInput(input); err != nil {
			fmt.Fprintln(os.Stderr, err)
		}
	}
}

func execInput(input string) error {

	input = strings.TrimSuffix(input, "\n")

	args := strings.Split(input, " ")

	switch args[0] {
	case "cd":
		// 'cd' to home dir with empty path not yet supported.
		if len(args) < 2 {
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

func promt() {
	path, err := os.Getwd()
	check(err)
	path = strings.Replace(string(path), "/home/anurag", "~", 1)
	fmt.Print(path, " > ")
}

func main() {
	// disable input buffering
	exec.Command("stty", "-F", "/dev/tty", "cbreak", "min", "1").Run()
	var b []byte = make([]byte, 1)
	promt()
	for {
		// disble chacter display on screen
		exec.Command("stty", "-F", "/dev/tty", "-echo").Run()
		LINE_NUMBER = total_lines() + 1

		os.Stdin.Read(b)
		decisionTree(b)
	}
}
