package main

import (
	"fiber/utils"
	"fmt"
	"github.com/gobuffalo/flect"
	"github.com/gobuffalo/plush/v5"
	"os"
	"path/filepath"
)

// ========== 模板数据结构 ==========

type HandlerTemplateData struct {
	PackageName      string
	HandlerName      string
	LowerHandlerName string
	RoutePath        string
	StructName       string
	LowerStructName  string
}

type ServiceTemplateData struct {
	PackageName     string
	StructName      string
	LowerStructName string
}

func GenAPI(name string) {
	lowerHandleName := name
	handlerName := flect.Capitalize(name)
	routePath := name
	packageName := name

	data := HandlerTemplateData{
		PackageName:      packageName,
		HandlerName:      handlerName,
		LowerHandlerName: lowerHandleName,
		RoutePath:        routePath,
	}

	// 读取模板文件
	tmplBytes, err := os.ReadFile("gen/tmpl/templates/api.tmpl")
	if err != nil {
		fmt.Errorf("❌ 读取 handler 模板失败: %v", err)
	}
	// 渲染模板并写入文件
	filePath := filepath.Join(utils.AppPath, "api", name)
	outPath := filepath.Join(utils.AppPath, "api", name, fmt.Sprint(name, ".go"))

	err = os.MkdirAll(filePath, os.ModePerm)
	if err != nil {
		fmt.Errorf("❌ 创建目录失败: %v", err)
	}
	tmplStr := string(tmplBytes)
	generateFile(tmplStr, data, outPath)
}
func GenService(name string) {

	structName := flect.Capitalize(name)
	lowerName := name

	data := HandlerTemplateData{
		PackageName:     lowerName,
		StructName:      structName,
		LowerStructName: lowerName,
	}
	// 读取模板文件
	tmplBytes, err := os.ReadFile("gen/tmpl/templates/service.tmpl")
	if err != nil {
		fmt.Errorf("❌ 读取 handler 模板失败: %v", err)
	}
	// 渲染模板并写入文件
	filePath := filepath.Join(utils.AppPath, "service", name)
	outPath := filepath.Join(utils.AppPath, "service", name, fmt.Sprint(name, ".go"))

	err = os.MkdirAll(filePath, os.ModePerm)
	if err != nil {
		fmt.Errorf("❌ 创建目录失败: %v", err)
	}
	tmplStr := string(tmplBytes)
	generateFile(tmplStr, data, outPath)
}

// generateAPIFile 使用 api.tmpl 模板生成 API  文件
func generateFile(tmplStr string, data HandlerTemplateData, outputFilePath string) {
	// 使用 plush 渲染模板
	ctx := plush.NewContext()
	ctx.Set("PackageName", data.PackageName)
	ctx.Set("HandlerName", data.HandlerName)
	ctx.Set("LowerCamelName", data.LowerHandlerName)
	ctx.Set("RoutePath", data.RoutePath)

	ctx.Set("LowerStructName", data.LowerStructName)
	//ctx.Set("LowerHandlerName", data.LowerHandlerName)
	ctx.Set("StructName", data.StructName)

	rendered, err := plush.Render(tmplStr, ctx)
	if err != nil {
		fmt.Errorf("❌ 渲染 handler 模板失败: %v", err)
	}
	// 写入到目标文件
	err = os.WriteFile(outputFilePath, []byte(rendered), 0644)
	if err != nil {
		fmt.Errorf("❌ 写入 handler 文件失败: %v", err)
	}
	fmt.Printf("✅ 已生成 Handler 文件: %s\n", outputFilePath)
}
