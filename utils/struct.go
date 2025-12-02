package utils

import (
	"reflect"
	"strings"
)

// 结构体转map
func Struct2map(req interface{}, tagKey string, skipZero bool) map[string]interface{} {
	reqMap := make(map[string]interface{})
	reqValue := reflect.Indirect(reflect.ValueOf(req))
	reqType := reqValue.Type()
	for i := 0; i < reqValue.NumField(); i++ {
		// 跳过零值
		if skipZero {
			if kind := reqValue.Field(i).Kind(); kind == reflect.Slice || kind == reflect.String {
				if reqValue.Field(i).IsZero() {
					continue
				}
			}
		}

		// 获取 tag 作为 map 的 key
		tag := reqType.Field(i).Tag.Get(tagKey)
		if tag == "" {
			continue
		}

		// 只取第一个值(kratos 框架的 json 上有多个值，用逗号分隔)
		if index := strings.Index(tag, ","); index != -1 {
			reqMap[tag[:index]] = reqValue.Field(i).Interface()
		} else {
			reqMap[tag] = reqValue.Field(i).Interface()
		}
	}

	return reqMap
}
