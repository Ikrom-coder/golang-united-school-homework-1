package main

import (
	"fmt"
	"testing"
)

func TestAdd(t *testing.T) {
	result := Hello()
	fmt.Println(result)
	if result != "Hello 🗺️ !" {
		t.Errorf("Hello function is PASSED, Expected %s, got %s\n", "Hello 🗺️ !", result)
	} else {
		t.Logf("Hello function is PASSED, Expected %s, got %s\n", "Hello 🗺️ !", result)
	}
}
