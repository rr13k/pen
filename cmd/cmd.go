package cmd

import (
	"fmt"

	"github.com/rr13k/pen/structure"
	"github.com/spf13/cobra"
)

var echoTimes int

var cmdApp = &cobra.Command{
	Use:   "new",
	Short: "new app",
	Long:  `快速构建pen框架脚手架`,
	Args:  cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
		var appName string
		var unitTestFlag string
		var unitTest bool

		fmt.Print("app name:")
		fmt.Scan(&appName)

		fmt.Print("use unit test (y/n):")
		fmt.Scan(&unitTestFlag)

		if unitTestFlag == "y" {
			unitTest = true
		}
		fmt.Println("name:", appName, "unit test", unitTest)

		structure.Run(&structure.AppConfig{
			Name: appName,
		})
	},
}

func Cli() {
	var rootCmd = &cobra.Command{Use: "pen"}
	rootCmd.AddCommand(cmdApp)
	rootCmd.Execute()
}
