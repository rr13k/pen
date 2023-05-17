package cmd

import (
	"fmt"
	"testing"

	"github.com/rr13k/pen/structure"
)

func Test_App(t *testing.T) {
	cmdApp.Run(nil, []string{"nihao"})
}

func Test_nihao2(t *testing.T) {
	fmt.Println("niaho")
}

// 测试分支下载
func Test_DownDst(t *testing.T) {
	structure.Run(&structure.AppConfig{
		Name:   "Test_DownDst",
		Branch: "gorm",
	})
}

// 测试模型生成, 在模版中使用自定义函数，还是没成功
func Test_module(t *testing.T) {
	// GenerationModel("/Users/zhouyuan11/work/pen/pen-test/internal/app/models/ok.go")
	GenerationModel("/Users/zhouyuan11/work/pen/pen-test/internal/app/models/user.go")
	// GenerationModel("/Users/zhouyuan11/work/pen/pen-test/internal/app/models/kele.go")
}
