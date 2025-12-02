package gormquery

import (
	"fmt"
	"strings"

	"gorm.io/gorm"
)

// 处理 select
// qeury map[表]字段用,符号分隔
func Select(query map[string]string) string {
	spe := ", "
	newQuery := ""
	for table, fieldStr := range query {
		for _, field := range strings.Split(fieldStr, ",") {
			if newField := strings.TrimSpace(field); newField != "" {
				newQuery += table + "." + newField + spe
			}
		}
	}
	return strings.TrimRight(newQuery, spe)
}

// 处理 where
// req := Struct2map(req) 请求的数据
// where := map[string]string{"id":"=","name":"LIKE"}
func Where(tx *gorm.DB, req map[string]interface{}, where map[string]string) *gorm.DB {
	for k, v := range where {
		// 解析字段(id,table.name,cid as table.id)
		reqKey := k
		whereKey := k
		if index := strings.Index(k, " as "); index != -1 {
			reqKey = k[:index]
			whereKey = k[index+4:]
		} else if index := strings.Index(k, "."); index != -1 {
			reqKey = k[index+1:]
			whereKey = k
		}
		// 解析操作符号(=,<>,>,>=,<,<=,IN,NOT IN,LIKE)
		if reqVal, ok := req[reqKey]; ok {
			switch v {
			case "=", "<>", ">", ">=", "<", "<=", "IN", "NOT IN":
				tx = tx.Where(fmt.Sprintf("%s %s ?", whereKey, v), reqVal)
			case "LIKE":
				tx = tx.Where(fmt.Sprintf("%s %s ?", whereKey, v), fmt.Sprintf("%%%s%%", reqVal))
			}
		}
	}
	return tx
}

// 处理 FIND_IN_SET
// fieldWhere := g.FindInSet("f", "1,2", "or")
// tx = tx.Where(fieldWhere[0], fieldWhere[1:]...)
func FindInSet(field string, value string, operator string) []interface{} {
	operator = " " + strings.TrimSpace(operator) + " "

	whereStr := ""
	valueSlice := strings.Split(value, ",")
	for i := 1; i <= len(valueSlice); i++ {
		whereStr += fmt.Sprintf("FIND_IN_SET(?,%s)%s", field, operator)
	}

	valueWhere := []interface{}{"(" + strings.TrimRight(whereStr, operator) + ")"}
	for _, v := range valueSlice {
		valueWhere = append(valueWhere, v)
	}

	return valueWhere
}
