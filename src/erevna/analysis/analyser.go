package analysis

import (
	"strings"
	"text/scanner"
)

type Analyser struct {
	
}

func ParseText(data string) scanner.Scanner {
	
	var src = strings.NewReader(data)
	var s scanner.Scanner
	
	s.Init(src)
	
	return s
} 