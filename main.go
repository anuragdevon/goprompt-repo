package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"os/exec"
	"strings"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func editGashHistory(input string) {
	f, err := os.OpenFile("gash_history.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, os.ModePerm)
	check(err)
	defer f.Close()
	_, errW := f.WriteString(input)
	check(errW)
}

func readGashHistory() {
	f, err := os.OpenFile("gash_history.log", os.O_RDONLY, os.ModePerm)
	check(err)
	defer f.Close()
	// count := 0

	rd := bufio.NewReader(f)
	// for _, err := rd.ReadString('\n'); err != io.EOF; _, err = rd.ReadString('\n') {
	// 	// lastLineSize := len(line)
	// 	// fmt.Print(lastLineSize)
	// 	count += 1
	// 	// break
	// }
	i := 0
	lineNumber := 2
	for line, err := rd.ReadString('\n'); err != io.EOF; line, err = rd.ReadString('\n') {
		i += 1
		if lineNumber == i {
			fmt.Printf("%s", line)

		}
	}
}

func main() {
	// disable input buffering
	exec.Command("stty", "-F", "/dev/tty", "cbreak", "min", "1").Run()

	var b []byte = make([]byte, 1)
	var c []byte = make([]byte, 1)
	var d []byte = make([]byte, 1)
	var con []byte = make([]byte, 1)

	for {
		// disble chacter display on screen
		exec.Command("stty", "-F", "/dev/tty", "-echo").Run()
		path, err := os.Getwd()
		if err != nil {
			fmt.Println(err)
		}
		path = strings.Replace(string(path), "/home/anurag", "~", 1)
		fmt.Print(path, " > ")
		os.Stdin.Read(b)
		if string(b) == string(byte(27)) {
			os.Stdin.Read(c)
			os.Stdin.Read(d)
			if string(c) == string(byte(91)) {
				if string(d) == string(byte(65)) {
					// read history
					readGashHistory()
					os.Stdin.Read(con)
					if string(b) == "\n" {
						fmt.Print("EXECUTING...")
					}

				} else if string(d) == string(byte(66)) {
					// read latest
					readGashHistory()
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
