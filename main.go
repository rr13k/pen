package main

import (
	"bufio"

	"github.com/pterm/pterm"
	"github.com/rr13k/pen/cmd"
)

func main() {
	cmd.Cli()
	// Test_F5()

}

func boolToText(b bool) string {
	if b {
		return pterm.Green("Yes")
	}
	return pterm.Red("No")
}

// 测试刷新命令行显示
func Test_F5() {

	result, _ := pterm.DefaultInteractiveTextInput.WithMultiLine(false).Show("app name:")
	pterm.Println() // Blank line
	pterm.Info.Printfln("You answered: %s", result)

	var options = []string{
		"1. only api",
		"2. api + gorm",
	}
	selectedOption, _ := pterm.DefaultInteractiveSelect.WithOptions(options).Show()
	pterm.Info.Printfln("Selected option: %s", pterm.Green(selectedOption))
}

func flush(reader *bufio.Reader) {
	var i int
	for i = 0; i < reader.Buffered(); i++ {
		reader.ReadByte()
	}
}
