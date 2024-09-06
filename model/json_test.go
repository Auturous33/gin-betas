package model

import (
	"fmt"
	"testing"
)

// test ToJsonStr func
func TestToJsonStr(t *testing.T) {
	fmt.Println(ToJsonStr(struct {
		Name string `json:"name"`
		Age  int    `json:"age"`
	}{
		Name: "auturous",
		Age:  12,
	}))
}
