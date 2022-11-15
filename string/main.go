package main

import (
	"fmt"
	"strings"
)

func matchVersion(kcpVersion, templateVersion string) string {
	templateVerStart := strings.Index(templateVersion, "v1")
	templateVerEnd := strings.Index(templateVersion, "+")
	if templateVerStart == -1 {
		return templateVersion
	}
	if templateVerEnd == -1 {
		templateVerEnd = len(templateVersion)
	}
	KcpEnd := strings.Index(kcpVersion, "+")
	if templateVersion[templateVerStart:templateVerEnd] == kcpVersion[:KcpEnd] {
		return templateVersion
	}
	return templateVersion[:templateVerStart] + kcpVersion[:KcpEnd] + templateVersion[templateVerEnd:]
}
func main() {
	res := matchVersion("v1.22.9+vmware.1", "/dc0/vm/154RTM")
	fmt.Println(res)
}
