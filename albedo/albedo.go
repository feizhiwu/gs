package albedo

import (
	"crypto/md5"
	"encoding/json"
	"errors"
	"fmt"
	"math/rand"
	"strconv"
	"time"
)

func MakeUint(num interface{}) uint {
	switch num.(type) {
	case uint:
		return num.(uint)
	case int:
		return uint(num.(int))
	case uint8:
		return uint(num.(uint8))
	case float32:
		return uint(num.(float32))
	case float64:
		return uint(num.(float64))
	case string:
		i, _ := strconv.Atoi(num.(string))
		return uint(i)
	default:
		return 0
	}
}

func MakeInt(num interface{}) int {
	switch num.(type) {
	case int:
		return num.(int)
	default:
		return int(MakeUint(num))
	}
}

func MakeUint32(num interface{}) uint32 {
	switch num.(type) {
	case uint32:
		return num.(uint32)
	default:
		return uint32(MakeUint(num))
	}
}

func MakeFloat32(num interface{}) float32 {
	switch num.(type) {
	case float32:
		return num.(float32)
	case float64:
		return float32(num.(float64))
	case string:
		f, _ := strconv.ParseFloat(num.(string), 64)
		return float32(f)
	default:
		return 0
	}
}

func MakeString(str interface{}) string {
	switch str.(type) {
	case string:
		return str.(string)
	case int, uint, int8, uint8, int16, uint16, int32, uint32, int64, uint64:
		return strconv.Itoa(MakeInt(str))
	case float32:
		return strconv.FormatFloat(float64(str.(float32)), 'f', -1, 32)
	case float64:
		return strconv.FormatFloat(str.(float64), 'f', -1, 32)
	default:
		return ""
	}
}

func MakeJson(data interface{}) []byte {
	b, _ := json.Marshal(data)
	return b
}

func MakeStruct(data, res interface{}) {
	json.Unmarshal(MakeJson(data), &res)
}

func MakeMap(data string) map[string]interface{} {
	var m map[string]interface{}
	json.Unmarshal([]byte(data), &m)
	return m
}

func MakeMaps(data string) []map[string]interface{} {
	var ms []map[string]interface{}
	json.Unmarshal([]byte(data), &ms)
	return ms
}

func MakeUniqId() string {
	nano := time.Now().UnixNano()
	sum := md5.Sum([]byte(strconv.FormatInt(nano, 10)))
	return fmt.Sprintf("%x", sum)
}

func MakeError(err interface{}) error {
	return errors.New(MakeString(err))
}

func MakeSn() string {
	rand.Seed(time.Now().UnixNano())
	return time.Now().Format("20060102") + strconv.FormatInt(time.Now().Unix(), 10) + strconv.Itoa(rand.Intn(9999999-1000000)+1000000)
}
