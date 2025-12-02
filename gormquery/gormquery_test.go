package gormquery

import (
	"log"
	"testing"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

const TableNameConfig = "configs"

// Config 配置表
type Config struct {
	ID        int64          `gorm:"column:id;primaryKey;autoIncrement:true" json:"id"`
	Name      string         `gorm:"column:name;not null;comment:名称" json:"name"`                                  // 名称
	Var       string         `gorm:"column:var;not null;comment:变量名" json:"var"`                                   // 变量名
	Type      int32          `gorm:"column:type;not null;comment:类型:0=单行文本,1=多行文本,2=单选按钮,3=复选框,4=下拉框" json:"type"` // 类型:0=单行文本,1=多行文本,2=单选按钮,3=复选框,4=下拉框
	Option    string         `gorm:"column:option;comment:配置项(针对于:单选按钮,复选框,下拉框)" json:"option"`                    // 配置项(针对于:单选按钮,复选框,下拉框)
	Value     string         `gorm:"column:value;comment:配置值" json:"value"`                                        // 配置值
	Weight    int64          `gorm:"column:weight;not null;default:100;comment:排序(权重)" json:"weight"`              // 排序(权重)
	CreatedAt time.Time      `gorm:"column:created_at" json:"created_at"`
	UpdatedAt time.Time      `gorm:"column:updated_at" json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"column:deleted_at" json:"deleted_at"`
}

// TableName Config's table name
func (*Config) TableName() string {
	return TableNameConfig
}

func TestSelect(t *testing.T) {
	t.Log(Select(map[string]string{
		"a": "id,password",
		"b": "id, nick_name as nickname , test",
		"c": "id , test ,",
		"d": "id, test , ",
	}))
}

func TestWhere(t *testing.T) {
	// 连接数据库
	newLogger := logger.New(
		log.New(log.Writer(), "\r\n", log.LstdFlags), // 日志输出目标（标准输出）
		logger.Config{
			SlowThreshold:             time.Second, // 慢 SQL 阈值（1 秒）
			LogLevel:                  logger.Info, // 日志级别（Info 级别会打印所有 SQL）
			IgnoreRecordNotFoundError: true,        // 忽略 "record not found" 错误
			Colorful:                  true,        // 彩色输出
		},
	)
	dsn := "root:123456@tcp(127.0.0.1:3306)/test?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: newLogger,
	})
	if err != nil {
		panic(err)
	}

	// 请求数据
	reqMap := map[string]interface{}{
		"id":     6,
		"name":   "标题",
		"weight": "1",
		"type":   []string{"0", "1", "2", "3"},
		"other":  "hhh",
	}
	where := Where(db, reqMap, map[string]string{
		"id":     "=",
		"name":   "LIKE",
		"weight": ">=",
		"type":   "in",
	})
	var result []Config
	db.Model(&Config{}).Where(where).Find(&result)
	t.Log(result)
}

func TestFindInSet(t *testing.T) {
	t.Logf("%#v", FindInSet("f", "1,2,3", "or"))
	t.Logf("%#v", FindInSet("f", "1,2,3", " or"))
	t.Logf("%#v", FindInSet("f", "1,2,3", "or "))
	t.Logf("%#v", FindInSet("f", "1,2,3", " or "))
	t.Logf("%#v", FindInSet("f", "1,2,3", "and"))
}
