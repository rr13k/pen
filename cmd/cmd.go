package cmd

import (
	"fmt"
	"os"
	"path"
	"path/filepath"
	"strings"

	"github.com/pterm/pterm"
	"github.com/rr13k/pen/structure"
	"github.com/spf13/cobra"
)

var AppModelMap = map[string]string{
	"1": "gorm",
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
			"1. api + gorm",
		}
		modleStr, _ := pterm.DefaultInteractiveSelect.WithOptions(options).Show()
		branch := AppModelMap[strings.Split(modleStr, ".")[0]]

		structure.Run(&structure.AppConfig{
			Name:   appName,
			Branch: branch,
		})

		fmt.Println("start gen models...")
		cmdPath, err := filepath.Abs(os.Args[0])
		fmt.Println("cmdPath:", cmdPath, "os:", os.Args[0])
		if err != nil {
			panic(err)
		}

		GenerationModel(path.Join(filepath.Dir(cmdPath), appName, "internal", "app", "models", "user.go"))

		fmt.Printf("%s项目%s创建成功！can you: cd %s\n && go mod tidy", branch, appName, appName)
		// 生成基础模型文件
	},
}

var cmdModul = &cobra.Command{
	Use:   "modul [path]",
	Short: "by model file path generation http server",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("传入的文件路径为:", args[0])
		GenerationModel(args[0])
	},
}

func Cli() {
	var rootCmd = &cobra.Command{Use: "pen"}
	rootCmd.AddCommand(cmdApp, cmdModul)
	rootCmd.Version = "v0.0.1"
	rootCmd.Execute()
}
