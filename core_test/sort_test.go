package core_test
//
// import (
// 	"testing"
// )
//
// const TLA_NUMS = `
// [
// 	3,
// 	2,
// 	1
// ]`
//
// const TLA_OBJS = `
// [
// 	{
// 		"a": 3,
// 		"b": 3
// 	},
// 	{
// 		"a": 2,
// 		"b": 3
// 	},
// 	{
// 		"a": 1,
// 		"b": 3
// 	},
// 	{
// 		"a": 1,
// 		"b": 1
// 	}
// ]`
// // a[0], "", asc|desc -> sorts first found a
// // a,"", asc|desc -> sorts every a
// const TLA_NESTED_ARR = `
// [
// 	{
// 		"a": [3,2,1],
// 		"b": 10
// 	},
// 	{
// 		"a": [5,7,1],
// 		"b": 2
// 	}
// ]
// `
// const TLO_NESTED = `
// {
// 	"a" : {
// 		"b" : {
// 			"c" : [
// 				{
// 					"d" : "hello"
// 				},
// 				{
// 					"d" : "ben"
// 				}
// 			]
// 		}
// 	}
// }
// `
//
//
// func TestSortTLA_NUMS(t *testing.T) {
// 	var djson DJSON
// 	djson.UnmarshalJSON([]byte(TLA_NUMS))
// 	djson.serializable.Sort("", "", "asc")
// 	got, _ := djson.MarshalJSON()
// 	want := `[1,2,3]`
// 	if string(got) != want {
// 		t.Errorf("\ngot:\t%s\nwanted:\t%s", got, want)
// 	}
// }
// func TestSortTLA_OBJS(t *testing.T) {
// 	var djson DJSON
// 	djson.UnmarshalJSON([]byte(TLA_OBJS))
// 	djson.serializable.Sort("", "a", "asc")
// 	got, _ := djson.MarshalJSON()
// 	want := `[{"a":1,"b":3},{"a":1,"b":1},{"a":2,"b":3},{"a":3,"b":3}]`
// 	if string(got) != want {
// 		t.Errorf("\ngot:\t%s\nwanted:\t%s", got, want)
// 	}
// }
// func TestSortTLA_OBJS_MULT_KEYS(t *testing.T) {
// 	var djson DJSON
// 	djson.UnmarshalJSON([]byte(TLA_OBJS))
// 	djson.serializable.Sort("", "a;b", "asc")
// 	got, _ := djson.MarshalJSON()
// 	want := `[{"a":1,"b":1},{"a":1,"b":3},{"a":2,"b":3},{"a":3,"b":3}]`
// 	if string(got) != want {
// 		t.Errorf("\ngot:\t%s\nwanted:\t%s", got, want)
// 	}
// }
// func TestSortTLA_NESTED_ARR_ITH(t *testing.T) {
// 	var djson DJSON
// 	djson.UnmarshalJSON([]byte(TLA_NESTED_ARR))
// 	djson.serializable.Sort("a[0]", "", "asc")
// 	got, _ := djson.MarshalJSON()
// 	want := `[{"a":[1,2,3],"b":10},{"a":[5,7,1],"b":2}]`
// 	if string(got) != want {
// 		t.Errorf("\ngot:\t%s\nwanted:\t%s", got, want)
// 	}
// }
// func TestSortTLA_NESTED_ARR_EVERY(t *testing.T) {
// 	var djson DJSON
// 	djson.UnmarshalJSON([]byte(TLA_NESTED_ARR))
// 	djson.serializable.Sort("a", "", "asc")
// 	got, _ := djson.MarshalJSON()
// 	want := `[{"a":[1,2,3],"b":10},{"a":[1,5,7],"b":2}]`
// 	if string(got) != want {
// 		t.Errorf("\ngot:\t%s\nwanted:\t%s", got, want)
// 	}
// }
// func TestSortTLA_NESTED_ARR_ROOT(t *testing.T) {
// 	var djson DJSON
// 	djson.UnmarshalJSON([]byte(TLA_NESTED_ARR))
// 	djson.serializable.Sort("", "b", "asc")
// 	got, _ := djson.MarshalJSON()
// 	want := `[{"a":[1,5,7],"b":2},{"a":[1,2,3],"b":10}]`
// 	if string(got) != want {
// 		t.Errorf("\ngot:\t%s\nwanted:\t%s", got, want)
// 	}
// }
// func TestSortTLO_NESTED(t * testing.T) {
// 	var djson DJSON
// 	djson.UnmarshalJSON([]byte(TLO_NESTED))
// 	djson.serializable.Sort("", "", "asc")
// 	got, _ := djson.MarshalJSON()
// 	//want := `[{"a":[1,5,7],"b":2},{"a":[1,2,3],"b":10}]`
// 	want := ""
// 	if string(got) != want {
// 		t.Errorf("\ngot:\t%s\nwanted:\t%s", got, want)
// 	}
//
// }
