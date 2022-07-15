package zdpgo_type

import (
	"fmt"
	"testing"
)

func TestSet_basic(t *testing.T) {
	s := NewSet[int]()
	s.Add(1, 2, 3, 3, 2, 1)
	fmt.Println(s.ToSlice())

	s.Remove(2, 3)
	fmt.Println(s.ToSlice())
}
