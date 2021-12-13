package albedo

import (
	"fmt"
	"strings"
	"testing"
)

func Test(t *testing.T) {
	order := "sort desc,id asc,time desc"
	var columns []string
	arr := strings.Split(order, ",")
	for k, v := range arr {
		if strings.Contains(v, " ") {
			columns = strings.Split(v, " ")
			if !strings.Contains(columns[0], "->") {
				columns[0] = "`" + columns[0] + "`"
			}
			v = columns[0] + " " + columns[1]
		} else {
			if !strings.Contains(v, "->") {
				v = "`" + v + "`"
			}
		}
		if k == 0 {
			order = v
		} else {
			order += "," + v
		}
	}
	fmt.Println(order)
}
