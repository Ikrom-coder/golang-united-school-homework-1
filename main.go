package main

import (
	"fmt"
	"github.com/kyokomi/emoji/v2"
)

func Hello() string {
	return emoji.Sprint("Hello :world_map:!")
}

func main() {
	fmt.Println(Hello())
}
