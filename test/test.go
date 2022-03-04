package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	signalChanel := make(chan os.Signal, 1)
	signal.Notify(signalChanel,
		syscall.SIGHUP,
		syscall.SIGINT,
		syscall.SIGTERM,
		syscall.SIGQUIT)

	exit_chan := make(chan int)
	go func() {
		for {
			s := <-signalChanel
			fmt.Println("Signal Received: ", s)
			switch s {
			// kill -SIGHUP XXXX [XXXX - PID for your program]
			case syscall.SIGHUP:
				fmt.Println("Signal hang up triggered.")

				// kill -SIGINT XXXX or Ctrl+c  [XXXX - PID for your program]
			case syscall.SIGINT:
				fmt.Println("Signal interrupt triggered.")

				// kill -SIGTERM XXXX [XXXX - PID for your program]
			case syscall.SIGTERM:
				fmt.Println("Signal terminte triggered.")
				exit_chan <- 0

				// kill -SIGQUIT XXXX [XXXX - PID for your program]
			case syscall.SIGQUIT:
				fmt.Println("Signal quit triggered.")
				exit_chan <- 0

			default:
				fmt.Println("Unknown signal.")
				exit_chan <- 1
			}
		}
	}()
	exitCode := <-exit_chan
	os.Exit(exitCode)
}

// package main

// import (
// 	"fmt"
// 	"os"
// 	"os/exec"
// )

// func main() {
// 	fmt.Println("Test 1")
// 	fmt.Println("Test 2")
// 	fmt.Print("\033[1A\033[K") //one up, remove line (should work after the newline of the Println)
// }
// func main() {
// 	// disable input buffering
// 	exec.Command("stty", "-F", "/dev/tty", "cbreak", "min", "1").Run()
// 	// do not display entered characters on the screen
// 	exec.Command("stty", "-F", "/dev/tty", "-echo").Run()

// 	var b []byte = make([]byte, 1)
// 	for {
// 		os.Stdin.Read(b)
// 		fmt.Println("I got the byte", b, "("+string(b)+")")
// 	}
// }

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
