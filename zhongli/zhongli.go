package zhongli

import (
	"github.com/feizhiwu/gs/albedo"
	"strconv"
)

type Rule struct {
	Key   interface{}
	Value interface{}
}

type Order struct {
	List []Rule
	Mode string
}

func (o Order) Swap(i, j int) {
	o.List[i], o.List[j] = o.List[j], o.List[i]
}

func (o Order) Len() int {
	return len(o.List)
}

func (o Order) Less(i, j int) bool {
	if o.Mode == "value" {
		if isNum(albedo.MakeString(o.List[i].Value)) {
			return albedo.MakeInt(o.List[i].Value) < albedo.MakeInt(o.List[j].Value)
		} else {
			return albedo.MakeString(o.List[i].Value) < albedo.MakeString(o.List[j].Value)
		}
	} else {
		if isNum(albedo.MakeString(o.List[i].Key)) {
			return albedo.MakeInt(o.List[i].Key) < albedo.MakeInt(o.List[j].Key)
		} else {
			return albedo.MakeString(o.List[i].Key) < albedo.MakeString(o.List[j].Key)
		}
	}
}

func isNum(s string) bool {
	_, err := strconv.ParseFloat(s, 64)
	return err == nil
}
