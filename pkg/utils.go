package pkg

import (
	"runtime"
	"strings"
)

type RuntimeVersion struct {
	Real  string
	Str   string
	Arr   []string
	Major string
	Minor string
	Patch string
}

func (r *RuntimeVersion) ReadVersion() *RuntimeVersion {
	r.Real = runtime.Version()
	r.Str = strings.Replace(r.Real, "go", "", 1)
	r.Arr = strings.Split(r.Str, ".")
	r.Major = r.Arr[0]
	r.Minor = r.Arr[1]
	r.Patch = r.Arr[2]
	return r
}

func (r *RuntimeVersion) String() string {
	return r.Str
}
