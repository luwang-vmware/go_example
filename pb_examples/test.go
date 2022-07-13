package main

import (
	"encoding/json"
	"fmt"

	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/types/known/structpb"
)

func main() {
	m := map[string]interface{}{
		"foo": "bar",
		"baz": 123,
	}
	b, _ := json.Marshal(m)
	s := &structpb.Struct{}
	_ = protojson.Unmarshal(b, s)
	fmt.Println(s)

	t, _ := structpb.NewStruct(m)
	fmt.Println(t)

	u, _ := structpb.NewValue(m)
	fmt.Println(u)

	v := &structpb.Value{}
	_ = protojson.Unmarshal(b, v)
	fmt.Println(v)
}
