// Jason prints the urls of a firefox bookmarks file(s).
package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/karlek/jason"
)

func init() {
	flag.Usage = usage
}

func usage() {
	fmt.Fprintf(os.Stderr, "Usage: %s [FILE],,,\n", os.Args[0])
	flag.PrintDefaults()
	os.Exit(1)
}

func main() {
	flag.Parse()
	if flag.NArg() < 1 {
		flag.Usage()
	}
	for _, filePath := range flag.Args() {
		err := readJason(filePath)
		if err != nil {
			log.Fatalln(err)
		}
	}
}

func readJason(filePath string) (err error) {
	obj, err := jason.Open(filePath)
	if err != nil {
		return err
	}

	for _, bookmark := range obj.Bookmarks() {
		fmt.Println(bookmark)
	}
	return nil
}
