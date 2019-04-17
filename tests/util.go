package tests

import (
	"fmt"
)

//formatJSON and print beautiful
func formatJSON(f interface{}) {

	if v, ok := f.([]interface{}); ok {
		for _, vv := range v {
			formatJSON(vv)
		}
		return
	}

	for k, v := range f.(map[string]interface{}) {
		switch vv := v.(type) {
		case int:
			fmt.Println(k, "==>", vv)
		case string:
			fmt.Println(k, "==>", vv)
		case float64:
			fmt.Println(k, "==>", vv)
		}
	}
}

// decode json data to list
func decodeList(f interface{}) []map[string]interface{} {
	ret := make([]map[string]interface{}, 0)
	if v, ok := f.([]interface{}); ok {
		for _, vv := range v {
			ret = append(ret, vv.(map[string]interface{}))
		}
	}

	return ret
}

// decode json data to map
func decodeMap(f interface{}) map[string]interface{} {
	var ret map[string]interface{}
	if v, ok := f.(map[string]interface{}); ok {
		ret = v
	}

	return ret
}
