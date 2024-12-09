package core

import "fmt"

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

func (c *JSONCore) sort(src interface{}, query string, key string, dir string) {
	target_arr, esc := c.parser.ParseSortQuery(query)
	//idx, digits,err := parse_idx_operator(query, esc)
	_, digits,err := c.parser.ParseIdxOperator(query, esc)
	if err != nil {
		fmt.Println(err)
		return
	}

	target_arr = target_arr[:len(target_arr) - digits]
	fmt.Println(target_arr)

	var found []interface{}
	arr, ok := src.([]interface{})
	if ok {
		for i := range arr {
			c.dive(arr[i], target_arr, &found)
		}
	}
	obj, ok := src.(map[string]interface{})
	if ok {
		for k := range obj {
			c.dive(obj[k], target_arr, &found)
		}
	}
	// if idx < len(found) && idx >= 0 {
	// 	fmt.Println(found[idx])
	// } else {
	// 	fmt.Println(found)
	// }
}

func (c *JSONCore) Sort(arr_key string, by_key string, order string) {
	c.sort(c.subject, arr_key, by_key, order)
}
