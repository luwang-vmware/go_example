package main

import (
	"fmt"
	"google.golang.org/protobuf/types/known/timestamppb"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func main() {
	var t metav1.Time
	dd := timestamppb.New(t.Time)
	fmt.Println(dd.Seconds) // output -62135596800
}
