package structure

import (
	"fmt"
	"html/template"
	"os"

	"github.com/rr13k/pen/suger"
	"github.com/rr13k/pen/toolkit/git"
	"github.com/rr13k/pen/toolkit/log"
)

type AppConfig struct {
	Name string
}

type DefaultStructureConfig struct {
	TempAddr string
}

var defaultStructureConfig = &DefaultStructureConfig{
	TempAddr: "https://github.com/rr13k/pen-tmpl",
}

// 解析
func Parse(app *AppConfig) {
	suger.DirsFiles(app.Name, suger.DirsFilesConfig{
		FillterDir: []string{".git"},
		FileCall: func(name string, path string) {
			if len(name) > 5 && name[len(name)-5:] == ".tmpl" {
				tpl, err := template.ParseFiles(path)
				if err != nil {
					fmt.Println(err)
				}

				outFile := path[:len(path)-5]
				file, err := os.OpenFile(outFile, os.O_CREATE|os.O_WRONLY, 0755)
				if err != nil {
					panic(err)
				}
				//渲染输出
				err = tpl.Execute(file, app)
				if err != nil {
					panic(err)
				}
				// 删除模版文件
				os.Remove(path)
			}
		},
	})
}

// 克隆项目并按照模版解析
func Run(appConfig *AppConfig) {
	// clone 模版项目
	err := git.CloneORPullRepo(defaultStructureConfig.TempAddr, appConfig.Name)
	if err != nil {
		log.Error("beego pro git clone or pull repo error, err: %s", err.Error())
		return
	}

	Parse(appConfig)
}
