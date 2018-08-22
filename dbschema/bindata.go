// Code generated by go-bindata.
// sources:
// migrations/001_base.sql
// migrations/002_add_buzz_type.sql
// migrations/003_add_created_at_updated_at.sql
// migrations/test/test.sql
// DO NOT EDIT!

package dbschema

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

var _migrations001_baseSql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x6c\xcf\xcd\x4a\x87\x40\x14\x05\xf0\xfd\x7d\x8a\xb3\x53\x29\x37\x81\xab\x56\xa3\x5e\x69\x6a\xd2\x98\x19\x43\x57\x61\x39\xc5\x40\x7e\xa0\x46\x60\xf4\xee\x91\x51\xb4\xf8\xdf\xed\xfd\x71\x0e\x27\x8e\x71\x36\xf8\x97\xa5\xdb\x1c\xea\x99\x32\xcd\xc2\x32\xac\x48\x15\x43\x16\x28\x2b\x0b\x6e\xa4\xb1\x06\x8f\x6f\xfb\xfe\xf0\xec\x5c\xbf\x22\x24\xc0\xf7\xf8\x3b\xc3\x5a\x0a\x85\x3b\x2d\x6f\x85\x6e\x71\xc3\xed\x39\x01\x63\x37\xb8\x5f\x71\x2f\x74\x76\x25\x74\x78\x91\x24\xd1\xf7\xaf\x77\xeb\xd3\xe2\xe7\xcd\x4f\x23\x2c\x37\x16\x47\x53\x59\x2b\x85\x9c\x0b\x51\x2b\x8b\x20\xf8\x81\x5b\xe7\x5f\xd7\x23\xe4\xda\x54\x65\x7a\x02\x7e\x7c\x06\x14\x5d\x12\xfd\xdf\x92\x4f\xef\x23\xd1\x57\x00\x00\x00\xff\xff\x36\x00\x94\xb4\xde\x00\x00\x00")

func migrations001_baseSqlBytes() ([]byte, error) {
	return bindataRead(
		_migrations001_baseSql,
		"migrations/001_base.sql",
	)
}

func migrations001_baseSql() (*asset, error) {
	bytes, err := migrations001_baseSqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "migrations/001_base.sql", size: 222, mode: os.FileMode(420), modTime: time.Unix(1534940482, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _migrations002_add_buzz_typeSql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xd2\xd5\x55\xd0\xce\xcd\x4c\x2f\x4a\x2c\x49\x55\x08\x2d\xe0\x72\x0e\x72\x75\x0c\x71\x55\x08\x89\x0c\x70\x55\x28\x49\x2d\x2e\x89\x2f\xa9\x2c\x48\x55\x70\x0c\x56\x70\xf5\x0b\xf5\x55\xd0\x50\x07\x89\x39\xaa\xeb\x28\x80\x19\x4e\x30\x86\x33\x8c\xe1\xa2\xae\x69\xcd\xe5\xe8\x13\xe2\x1a\xa4\x10\xe2\xe8\xe4\xe3\xaa\x90\x54\x5a\x55\x15\x9f\x96\x9a\x9a\x52\xac\xe0\xe8\xe2\xa2\xe0\xec\xef\x13\xea\xeb\x07\x11\x05\x9b\x0c\xb7\xc3\x9a\x8b\x0b\xd9\x29\x2e\xf9\xe5\x79\x5c\x2e\x41\xfe\x01\x50\x73\x3c\xdd\x14\x5c\x23\x3c\x83\x43\x82\x91\x4c\xb4\x86\x2a\x40\x71\xab\x35\x20\x00\x00\xff\xff\x33\x9d\xd0\xdc\xd1\x00\x00\x00")

func migrations002_add_buzz_typeSqlBytes() ([]byte, error) {
	return bindataRead(
		_migrations002_add_buzz_typeSql,
		"migrations/002_add_buzz_type.sql",
	)
}

func migrations002_add_buzz_typeSql() (*asset, error) {
	bytes, err := migrations002_add_buzz_typeSqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "migrations/002_add_buzz_type.sql", size: 209, mode: os.FileMode(420), modTime: time.Unix(1534940482, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _migrations003_add_created_at_updated_atSql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xd2\xd5\x55\xd0\xce\xcd\x4c\x2f\x4a\x2c\x49\x55\x08\x2d\xe0\x72\xf4\x09\x71\x0d\x52\x08\x71\x74\xf2\x71\x55\x48\x2a\xad\xaa\x8a\x4f\x4b\x4d\x4d\x29\x56\x70\x74\x71\x51\x70\xf6\xf7\x09\xf5\xf5\x53\x48\x2e\x4a\x4d\x2c\x49\x4d\x89\x4f\x2c\x51\xc8\xcc\x2b\x49\x4d\x4f\x2d\xb2\x26\x42\x57\x69\x41\x0a\x86\x2e\x2e\x64\xbb\x5d\xf2\xcb\xf3\x00\x01\x00\x00\xff\xff\x03\x17\x35\x92\x8c\x00\x00\x00")

func migrations003_add_created_at_updated_atSqlBytes() ([]byte, error) {
	return bindataRead(
		_migrations003_add_created_at_updated_atSql,
		"migrations/003_add_created_at_updated_at.sql",
	)
}

func migrations003_add_created_at_updated_atSql() (*asset, error) {
	bytes, err := migrations003_add_created_at_updated_atSqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "migrations/003_add_created_at_updated_at.sql", size: 140, mode: os.FileMode(420), modTime: time.Unix(1534928685, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _migrationsTestTestSql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xb4\xd6\xbd\x0a\xc2\x30\x14\x47\xf1\xdd\xa7\xc8\x56\x85\x0c\xfe\xef\xf5\x13\x27\x9f\x44\xac\x46\xa8\x48\x0b\x26\x75\xa8\xf8\xee\x62\xc1\x59\x10\xce\x92\x90\x0c\x67\xfb\x71\x6f\xd3\xe6\x74\x2f\xa1\x69\x4b\x17\xea\x7e\x18\x0e\x97\x94\xce\x39\x3c\x8e\xb7\x3e\xe5\x30\x55\xac\x4a\xca\xa5\x8a\x61\xbc\x43\x57\x5f\xd3\x69\x7c\x3e\x5f\xdf\xcf\x7d\x15\xe5\x32\xff\x9c\xae\xd9\x6e\xf2\xa3\x69\x40\xd3\x81\xe6\x02\x68\x2e\x81\xe6\x0a\x68\xae\x81\xe6\x06\x68\x6e\x81\xa6\xe6\x44\x94\x90\x24\x82\x92\x08\x4b\x22\x30\x89\xd0\x24\x82\x93\x08\x4f\x22\x40\x89\x10\x65\x84\x28\x43\x66\x13\x21\xca\x08\x51\x46\x88\x32\x42\x94\x11\xa2\x8c\x10\x65\x84\x28\x23\x44\x39\x21\xca\x09\x51\x8e\xac\x7b\xff\x88\x7a\x07\x00\x00\xff\xff\x6b\x16\xe8\x10\xeb\x0a\x00\x00")

func migrationsTestTestSqlBytes() ([]byte, error) {
	return bindataRead(
		_migrationsTestTestSql,
		"migrations/test/test.sql",
	)
}

func migrationsTestTestSql() (*asset, error) {
	bytes, err := migrationsTestTestSqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "migrations/test/test.sql", size: 2795, mode: os.FileMode(420), modTime: time.Unix(1534940010, 0)}
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
	"migrations/001_base.sql":                      migrations001_baseSql,
	"migrations/002_add_buzz_type.sql":             migrations002_add_buzz_typeSql,
	"migrations/003_add_created_at_updated_at.sql": migrations003_add_created_at_updated_atSql,
	"migrations/test/test.sql":                     migrationsTestTestSql,
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
	"migrations": &bintree{nil, map[string]*bintree{
		"001_base.sql":                      &bintree{migrations001_baseSql, map[string]*bintree{}},
		"002_add_buzz_type.sql":             &bintree{migrations002_add_buzz_typeSql, map[string]*bintree{}},
		"003_add_created_at_updated_at.sql": &bintree{migrations003_add_created_at_updated_atSql, map[string]*bintree{}},
		"test": &bintree{nil, map[string]*bintree{
			"test.sql": &bintree{migrationsTestTestSql, map[string]*bintree{}},
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
