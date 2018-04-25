package properties

import (
	"strconv"
	"fmt"
	h "app/helpers"
)

var data map[string]string = make(map[string]string)

func Set(name string, val string) string {
	data[name] = val
	return data[name]
}

func Get(name string) string {
	if val, ok := data[name]; ok {
		return val
	}
	return ""
}

func Has(name string) bool {
	if _, ok := data[name]; ok {
		return true
	}
	return false
}

func Del(name string) bool {
	if Has(name) {
		delete(data, name)
		return true
	}
	return false
}

func GetInt(name string) int {
	if Has(name) {
		u, err := strconv.Atoi(Get(name))
		if err != nil {
			return 0
		}
		return u
	}
	return 0
}

func Dump() {
	for key, val := range data {
		//fmt.Println("kekus")
		h.Log(fmt.Sprintf("properties -> %s = %s", key, val), "debug")
	}
}


