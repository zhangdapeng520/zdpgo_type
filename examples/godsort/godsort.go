package main

import "github.com/zhangdapeng520/zdpgo_type/utils"

// SortExample to demonstrate basic usage of basic sort
func main() {
	strings := []string{}                       // []
	strings = append(strings, "d")              // ["d"]
	strings = append(strings, "a")              // ["d","a"]
	strings = append(strings, "b")              // ["d","a",b"
	strings = append(strings, "c")              // ["d","a",b","c"]
	utils.Sort(strings, utils.StringComparator) // ["a","b","c","d"]
}
