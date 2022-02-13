package pkg

import (
	"runtime"
	"strings"
)

func RuntimeVersionMap() map[string]string {
	version := make(map[string]string, 5)
	version["runtime"] = runtime.Version()
	version["string"] = strings.Replace(version["runtime"], "go", "", 1)
	goVersionArr := strings.Split(version["string"], ".")
	version["major"] = goVersionArr[0]
	version["minor"] = goVersionArr[1]
	version["patch"] = goVersionArr[2]
	return version
}
