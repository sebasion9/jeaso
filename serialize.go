package main

import (
	"encoding/json"
)

type JSONSerializable interface {
	Sortable
	serialize() ([]byte, error)
}

type Arr struct {
	arr []interface{}
}

type Obj struct {
	obj map[string]interface{}
}

func (a *Arr) serialize() ([]byte, error) {
	return json.Marshal(a.arr)
}

func (o *Obj) serialize() ([]byte, error) {
	return json.Marshal(o.obj)
}

type DJSON struct {
	serializable JSONSerializable
}

func (dj *DJSON) UnmarshalJSON(b []byte) error {
	var err error = nil
	// unmarshal to array or obj
	arr := make([]interface{}, 1)
	obj := make(map[string]interface{})
	err = json.Unmarshal(b, &arr)

	if err != nil {
		// failed to parse into array
		err = json.Unmarshal(b, &obj)
		if err != nil {
			// failed to parse into array and object
			return err
		}
	}
	if arr[0] != nil {
		dj.serializable = &Arr{arr}
	}
	if len(obj) != 0 {
		dj.serializable = &Obj{obj}
	}
	return err
}

func (dj *DJSON) MarshalJSON() ([]byte, error) {
	return dj.serializable.serialize()
}
