package zhongli

import (
	"fmt"
	"sort"
	"testing"
)

func Test(t *testing.T) {
	m := map[string]string{
		"1": "111",
		"2": "333",
		"3": "222",
	}
	var o Order
	o.Mode = "value"
	for k, v := range m {
		o.List = append(o.List, Rule{k, v})
	}
	sort.Sort(o)
	fmt.Println(o.List)
}
