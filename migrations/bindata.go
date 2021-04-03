package migrations

import (
	"bytes"
	"compress/gzip"
	"fmt"
	"io"
	"strings"
)

func bindata_read(data []byte, name string) ([]byte, error) {
	gz, err := gzip.NewReader(bytes.NewBuffer(data))
	if err != nil {
		return nil, fmt.Errorf("Read %q: %v", name, err)
	}

	var buf bytes.Buffer
	_, err = io.Copy(&buf, gz)
	gz.Close()

	if err != nil {
		return nil, fmt.Errorf("Read %q: %v", name, err)
	}

	return buf.Bytes(), nil
}

var __000001_init_down_cql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x4a\x29\xca\x2f\x50\xc8\xcc\x4b\x49\xad\x50\x28\xc8\x4c\x2e\x29\x8e\x07\x91\xf1\x99\x29\xf1\x60\x31\x6b\x2e\xb0\x7c\x49\x62\x52\x4e\xaa\x42\x5a\x51\x62\x6e\x6a\xb1\x1e\x58\x99\x35\x20\x00\x00\xff\xff\xa0\xcb\x90\x1a\x38\x00\x00\x00")

func _000001_init_down_cql() ([]byte, error) {
	return bindata_read(
		__000001_init_down_cql,
		"000001_init.down.cql",
	)
}

var __000001_init_up_cql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x6c\x8f\x41\x4e\xc4\x30\x0c\x45\xf7\x39\xc5\x5f\x4e\xa5\x8a\x0b\x70\x98\xca\x49\x3c\x60\xd1\x24\x95\xed\x00\x73\x7b\x34\x34\xa5\x5d\xe0\xed\xd3\xfb\x7a\x4e\xca\xe4\x0c\xa7\xb8\x32\xe4\x8e\xda\x1c\xfc\x2d\xe6\x86\xbb\x52\x61\x7b\xd9\x24\xb9\xe1\x16\x00\x20\x51\x61\xa5\x45\x32\xc6\x49\xf5\x79\x27\xbf\x3b\x79\x21\x1f\xc4\xa5\xb0\x39\x95\x6d\xe7\xcf\x95\x8b\x07\xf4\x2e\x79\xa8\xad\x3a\x57\x3f\x51\x5c\x5b\x1c\x96\xf2\xa7\xb4\x6e\xcb\xa1\x9f\x96\xb1\x99\xb4\xfa\xdf\xe0\xa6\x52\x48\x1f\xf8\xe0\x07\x6e\x7f\xc9\xf3\xa5\x71\x3e\x7a\xa6\x30\xe1\x4b\xfc\x1d\x69\xed\xe6\xac\x52\xdf\xd0\x34\xb3\x22\x3e\xdd\xf3\xa9\xcc\x96\xa6\xd7\x10\x7e\x02\x00\x00\xff\xff\xb1\x3e\x24\x44\x30\x01\x00\x00")

func _000001_init_up_cql() ([]byte, error) {
	return bindata_read(
		__000001_init_up_cql,
		"000001_init.up.cql",
	)
}

var _bindata_go = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x01\x00\x00\xff\xff\x00\x00\x00\x00\x00\x00\x00\x00")

func bindata_go() ([]byte, error) {
	return bindata_read(
		_bindata_go,
		"bindata.go",
	)
}

// Asset loads and returns the asset for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
func Asset(name string) ([]byte, error) {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[cannonicalName]; ok {
		return f()
	}
	return nil, fmt.Errorf("Asset %s not found", name)
}

// AssetNames returns the names of the assets.
func AssetNames() []string {
	names := make([]string, 0, len(_bindata))
	for name := range _bindata {
		names = append(names, name)
	}
	return names
}

// _bindata is a table, holding each asset generator, mapped to its name.
var _bindata = map[string]func() ([]byte, error){
	"000001_init.down.cql": _000001_init_down_cql,
	"000001_init.up.cql": _000001_init_up_cql,
	"bindata.go": bindata_go,
}
// AssetDir returns the file names below a certain
// directory embedded in the file by go-bindata.
// For example if you run go-bindata on data/... and data contains the
// following hierarchy:
//     data/
//       foo.txt
//       img/
//         a.png
//         b.png
// then AssetDir("data") would return []string{"foo.txt", "img"}
// AssetDir("data/img") would return []string{"a.png", "b.png"}
// AssetDir("foo.txt") and AssetDir("notexist") would return an error
// AssetDir("") will return []string{"data"}.
func AssetDir(name string) ([]string, error) {
	node := _bintree
	if len(name) != 0 {
		cannonicalName := strings.Replace(name, "\\", "/", -1)
		pathList := strings.Split(cannonicalName, "/")
		for _, p := range pathList {
			node = node.Children[p]
			if node == nil {
				return nil, fmt.Errorf("Asset %s not found", name)
			}
		}
	}
	if node.Func != nil {
		return nil, fmt.Errorf("Asset %s not found", name)
	}
	rv := make([]string, 0, len(node.Children))
	for name := range node.Children {
		rv = append(rv, name)
	}
	return rv, nil
}

type _bintree_t struct {
	Func func() ([]byte, error)
	Children map[string]*_bintree_t
}
var _bintree = &_bintree_t{nil, map[string]*_bintree_t{
	"000001_init.down.cql": &_bintree_t{_000001_init_down_cql, map[string]*_bintree_t{
	}},
	"000001_init.up.cql": &_bintree_t{_000001_init_up_cql, map[string]*_bintree_t{
	}},
	"bindata.go": &_bintree_t{bindata_go, map[string]*_bintree_t{
	}},
}}
