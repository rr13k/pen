package structure

import (
	"fmt"
	"testing"

	"github.com/rr13k/pen/toolkit/git"
)

func TestMain(m *testing.T) {
	fmt.Println("run TestMain")
	Run(&AppConfig{
		Name: "hao123",
	})
}

func Test_Run(t *testing.T) {

	Run(&AppConfig{
		Name: "hao123",
		Orm:  true,
	})
}

// 测试克隆模版项目
func Test_Clone_temp(t *testing.T) {

	appConfig := &AppConfig{
		Name:   "hao123",
		Orm:    true,
		Branch: "gorm",
	}

	err := git.CloneORPullRepo(defaultStructureConfig.TempAddr, appConfig.Branch, appConfig.Name)

	if err != nil {
		fmt.Println("Test_Clone_temp error:", err)
	}
}

// 测试模版内容解析
func Test_Parse(t *testing.T) {
	appConfig := &AppConfig{
		Name:   "hao123",
		Orm:    true,
		Branch: "gorm",
	}

	Parse(appConfig)
}

// 测试刷新命令行显示
func Test_F5(t *testing.T) {

}
