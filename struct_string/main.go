package main

import (
	"encoding/json"
	"fmt"
	"time"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/types"
)

type SS struct {
	Name string
}
type money struct {
	Base     string  `json:"base"`
	Currency string  `json:"currency"`
	Amount   float32 `json:"amount"`
}

type info struct {
	Data money
}

// Status info of MD.
type Status struct {
	Name              string      `json:"name"`
	Phase             string      `json:"phase"`
	UID               types.UID   `json:"uid"`
	CreationTimestamp metav1.Time `json:"creationTimestamp"`
	UpdateTimestamp   metav1.Time `json:"updateTimestamp"`
	ResourceVersion   string      `json:"resourceVersion"`
}

type NewStatus struct {
	Name              string      `json:"name"`
	Phase             string      `json:"phase"`
	UID               types.UID   `json:"uid"`
	CreationTimestamp metav1.Time `json:"creationTimestamp"`
	UpdateTimestamp   metav1.Time `json:"updateTimestamp"`
	ResourceVersion   string      `json:"resourceVersion"`
}

func (s *Status) test() {
	fmt.Println("-------")
}

type NPStatus map[string]NewStatus

type ss string

func main() {
	str := `{"data":{"base":"BTC","currency":"USD","amount":4225.87}}`

	var i info

	if err := json.Unmarshal([]byte(str), &i); err != nil {
		fmt.Println("ugh: ", err)
	}

	fmt.Println("info: ", i)
	fmt.Println("currency: ", i.Data.Currency)

	bb, _ := json.Marshal(i)
	fmt.Println(string(bb))

	ss := NewStatus{
		Name:              "test",
		Phase:             "running",
		UID:               "helloIamuniq",
		CreationTimestamp: metav1.Time{Time: time.Now()},
		UpdateTimestamp:   metav1.Time{Time: time.Now()},
		ResourceVersion:   "11111",
	}

	fmt.Println("%#v", ss)

	maps := NPStatus{"aa": ss, "bb": ss}

	fmt.Println("------------")
	fmt.Println("%#v", maps)

	nn, _ := json.Marshal(maps)
	fmt.Println("------------")
	fmt.Println(string(nn))

	var maps2 NPStatus
	if err := json.Unmarshal(nn, &maps2); err != nil {
		fmt.Println("ugh: ", err)
	}
	fmt.Println("%+v", maps2)

	//ss2 := "\"name\":\"demo-cluster-demo-pool-srjf6\",\"phase\":\"Running\",\"uid\":\"cd4ed823-aecb-4131-b489-ff9ba3a9dc25\",\"creationTimestamp\":\"2022-08-18T02:08:12Z\",\"updateTimestamp\":\"2022-08-23T02:35:16Z\",\"resourceVersion\":\"8909745\""
	ss2 := "{\"name\":\"demo-cluster-node-pool-1-z6b46\",\"phase\":\"Running\",\"uid\":\"1bbbba48-2cf9-4389-a6db-6b5547ff9c46\",\"creationTimestamp\":\"2022-08-17T10:16:22Z\",\"updateTimestamp\":\"2022-08-19T23:23:05Z\",\"resourceVersion\":\"5538466\"}"
	fmt.Println(ss2)
	fmt.Println([]byte(ss2))

	var ss2Status NewStatus
	if err := json.Unmarshal([]byte(ss2), &ss2Status); err != nil {
		fmt.Println("ugh: ", err)
	}
	fmt.Printf("%+n", ss2Status)

	//maps1 := make(map[string]*SS)

	fmt.Println("-------")
	var tt *SS
	//tt = maps1["test"]
	fmt.Println(tt)
	fmt.Println(tt == nil)
}
