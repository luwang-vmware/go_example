package main

import (
	"fmt"
)

type VarType struct {
	Name  string
	Value int
}

func getVarNames(variables []*VarType) []string {
	namesArr := make([]string, len(variables))
	//namesArr := []string{}
	// for _, item := range variables {
	// 	namesArr = append(namesArr, item.Name)
	// }
	for idx, item := range variables {
		namesArr[idx] = item.Name
	}
	fmt.Println(len(namesArr))
	return namesArr
}

func SetMissingVariables(source []*VarType, sub []*VarType) []*VarType {
	var addedVariables []*VarType

	searchVarNames := getVarNames(source)
	var found bool
	for idx, item := range sub {
		fmt.Println(idx)
		found = false
		for _, name := range searchVarNames {
			if item.Name == name {
				fmt.Printf("%v already exists\n", item.Name)
				found = true
				break
			}
		}
		if found {
			continue
		}
		fmt.Printf("found missing %v\n", item.Name)
		defaultValue := 0
		newAdded := &VarType{Name: item.Name, Value: defaultValue}
		addedVariables = append(addedVariables, newAdded)
	}
	return addedVariables
}

func main() {
	source1 := make([]*VarType, 4)
	source1[0] = &VarType{Name: "aaa", Value: 0}
	//fmt.Printf("%+v", source1)
	source := []*VarType{
		{Name: "a", Value: 1},
		{Name: "b", Value: 2},
		{Name: "d", Value: 4},
		{Name: "f", Value: 1},
	}
	//fmt.Println(source)
	sub := []*VarType{
		{Name: "ee", Value: 1},
		{Name: "a", Value: 2},
		{Name: "h", Value: 1},
		{Name: "d", Value: 1},
	}
	
	//fmt.Println(sub)

	out := SetMissingVariables(source, sub)
	fmt.Println(out)
}
