// Code generated by go-bindata.
// sources:
// templates/html/inventory.html
// DO NOT EDIT!

package html

import (
	"bytes"
	"compress/gzip"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"time"
)

func bindataRead(data []byte, name string) ([]byte, error) {
	gz, err := gzip.NewReader(bytes.NewBuffer(data))
	if err != nil {
		return nil, fmt.Errorf("Read %q: %v", name, err)
	}

	var buf bytes.Buffer
	_, err = io.Copy(&buf, gz)
	clErr := gz.Close()

	if err != nil {
		return nil, fmt.Errorf("Read %q: %v", name, err)
	}
	if clErr != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}

type asset struct {
	bytes []byte
	info  os.FileInfo
}

type bindataFileInfo struct {
	name    string
	size    int64
	mode    os.FileMode
	modTime time.Time
}

func (fi bindataFileInfo) Name() string {
	return fi.name
}
func (fi bindataFileInfo) Size() int64 {
	return fi.size
}
func (fi bindataFileInfo) Mode() os.FileMode {
	return fi.mode
}
func (fi bindataFileInfo) ModTime() time.Time {
	return fi.modTime
}
func (fi bindataFileInfo) IsDir() bool {
	return false
}
func (fi bindataFileInfo) Sys() interface{} {
	return nil
}

var _templatesHtmlInventoryHtml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x94\x55\xef\x6f\xdb\x36\x10\xfd\x5c\xff\x15\x37\x2d\xc0\x36\xa0\x92\xda\x0c\xcd\x16\x57\x12\x30\x64\x0d\x16\xa0\x5b\xb7\x25\xc1\xb6\x8f\x27\xe9\x6c\xdd\x46\x91\x0a\x79\xb6\xeb\x08\xfe\xdf\x07\x52\x52\x6c\x27\xe9\x7e\x7c\xb1\x44\xea\xdd\x7b\xe4\xe3\x3b\x3a\xfb\xec\xfb\x0f\x17\x37\x7f\xfc\xfc\x0e\x1a\x69\x55\x31\xcb\xfc\x03\x14\xea\x65\x1e\x91\x8e\x8a\x19\x40\xd6\x10\xd6\xc5\xec\x05\x40\x26\x2c\x8a\x8a\xdf\x1a\xf3\x85\x83\x0f\x1a\x2e\xd9\x3a\x81\xeb\x5f\xde\xb3\x10\x5c\xe9\x35\x69\x31\x76\x9b\xa5\x03\x2c\x54\xb4\x24\x08\x8d\x48\x17\xd3\xdd\x8a\xd7\x79\x74\x61\xb4\x90\x96\xf8\x66\xdb\x51\x04\xd5\x30\xca\x23\xa1\x8f\x92\x7a\xed\xb7\x50\x35\x68\x1d\x49\x7e\x7b\x73\x19\x7f\x1b\x1d\xd0\x68\x6c\x29\x8f\x2c\x2d\xc8\x5a\xb2\x07\xc5\xc6\xf2\x92\x75\xf4\x09\xc5\xdf\xe3\xdb\xef\xe2\x0b\xd3\x76\x28\x5c\xaa\x43\xd1\xab\x77\xf9\x79\x04\xe9\x13\x09\xec\x3a\x45\x71\x6b\x4a\x56\x14\x6f\xa8\x8c\xb1\xeb\xe2\x0a\x3b\x3c\x2e\xdf\x92\xfb\xcf\xd5\x4e\x50\x56\x2e\x2e\xd1\xc6\x4e\xb6\x47\x34\xa5\xc2\xea\xaf\xe7\x88\x7e\x40\x5d\x37\xa4\xea\x4b\xcb\xa4\x6b\xb5\x3d\xb4\xcb\xae\xe8\xb9\x92\x35\xd3\xa6\x33\x56\x0e\xa0\x1b\xae\xa5\xc9\x6b\x5a\x73\x45\x71\x18\xbc\x04\xd6\x2c\x8c\x2a\x76\x15\x2a\xca\x5f\xbf\x84\x16\x3f\x72\xbb\x6a\x0f\x26\x58\x1f\x4f\xac\x1c\xd9\x30\xf2\x26\xe4\xda\x3c\xa8\x87\xed\x80\x6c\x3b\x1a\x4f\xb1\x72\x6e\x38\x0a\x80\xd2\xd4\x5b\xe8\x87\xf7\x85\xd1\x12\x2f\xb0\x65\xb5\x9d\x83\x43\xed\x62\x47\x96\x17\x6f\x87\xaf\x8a\x35\xc5\x0d\xf1\xb2\x91\x39\xbc\x4e\xbe\xa6\x76\xfc\xb0\x9b\x4d\x80\x89\xa8\x45\xbb\x64\x1d\x97\x46\xc4\xb4\x1e\xfc\xe6\x00\x1c\x1e\xc3\x6f\x65\x6a\x9a\x6a\x5e\x04\xf5\xcd\xc8\xff\xcd\xab\x57\x63\xc1\x8b\xca\x28\x63\xe7\xf0\xf9\xf9\xf9\xf9\x23\xc1\xa4\x26\x41\x56\xee\x91\xac\x98\x6e\x0e\x47\x92\x23\xdc\x35\x78\xfa\xe6\x6c\x42\x4f\xbc\x67\x67\x67\x87\x6b\xcb\xd2\x60\x57\x68\xab\x74\xe8\x2b\x80\xcc\xdb\x54\xcc\x66\x1e\x95\x35\xa7\xff\xde\x5f\xcd\xe9\x84\xee\x8a\x9b\x86\x1d\x2c\x58\x11\x6c\xd0\x41\x86\xd0\x58\x5a\xe4\x91\x6f\x00\x37\x4f\xd3\x25\x4b\xb3\x2a\x93\xca\xb4\xe9\xa6\x31\xce\xe8\x85\xe7\x4c\x97\x26\x3e\x18\xc6\xee\x4e\xb1\x50\x5a\x2a\x53\xa6\x2d\x3a\x21\x9b\x56\x6d\x9d\x6e\xcc\x62\xfc\x14\xf3\xa4\x9e\x2c\x4d\x54\x2c\x49\x93\x45\xa1\x1a\xca\x2d\x58\x53\x1a\x71\x59\x8a\x05\x18\x0d\x59\xb0\xbd\x52\xe8\x5c\x1e\xd5\x28\x14\x15\x7d\x0f\xc9\x7b\x74\xf2\xa3\xa9\x79\xc1\x54\xc3\x6e\x97\xa5\x1e\x56\x24\x70\xd3\x90\x25\x60\x07\xa8\x9c\x01\x84\x16\xab\xc6\x47\xc1\x12\xd6\x3e\x69\xb0\x26\xeb\xd8\x68\xc0\x35\x72\xc8\x1e\xa0\xec\xb7\xb9\x5f\xd6\x9f\xce\xe8\xa8\x38\x1e\xfb\x35\x25\x59\xda\x4d\x6e\xad\x54\x11\x5e\xfa\x1e\x2c\xea\x25\xc1\x09\xbf\x84\x13\x0b\xf3\x1c\x92\x5f\xc9\xf7\x8d\x83\xdd\x6e\x04\x2b\x1e\xc0\x21\xe2\xd6\xe8\x65\xf1\xa0\xda\xf7\x70\x62\x93\x9f\xb0\x25\xd8\xed\x92\xf2\xfe\x34\xec\x71\x3f\xe5\x75\xfd\x49\x87\x2a\xc8\x5c\x8b\x4a\x4d\x96\x0c\x31\x99\x0a\xae\xc3\xc8\x5f\x4e\x96\x9c\x1b\xad\x09\xf8\x22\x2b\xad\x6f\xb1\x71\x09\x35\xaf\x1f\x4c\x1d\x72\xe9\x7b\xac\xef\x81\x17\x40\x77\x9e\xea\xc2\xac\xb4\xc0\x6b\xd8\xed\x42\x24\x6a\x14\x2c\xd1\x51\xb8\x08\x90\xb5\x3b\x3e\x99\xca\xa3\xa3\xc2\x68\x1a\x8f\x02\x1e\xa5\xce\x52\x65\x6c\xdd\xf7\x40\xca\xd1\xff\x23\x1d\xb6\x16\xd6\x73\x2d\x96\xf5\x72\x7f\xe0\xcf\xab\x38\x2f\xa3\xfd\xe6\x01\x75\xed\xd3\x70\x44\xeb\xf8\x9e\x1e\x0c\xe3\x7b\xda\xdb\xf5\x84\xbe\xda\x3b\xe9\x99\xfe\x91\xe6\x49\xf1\x4a\xef\xcb\x93\x23\xfb\xaf\x24\x74\x57\x65\x29\x84\xfe\xf9\x8c\x9f\xd8\xa3\x94\x3f\xa1\xf7\x0b\xfa\x52\x1a\x0a\x2e\x02\x8b\x23\xb5\xf8\x2a\xf0\x2a\x74\x02\xab\xae\x0e\xe4\x9f\xec\xa1\x91\xff\x36\xe0\x1e\xb3\x27\x53\x50\xd2\x9a\xd7\xc3\xb2\xb3\x74\x8a\xf0\x83\xbd\xe3\xbc\xef\x83\x70\x03\x0d\x17\x4f\x96\x0e\xff\xfd\x7f\x07\x00\x00\xff\xff\x1e\x22\x0f\xa9\x0c\x08\x00\x00")

func templatesHtmlInventoryHtmlBytes() ([]byte, error) {
	return bindataRead(
		_templatesHtmlInventoryHtml,
		"templates/html/inventory.html",
	)
}

func templatesHtmlInventoryHtml() (*asset, error) {
	bytes, err := templatesHtmlInventoryHtmlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "templates/html/inventory.html", size: 2060, mode: os.FileMode(420), modTime: time.Unix(1513817283, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

// Asset loads and returns the asset for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
func Asset(name string) ([]byte, error) {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[cannonicalName]; ok {
		a, err := f()
		if err != nil {
			return nil, fmt.Errorf("Asset %s can't read by error: %v", name, err)
		}
		return a.bytes, nil
	}
	return nil, fmt.Errorf("Asset %s not found", name)
}

// MustAsset is like Asset but panics when Asset would return an error.
// It simplifies safe initialization of global variables.
func MustAsset(name string) []byte {
	a, err := Asset(name)
	if err != nil {
		panic("asset: Asset(" + name + "): " + err.Error())
	}

	return a
}

// AssetInfo loads and returns the asset info for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
func AssetInfo(name string) (os.FileInfo, error) {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[cannonicalName]; ok {
		a, err := f()
		if err != nil {
			return nil, fmt.Errorf("AssetInfo %s can't read by error: %v", name, err)
		}
		return a.info, nil
	}
	return nil, fmt.Errorf("AssetInfo %s not found", name)
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
var _bindata = map[string]func() (*asset, error){
	"templates/html/inventory.html": templatesHtmlInventoryHtml,
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
	for childName := range node.Children {
		rv = append(rv, childName)
	}
	return rv, nil
}

type bintree struct {
	Func     func() (*asset, error)
	Children map[string]*bintree
}
var _bintree = &bintree{nil, map[string]*bintree{
	"templates": &bintree{nil, map[string]*bintree{
		"html": &bintree{nil, map[string]*bintree{
			"inventory.html": &bintree{templatesHtmlInventoryHtml, map[string]*bintree{}},
		}},
	}},
}}

// RestoreAsset restores an asset under the given directory
func RestoreAsset(dir, name string) error {
	data, err := Asset(name)
	if err != nil {
		return err
	}
	info, err := AssetInfo(name)
	if err != nil {
		return err
	}
	err = os.MkdirAll(_filePath(dir, filepath.Dir(name)), os.FileMode(0755))
	if err != nil {
		return err
	}
	err = ioutil.WriteFile(_filePath(dir, name), data, info.Mode())
	if err != nil {
		return err
	}
	err = os.Chtimes(_filePath(dir, name), info.ModTime(), info.ModTime())
	if err != nil {
		return err
	}
	return nil
}

// RestoreAssets restores an asset under the given directory recursively
func RestoreAssets(dir, name string) error {
	children, err := AssetDir(name)
	// File
	if err != nil {
		return RestoreAsset(dir, name)
	}
	// Dir
	for _, child := range children {
		err = RestoreAssets(dir, filepath.Join(name, child))
		if err != nil {
			return err
		}
	}
	return nil
}

func _filePath(dir, name string) string {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	return filepath.Join(append([]string{dir}, strings.Split(cannonicalName, "/")...)...)
}
