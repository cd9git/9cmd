//

package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
)

const whsiz int = 48

var wheel = [whsiz]float64{
	2, 10, 2, 4, 2, 4, 6, 2, 6, 4,
	2, 4, 6, 6, 2, 6, 4, 2, 6, 4,
	6, 8, 4, 2, 4, 2, 4, 8, 6, 4,
	6, 2, 4, 6, 2, 6, 6, 4, 2, 4,
	6, 2, 6, 4, 2, 4, 2, 10,
}

func main() {

	if len(os.Args) > 1 {
		for _, arg := range os.Args[1:] {
			n, err := strconv.ParseFloat(arg, 64)
			if err != nil {
				log.Fatal(err)
			}
			factor(n)
		}
	} else {
		reader := bufio.NewReader(os.Stdin)
		for {
			l, _ := reader.ReadString('\n')
			l = strings.TrimRight(l, "\r\n")
			if len(l) == 0 {
				break
			} else {
				n, err := strconv.ParseFloat(l, 64)
				if err != nil {
					log.Fatal(err)
				}
				if n > 0 {
					factor(n)
				}
			}
		}
	}
}

func factor(n float64) {
	var quot float64
	fmt.Printf("%.0f\n", n)
	if n == 0 {
		os.Exit(0)
	}
	s := math.Sqrt(n) + 1

	for modf(n/2, &quot) == 0 {
		fmt.Println("     2")
		n = quot
		s = math.Sqrt(n) + 1
	}
	for modf(n/3, &quot) == 0 {
		fmt.Println("     3")
		n = quot
		s = math.Sqrt(n) + 1
	}
	for modf(n/5, &quot) == 0 {
		fmt.Println("     5")
		n = quot
		s = math.Sqrt(n) + 1
	}
	for modf(n/7, &quot) == 0 {
		fmt.Println("     7")
		n = quot
		s = math.Sqrt(n) + 1
	}
	d := 1.0
	for i := 1; ; {
		d += wheel[i]
		for modf(n/d, &quot) == 0 {
			fmt.Printf("     %.0f\n", d)
			n = quot
			s = math.Sqrt(n) + 1
		}
		i++
		if i >= whsiz {
			i = 0
			if d > s {
				break
			}
		}
	}

	if n > 1 {
		fmt.Printf("     %.0f\n", n)
	}
	fmt.Printf("\n")
}

func modf(n float64, intp *float64) float64 {
	int, frac := math.Modf(n)
	*intp = int
	return frac
}
