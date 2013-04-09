// Package jason parses firefox's JSON bookmark files.
package jason

import (
	"encoding/json"
	"io/ioutil"
)

type object struct {
	Title    string
	Uri      string
	Type     string
	Children []object
}

// Open opens the provided JSON bookmark file and returns the parsed bookmarks.
func Open(filePath string) (obj *object, err error) {
	buf, err := ioutil.ReadFile(filePath)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(buf, &obj)
	if err != nil {
		return nil, err
	}
	return obj, nil
}

// Bookmarks returns all bookmarks in obj.
func (obj object) Bookmarks() (bookmarks []string) {
	var traverse func(obj object)
	traverse = func(obj object) {
		// If obj.Type is "text/x-moz-place", it's a bookmark.
		if obj.Type == "text/x-moz-place" {
			bookmarks = append(bookmarks, obj.Uri)
		}
		for _, o := range obj.Children {
			traverse(o)
		}
	}
	traverse(obj)
	return bookmarks
}
