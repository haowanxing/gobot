package to

import (
	"log"
	"strconv"
)

func Int64(in interface{}) int64 {
	var r int64 = 0
	switch t := in.(type) {
	case string:
		r, _ = strconv.ParseInt(t, 10, 64)
	case []byte:
		r, _ = strconv.ParseInt(string(t), 10, 64)
	case int64:
		r = t
	case int32:
		r = int64(t)
	case int:
		r = int64(t)
	default:
		log.Printf("to.Int64 not match type: %v - %v", in, t)
	}
	return r
}

func String(in interface{}) string {
	var s string = ""
	switch t := in.(type) {
	case string:
		s = t
	case []byte:
		s = string(t)
	case rune:
		s = String(t)
	case int64:
		s = strconv.FormatInt(t, 10)
	case int:
		s = strconv.Itoa(t)
	default:
		log.Printf("to.Int64 not match type: %v - %v", in, t)
	}
	return s
}

func Int(in interface{}) int {
	var i int = 0
	switch t := in.(type) {
	case int:
		i = t
	case int64:
		i = int(t)
	case int32:
		i = int(t)
	case string:
		i, _ = strconv.Atoi(t)
	case []byte:
		i, _ = strconv.Atoi(string(t))
	default:
		log.Printf("to.Int64 not match type: %v - %v", in, t)
	}
	return i
}
