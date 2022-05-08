// date prints current date and time
// synopsis: date [-un] [seconds]
package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"strconv"
	"time"
)

func main() {
	fs := flag.NewFlagSet("date", flag.ContinueOnError)
	nflg := fs.Bool("n", false, "a bool")
	uflg := fs.Bool("u", false, "a bool")
	if err := fs.Parse(os.Args[1:]); err != nil {
		log.Fatal(err)
	}
	now := time.Now()

	//default:	fprint(2, "usage: date [-un] [seconds]\n"); exits("usage");
	if len(fs.Args()) == 1 {
		n, err := strconv.ParseInt(fs.Args()[0], 10, 64)
		if err != nil {
			log.Fatal(err)
		} else {
			now = time.Unix(n, 0)
		}
	} else {
		now = time.Now()
	}

	if *nflg {
		fmt.Println(now.Unix())
	} else if *uflg {
		loc, _ := time.LoadLocation("UTC")
		fmt.Println(now.In(loc))
	} else {
		fmt.Println(now)
	}
}
