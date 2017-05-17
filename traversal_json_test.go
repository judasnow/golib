package golib

import (
	"testing"
	"fmt"
)

type JsonValue map[string]interface{}

func TestTraversalJson(t *testing.T) {
	j := map[string]interface{}{
		"foo": "bar",
		"baz": map[string]interface{}{"bazbaz": "baz"},
	}
	TraversalJson(j, func(data interface{}) {
		fmt.Println(data)
	})
}
