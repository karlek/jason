// A program to read Firefox's JSON bookmark files.
package jason

import "encoding/json"
import "fmt"
import "io/ioutil"

// The main wrapping object
type FFBookmarks struct {
	_        string   `json: "title"`
	_        int      `json: "id"`
	_        int64    `json: "dateAdded"`
	_        int64    `json: "lastModified"`
	_        string   `json: "type"`
	_        string   `json: "root"`
	Children []folder `json: "children"`
}

// A folder containing several bookmarks or folders
type folder struct {
	_        int           `json: "index"`
	_        string        `json: "title"`
	_        int           `json: "id"`
	_        int           `json: "parent"`
	_        int64         `json: "dateAdded"`
	_        int64         `json: "lastModified"`
	_        []annos       `json: "annos"`
	_        string        `json: "type"`
	_        string        `json: "root"`
	Children []interface{} `json: "children"`
}

// Unknown what this object stores
type annos struct {
	_ string  `json: "name"`
	_ int     `json: "flags"`
	_ int     `json: "expires"`
	_ *string `json: "mimeType"`
	_ int     `json: "type"`
	_ string  `json: "value"`
}

// A bookmark object
type bookmark struct {
	_   string  `json: "title"`
	_   int     `json: "id"`
	_   int     `json: "parent"`
	_   int64   `json: "dateAdded"`
	_   int64   `json: "lastModified"`
	_   []annos `json: "annos"`
	_   string  `json: "type"`
	Uri string  `json: "uri"`
	_   string  `json: "keyword"`
}

// Open reads a firefox json bookmark file into a ffb struct
func Open(fileName string) (ffb *FFBookmarks, err error) {
	buf, err := ioutil.ReadFile(fileName)
	if err != nil {
		return nil, err
	}

	ffb = new(FFBookmarks)
	err = json.Unmarshal(buf, &ffb)
	if err != nil {
		return nil, err
	}

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
