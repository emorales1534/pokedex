package main

import (
//	"fmt"
	"strings"
)

func cleanInput(text string) []string {
	lower := strings.ToLower(text)
	res := strings.Fields(lower)
	return res
}
