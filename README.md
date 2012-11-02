Jason
-----
This program will read a firefox JSON bookmark file and print the urls.

Syntax - Print
--------------

	ffb, err := jason.Open("fileName.json")
	if err != nil {
		// Handle error
	}
	ffb.Print()

Installation - Library
------------
Use `go get github.com/karlek/jason`
   
   go get github.com/karlek/jason

Installation - Program
------------
Use `go install github.com/karlek/jason/cmd/jason`
   
   go install github.com/karlek/jason/cmd/jason


API documentation
-----------------
http://go.pkgdoc.org/github.com/karlek/jason

Public domain
-------------
I hereby release this code into the [public domain](https://creativecommons.org/publicdomain/zero/1.0/).
