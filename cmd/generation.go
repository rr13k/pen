package cmd

import (
	"fmt"
	"os"
	"path"
	"path/filepath"
	"strings"
	"text/template"

	"github.com/rr13k/pen/suger"
)

type AppGenConfig struct {
	Name       string             `json:"name"`        // 应用名称
	StructName string             `json:"struct_name"` // 结构体名称
	Fields     []suger.StructCode `json:"fields"`      // 结构体字段
	PrimaryKey string             `json:"primary_key"` // 主键
}

/*
*
代码生成器
*/
func GenerationModel(moduleFilePath string) {
	structs := suger.ExtractStructs(moduleFilePath, true)

	modulePath := path.Dir(moduleFilePath)
	appPath := path.Dir(modulePath)

	fmt.Println("应用路径", appPath, "模块路径", modulePath)

	// 生成pen模型
	genModulePath := path.Join(modulePath, "pen_models")
	if !CheckFolderExist(genModulePath) {
		os.Mkdir(genModulePath, os.ModePerm)
	}

	// 写入模型文件
	for name, fields := range structs {

		app := &AppGenConfig{
			Name:       path.Base(GetParentDir(appPath, 2)),
			StructName: name,
			Fields:     fields,
			PrimaryKey: getOrmPrimaryKey(fields),
		}

		outPenModuleFile := path.Join(genModulePath, fmt.Sprintf("%s.go", ToLowerFirst(app.StructName)))

		if !CheckFolderExist(outPenModuleFile) {
			//渲染输出
			err := renderTemplate("model.temp", app, outPenModuleFile)
			if err != nil {
				os.RemoveAll(genModulePath)
				panic(err)
			}

			fmt.Println("pen module file:", outPenModuleFile, ". gen success～")
		}

	}

	// 生成server服务
	serverPath := path.Join(appPath, "servers")

	for name, fields := range structs {
		app := &AppGenConfig{
			Name:       path.Base(GetParentDir(appPath, 2)),
			StructName: name,
			Fields:     fields,
			PrimaryKey: getOrmPrimaryKey(fields),
		}

		appServerDir := path.Join(serverPath, fmt.Sprintf("%s_servers", ToLowerFirst(name)))
		if !CheckFolderExist(appServerDir) {
			os.Mkdir(appServerDir, os.ModePerm)
		}

		// os.RemoveAll(appServerDir) // 快速测试使用-删除文件夹

		// 生成供用户使用的server文件
		appUserFile := path.Join(appServerDir, fmt.Sprintf("%s.go", ToLowerFirst(name)))
		if !CheckFolderExist(appUserFile) {
			// gen user use server
			err := renderTemplate("u_server.temp", app, appUserFile)
			if err != nil {
				panic(err)
			}
		}

		penAppServerDir := path.Join(appServerDir, fmt.Sprintf("pen_%s_server", ToLowerFirst(app.StructName)))

		if !CheckFolderExist(penAppServerDir) {
			os.Mkdir(penAppServerDir, os.ModePerm)
		}

		outPenServerFile := path.Join(penAppServerDir, fmt.Sprintf("%s.go", ToLowerFirst(app.StructName)))

		// 结构体， 模版， 输出文件
		err := renderTemplate("server.temp", app, outPenServerFile)
		if err != nil {
			os.RemoveAll(penAppServerDir)
			panic(err)
		}
		fmt.Println("pen server file:", outPenServerFile, ". gen success～")

	}

	// 生成router路由
	routerPath := path.Join(appPath, "api", "handlers")
	penAppRouterDir := path.Join(routerPath, "pen_handler")

	// os.RemoveAll(penAppRouterDir) // 快速测试使用-删除文件夹
	if !CheckFolderExist(penAppRouterDir) {
		os.Mkdir(penAppRouterDir, os.ModePerm)
	}

	for name, fields := range structs {
		app := &AppGenConfig{
			Name:       path.Base(GetParentDir(appPath, 2)),
			StructName: name,
			Fields:     fields,
			PrimaryKey: getOrmPrimaryKey(fields),
		}

		outPenRouterFile := path.Join(penAppRouterDir, fmt.Sprintf("%s.go", ToLowerFirst(app.StructName)))

		if !CheckFolderExist(outPenRouterFile) {
			// 结构体， 模版， 输出文件
			err := renderTemplate("router.temp", app, outPenRouterFile)
			if err != nil {
				os.RemoveAll(penAppRouterDir)
				panic(err)
			}
			fmt.Println("pen router file:", outPenRouterFile, ". gen success～")
		}
	}

}

// 检查文件夹是否存在
func CheckFolderExist(path string) bool {
	if _, err := os.Stat(path); err == nil {
		return true // 文件夹存在
	} else {
		return false // 文件夹不存在
	}
}

// 检查是否有time.Time类型
func HasTime(s string) bool {
	return s == "time.Time"
}

// 后续调整为支持多种类型自动引入
func autoImport(StructCodes []suger.StructCode) string {
	for _, v := range StructCodes {
		if v.Type == "time.Time" {
			return `"time"`
		}
	}
	return ""
}

// 首字母小写
func toLowerFiristChar(input string) string {
	if len(input) == 0 {
		return input
	}
	return strings.ToLower(string(input[0])) + input[1:]
}

func getServerPenStruceName(name string) string {
	return fmt.Sprintf("pen_%s_server", toLowerFiristChar(name))
}

// 获取orm主键,需要根据主建生成删除代码
func getOrmPrimaryKey(StructCodes []suger.StructCode) string {
	for _, v := range StructCodes {
		for _, g := range v.Gorm {
			if g == "primary_key" {
				return v.Key
			}
		}
	}

	panic("not found orm primary key")
}

// 过滤掉默认字段, 如更新时间，创建时间等
func filterDefaultField(StructCodes []suger.StructCode) []suger.StructCode {

	var _filterStructCodes = make([]suger.StructCode, 0)
	for _, v := range StructCodes {
		gorms := strings.Join(v.Gorm, ",")
		if strings.Contains(gorms, "current_timestamp") || strings.Contains(gorms, "autoUpdateTime") {
			continue
		}
		_filterStructCodes = append(_filterStructCodes, v)
	}

	return _filterStructCodes
}

// 渲染模版根据结构体生成文件
func renderTemplate(tempName string, app interface{}, outPath string) error {
	// tempsDir, _ := GetTempsDir()
	// _filePath := path.Join(tempsDir, tempName)

	_filePath, err := CmdEmbedTempsContent.ReadFile(fmt.Sprintf("temps/%s", tempName))
	if err != nil {
		fmt.Println("read file err:", err)
	}

	funcMap := template.FuncMap{
		"autoImport":             autoImport,
		"toLowerFiristChar":      toLowerFiristChar,
		"getServerPenStruceName": getServerPenStruceName,
		"getOrmPrimaryKey":       getOrmPrimaryKey,
		"filterDefaultField":     filterDefaultField,
	}
	tmpl, err := template.New(tempName).Funcs(funcMap).Parse(string(_filePath))

	if err != nil {
		return fmt.Errorf("failed to parse template: %v", err)
	}

	file, err := os.OpenFile(outPath, os.O_CREATE|os.O_WRONLY, 0755)
	if err != nil {
		return err
	}

	//渲染输出
	return tmpl.Execute(file, app)
}

// 字符串首字母小写
func ToLowerFirst(str string) string {
	if len(str) < 1 {
		return str
	}
	return strings.ToLower(str[:1]) + str[1:]
}

// 获取父级目录，level为父级目录层级
func GetParentDir(path string, level int) string {
	for i := 0; i < level; i++ {
		path = filepath.Dir(path)
	}
	return path
}

// 获取模版目录
func GetTempsDir() (string, error) {
	tmpDir := os.TempDir()

	fmt.Println("tempdir", tmpDir)
	exe, err := os.Executable()
	if err != nil {
		fmt.Println("Failed to get executable path:", err)
		return "", err
	}

	// 获取可执行文件所在相对目录
	root := filepath.Dir(exe)

	if filepath.Base(root) == "cmd" {
		root = filepath.Dir(root)
	}
	fmt.Println("Root path:", root)

	return path.Join(root, "temps"), nil
}
