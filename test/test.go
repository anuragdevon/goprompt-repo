package main

import (
	"fmt"
	"os"
	"os/exec"
)

func main() {
	// disable input buffering
	exec.Command("stty", "-F", "/dev/tty", "cbreak", "min", "1").Run()
	// do not display entered characters on the screen
	exec.Command("stty", "-F", "/dev/tty", "-echo").Run()

	var b []byte = make([]byte, 1)
	for {
		os.Stdin.Read(b)
		fmt.Println("I got the byte", b, "("+string(b)+")")
	}
}

// package main

// func inapp(input string) string {
// 	args := strings.Split(input, " ")
// 	cmd := exec.Command(args[0], args[1:]...)
// 	stdout, err := cmd.Output()
// 	check(err)
// 	return string(stdout)
// }
// func main() {
// 	app := "whoami"

// 	cmd := exec.Command(app)
// 	stdout, err := cmd.Output()
// 	user := string(stdout)
// 	if err != nil {
// 		fmt.Println(err.Error())
// 		return
// 	}

// 	// Print the output
// 	fmt.Print(user)
// }

// for _, err := rd.ReadString('\n'); err != io.EOF; _, err = rd.ReadString('\n') {
// 	// lastLineSize := len(line)
// 	// fmt.Print(lastLineSize)
// 	count += 1
// 	// break
// }
