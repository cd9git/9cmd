// cat catenates files
// synopsis : cat [file ...]
package main

import (
	"io"
	"log"
	"os"
)

func main() {
	if len(os.Args) == 1 {
		_, err := io.Copy(os.Stdout, os.Stdin)
		if err != nil {
			log.Fatal(err)
		}
	} else {
		for _, fname := range os.Args[1:] {
			f, err := os.Open(fname)
			if err != nil {
				log.Fatal(err)
			}
			_, err = io.Copy(os.Stdout, f)
			if err != nil {
				log.Fatal(err)
			}
		}
	}
}
