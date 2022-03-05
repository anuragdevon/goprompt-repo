package history

import (
	"bufio"
	"fmt"
	"io"
	"os"

	"gash/globals"
)

func check(e error) {
	if e != nil {
		// panic(e)
		fmt.Println(e)
	}
}

func EditGashHistory(input string) {
	f, err := os.OpenFile(globals.HIST_FILE, os.O_APPEND|os.O_CREATE|os.O_WRONLY, os.ModePerm)
	check(err)
	defer f.Close()
	_, errW := f.WriteString(input)
	check(errW)
}

func ReadGashHistory(lineNumber int) string {
	f, err := os.OpenFile(globals.HIST_FILE, os.O_RDONLY, os.ModePerm)
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

func FileLines() int {
	f, err := os.OpenFile(globals.HIST_FILE, os.O_RDONLY, os.ModePerm)
	check(err)
	defer f.Close()
	count := 0

	rd := bufio.NewReader(f)
	for _, err := rd.ReadString('\n'); err != io.EOF; _, err = rd.ReadString('\n') {
		count += 1
	}
	return count
}
