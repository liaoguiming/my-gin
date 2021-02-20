package utils

import (
	"encoding/json"
	"liao/global"
	"reflect"
	"time"
)

func SetOrGetRedisKey(key string, data interface{}, Cache bool) interface{} {
	if Cache == false {
		return data
	}
	str, _ := global.REDIS.Get(key).Result()
	if str == "" || str == "null" {
		info, _ := json.Marshal(data)
		if string(info) != "" {
			global.REDIS.Set(key, string(info), 86400*time.Second)
		}
	} else {
		_ = json.Unmarshal([]byte(str), &data)
	}
	return data
}

// 利用反射将结构体转化为map
func StructToMap(obj interface{}) map[string]interface{} {
	obj1 := reflect.TypeOf(obj)
	obj2 := reflect.ValueOf(obj)

	var data = make(map[string]interface{})
	for i := 0; i < obj1.NumField(); i++ {
		data[obj1.Field(i).Name] = obj2.Field(i).Interface()
	}
	return data
}
