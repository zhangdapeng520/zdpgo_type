package zdpgo_type

import "encoding/json"

// GetMap 获取map
func GetMap(keyValues ...interface{}) map[string]interface{} {
	data := make(map[string]interface{})

	// 极端情况
	if keyValues == nil {
		return data
	}

	// 构造map
	for i := 0; i < len(keyValues); i += 2 {
		keyI := keyValues[i]
		value := keyValues[i+1]
		key := keyI.(string)
		data[key] = value
	}

	// 返回
	return data
}

// GetArrMap 获取map数组
func GetArrMap(maps ...map[string]interface{}) []map[string]interface{} {
	return maps
}

// GetArrString 获取字符串数组
func GetArrString(args ...string) []string {
	return args
}

// MapToJson 将map对象转换为json字符串
func MapToJson(mapObj map[string]interface{}) (string, error) {
	value, err := json.Marshal(mapObj)
	if err != nil {
		return "", err
	}
	return string(value), nil
}

// JsonToObj 将JSON字符串转换为obj对象
func JsonToObj(jsonStr string, obj interface{}) error {
	err := json.Unmarshal([]byte(jsonStr), obj)
	return err
}
