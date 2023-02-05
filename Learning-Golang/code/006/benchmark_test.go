package benchmark

import (
	"bytes"
	"strings"
	"testing"
)

func BenchmarkConcat(b *testing.B) {
	for n := 0; n < b.N; n++ {
		Concat("a", "b")
	}
}

func BenchmarkBytesBuffer(b *testing.B) {
	var buf bytes.Buffer

	for n := 0; n < b.N; n++ {
		buf.WriteString("a")
	}
}

func BenchmarkStringsBuilder(b *testing.B) {
	var builder strings.Builder
	for n := 0; n < b.N; n++ {
		builder.WriteString("a")
	}

}
