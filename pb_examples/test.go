package main

import (
	"encoding/json"
	"fmt"

	//	nodepool "gitlab.eng.vmware.com/olympus/api/pkg/proto/manage/v1alpha1/managementcluster/provisioner/tanzukubernetescluster/nodepool"
	nodepool "gitlab.eng.vmware.com/olympus/api/pkg/proto/manage/v1alpha1/managementcluster/provisioner/tanzukubernetescluster/nodepool"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/types/known/structpb"
	"luwang.com/goexamples/pb"
)

const npJSON = `
{
    "spec": {"failure_domain1": "failureDomain", "class": "np-1", "replicas": 1},
	"info": {"name": "md-0"},
}
`

type PhoneNumber struct {
	Number string `json:"number,omitempty"`
	Type   int    `json:"type,omitempty"`
}
type person2 struct {
	Name   string        `json:"name,omitempty"`
	Id     int           `json:"id,omitempty"`
	Email  string        `json:"email,omitempty"`
	Phones []PhoneNumber `json:"phones,omitempty"`
}

type Definition struct {
	Spec *nodepool.Spec `json:"spec,omitempty"`
	Info *nodepool.Info `json:"info,omitempty"`
}

const per = `
{
	"name": "lu",
	"id": 123,
	"email": "luwang@vmware.com",
	"phones": [
		{"number": "111111", "type": 2},
		{"number": "222222", "type": 1},
	]
}
`
const phone = `
{
	"number": "123456",
	"type": 1
}
`

func main() {
	m := map[string]interface{}{
		"foo": "bar",
		"baz": 123,
	}
	fmt.Println(m)
	// map[baz:123 foo:bar]
	b, _ := json.Marshal(m)
	s := &structpb.Struct{}
	_ = protojson.Unmarshal(b, s)
	fmt.Println(s)
	// fields:{key:"baz" value:{number_value:123}} fields:{key:"foo" value:{string_value:"bar"}}

	t, _ := structpb.NewStruct(m)
	fmt.Println(t)
	// fields:{key:"baz" value:{number_value:123}} fields:{key:"foo" value:{string_value:"bar"}}

	u, _ := structpb.NewValue(m)
	fmt.Println(u)
	// struct_value:{fields:{key:"baz" value:{number_value:123}} fields:{key:"foo" value:{string_value:"bar"}}}

	v := &structpb.Value{}
	_ = protojson.Unmarshal(b, v)
	fmt.Println(v)
	// struct_value:{fields:{key:"baz" value:{number_value:123}} fields:{key:"foo" value:{string_value:"bar"}}}

	fmt.Println("---")
	p := &pb.Person{}
	_ = protojson.Unmarshal([]byte(per), p)
	fmt.Println(p.Email, p.Id, p.Name)
	fmt.Println(p)
	fmt.Println(p.GetPhones())
	//
	// for p1 := range p.Phones {
	// 	fmt.Println(p1)
	// }
	fmt.Println("---1")

	//fmt.Println(per)
	var pp person2
	_ = json.Unmarshal([]byte(per), &pp)
	fmt.Println(pp)
	fmt.Println("---2")

	var ph pb.Person_PhoneNumber
	_ = protojson.Unmarshal([]byte(phone), &ph)
	fmt.Println(ph.String())
	fmt.Println("---3")

	var ph2 PhoneNumber
	_ = json.Unmarshal([]byte(phone), &ph2)
	fmt.Println(ph2)
	fmt.Println("---4")

	bytes := []byte(npJSON)
	var temp interface{}
	_ = json.Unmarshal(bytes, temp)
	fmt.Println(temp)

	//fmt.Println(npJSON)

	v2 := &structpb.Struct{}
	_ = protojson.Unmarshal(bytes, v2)
	fmt.Println(v2)

	// var np *nodepool.Definition
	// err := protojson.Unmarshal(bytes, np)
	// fmt.Println(err)
	// fmt.Println(np)

	var def *Definition

	_ = json.Unmarshal([]byte(npJSON), def)
	//fmt.Println(err)
	fmt.Println(def)
}
