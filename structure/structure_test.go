package structure

import (
	"fmt"
	"testing"
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

// 测试刷新命令行显示
func Test_F5(t *testing.T) {

}
