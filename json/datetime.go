package json

import (
	"fmt"
	"time"
)

type JsonDate time.Time // json日期类型

// 自动将time.Time转换为json格式的方法
func (j JsonDate) MarshalJSON() ([]byte, error) {
	var stmp = fmt.Sprintf("\"%s\"", time.Time(j).Format("2006-01-02"))
	return []byte(stmp), nil
}

// 根据uin64类型的数返回JsonDate对象
func NewJsonDateUint64(date uint64) JsonDate {
	return JsonDate(time.Unix(int64(date), 0))
}
