package prompt

import (
	"fmt"
	"os"
	"strings"

	"github.com/ttacon/chalk"
)

func check(e error) {
	if e != nil {
		// panic(e)
		fmt.Println(e)
	}
}

func Prompt() {
	hostname, err := os.Hostname()
	// os.
	check(err)
	path, err := os.Getwd()
	check(err)
	path = strings.Replace(string(path), "/home/anurag", "~", 1)
	CyanSt := chalk.Cyan.NewStyle().WithTextStyle(chalk.Bold).WithBackground(chalk.ResetColor)
	GreenSt := chalk.Blue.NewStyle().WithTextStyle(chalk.Italic).WithBackground(chalk.ResetColor)
	MagentaSt := chalk.Magenta.NewStyle().WithTextStyle(chalk.Bold).WithBackground(chalk.ResetColor)
	fmt.Print(CyanSt.Style(path), "(", GreenSt.Style(hostname), ")", MagentaSt.Style(" > "))
}
