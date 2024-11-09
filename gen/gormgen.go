package main

import (
	yamlgen "github.com/we7coreteam/gorm-gen-yaml"
	"gorm.io/driver/mysql"
	"gorm.io/gen"
	"gorm.io/gen/field"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

func main() {
	g := gen.NewGenerator(gen.Config{
		OutPath:      "./gen/dao",
		Mode:         gen.WithoutContext | gen.WithDefaultQuery | gen.WithQueryInterface, // generate mode
		ModelPkgPath: "./gen/model",                                                      // 定义 model 文件输出目录
		// 生成 gorm 标签的字段类型属性

		// WithDefaultQuery 生成默认查询结构体(作为全局变量使用), 即`Q`结构体和其字段(各表模型)
		// WithoutContext 生成没有context调用限制的代码供查询
		// WithQueryInterface 生成interface形式的查询代码(可导出), 如`Where()`方法返回的就是一个可导出的接口类型

		// 生成 gorm 标签的字段索引属性
		FieldWithIndexTag: true,
		// 生成 gorm 标签的字段类型属性
		FieldWithTypeTag: true,
	})

	gormdb, _ := gorm.Open(mysql.Open("root:root@(127.0.0.1:3306)/likeadmin?charset=utf8mb4&parseTime=True&loc=Local"),
		&gorm.Config{NamingStrategy: schema.NamingStrategy{TablePrefix: "la_"}})
	//	reuse your gorm db
	//gormdb *gorm.DB
	g.UseDB(gormdb)

	// 自定义字段的数据类型
	// 统一数字类型为int64,兼容protobuf和thrift
	dataMap := map[string]func(detailType gorm.ColumnType) (dataType string){
		"tinyint":   func(detailType gorm.ColumnType) (dataType string) { return "uint8" },
		"smallint":  func(detailType gorm.ColumnType) (dataType string) { return "int64" },
		"mediumint": func(detailType gorm.ColumnType) (dataType string) { return "int64" },
		"bigint":    func(detailType gorm.ColumnType) (dataType string) { return "int32" },
		"int":       func(detailType gorm.ColumnType) (dataType string) { return "int32" },
		"timestamp": func(detailType gorm.ColumnType) (dataType string) { return "LocalTime" }, // 自定义时间
		//"decimal":   func(detailType gorm.ColumnType) (dataType string) { return "Decimal" },   // 金额类型全部转换为第三方库,github.com/shopspring/decimal
		"decimal": func(detailType gorm.ColumnType) (dataType string) { return "float32" },
	}
	// 要先于`ApplyBasic`执行
	g.WithDataTypeMap(dataMap) // 自定义字段的数据类型
	createTimeGormTag := gen.FieldGORMTag("create_time", func(tag field.GormTag) field.GormTag {
		tag.Append("autoCreateTime")
		return tag
	})
	updateTimeGormTag := gen.FieldGORMTag("update_time", func(tag field.GormTag) field.GormTag {
		tag.Append("autoUpdateTime")
		return tag
	})
	deleteTimeGormTag := gen.FieldGORMTag("delete_time", func(tag field.GormTag) field.GormTag {
		tag.Append("DEFAULT", "0")
		return tag
	})
	// 2021-07-30 16:54:18.470685+08:00 日期型
	//softDeleteField := gen.FieldType("delete_time", "gorm.DeletedAt")
	// 时间戳
	softDeleteField := gen.FieldType("delete_time", "soft_delete.DeletedAt")

	modelOpts := []gen.ModelOpt{softDeleteField, createTimeGormTag, updateTimeGormTag, deleteTimeGormTag}
	// 生成 yaml 文件
	yamlgen.NewYamlGenerator("./gen/gen.yaml").UseGormGenerator(g).Generate(modelOpts...)
	g.Execute()
}
