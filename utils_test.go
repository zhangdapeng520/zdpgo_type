package zdpgo_type

import "testing"

func TestMapToJson(t *testing.T) {
	m := GetMap("a", 11, "b", 22)

	// 转换为json字符串
	jsonStr, err := MapToJson(m)
	if err != nil {
		panic(err)
	}
	t.Log(jsonStr)

	// 转换为obj对象
	obj := struct {
		A int `json:"a"`
		B int `json:"b"`
	}{}
	err = JsonToObj(jsonStr, &obj)
	if err != nil {
		panic(err)
	}
	t.Log(obj.A, obj.B)
}
