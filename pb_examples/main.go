package main

import (
	"fmt"

	"github.com/VirtusLab/render/renderer"
	"github.com/VirtusLab/render/renderer/parameters"
	_struct "github.com/golang/protobuf/ptypes/struct"
	"sigs.k8s.io/yaml"
)

const remoteResourceMetaTemplate = `
apiVersion: cluster.x-k8s.io/v1beta1
kind: Cluster
metadata:
  name: {{ quote .name }}
  namespace: {{ quote .namespace }}
  uid: {{ quote .uid }}
---
apiVersion: cluster.x-k8s.io/v1beta1
kind: Cluster
metadata:
  name: luwang
`

func render(template string, opts []string, params parameters.Parameters) (string, error) {
	return renderer.New(
		renderer.WithOptions(opts...),
		renderer.WithParameters(params),
		renderer.WithSprigFunctions(),
		renderer.WithExtraFunctions(),
	).Render(template)
}

func main() {
	params := parameters.Parameters{
		"name":      "test",
		"namespace": "testns",
		"uid":       "1111",
	}

	y, err := render(remoteResourceMetaTemplate, nil, params)
	if err != nil {
		fmt.Println("error")
	}
	out, _ := yaml.YAMLToJSON([]byte(y))
	fmt.Println([]byte(y))
	fmt.Println("-------")
	fmt.Println(out)

	var resource _struct.Struct
	if err := resource.UnmarshalJSON(out); err != nil {
		fmt.Println("error here")
	}
	fmt.Println(resource)

}
