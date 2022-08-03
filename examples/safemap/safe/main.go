package main

import (
	"fmt"
	"github.com/zhangdapeng520/zdpgo_type/maps/safemap"
)

func main() {
	smap := new(safemap.SafeMap[string, string])
	smap.Set("test.py", "xxx")
	fmt.Println(smap.Get("test.py"))
}
