package benchmark

import (
	"strings"
)

func Concat(a, b string) string {
	// 1
	//return fmt.Sprintf("%s%s", a, b)

	// 2
	builder := strings.Builder{}

	builder.WriteString(a)

	return builder.String()
}
