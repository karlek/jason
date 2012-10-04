//A program to read Firefox's JSON bookmark files.
package main

import "fmt"
import "log"
import "encoding/json"
import "io/ioutil"

type FFBookmarks struct {
	Title        string
	Id           int
	DateAdded    int64
	LastModified int64
	Type         string
	Root         string
	Children     []Folder
}

type Folder struct {
	Title        string
	Id           int
	Parent       int
	DateAdded    int64
	LastModified int64
	Annos        []Annos
	Type         string
	Root         string
	Children     []Bookmark
}

type Annos struct {
	Name     string
	Flags    int
	Expires  int
	MimeType string
	Type     int
	Value    string
}

type Bookmark struct {
	Title        string
	Id           int
	Parent       int
	DateAdded    int64
	LastModified int64
	Annos        []Annos
	Type         string
	Uri          string
	Keyword      string
}

//Error wrapper
func main() {
	err := readFFB()
	if err != nil {
		log.Fatalln(err)
	}
}

func readFFB() (err error) {
	buf, err := ioutil.ReadFile("a.json")
	if err != nil {
		return err
	}

	var ffb FFBookmarks
	json.Unmarshal(buf, &ffb)

	ffb.PrintUri()

	return nil
}

//PrintUri prints all bookmarks in ffb 
func (ffb FFBookmarks) PrintUri() {
	for _, child := range ffb.Children {
		for _, nested := range child.Children {
			if nested.Uri != "" {
				fmt.Println(nested.Uri)
			}
		}
	}
}
