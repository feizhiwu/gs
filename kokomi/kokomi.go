package kokomi

import (
	"github.com/feizhiwu/gs/albedo"
	"github.com/jinzhu/gorm"
	"reflect"
	"strconv"
	"strings"
	"unsafe"
)

type query struct {
	db    **gorm.DB
	data  map[string]interface{}
	page  uint
	limit uint
	field string
	key   string
	num   uint16
}

func NewQuery(db **gorm.DB, data map[string]interface{}) query {
	var page, limit uint
	if albedo.MakeUint(data["page"]) != 0 {
		page = albedo.MakeUint(data["page"])
	} else {
		page = 1
	}
	if albedo.MakeUint(data["limit"]) != 0 {
		limit = albedo.MakeUint(data["limit"])
	} else {
		limit = 20
	}
	newQuery := query{db: db, data: data, page: page, limit: limit}
	if data["order"] != nil {
		order := albedo.MakeString(data["order"])
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
		*newQuery.db = (*newQuery.db).Order(order)
	} else {
		*newQuery.db = (*newQuery.db).Order("id desc")
	}
	return newQuery
}

func (q *query) Eq(args ...string) {
	if q.checkKey(args...) {
		*q.db = (*q.db).Where(q.field+" = ?", q.data[q.key])
	}
}

func (q *query) Gt(args ...string) {
	if q.checkKey(args...) {
		*q.db = (*q.db).Where(q.field+" > ?", q.data[q.key])
	}
}

func (q *query) Gte(args ...string) {
	if q.checkKey(args...) {
		*q.db = (*q.db).Where(q.field+" >= ?", q.data[q.key])
	}
}

func (q *query) Lt(args ...string) {
	if q.checkKey(args...) {
		*q.db = (*q.db).Where(q.field+" < ?", q.data[q.key])
	}
}

func (q *query) Lte(args ...string) {
	if q.checkKey(args...) {
		*q.db = (*q.db).Where(q.field+" <= ?", q.data[q.key])
	}
}

func (q *query) Like(args ...string) {
	if q.checkKey(args...) {
		*q.db = (*q.db).Where(q.field+" like ?", "%"+q.data[q.key].(string)+"%")
	}
}

// AwLike after wildcard后通配
func (q *query) AwLike(args ...string) {
	if q.checkKey(args...) {
		*q.db = (*q.db).Where(q.field+" like ?", q.data[q.key].(string)+"%")
	}
}

func (q *query) In(args ...string) {
	if q.checkKey(args...) {
		*q.db = (*q.db).Where(q.field+" in (?)", q.data[q.key])
	}
}

func (q *query) NotIn(args ...string) {
	if q.checkKey(args...) {
		*q.db = (*q.db).Where(q.field+" not in (?)", q.data[q.key])
	}
}

func (q *query) EqZero(args ...string) {
	q.checkKey(args...)
	*q.db = (*q.db).Where(q.field + " = 0")
}

func (q *query) GtZero(args ...string) {
	q.checkKey(args...)
	*q.db = (*q.db).Where(q.field + " > 0")
}

func (q *query) IsEmpty(args ...string) {
	if q.checkKey(args...) {
		parseBool, _ := strconv.ParseBool(albedo.MakeString(q.data[q.key]))
		if parseBool {
			*q.db = (*q.db).Where(q.field + " <> ''")
		} else {
			*q.db = (*q.db).Where(q.field + " is null or " + q.field + " = ''")
		}
	}
}

func (q *query) Null(args ...string) {
	q.checkKey(args...)
	*q.db = (*q.db).Where(q.field + " is null")
}

func (q *query) NotNull(args ...string) {
	q.checkKey(args...)
	*q.db = (*q.db).Where(q.field + " is not null")
}

//wildcard 通配
func (q *query) Wc(args ...string) {
	if q.checkKey(args...) {
		if q.data[q.key] == "*" {
			*q.db = (*q.db).Where(q.field + " > 0")
		} else {
			*q.db = (*q.db).Where(q.field+" = ?", q.data[q.key])
		}
	}
}

// Raw 原生where语句
func (q *query) Raw(query interface{}, args ...interface{}) {
	var ok bool
	var value interface{}
	var values []interface{}
	if len(args) > 0 {
		for _, v := range args {
			value, ok = q.data[v.(string)]
			values = append(values, value)
		}
		if ok {
			*q.db = (*q.db).Where(query, values...)
		}
	}
}

func (q *query) Pages(value interface{}) *query {
	t := reflect.TypeOf(value).Elem()
	v := reflect.ValueOf(value)
	for i := 0; i < t.NumField(); i++ {
		name := t.Field(i).Name
		switch name {
		case "Page":
			*(*uint)(unsafe.Pointer(v.Elem().FieldByName(name).Addr().Pointer())) = q.page
		case "Limit":
			*(*uint)(unsafe.Pointer(v.Elem().FieldByName(name).Addr().Pointer())) = q.limit
		case "Count":
			var count uint
			(*q.db).Count(&count)
			*(*uint)(unsafe.Pointer(v.Elem().FieldByName(name).Addr().Pointer())) = count
		}
	}
	return q
}

func (q *query) List(value interface{}) *query {
	(*q.db).Limit(q.limit).Offset(getOffset(q.page, q.limit)).Find(value)
	return q
}

func (q *query) checkKey(args ...string) bool {
	arr := strings.Split(args[0], ".")
	if len(args) == 1 {
		if len(arr) == 1 {
			q.field = "`" + args[0] + "`"
			q.key = arr[0]
		} else {
			q.field = args[0]
			q.key = arr[1]
		}
	} else {
		if len(arr) == 1 {
			q.field = "`" + args[0] + "`"
		} else {
			q.field = args[0]
		}
		q.key = args[1]
	}
	if q.data[q.key] == nil {
		return false
	}
	q.num++
	return true
}

func getOffset(page interface{}, limit uint) uint {
	num := albedo.MakeUint(page) - 1
	if num < 0 {
		num = 0
	}
	return num * limit
}