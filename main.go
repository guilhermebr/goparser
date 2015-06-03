package main

import (
	"flag"
	"fmt"
	"github.com/guilhermebr/goparser/parser"
)

func main() {
	flag.Parse()
	if len(flag.Args()) != 1 {
		fmt.Println("miss argument: url to parse")
		return
	}
	arg := flag.Arg(0)
	parser.Parse(arg)
}
