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

func Test_LI(t *testing.T) {

	// user := AppConfig{
	// 	Name: "hao123",
	// }

	// DirsFiles(user.Name, DirsFilesConfig{
	// 	FillterDir: []string{".git"},
	// 	FileCall: func(name string, path string) {
	// 		if len(name) > 5 && name[len(name)-5:] == ".tmpl" {
	// 			fmt.Println("匹配的文件", path)
	// 			tpl, err := template.ParseFiles(path) // "hao123/README.md.tmpl"
	// 			if err != nil {
	// 				fmt.Println(err)
	// 			}

	// 			outFile := path[:len(path)-5]
	// 			file, err := os.OpenFile(outFile, os.O_CREATE|os.O_WRONLY, 0755)
	// 			if err != nil {
	// 				panic(err)
	// 			}
	// 			//渲染输出
	// 			err = tpl.Execute(file, user)
	// 			if err != nil {
	// 				panic(err)
	// 			}
	// 			// 删除模版文件
	// 			os.Remove(path)
	// 		}
	// 	},
	// })
}
