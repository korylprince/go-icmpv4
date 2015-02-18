package icmpv4

import (
	"bytes"
	"compress/gzip"
	"fmt"
	"io"
	"strings"
	"os"
	"time"
	"io/ioutil"
	"path"
	"path/filepath"
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

type asset struct {
	bytes []byte
	info  os.FileInfo
}

type bindata_file_info struct {
	name string
	size int64
	mode os.FileMode
	modTime time.Time
}

func (fi bindata_file_info) Name() string {
	return fi.name
}
func (fi bindata_file_info) Size() int64 {
	return fi.size
}
func (fi bindata_file_info) Mode() os.FileMode {
	return fi.mode
}
func (fi bindata_file_info) ModTime() time.Time {
	return fi.modTime
}
func (fi bindata_file_info) IsDir() bool {
	return false
}
func (fi bindata_file_info) Sys() interface{} {
	return nil
}

var _test_data_0_dat = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\xe2\x60\xa8\x74\x4b\x0f\x61\x60\xfc\xb1\xfe\x71\x08\x03\x10\x94\xf6\xb0\x83\x28\x06\x01\x41\x21\x61\x11\x51\x31\x71\x09\x49\x29\x69\x19\x59\x39\x79\x05\x45\x25\x65\x15\x55\x35\x75\x0d\x4d\x2d\x6d\x1d\x5d\x3d\x7d\x03\x43\x23\x63\x13\x53\x33\x73\x40\x00\x00\x00\xff\xff\x03\x72\xe7\x3e\x40\x00\x00\x00")

func test_data_0_dat_bytes() ([]byte, error) {
	return bindata_read(
		_test_data_0_dat,
		"test_data/0.dat",
	)
}

func test_data_0_dat() (*asset, error) {
	bytes, err := test_data_0_dat_bytes()
	if err != nil {
		return nil, err
	}

	info := bindata_file_info{name: "test_data/0.dat", size: 64, mode: os.FileMode(420), modTime: time.Unix(1424287638, 0)}
	a := &asset{bytes: bytes, info:  info}
	return a, nil
}

var _test_data_1_dat = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x62\x60\xd0\x5a\x74\x35\x96\x81\x6c\x00\x08\x00\x00\xff\xff\xfe\x6c\x09\xae\x40\x00\x00\x00")

func test_data_1_dat_bytes() ([]byte, error) {
	return bindata_read(
		_test_data_1_dat,
		"test_data/1.dat",
	)
}

func test_data_1_dat() (*asset, error) {
	bytes, err := test_data_1_dat_bytes()
	if err != nil {
		return nil, err
	}

	info := bindata_file_info{name: "test_data/1.dat", size: 64, mode: os.FileMode(420), modTime: time.Unix(1424287638, 0)}
	a := &asset{bytes: bytes, info:  info}
	return a, nil
}

var _test_data_2_dat = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x62\x66\xfe\x67\xc3\x00\x04\xae\x0c\x0c\x09\x0c\x0c\x0e\x0c\x96\x82\x0e\x0b\x73\x19\x9c\xb8\xe2\x3b\x3e\x25\x33\x98\xd6\xe7\x33\xf8\x3c\xde\x1e\xb6\xbb\xb1\x81\x81\x91\x81\x09\xa4\x8e\x25\xb3\xa0\xcc\x84\x33\x33\x39\x31\x2f\x23\xb1\x2a\xb3\x80\x39\x39\x3f\x97\x01\x28\xc7\x78\x80\x07\x44\x02\x99\x3a\x0c\x2c\x37\x5c\xff\xa4\xa2\xf1\x53\x00\x01\x00\x00\xff\xff\x38\x81\x8a\x4d\x68\x00\x00\x00")

func test_data_2_dat_bytes() ([]byte, error) {
	return bindata_read(
		_test_data_2_dat,
		"test_data/2.dat",
	)
}

func test_data_2_dat() (*asset, error) {
	bytes, err := test_data_2_dat_bytes()
	if err != nil {
		return nil, err
	}

	info := bindata_file_info{name: "test_data/2.dat", size: 104, mode: os.FileMode(420), modTime: time.Unix(1424287638, 0)}
	a := &asset{bytes: bytes, info:  info}
	return a, nil
}

var _test_data_3_dat = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\xe2\x60\xf8\xb8\xec\x81\x35\x03\x43\xe0\xf2\x6b\x1d\x0c\x4c\x42\x2f\x38\x38\xb9\xb8\x79\x78\xf9\xf8\x05\x04\x85\x84\x45\x44\xc5\xc4\x25\x24\xa5\xa4\x65\x64\xe5\xe4\x15\x14\x95\x94\x55\x54\xd5\xd4\x35\x34\xb5\xb4\x75\x74\xf5\xf4\x0d\x0c\x8d\x8c\x4d\x4c\xcd\xcc\x01\x01\x00\x00\xff\xff\x17\xe5\x0f\xd6\x40\x00\x00\x00")

func test_data_3_dat_bytes() ([]byte, error) {
	return bindata_read(
		_test_data_3_dat,
		"test_data/3.dat",
	)
}

func test_data_3_dat() (*asset, error) {
	bytes, err := test_data_3_dat_bytes()
	if err != nil {
		return nil, err
	}

	info := bindata_file_info{name: "test_data/3.dat", size: 64, mode: os.FileMode(420), modTime: time.Unix(1424287638, 0)}
	a := &asset{bytes: bytes, info:  info}
	return a, nil
}

var _test_data_4_dat = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x62\x60\x30\x58\x78\xdb\x9a\x81\x31\x70\xf9\xb5\x46\x06\xe6\x07\x9f\x38\x38\xb9\xb8\x79\x78\xf9\xf8\x05\x04\x85\x84\x45\x44\xc5\xc4\x25\x24\xa5\xa4\x65\x64\xe5\xe4\x15\x14\x95\x94\x55\x54\xd5\xd4\x35\x34\xb5\xb4\x75\x74\xf5\xf4\x0d\x0c\x8d\x8c\x4d\x4c\xcd\xcc\x01\x01\x00\x00\xff\xff\xf4\xeb\xbb\xd5\x40\x00\x00\x00")

func test_data_4_dat_bytes() ([]byte, error) {
	return bindata_read(
		_test_data_4_dat,
		"test_data/4.dat",
	)
}

func test_data_4_dat() (*asset, error) {
	bytes, err := test_data_4_dat_bytes()
	if err != nil {
		return nil, err
	}

	info := bindata_file_info{name: "test_data/4.dat", size: 64, mode: os.FileMode(420), modTime: time.Unix(1424287638, 0)}
	a := &asset{bytes: bytes, info:  info}
	return a, nil
}

var _test_data_5_dat = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\xe2\x66\x30\x50\x62\x00\x02\x57\x06\xd6\x3b\x0c\x0c\x0e\x0c\x8c\x82\x76\x27\x0f\xac\x48\x99\xc0\x01\x04\xad\xa9\x6b\xe7\xb0\x9e\x38\x5c\xcc\xc8\x00\x01\x47\xb6\x3f\x0e\x01\xd1\x42\x93\x79\x19\x46\xc1\x88\x03\x80\x00\x00\x00\xff\xff\xe8\xa3\x5c\x25\x2b\x02\x00\x00")

func test_data_5_dat_bytes() ([]byte, error) {
	return bindata_read(
		_test_data_5_dat,
		"test_data/5.dat",
	)
}

func test_data_5_dat() (*asset, error) {
	bytes, err := test_data_5_dat_bytes()
	if err != nil {
		return nil, err
	}

	info := bindata_file_info{name: "test_data/5.dat", size: 555, mode: os.FileMode(420), modTime: time.Unix(1424287638, 0)}
	a := &asset{bytes: bytes, info:  info}
	return a, nil
}

var _test_data_6_dat = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x62\x66\xfc\xf3\x8f\x01\x08\x5c\x19\x18\x42\x1e\xa4\x39\x30\xd8\x33\xaa\xf3\x1c\x58\x91\x32\x81\xeb\x24\x8b\x29\x07\x43\xb0\x78\xde\x35\x06\x36\xf1\x1d\x8f\x43\x40\x6a\x0a\x75\xb8\x41\x14\x83\x80\xa0\x90\xb0\x88\xa8\x98\xb8\x84\xa4\x94\xb4\x8c\xac\x9c\xbc\x82\xa2\x92\xb2\x8a\xaa\x9a\xba\x86\xa6\x96\xb6\x8e\xae\x9e\xbe\x81\xa1\x91\xb1\x89\xa9\x99\x39\x20\x00\x00\xff\xff\xd9\x81\x7e\x5a\x5c\x00\x00\x00")

func test_data_6_dat_bytes() ([]byte, error) {
	return bindata_read(
		_test_data_6_dat,
		"test_data/6.dat",
	)
}

func test_data_6_dat() (*asset, error) {
	bytes, err := test_data_6_dat_bytes()
	if err != nil {
		return nil, err
	}

	info := bindata_file_info{name: "test_data/6.dat", size: 92, mode: os.FileMode(420), modTime: time.Unix(1424287638, 0)}
	a := &asset{bytes: bytes, info:  info}
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
	"test_data/0.dat": test_data_0_dat,
	"test_data/1.dat": test_data_1_dat,
	"test_data/2.dat": test_data_2_dat,
	"test_data/3.dat": test_data_3_dat,
	"test_data/4.dat": test_data_4_dat,
	"test_data/5.dat": test_data_5_dat,
	"test_data/6.dat": test_data_6_dat,
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
	Func func() (*asset, error)
	Children map[string]*_bintree_t
}
var _bintree = &_bintree_t{nil, map[string]*_bintree_t{
	"test_data": &_bintree_t{nil, map[string]*_bintree_t{
		"0.dat": &_bintree_t{test_data_0_dat, map[string]*_bintree_t{
		}},
		"1.dat": &_bintree_t{test_data_1_dat, map[string]*_bintree_t{
		}},
		"2.dat": &_bintree_t{test_data_2_dat, map[string]*_bintree_t{
		}},
		"3.dat": &_bintree_t{test_data_3_dat, map[string]*_bintree_t{
		}},
		"4.dat": &_bintree_t{test_data_4_dat, map[string]*_bintree_t{
		}},
		"5.dat": &_bintree_t{test_data_5_dat, map[string]*_bintree_t{
		}},
		"6.dat": &_bintree_t{test_data_6_dat, map[string]*_bintree_t{
		}},
	}},
}}

// Restore an asset under the given directory
func RestoreAsset(dir, name string) error {
        data, err := Asset(name)
        if err != nil {
                return err
        }
        info, err := AssetInfo(name)
        if err != nil {
                return err
        }
        err = os.MkdirAll(_filePath(dir, path.Dir(name)), os.FileMode(0755))
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

// Restore assets under the given directory recursively
func RestoreAssets(dir, name string) error {
        children, err := AssetDir(name)
        if err != nil { // File
                return RestoreAsset(dir, name)
        } else { // Dir
                for _, child := range children {
                        err = RestoreAssets(dir, path.Join(name, child))
                        if err != nil {
                                return err
                        }
                }
        }
        return nil
}

func _filePath(dir, name string) string {
        cannonicalName := strings.Replace(name, "\\", "/", -1)
        return filepath.Join(append([]string{dir}, strings.Split(cannonicalName, "/")...)...)
}

