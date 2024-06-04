package gorm_ex

import "encoding/json"

// Map2MysqlJsonString map转换成合法的mysql json类型数据
// nil/empty => "{}"
func Map2MysqlJsonString(m map[string]interface{}) string {
	if m == nil {
		return "{}"
	}
	j, _ := json.Marshal(m)
	return string(j)
}
