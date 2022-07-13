package cmd

import (
	"fmt"
	"strings"

	"github.com/pterm/pterm"
	"github.com/rr13k/pen/structure"
	"github.com/spf13/cobra"
)

var echoTimes int

var AppModelMap = map[string]string{
	"1": "api",
	"2": "gorm",
}

var cmdApp = &cobra.Command{
	Use:   "new",
	Short: "new app",
	Long:  `快速构建pen框架脚手架`,
	Args:  cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
		var appName string

		appName, _ = pterm.DefaultInteractiveTextInput.WithMultiLine(false).Show("app name")
		pterm.Println() // Blank line

		var options = []string{
			"1. only api",
			"2. api + gorm",
		}
		modleStr, _ := pterm.DefaultInteractiveSelect.WithOptions(options).Show()
		branch := AppModelMap[strings.Split(modleStr, ".")[0]]

		structure.Run(&structure.AppConfig{
			Name:   appName,
			Branch: branch,
		})

		fmt.Println(fmt.Sprintf("%s项目%s 创建成功! can you: cd %s", branch, appName, appName))
	},
}

func Cli() {
	var rootCmd = &cobra.Command{Use: "pen"}
	rootCmd.AddCommand(cmdApp)
	rootCmd.Execute()
}
