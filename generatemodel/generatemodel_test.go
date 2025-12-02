package generatemodel

import (
	"os"
	"testing"

	"gorm.io/driver/mysql"
	"gorm.io/gen"
	"gorm.io/gorm"
)

func TestGenerateModel(t *testing.T) {
	// 连接数据库
	dsn := "root:123456@tcp(127.0.0.1:3306)/test?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	// 生成 model
	dir, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	g := gen.NewGenerator(gen.Config{
		OutPath:      dir + "/internal/query",
		ModelPkgPath: dir + "/internal/model",
	})
	g.UseDB(db)
	err = GenerateModel(db, g, "./configs/generate-model.yml")
	if err != nil {
		panic(err)
	}

	// 不生成 query
	keys := make([]string, 0, len(g.Data))
	for k := range g.Data {
		keys = append(keys, k)
	}
	for _, v := range keys {
		delete(g.Data, v)
	}

	g.Execute()
}
