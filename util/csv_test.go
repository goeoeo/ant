package util

import (
	"fmt"
	"testing"
)

func TestNewCsv(t *testing.T) {
	csv := NewCsv("id", "name")

	for i := 0; i <= 10; i++ {
		if i == 5 {
			csv.Wr.Flush()
		}
		csv.Row("a", "b")
	}

	fmt.Println(string(csv.Bytes()))
}
