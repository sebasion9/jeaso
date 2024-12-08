package main

import "fmt"
type Sortable interface {
	Sort(string, string, string)
	Reverse()
}

func dive(step interface{}, target string, found *[]interface{}) {
	obj, ok := step.(map[string]interface{})
	if ok {
		for v := range obj {
			dive(obj[v], target, found)
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
		dive(arr[i], target, found)
	}
}

func sort(src interface{}, target_arr string, key string, dir string) {
	var found []interface{}
	arr, ok := src.([]interface{})
	if ok {
		for i := range arr {
			dive(arr[i], target_arr, &found)
		}
	}
	obj, ok := src.(map[string]interface{})
	if ok {
		for k := range obj {
			dive(obj[k], target_arr, &found)
		}
	}
	fmt.Println(found)
}

func (a *Arr) Sort(target_arr string, key string, dir string) {
	// same is for obj, but here is top level array to consider
	sort(a.arr, target_arr, key, dir)
}

func (o *Obj) Sort(target_arr string, key string, dir string) {
	sort(o.obj, target_arr, key, dir)
}

func (a *Arr) Reverse() {

}

func (o *Obj) Reverse() {

}
