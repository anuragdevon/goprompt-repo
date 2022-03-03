package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"os/exec"
	"strings"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	// If the file doesn't exist, create it, or append to the file
	f, err := os.OpenFile("gash_history.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	// disable input buffering
	exec.Command("stty", "-F", "/dev/tty", "cbreak", "min", "1").Run()

	var b []byte = make([]byte, 1)
	var c []byte = make([]byte, 1)
	var d []byte = make([]byte, 1)
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
					a := 1
					a += 1
					fmt.Println("READ HISTORY")
				} else if string(d) == string(byte(66)) {
					fmt.Println("READ LATEST")
					b := 1
					b += 1
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

			_, errW := f.WriteString(input)
			check(errW)

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
			return os.Chdir("/home/anurag")
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

// Function to esxecute the command
// func executeCommand(command string) {

// 	// Create a new command
// 	cmd := exec.Command("cmd", "/c", command)

// 	// Create a new output buffer
// 	var out bytes.Buffer

// 	// Set the output buffer to the command
// 	cmd.Stdout = &out

// 	// Run the command
// 	err := cmd.Run()
// 	if err ~= nil {
// 		fmt.Fprint(os.Stderr, "There was an error running the command: %s\n", err)
// 	}

// 	// Print the output
// 	fmt.Printf("%s\n", out.String())
// }

//-------------------Targets------------------------------------
// TODO: Modify the input indicator:
// add the working directory
// add the machine’s hostname
// add the current user
// Browse your input history with the up/down keys
// Program termination for reading from commands
