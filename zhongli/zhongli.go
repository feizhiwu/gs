package zhongli

import "strconv"

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
	if o.Mode == "key" {
		return makeInt(o.List[i].Key) < makeInt(o.List[j].Key)
	}
	return makeInt(o.List[i].Value) < makeInt(o.List[j].Value)
}

func makeInt(num interface{}) int {
	switch num.(type) {
	case int:
		return num.(int)
	case uint:
		return int(num.(uint))
	case float32:
		return int(num.(float32))
	case float64:
		return int(num.(float64))
	case string:
		i, _ := strconv.Atoi(num.(string))
		return i
	default:
		return 0
	}
}
