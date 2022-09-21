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

// 判断两个二维数组是否相等
func TestIsEqualsDoubleArrString(t *testing.T) {
	testData := []struct {
		arg1     [][]string
		arg2     [][]string
		expected bool
	}{
		{nil, nil, true},
		{nil, [][]string{{"a"}, {"b"}}, false},
		{[][]string{{"a"}, {"b"}}, nil, false},
		{[][]string{{"a"}, {"b"}}, [][]string{{"a"}}, false},
		{[][]string{{"a"}, {"b"}}, [][]string{{"a"}, {"b"}}, true},
		{[][]string{{"a"}, {"b"}}, [][]string{{"a", "b"}, {"b"}}, false},
		{[][]string{{"a"}, {"b"}}, [][]string{{"a", "b"}, {"b", "c"}}, false},
		{[][]string{{"a", "b"}, {"c", "b"}}, [][]string{{"a", "b"}, {"b", "c"}}, false},
		{[][]string{{"a", "b"}, {"b", "c"}}, [][]string{{"a", "b"}, {"b", "c"}}, true},
	}

	for _, tt := range testData {
		excepted := IsEqualsDoubleArrString(tt.arg1, tt.arg2)
		if excepted != tt.expected {
			t.Errorf("%v %v : expected %v, got %v", tt.arg1, tt.arg2, tt.expected, excepted)
		}
	}
}

// 判断两个字符串数组是否相等
func TestIsEqualsArrString(t *testing.T) {
	testData := []struct {
		arg1     []string
		arg2     []string
		expected bool
	}{
		{nil, nil, true},
		{nil, []string{"a", "b"}, false},
		{[]string{"a", "b"}, nil, false},
		{[]string{"a", "b"}, []string{"a"}, false},
		{[]string{"a", "b"}, []string{"a", "b"}, true},
		{[]string{"a", "b"}, []string{"a", "b", "b"}, false},
		{[]string{"a", "b"}, []string{"a", "b", "b", "c"}, false},
		{[]string{"a", "b", "c", "b"}, []string{"a", "b", "b", "c"}, false},
		{[]string{"a", "b", "b", "c"}, []string{"a", "b", "b", "c"}, true},
	}

	for _, tt := range testData {
		excepted := IsEqualsArrString(tt.arg1, tt.arg2)
		if excepted != tt.expected {
			t.Errorf("expected %v, got %v", excepted, tt.expected)
		}
	}
}

// 测试将一维数组拆分为二维数组
func TestSplitArrString(t *testing.T) {
	testData := []struct {
		arg      []string
		lines    int
		expected [][]string
	}{
		{[]string{"a", "b", "c"}, 2, [][]string{{"a", "b"}, {"b", "c"}}},
		{[]string{"a", "b"}, 2, [][]string{{"a", "b"}}},
		{[]string{"a"}, 2, [][]string{{"a"}}},
		{nil, 2, [][]string{}},
		{[]string{"a"}, 0, [][]string{{"a"}}},
		{[]string{"a", "b"}, 0, [][]string{{"a"}, {"b"}}},
	}

	for _, tt := range testData {
		result := SplitArrString(tt.arg, tt.lines)
		if !IsEqualsDoubleArrString(result, tt.expected) {
			t.Errorf("expected %v but got %v", tt.expected, result)
		}
	}
}

// 测试字符串中相同的连续子串
func TestGetStringSameSlice(t *testing.T) {
	testData := []struct {
		str1     string
		str2     string
		splitSep string
		lines    int
		expected [][]int
	}{
		{"a b c", "a b c d e", " ", 2,
			[][]int{
				{0, 1, 0, 1},
				{1, 2, 1, 2},
			},
		},
		{"a b c d e f g", "a b c d e h i j k l m n", " ", 4,
			[][]int{
				{0, 3, 0, 3},
				{1, 4, 1, 4},
			},
		},
	}

	for _, tt := range testData {
		result := GetStringSameSlice(tt.str1, tt.str2, tt.splitSep, tt.lines)
		if !IsEqualsDoubleArrInt(result, tt.expected) {
			t.Errorf("expected %v, got %v", tt.expected, result)
		}
	}
}

// 测试二维整数数组合并
func TestMergeSerialSliceFromDoubleArrInt(t *testing.T) {
	testData := []struct {
		arr      [][]int
		expected [][]int
	}{
		{
			nil,
			[][]int{},
		},
		{
			[][]int{},
			[][]int{},
		},
		{
			[][]int{{0, 1, 0, 1}},
			[][]int{{0, 1, 0, 1}},
		},
		{
			[][]int{{0, 1, 0, 1}, {1, 2, 1, 2}},
			[][]int{{0, 2, 0, 2}},
		},
		{
			[][]int{{0, 3, 0, 3}, {1, 4, 1, 4}},
			[][]int{{0, 4, 0, 4}},
		},
		{
			[][]int{{0, 3, 0, 3}, {1, 4, 1, 4}, {5, 8, 10, 13}},
			[][]int{{0, 4, 0, 4}, {5, 8, 10, 13}},
		},
		{
			[][]int{{0, 3, 0, 3}, {1, 4, 1, 4}, {5, 8, 10, 13}, {7, 10, 12, 15}},
			[][]int{{0, 4, 0, 4}, {5, 10, 10, 15}},
		},
		{
			[][]int{{0, 9, 0, 9}, {1, 4, 1, 4}, {5, 8, 10, 13}, {7, 10, 12, 15}},
			[][]int{{0, 4, 0, 4}, {5, 10, 10, 15}},
		},
	}

	for _, tt := range testData {
		result := MergeSerialSliceFromDoubleArrInt(tt.arr)
		if !IsEqualsDoubleArrInt(result, tt.expected) {
			t.Errorf("expected %v, got %v", tt.expected, result)
		}
	}
}
func TestMergeSerialSliceFromDoubleArrInt1(t *testing.T) {
	// arr := [][]int{{329, 338, 286, 295}, {329, 338, 289, 298}, {329, 338, 323, 332}, {329, 338, 326, 335}, {329, 338, 329, 338}, {329, 338, 332, 341}, {329, 338, 335, 344}, {330, 339, 284, 293}, {330, 339, 287, 296}, {330, 339, 324, 333}, {330, 339, 327, 336}, {330, 339, 330, 339}, {330, 339, 333, 342}, {331, 340, 285, 294}, {331, 340, 288, 297}, {331, 340, 325, 334}, {331, 340, 328, 337}, {331, 340, 331, 340}, {331, 340, 334, 343}, {332, 341, 283, 292}, {332, 341, 286, 295}, {332, 341, 289, 298}, {332, 341, 323, 332}, {332, 341, 326, 335}, {332, 341, 329, 338}, {332, 341, 332, 341}, {332, 341, 335, 344}, {333, 342, 284, 293}, {333, 342, 287, 296}, {333, 342, 324, 333}, {333, 342, 327, 336}, {333, 342, 330, 339}, {333, 342, 333, 342}, {334, 343, 285, 294}, {334, 343, 288, 297}, {334, 343, 325, 334}, {334, 343, 328, 337}, {334, 343, 331, 340}, {334, 343, 334, 343}, {335, 344, 283, 292}, {335, 344, 286, 295}, {335, 344, 289, 298}, {335, 344, 323, 332}, {335, 344, 326, 335}, {335, 344, 329, 338}, {335, 344, 332, 341}, {335, 344, 335, 344}, {336, 345, 336, 345}}
	arr := [][]int{{0, 3, 0, 3}, {1, 4, 1, 4}, {5, 8, 10, 13}}
	result := MergeSerialSliceFromDoubleArrInt(arr)
	t.Log(result)
}

// 测试驼峰转蛇形
func TestStringCamelToSnake(t *testing.T) {
	testData := []struct {
		s        string
		expected string
	}{
		{"Abc", "abc"},
		{"AbcDDD", "abc_d_d_d"},
		{"helloWorld", "hello_world"},
	}

	for _, tt := range testData {
		snake := StringCamelToSnake(tt.s)
		if snake != tt.expected {
			t.Errorf("expected %v got %v", tt.expected, snake)
		}
	}
}

// 测试蛇形转驼峰
func TestStringSnakeToCamel(t *testing.T) {
	testData := []struct {
		s        string
		expected string
	}{
		{"abc", "Abc"},
		{"abc_d_d_d", "AbcDDD"},
		{"hello_world", "HelloWorld"},
	}

	for _, tt := range testData {
		snake := StringSnakeToCamel(tt.s)
		if snake != tt.expected {
			t.Errorf("expected %v got %v", tt.expected, snake)
		}
	}
}

// 测试蛇形转小驼峰
func TestStringSnakeToCamelSmall(t *testing.T) {
	testData := []struct {
		s        string
		expected string
	}{
		{"abc", "abc"},
		{"abc_d_d_d", "abcDDD"},
		{"hello_world", "helloWorld"},
	}

	for _, tt := range testData {
		snake := StringSnakeToCamelSmall(tt.s)
		if snake != tt.expected {
			t.Errorf("expected %v got %v", tt.expected, snake)
		}
	}
}
