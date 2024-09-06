package model

import "encoding/json"

// encode val to json string
func ToJsonStr(val any) string {
	bytes, _ := json.Marshal(val)
	return string(bytes)
}
