package main

import (
	"fmt"
	"go-utils/v1/stringUtils"
)

func main() {
	r := stringUtils.Contains("abcd", "abcx", "cdx", "d")
	fmt.Printf("%+v\n", r)
}
