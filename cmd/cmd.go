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
			// "1. only api",
			"1. api + gorm",
		}
		modleStr, _ := pterm.DefaultInteractiveSelect.WithOptions(options).Show()
		branch := AppModelMap[strings.Split(modleStr, ".")[0]]

		structure.Run(&structure.AppConfig{
			Name:   appName,
			Branch: branch,
		})

		fmt.Printf("%s项目%s创建成功！can you: cd %s\n", branch, appName, appName)
		// 生成基础模型文件

		exe, err := os.Executable()
		if err != nil {
			fmt.Println("Failed to get executable path:", err)
		}

		// 获取可执行文件所在相对目录
		root := filepath.Dir(exe)

		fmt.Println("root:", root)

		GenerationModel(path.Join(root, "appName", "internal", "app", "models", "user.go"))
	},
}

var cmdModul = &cobra.Command{
	Use:   "modul [path]",
	Short: "Print the given path",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("传入的文件路径为:", args[0])
		GenerationModel(args[0])
	},
}

func Cli() {
	var rootCmd = &cobra.Command{Use: "pen"}
	rootCmd.AddCommand(cmdApp, cmdModul)
	rootCmd.Execute()
}
