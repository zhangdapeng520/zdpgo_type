package zdpgo_type

import (
	"encoding/json"
	"strings"
)

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

// SplitArrString 将一维数组按照指定量级拆分为二维数组
func SplitArrString(arr []string, lines int) [][]string {
	// 处理数组为nil
	if arr == nil {
		return [][]string{}
	}

	// 行数小于等于0
	if lines <= 0 {
		lines = 1
	}

	// 处理数组长度小于分割行数
	if len(arr) <= lines {
		return [][]string{arr}
	}

	// 切割
	var result [][]string
	for i := lines; i <= len(arr); i++ {
		start := i - lines
		result = append(result, arr[start:start+lines])
	}
	return result
}

// IsEqualsDoubleArrString 判断两个二维字符串数组是否相等
func IsEqualsDoubleArrString(arr1, arr2 [][]string) bool {
	// 两个都为空相等
	if arr1 == nil && arr2 == nil {
		return true
	}

	// 单独一个为空，不相等
	if arr1 == nil || arr2 == nil {
		return false
	}

	// 判断长度
	if len(arr1) != len(arr2) {
		return false
	}

	// 判断元素
	for i := 0; i < len(arr1); i++ {
		arr11 := arr1[i]
		arr22 := arr2[i]

		// 判断空值
		if arr11 == nil || arr22 == nil {
			return false
		}

		// 判断长度
		if len(arr11) != len(arr22) {
			return false
		}

		// 判断元素
		for j := 0; j < len(arr11); j++ {
			if arr11[j] != arr22[j] {
				return false
			}
		}
	}

	return true
}

// IsEqualsArrString 判断两个字符串数组是否相等
func IsEqualsArrString(arr1, arr2 []string) bool {
	// 如果都为空，则相等
	if arr1 == nil && arr2 == nil {
		return true
	}

	// 单独一个为空，不相等
	if arr1 == nil || arr2 == nil {
		return false
	}

	// 判断长度
	if len(arr1) != len(arr2) {
		return false
	}

	// 判断元素
	for i := 0; i < len(arr1); i++ {
		// 判断元素
		for j := 0; j < len(arr1); j++ {
			if arr1[j] != arr2[j] {
				return false
			}
		}
	}

	return true
}

// GetStringSameSlice 获取字符串中相同的连续子串切片
// @param str1 目标字符串
// @param str2 要比对的字符串
// @param slitSep 切割字符串，一般是空格或者换行符
// @param lines 切割精度，按照多少行进行切割
// @return [][]int 返回相同的子串的索引切片，可能有多组，每组四个元素，代表[str2开始索引，str2结束索引，str1开始索引，str1结束索引]
func GetStringSameSlice(str1, str2, splitSep string, lines int) [][]int {
	arr1 := strings.Split(str1, splitSep)
	arr2 := strings.Split(str2, splitSep)

	// 拆分数组
	arr11 := SplitArrString(arr1, lines)
	arr22 := SplitArrString(arr2, lines)

	// 遍历数组2，查找相同的行
	var result [][]int
	for i := 0; i < len(arr22); i++ {
		for j := 0; j < len(arr11); j++ {
			if IsEqualsArrString(arr22[i], arr11[j]) {
				result = append(result, []int{i, i + lines - 1, j, j + lines - 1})
			}
		}
	}

	// 返回
	return result
}

// IsEqualsDoubleArrInt 判断两个二维整数数组是否相等
func IsEqualsDoubleArrInt(arr1, arr2 [][]int) bool {
	// 两个都为空相等
	if arr1 == nil && arr2 == nil {
		return true
	}

	// 单独一个为空，不相等
	if arr1 == nil || arr2 == nil {
		return false
	}

	// 判断长度
	if len(arr1) != len(arr2) {
		return false
	}

	// 判断元素
	for i := 0; i < len(arr1); i++ {
		arr11 := arr1[i]
		arr22 := arr2[i]

		// 判断空值
		if arr11 == nil || arr22 == nil {
			return false
		}

		// 判断长度
		if len(arr11) != len(arr22) {
			return false
		}

		// 判断元素
		for j := 0; j < len(arr11); j++ {
			if arr11[j] != arr22[j] {
				return false
			}
		}
	}

	return true
}

// MergeSerialSliceFromDoubleArrInt 从二维整数数组中，合并连续的切片
// 示例：[][]int{{0, 1, 0, 1}, {1, 2, 1, 2}}
// 其中第一个元素的0-1和第二个元素的1-2是连续的
// 合并为：{0, 2, 0, 2}
func MergeSerialSliceFromDoubleArrInt(arr [][]int) [][]int {
	var result = make([][]int, 0, 0)

	// 为空
	if arr == nil {
		return result
	}

	// 没有元素
	if len(arr) == 0 {
		return result
	}

	// 不满足元素为4个
	if len(arr[0]) != 4 {
		return result
	}

	// 就1个元素
	if len(arr) == 1 {
		result = append(result, arr[0])
		return result
	}

	// 开始合并
	var count int
	for {
		arr, count = mergeArr(arr)
		if count == 0 {
			break
		}
	}

	// 返回
	return arr
}

func mergeArr(arr [][]int) ([][]int, int) {
	if len(arr) <= 1 {
		return arr, 0
	}

	var count = 0
	var i = 1
	var isMerge = false
	var tempArr [][]int
	for i < len(arr) {
		if arr[i][0] <= arr[i-1][1] {
			tempArr = append(tempArr, []int{arr[i-1][0], arr[i][1], arr[i-1][2], arr[i][3]})
			count++
			isMerge = true
			i++
		} else {
			if !isMerge {
				tempArr = append(tempArr, []int{arr[i-1][0], arr[i-1][1], arr[i-1][2], arr[i-1][3]})
			}
			if i == len(arr)-1 {
				tempArr = append(tempArr, []int{arr[i][0], arr[i][1], arr[i][2], arr[i][3]})
			}
			i++
		}
	}
	return tempArr, count
}

// StringCamelToSnake 驼峰转蛇形
// @param s 需要转换的字符串
// @return 蛇形字符串
func StringCamelToSnake(s string) string {
	data := make([]byte, 0, len(s)*2)
	j := false
	num := len(s)
	for i := 0; i < num; i++ {
		d := s[i]
		// or通过ASCII码进行大小写的转化
		// 65-90（A-Z），97-122（a-z）
		//判断如果字母为大写的A-Z就在前面拼接一个_
		if i > 0 && d >= 'A' && d <= 'Z' && j {
			data = append(data, '_')
		}
		if d != '_' {
			j = true
		}
		data = append(data, d)
	}

	//ToLower把大写字母统一转小写
	return strings.ToLower(string(data[:]))
}

// StringSnakeToCamel 蛇形字符串转驼峰
// @param s 需要转换的字符串
// @return 驼峰字符串
func StringSnakeToCamel(s string) string {
	data := make([]byte, 0, len(s))
	j := false
	k := false
	num := len(s) - 1
	for i := 0; i <= num; i++ {
		d := s[i]
		if k == false && d >= 'A' && d <= 'Z' {
			k = true
		}
		if d >= 'a' && d <= 'z' && (j || k == false) {
			d = d - 32
			j = false
			k = true
		}
		if k && d == '_' && num > i && s[i+1] >= 'a' && s[i+1] <= 'z' {
			j = true
			continue
		}
		data = append(data, d)
	}
	return string(data[:])
}

// StringSnakeToCamelSmall 蛇形字符串转小驼峰
// @param s 需要转换的字符串
// @return 驼峰字符串
func StringSnakeToCamelSmall(s string) string {
	camel := StringSnakeToCamel(s)
	firstChar := camel[0]
	firstStr := strings.ToLower(string(firstChar))
	return firstStr + camel[1:]
}
