package history

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func check(e error) {
	if e != nil {
		// panic(e)
		fmt.Println(e)
	}
}

var HIST_FILE string = "/home/anurag/docs/gash/history/gash_history.log"

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
