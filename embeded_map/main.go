package main

import (
	"encoding/json"
	"fmt"
)

type Item map[string]string
type AggreateItems map[string]Item

type NestedMap map[string]map[string]string

func ParseData(data string) NestedMap {
	// var outerMap map[string]string
	// // Ignore error here as error should never be happened here.
	// _ = json.Unmarshal([]byte(data), &outerMap)

	// nestedMap := make(NestedMap)

	// for name, value := range outerMap {
	// 	var item map[string]string
	// 	// Ignore error here as error should never be happened here.
	// 	_ = json.Unmarshal([]byte(value), &item)
	// 	nestedMap[name] = item
	// }
	nestedMap := make(NestedMap)
	if data != "" {
		// Ignore error here as error should never be happened here.
		_ = json.Unmarshal([]byte(data), &nestedMap)
	}
	return nestedMap
}

func (nm NestedMap) Get(name string) map[string]string {
	return nm[name]
}

func (nm NestedMap) Set(name string, m map[string]string) {
	nm[name] = m
}

func (nm NestedMap) Delete(name string) {
	if len(nm) == 0 {
		return
	}
	delete(nm, name)
}

func (nm NestedMap) JSON() string {
	if len(nm) == 0 {
		return ""
	}

	// Ignore error here as error should never be happened here.
	dt, _ := json.Marshal(nm)

	return string(dt)
}

func main() {
	// nodepool1
	item1 := make(map[string]string)
	item1["key1-1"] = "value11"
	item1["key1-2"] = "value12"
	bytes1, _ := json.Marshal(item1)

	//nodepool2
	item2 := make(map[string]string)
	item2["key2-1"] = "value21"
	item2["key2-2"] = "value22"
	bytes2, _ := json.Marshal(item2)

	// nodepool
	aItem := Item{
		"item1": string(bytes1),
		"item2": string(bytes2),
	}
	bytesa, _ := json.Marshal(aItem)
	fmt.Println(string(bytesa))

	// // save as aItem
	// bItem := make(map[string]string)
	// bItem["item1"] = string(bytes1)
	// bItem["item2"] = string(bytes2)
	// bytesb, _ := json.Marshal(bItem)
	// fmt.Println(string(bytesb))

	outerItem := Item{
		"name":     "luwang",
		"age":      "37",
		"nodepool": string(bytesa),
	}

	bytes, _ := json.Marshal(outerItem)
	fmt.Println(string(bytes))

	var output Item
	_ = json.Unmarshal(bytes, &output)
	fmt.Printf("------%+v", output)

	var test map[string]string
	b, err := json.Marshal(test)
	fmt.Println(err)
	fmt.Println(string(b))

	err = json.Unmarshal([]byte(""), &test)
	fmt.Println(err)
	fmt.Println(test)

	m := NestedMap{
		"poo1": map[string]string{
			"type":  "qe",
			"size":  "small",
			"owner": "qe",
		},
		"poo2": map[string]string{
			"type":  "dev",
			"size":  "small",
			"owner": "dev",
		},
		"poo3": map[string]string{
			"type":  "dev",
			"size":  "big",
			"owner": "dev",
		},
	}
	fmt.Println(m)
	bb, _ := json.Marshal(m)
	fmt.Println(string(bb))

	newM := ParseData(string(bb))
	newM.Set("pool4", map[string]string{"type": "new"})
	newM.Delete("pool5")
	fmt.Println(newM)
	fmt.Println(newM.JSON())

	mm := make(map[string]string)
	delete(mm, "whjo")
	fmt.Println(mm)
}
