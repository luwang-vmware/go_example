package main

import (
	"fmt"
	"reflect"

	"google.golang.org/protobuf/types/known/structpb"
)

//https://stackoverflow.com/questions/38748098/golang-type-switch-how-to-match-a-generic-slice-array-map-chan

type Test struct {
	Name  string
	Value []int
}

func interface_type(x interface{}) {
	switch v := x.(type) {
	case nil:
		fmt.Println("x is nil") // here v has type interface{}
	case int:
		fmt.Println("x is", v) // here v has type int
	case bool, string:
		fmt.Println("x is bool or string") // here v has type interface{}
	case []interface{}:
		fmt.Println("x is slice of any type") // here v has type []interface{}
	case []int:
		fmt.Println("x is slice of int") // here v has type []interface{}
	default:
		fmt.Println("type unknown") // here v has type interface{}
	}
}

func reflect_type(x interface{}) {
	v := reflect.ValueOf(x)
	switch v.Kind() {
	case reflect.Bool:
		fmt.Printf("bool: %v\n", v.Bool())
	case reflect.Int, reflect.Int8, reflect.Int32, reflect.Int64:
		fmt.Printf("int: %v\n", v.Int())
	case reflect.Uint, reflect.Uint8, reflect.Uint32, reflect.Uint64:
		fmt.Printf("int: %v\n", v.Uint())
	case reflect.Float32, reflect.Float64:
		fmt.Printf("float: %v\n", v.Float())
	case reflect.String:
		fmt.Printf("string: %v\n", v.String())
	case reflect.Slice:
		fmt.Printf("slice: len=%d, %v\n", v.Len(), v.Interface())
	case reflect.Map:
		fmt.Printf("map: %v\n", v.Interface())
	case reflect.Chan:
		fmt.Printf("chan %v\n", v.Interface())
	default:
		fmt.Println(x)
	}
	fmt.Println(reflect.TypeOf(x))
	default_value := reflect.Zero(reflect.TypeOf(x))
	fmt.Println(default_value)
}
func main() {
	// var x1 interface{} = []int{1, 2, 3}
	// interface_type(x1)
	// reflect_type(x1)

	var x interface{} = map[string]int{
		"1": 1,
		"2": 2,
	}
	reflect_type(x) //map[string]int, map[]
	var y interface{} = Test{
		Name:  "luwang",
		Value: []int{1, 2, 3},
	}
	reflect_type(y) //main.Test , { []}

	z := []int{1, 2, 3}
	reflect_type(z) // []int, []
	// var x interface{} = map[string][]int{
	// 	"name": [
	// 		1,
	// 		2,
	// 	],
	// 	"age": [
	// 		10,
	// 		20,
	// 	],
	// }

	//var x interface{} = "foo"

	s := make([]int, 3)
	s[0] = 1
	s[1] = 2
	s[2] = 3

	v, e := structpb.NewValue(s)
	fmt.Println(v)
	fmt.Println(e)
	v, e = structpb.NewValue([]interface{}{1, 2, 3})
	fmt.Println(v)
	fmt.Println(e)

}
