//echo prints arguments
//synopsis : echo [-n] [ args ... ]
package main

import (
	"flag"
	"fmt"
	"strings"
)

var n = flag.Bool("n", false, "a bool")

func main() {
	flag.Parse()
	s := strings.Join(flag.Args(), " ")
	fmt.Print(s)
	if !*n {
		fmt.Print("\n")
	}
}
