// A program to read Firefox's JSON bookmark files.
package jason

import "encoding/json"
import "fmt"
import "io/ioutil"

// The main wrapping object
type FFBookmarks struct {
	Title        string
	Id           int
	DateAdded    int64
	LastModified int64
	Type         string
	Root         string
	Children     []folder
}

// A folder containing several bookmarks or folders
type folder struct {
	Index        int
	Title        string
	Id           int
	Parent       int
	DateAdded    int64
	LastModified int64
	Annos        []annos
	Type         string
	Root         string
	Children     []interface{}
}

// Unknown what this object stores
type annos struct {
	Name     string
	Flags    int
	Expires  int
	MimeType string
	Type     int
	Value    string
}

// A bookmark object
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

// Open reads a firefox json bookmark file into a ffb struct
func Open(fileName string) (ffb *FFBookmarks, err error) {
	buf, err := ioutil.ReadFile(fileName)
	if err != nil {
		return nil, err
	}

	ffb = new(FFBookmarks)
	json.Unmarshal(buf, &ffb)

	return ffb, nil
}

// PrintUri prints all bookmarks in ffb 
func (ffb *FFBookmarks) Print() {
	for _, fol := range ffb.Children {
		for _, folOrBm := range fol.Children {
			switch unk := folOrBm.(type) {
			case map[string]interface{}:
				traverse(unk)
			}
		}
	}
}

// If unk is a bookmark print it otherwise traverse the folders children
// in search for bookmarks.
func traverse(unk map[string]interface{}) {
	found, ok := unk["uri"]
	if !ok {
		switch children := unk["children"].(type) {
		case []interface{}:
			for _, interChild := range children {
				switch child := interChild.(type) {
				case map[string]interface{}:
					traverse(child)
				}
			}
		}
	} else {
		fmt.Println(found)
	}
}
