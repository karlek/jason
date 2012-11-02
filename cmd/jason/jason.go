//Example program for using jason
package main

import "flag"
import "fmt"
import "log"
import "os"

import "github.com/karlek/jason"

func init() {
	flag.Usage = usage
	flag.Parse()

	if flag.NArg() < 1 {
		usage()
	}
}

func usage() {
	fmt.Fprintf(os.Stderr, "Usage: %s [File(s)]\n", os.Args[0])
	flag.PrintDefaults()
	os.Exit(-1)
}

func main() {
	for _, filePath := range flag.Args() {
		err := readJason(filePath)
		if err != nil {
			log.Fatalln(err)
		}
	}

}

func readJason(filePath string) (err error) {
	ffb, err := jason.Open(filePath)
	if err != nil {
		return err
	}

	ffb.Print()

	return nil
}
