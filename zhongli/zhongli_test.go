package zhongli

import (
	"fmt"
	"sort"
	"testing"
)

func Test(t *testing.T) {
	m := map[string]string{
		"az": "111",
		"ba": "333",
		"cb": "222",
	}
	var o Order
	o.Mode = "key"
	for k, v := range m {
		o.List = append(o.List, Rule{k, v})
	}
	sort.Sort(o)
	fmt.Println(o.List)
}
