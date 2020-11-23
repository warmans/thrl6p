package flag

import (
	goflag "flag"
	"fmt"
	"os"
	"strings"
)

func StringVarEnv(s *string, prefix string, name string, value string, usage string) {
	goflag.StringVar(s, name, value, usage)
	stringFromEnv(s, prefix, name)
}

func stringFromEnv(p *string, prefix, name string) {
	val := os.Getenv(fmt.Sprintf("%s_%s", strings.ToUpper(prefix), strings.ToUpper(strings.Replace(name, "-", "_", -1))))
	if val == "" {
		return
	}
	valPtr := &val
	*p = *valPtr
}

func Parse() {
	goflag.Parse()
}
