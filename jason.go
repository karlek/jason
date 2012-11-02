//A program to read Firefox's JSON bookmark files.
package jason

import "encoding/json"
import "fmt"
import "io/ioutil"

//The main wrapping object
type ffBookmarks struct {
	Title        string
	Id           int
	DateAdded    int64
	LastModified int64
	Type         string
	Root         string
	Children     []folder
}

//A folder containing several bookmarks
type folder struct {
	Title        string
	Id           int
	Parent       int
	DateAdded    int64
	LastModified int64
	Annos        []annos
	Type         string
	Root         string
	Children     []bookmark
}

//Unknown what this object stores
type annos struct {
	Name     string
	Flags    int
	Expires  int
	MimeType string
	Type     int
	Value    string
}

//A bookmark object
type bookmark struct {
	Title        string
	Id           int
	Parent       int
	DateAdded    int64
	LastModified int64
	Annos        []annos
	Type         string
	Uri          string
	Keyword      string
}

//Open reads a firefox json bookmark file into a ffb struct
func Open(fileName string) (ffb *ffBookmarks, err error) {
	buf, err := ioutil.ReadFile(fileName)
	if err != nil {
		return nil, err
	}

	ffb = new(ffBookmarks)
	json.Unmarshal(buf, &ffb)

	return ffb, nil
}

//PrintUri prints all bookmarks in ffb 
func (ffb ffBookmarks) Print() {
	for _, fol := range ffb.Children {
		for _, bm := range fol.Children {
			if bm.Uri != "" {
				fmt.Println(bm.Uri)
			}
		}
	}
}
