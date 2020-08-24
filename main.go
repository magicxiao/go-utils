package main

import (
	"fmt"

	"github.com/magicxiao/go-utils/stringUtils"
)

func main() {
	r := stringUtils.Contains("abcd", "abcx", "cdx", "d")
	fmt.Printf("%+v\n", r)
}
