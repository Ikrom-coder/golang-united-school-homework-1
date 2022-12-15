package main

import (
	"fmt"
	"testing"
)

func TestAdd(t *testing.T) {
	res := "Hello 🗺️ !"
	result := Hello()
	fmt.Println(result)
	if result != res {
		t.Errorf("Hello function is not PASSED, Expected %s, got %s\n", res, result)
	} else {
		t.Logf("Hello function is PASSED, Expected %s, got %s\n", res, result)
	}
}
