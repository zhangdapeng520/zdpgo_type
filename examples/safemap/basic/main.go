package main

/*
@Time : 2022/6/6 20:05
@Author : 张大鹏
@File : main.go
@Software: Goland2021.3.1
@Description:
*/

import (
	"fmt"
	. "github.com/zhangdapeng520/zdpgo_type/maps/safemap"
)

func main() {
	// 创建map1
	m := new(Map[string, int])

	// 创建map2
	var m1 = new(Map[string, int])
	m1.Set("b", 11)
	m1.Set("c", 22)

	// 合并两个map
	m.Merge(m1)
	m.Set("a", 1)

	// 求和
	sum := 0
	m.Range(func(k string, v int) bool {
		fmt.Println(k, v)
		sum += v
		return true
	})
	fmt.Println(sum)

	// 过滤：只保留大于10的
	m.Range(func(k string, v int) bool {
		if v <= 10 {
			m.Delete(k)
		}
		return true
	})
	fmt.Println(m)
}
