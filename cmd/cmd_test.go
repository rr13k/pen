package cmd

import (
	"fmt"
	"testing"

	"github.com/rr13k/pen/structure"
)

func Test_App(t *testing.T) {
	cmdApp.Run(nil, []string{"nihao"})
}

func Test_nihao(t *testing.T) {
	fmt.Println("niaho")
}

// 测试分支下载
func Test_DownDst(t *testing.T) {
	structure.Run(&structure.AppConfig{
		Name:   "Test_DownDst",
		Branch: "gorm",
	})
}
