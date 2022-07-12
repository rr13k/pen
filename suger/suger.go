package suger

import (
	"fmt"
	"io/ioutil"
)

func InString(target string, str_array []string) bool {
	for _, item := range str_array {
		if item == target {
			return true
		}
	}
	return false
}

type DirsFilesConfig struct {
	FillterDir []string                       // 过滤不需要的目录 例: []string{".git"}
	FileCall   func(name string, path string) // 处理匹配的文件
}

// 目录过滤
func DirsFiles(dirname string, cfg DirsFilesConfig) {
	fileInfos, err := ioutil.ReadDir(dirname)
	if err != nil {
		fmt.Println(err)
	}
	for _, fi := range fileInfos {
		filename := fi.Name()
		pathName := dirname + "/" + filename
		if cfg.FileCall != nil {
			cfg.FileCall(filename, pathName)
		}
		if fi.IsDir() && !InString(filename, cfg.FillterDir) {
			DirsFiles(pathName, cfg)
		}
	}
}
