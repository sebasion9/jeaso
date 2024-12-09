package core

import (
	"fmt"
)

func (c *JSONCore) dive(step interface{}, target string, found *[]interface{}) {
	obj, ok := step.(map[string]interface{})
	if ok {
		for v := range obj {
			c.dive(obj[v], target, found)
			if v != target {
				continue;
			}
			arr, ok := obj[v].([]interface{})
			if ok {
				*found = append(*found, arr)
			}
		}
	}
	arr, ok := step.([]interface{})
	if !ok {
		return
	}
	for i := range arr {
		c.dive(arr[i], target, found)
	}
}

func (c *JSONCore) sort_enter(src interface{}, query string, by_key string, order string) {
	arr_key, idx, err  := c.parser.ParseKeyAndIdx(query)
	if err != nil {
		fmt.Println(err)
		return
	}

	var found []interface{}
	arr, ok := src.([]interface{})
	if ok {
		for i := range arr {
			c.dive(arr[i], arr_key, &found)
		}
	}
	obj, ok := src.(map[string]interface{})
	if ok {
		for k := range obj {
			c.dive(obj[k], arr_key, &found)
		}
	}
	// if found[i] array is map[string]interface{}, and by_key given, sort the array by key and order rule
	// if the array is of primitive type, sort by order rule
	// this shouldnt be done here, but in JSONCore.dive(), to modify the arrays in place
	if idx < len(found) && idx > 0 {
		// sort found[idx] array
		fmt.Println(found[idx])
	} else {
		// sort all arrays in found
		fmt.Println(found)
	}
}

func (c *JSONCore) sort(src []interface{}, by_key string, order string) []interface{} {
	return src
}

func (c *JSONCore) Sort(query string, by_key string, order string) {
	c.sort_enter(c.subject, query, by_key, order)
}
