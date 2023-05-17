package suger

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"os"
	"path"
	"path/filepath"
	"strings"
)

type StructCode struct {
	Key  string   // 结构体字段名称
	Type string   // 结构体名称
	Json string   // json标签信息
	Gorm []string // gorm标签信息（可能有多个）
}

// 从 Go 文件中提取结构体定义
func ExtractStructs(filePath string, deep bool) map[string][]StructCode {
	fset := token.NewFileSet()
	// 解析 Go 文件，生成抽象语法树（AST）
	node, err := parser.ParseFile(fset, filePath, nil, parser.ParseComments)
	if err != nil {
		panic(err)
	}

	var structMap = make(map[string][]StructCode)

	for _, decl := range node.Decls {
		var structCodes []StructCode
		genDecl, ok := decl.(*ast.GenDecl)
		if !ok || genDecl.Tok != token.TYPE {
			continue
		}

		for _, spec := range genDecl.Specs {
			typeSpec, ok := spec.(*ast.TypeSpec)
			importSpec, _ := spec.(*ast.ImportSpec)
			fmt.Println("import text:", importSpec)
			if !ok {
				continue
			}

			structType, ok := typeSpec.Type.(*ast.StructType)
			if !ok {
				continue
			}

			structName := typeSpec.Name.Name

			for _, field := range structType.Fields.List {
				var fieldName string
				for _, ident := range field.Names {
					fieldName = ident.Name
				}

				// 当模型继承了其他类型时,需要展开
				if fieldName == "" {
					if deep {
						privateStruceName := field.Type.(*ast.Ident).Name
						// 如果里面继承了结构体,则需要找到结构体并展开
						if !strings.Contains(privateStruceName, ".") { // 当不包含.时意味着结构体处于相同包
							goFiles, err := getGoFiles(path.Dir(filePath))
							// fileName :=  path.Base(filePath)
							if err != nil {
								panic(err)
							}

							for _, _gofile := range goFiles {
								// 当查找子结构时，不展开，避免循环递归
								_structs := ExtractStructs(_gofile, false)
								for _structName, _struct := range _structs {
									if _structName == privateStruceName {
										structCodes = append(structCodes, _struct...)
									}
								}
							}
						}
					}

					continue
					// 展开结构体
				}

				fieldType := fmt.Sprintf("%v", field.Type)

				var jsonTag, gormTags string
				if field.Tag != nil {
					tag := field.Tag.Value[1 : len(field.Tag.Value)-1]
					jsonTag = getTagValue(tag, "json")
					gormTags = getTagValue(tag, "gorm")
				}

				gormSplits := strings.Split(gormTags, ";")
				var gormValues []string
				for _, gormSplit := range gormSplits {
					if gormSplit != "" {
						gormValues = append(gormValues, strings.ReplaceAll(strings.TrimSpace(gormSplit), "\"", ""))
					}
				}

				structCodes = append(structCodes, StructCode{
					Key:  fieldName,
					Type: formatType(fieldType),
					Json: strings.ReplaceAll(jsonTag, "\"", ""),
					Gorm: gormValues,
				})
			}
			structMap[structName] = structCodes
		}
	}

	return structMap
}

func getGoFiles(folderPath string) ([]string, error) {
	var goFiles []string

	dir, err := os.ReadDir(folderPath)
	if err != nil {
		return nil, err
	}

	for _, file := range dir {
		if !file.IsDir() && filepath.Ext(file.Name()) == ".go" {
			goFiles = append(goFiles, filepath.Join(folderPath, file.Name()))
		}
	}

	return goFiles, nil
}

func formatType(fieldType string) string {
	switch fieldType {
	case "&{time Time}":
		return "time.Time"
	default:
		return fieldType
	}
}

func getTagValue(tag string, key string) string {
	_key := key + ":"
	gormIndex := strings.Index(tag, _key)
	if gormIndex >= 0 {
		gormStr := tag[gormIndex+len("gorm:\""):]
		gormStr = gormStr[:strings.Index(gormStr, "\"")]
		return gormStr
	}

	return ""
}

// 存储结构体字段信息
type FieldInfo struct {
	Name string
	Type string
}
