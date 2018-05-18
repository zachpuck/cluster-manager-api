// Code generated by go-bindata.
// sources:
// api/api.proto
// DO NOT EDIT!

package protobuf

import (
	"github.com/elazarl/go-bindata-assetfs"
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

var _apiProto = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xcc\x59\x5f\x6f\xdb\x38\x12\x7f\xf7\xa7\x18\xf8\xe5\x92\x43\x63\xb5\xe9\x1e\x70\x48\x10\xec\xf9\xdc\xde\xd6\x68\x9b\x04\x75\xb6\xc1\x3e\x05\x34\x35\x91\x79\x95\x48\x2e\x49\xd9\xf5\x15\xfd\xee\x87\x21\x29\x59\xb4\xe4\xa4\x97\x0d\xb0\xe7\x87\xc4\x22\x87\xc3\xf9\xcd\x3f\xfe\x28\x67\x19\xcc\x94\xde\x1a\x51\xac\x1c\x9c\xbe\x7c\xf5\x77\x58\xb0\xca\xd6\xb2\x80\xc5\x9b\x05\xcc\x4a\x55\xe7\x70\xc9\x9c\x58\x23\xcc\x54\xa5\x6b\x27\x64\x01\x37\xc8\x2a\x60\xb5\x5b\x29\x63\x27\xa3\x2c\x1b\x65\x19\x7c\x10\x1c\xa5\xc5\x1c\x6a\x99\xa3\x01\xb7\x42\x98\x6a\xc6\x57\xd8\xcc\xbc\x80\xcf\x68\xac\x50\x12\x4e\x27\x2f\xe1\x88\x04\xc6\x71\x6a\x7c\x7c\x4e\x2a\xb6\xaa\x86\x8a\x6d\x41\x2a\x07\xb5\x45\x70\x2b\x61\xe1\x5e\x94\x08\xf8\x95\xa3\x76\x20\x24\x70\x55\xe9\x52\x30\xc9\x11\x36\xc2\xad\xfc\x3e\x51\x0b\x59\x02\xbf\x45\x1d\x6a\xe9\x98\x90\xc0\x80\x2b\xbd\x05\x75\xdf\x15\x04\xe6\xa2\xd1\xf4\x59\x39\xa7\xcf\xb2\x6c\xb3\xd9\x4c\x98\x37\x78\xa2\x4c\x91\x95\x41\xd4\x66\x1f\xe6\xb3\xb7\x97\x8b\xb7\x27\xa7\x93\x97\x71\xd1\xaf\xb2\x44\x6b\xc1\xe0\xef\xb5\x30\x98\xc3\x72\x0b\x4c\xeb\x52\x70\xb6\x2c\x11\x4a\xb6\x01\x65\x80\x15\x06\x31\x07\xa7\xc8\xe8\x8d\x11\xe4\xb7\x17\x60\xd5\xbd\xdb\x30\x83\xa4\x26\x17\xd6\x19\xb1\xac\x5d\xe2\xb3\xc6\x44\x61\x13\x01\x25\x81\x49\x18\x4f\x17\x30\x5f\x8c\xe1\x9f\xd3\xc5\x7c\xf1\x82\x94\xdc\xce\x6f\xde\x5d\xfd\x7a\x03\xb7\xd3\x4f\x9f\xa6\x97\x37\xf3\xb7\x0b\xb8\xfa\x04\xb3\xab\xcb\x37\xf3\x9b\xf9\xd5\xe5\x02\xae\xfe\x05\xd3\xcb\xdf\xe0\xfd\xfc\xf2\xcd\x0b\x40\xe1\x56\x68\x00\xbf\x6a\x43\x08\x94\x01\x41\xde\xc4\xdc\xbb\x6e\x81\x98\x98\x70\xaf\x82\x49\x56\x23\x17\xf7\x82\x43\xc9\x64\x51\xb3\x02\xa1\x50\x6b\x34\x92\x32\x41\xa3\xa9\x84\xa5\xa8\x5a\x60\x32\x27\x35\xa5\xa8\x84\x63\xce\x0f\xf5\x70\x4d\x46\x23\xbb\x95\x8e\x7d\x85\x0b\x18\x6b\xa3\x9c\x7a\x3d\x3e\x1f\x8d\x34\xe3\x5f\x48\x31\x2f\x6b\xeb\xd0\xdc\x55\x4c\xb2\x02\xcd\x1d\xd3\xe2\x7c\x34\x12\x95\x56\xc6\xc1\xb8\x50\xaa\x28\x31\x63\x5a\x64\x4c\x4a\x15\x37\x99\x78\x35\xe3\xf3\x56\xcc\x3f\xf3\x93\x02\xe5\x89\xdd\xb0\xa2\x40\x93\x29\xed\x45\x07\x97\x8d\x46\x61\x16\x8e\x0a\xa3\xf9\xa4\x60\x0e\x37\x6c\x1b\xa6\xf9\x5d\x81\xf2\x2e\x6a\x99\x44\x2d\x13\xa5\x51\x32\x2d\xd6\xa7\xcd\xcc\x31\x5c\xc0\xb7\x11\x80\x90\xf7\xea\xcc\x7f\x03\x70\xc2\x95\x78\x06\xe3\x59\x80\x04\x1f\x03\x24\x98\x5e\xcf\xc7\xe7\x5e\x62\x1d\xca\xe1\x0c\xc6\xeb\x97\x93\x57\x93\x97\x71\x98\x2b\xe9\x18\x77\x8d\x1e\xfa\x48\x56\x91\xaa\x8f\x82\xaf\x18\x96\xf0\x19\x25\xfe\x47\xb0\x28\x4f\x9f\xda\x94\x67\x30\xa6\x4c\xb6\x67\x59\x56\x08\xb7\xaa\x97\x13\xae\xaa\x6c\xdd\x13\xc5\x8a\x09\x12\xae\xe2\xd4\x3f\x0a\x1a\x20\xe1\x28\xf4\x9d\xfe\xf9\x3f\xf8\xd5\xa1\x91\xac\xbc\xcb\x15\xb7\x8d\x3d\x07\xb7\xb2\xa1\x6f\x9c\x70\xc9\x5d\x16\x03\x79\x12\x03\x79\xc2\xb4\x88\xea\x73\xb4\xdc\x08\xef\x49\x82\xa4\x0c\x02\x5b\xaa\xda\xc1\x21\x47\x7d\x1f\x01\x58\xbe\xc2\x0a\xed\x19\xbc\xbb\xb9\xb9\x3e\xdf\x1f\x58\xd0\x08\x57\xd2\xd6\x7e\x68\x1c\xab\x91\xb6\xc8\xfe\x6d\x95\xf4\x6a\xb4\x51\x79\xcd\x0f\xcd\x7f\x3f\x1f\x8d\x2c\x9a\xb5\xe0\xd8\x1a\x12\xf0\x52\x61\x50\x95\x20\xbc\xc3\xb2\x54\x70\xab\x4c\x99\xc3\x22\xca\x9e\xc0\x46\x94\x25\x18\xd4\xc8\x1c\x30\xa0\xaa\xf7\x2d\xd2\x29\x9f\xf6\x14\x39\xda\x7a\x2d\x72\xcc\xbd\x3e\xa3\x79\xd0\x14\x14\x1d\xed\xbe\x7f\xb4\xc5\x31\x18\x74\xb5\x91\xb6\x3b\xfe\x09\x75\xb9\x3d\xee\xa4\x43\x9b\xaf\xbe\x1e\x26\x4c\x8b\x09\xc5\xa3\xc9\x42\xfa\x68\x65\x1d\x9c\xc1\xd8\x17\xcb\xfa\x55\xb6\x22\x6d\x1b\xd2\x36\x8e\x12\x4b\x95\x6f\xcf\x60\xfc\xd7\xf1\x2e\xe8\xc1\xd7\x5d\xc8\x5a\xe5\xc0\x55\x2d\x1d\x18\xb4\x5a\x51\x01\x03\xdc\x06\xc4\xf4\x9c\xef\x9a\xb0\xac\xab\x25\x1a\xea\xb4\x5a\xe5\x96\xba\x5e\x83\xdf\x6a\xc6\x07\x9c\xf0\x0b\xba\x6b\x95\xcf\xbc\xf6\xa3\xce\x43\xea\x86\xce\xc4\x53\xfc\x30\xec\x8d\x02\x9d\x56\xb9\x07\x36\x4e\x04\xc9\x29\xb0\xf3\xca\x90\x67\x3c\x7a\x0f\xc6\x9f\x65\xac\xe9\x59\x2d\xae\x99\x41\xe6\xb0\x49\xa2\xa3\xe4\x31\xc5\x96\x4c\xfd\x01\x74\x75\x02\x2e\xda\xf3\x34\x60\x06\x9d\x11\xb8\x0e\x07\x81\x75\xcc\xd5\x96\x42\xda\xa2\xa4\x26\x0f\xc2\x59\xf8\x52\x2f\x91\x2b\x79\x2f\x0a\x7f\x4e\x70\x25\x25\x72\x27\xd6\xc2\x6d\xbb\x11\x6e\xdd\xb0\xfb\xde\x8b\xef\x1f\x76\x40\x81\x0f\x3b\x60\x10\x69\x8e\x25\x3a\x1c\x88\xdf\x1b\x3f\xd1\x1a\x9e\x3c\xa6\xb6\x27\x53\x4f\x37\x3f\x5a\xf2\x3f\x23\x68\x63\xc5\xa0\x14\xd6\x51\x9c\xe2\x42\x3b\x10\x82\x0f\x24\x72\x94\x3e\x1f\x0a\x05\xcd\x3d\x77\x38\x32\xb2\xf1\x11\x44\x42\x5a\xc7\xca\x12\x8e\x94\x01\x83\xf1\xe9\x18\x9c\x28\xcb\x4e\x80\xae\x9b\xe2\xbb\xf1\xe3\x70\xb4\x37\x90\xa2\xda\x9b\x7c\xbe\x26\x12\xac\x7a\x5a\x99\x1d\x00\xba\xc2\xb2\x02\xbe\x62\xc6\x35\xd2\x37\x44\x81\xfd\x11\xb3\x44\x3a\x3a\x9d\xa9\xb9\x27\xe3\xc2\x17\x25\x89\xc2\x8a\x59\x60\xa5\x41\x96\x6f\x61\x89\x28\x21\x47\x5d\xaa\x2d\x76\x9a\xb3\xa5\x63\x88\x7a\x71\xeb\xc4\x79\xd8\xf3\x1d\x96\xd5\xcc\x6b\x39\xda\x1f\x49\xdd\xb8\x3f\xfb\x6c\xed\x8a\x30\x3f\xcd\x89\xb1\x6e\x5a\xb4\x7b\xde\xdb\xd5\x72\x07\xe4\xde\xc0\x50\x3d\x3f\x03\xc4\x7e\x45\xa7\x28\x5b\x38\xdf\x47\x23\x1f\xe4\x94\x5c\xd0\xcd\x02\xad\x1b\x55\x68\x2d\xb1\xe2\x84\x22\xc4\xad\xe8\x6a\x20\x8b\xc0\x2f\x2e\xe0\xd5\x79\x47\x55\x73\x54\x13\x07\xe9\xa8\x1d\x50\xe7\x11\xa6\x0a\x1b\xa1\x46\x67\xf3\x9c\x1e\xd0\x3b\x6e\x74\xd9\x1e\xf0\x4e\xc1\x3d\x3a\x1e\x12\xae\x25\x0e\x8d\xdc\x07\x64\x6b\x04\xac\xb4\xdb\x92\xe4\xef\x35\x9a\x2d\x50\x09\xb4\x0c\xc1\xee\xe3\x0a\x6a\x1f\x30\xa4\x6b\x3e\x99\xf2\x08\x03\xa1\x62\x4b\x77\x3c\xf6\x4b\x85\x74\xaf\x4f\xc3\x9a\xfd\xcd\xf6\x0f\xef\x14\x77\x73\xad\x6c\xce\x46\xa7\xa8\x4a\x5b\x6e\x10\xb9\x4e\x2f\x52\x6d\x69\xb7\xa4\xc8\xb4\x77\x2c\x4f\x4a\xbd\x44\xb2\xf5\x75\x94\x5b\x68\xe4\xbb\x45\x17\x70\x7a\xd8\xda\x3d\xe7\xdc\xae\xd0\xdf\xfd\x94\xf1\xd7\xeb\xae\xd9\x1b\x66\xbb\x46\xd3\x7d\xd6\xdf\xbc\x9b\x34\x0c\x35\xa9\x4a\x50\x5f\x7a\x00\x88\x22\xf4\xfc\x50\xa0\x44\xb3\x43\x12\x1d\x10\xd9\xc4\xbe\xd1\x09\x31\xf8\x11\xff\x96\x4a\x7d\xa1\x2b\xb3\x3e\x58\x07\x7d\xd5\x7b\xce\x98\xdb\x44\x6f\xcc\x14\xbb\xb5\x0e\xab\x83\x70\x6f\x57\xcc\xd1\xb5\x3c\x65\x46\x1d\x3d\x87\xc0\x0e\xac\xef\xb0\x27\xa7\x1a\xf2\xd4\xdc\x19\x06\xd4\x75\xe4\x2f\xe0\x75\x02\x72\x9f\x9f\x74\x43\xbe\xdb\x30\xea\xfc\x8b\x0d\x9e\x72\x2a\x1c\x26\x6a\xfb\xa8\x13\xfb\x24\x67\xb7\xc3\x4c\xd5\x65\x9e\xb8\xb2\x39\xa5\xa8\x21\x1f\xf4\xe4\x22\xf1\x5e\x37\xcd\x7e\x34\x55\x22\x79\x81\x6f\x87\xa7\x9f\x23\xe4\x1f\x06\x69\x95\xbf\xeb\x61\x0e\x9d\xcd\xe6\x0e\xab\x56\xaa\x5f\x97\x7b\x82\x0f\x25\xf9\x83\x2d\xe3\xc9\x29\xf8\x7d\xb4\xb3\xa6\xcf\x96\x3a\x01\x6d\x2a\x2d\x30\x2b\xbb\xf2\x01\x5e\x62\x43\x58\xfc\x9b\xa8\xee\x26\x8d\x3f\xbb\x56\x76\x8e\x84\xc3\x6a\x84\x3c\xd4\xee\x77\x15\xe3\xdf\x17\x4a\xe5\x31\x7a\x4d\xfe\x18\x0d\xef\xd4\x02\x7f\xca\xa0\xd6\x85\x61\x39\x65\x74\x57\x5f\x7c\xb5\x12\x8a\x25\xcd\x80\x68\x53\xcb\xfb\x4f\x36\x22\x6f\x46\x7f\x6e\x33\x34\x58\x2c\x88\x5c\xad\x31\x15\x65\x79\x25\x24\x68\x23\xd6\xa2\xc4\x02\xed\xcf\xbb\xf4\x69\x5e\x5f\x79\xb9\x0b\xf8\xa9\xef\x12\xb2\x81\x62\xe8\x3a\x4e\xf1\xaf\x0d\x9d\x8a\x8a\xa3\x7f\xdb\x2c\x8b\x88\xfc\xe4\xdd\xee\xe8\x82\x0b\xf8\x5b\x92\x64\x43\x3c\xb7\xd3\x0b\x58\x40\xaf\x74\xec\xcc\x60\x6b\xce\xd1\xda\xfb\xba\x7c\xb8\xe3\x45\xfd\x16\x36\x68\x10\x0a\xb1\x46\x39\xcc\x17\xd2\x9c\x1f\x60\x93\xcf\x9c\x66\x91\xfd\x5b\x74\x4e\xc8\x22\xd4\xe6\x2f\x74\xf2\x08\x1e\xa6\x16\x61\xa6\xd9\xa5\x9b\x59\x81\x0d\x2e\x06\x96\xee\xb8\x62\x20\xd7\xfb\xed\x76\x90\x07\xff\x89\x7e\xee\x13\xda\xff\x3b\x37\x5f\x36\xf7\x8e\x66\x97\x41\xc7\x0e\x91\xef\x3f\xd1\xaf\x83\x10\x07\x8f\xd6\x88\xbb\x2d\xcd\x07\x68\xec\xc0\xea\xa6\x53\xa5\x3d\x6e\xb8\x97\x0d\x5a\xb8\x4b\xd8\x41\xeb\x64\xe7\x88\x09\x37\xa5\x0a\xa5\xfb\xe1\x53\x46\x76\xe9\x7d\x58\xef\xc7\x3b\xaf\x50\xf7\x5a\xef\x70\x2b\x4f\xb8\x88\xb7\xd6\xa0\x56\x56\x38\x65\x12\x0a\x42\xa3\x49\xd7\xee\x2f\x94\x07\xd2\xe9\xa7\x74\x0d\x33\x18\x33\x26\xfc\x14\x71\x24\xd1\x52\x3f\xdd\xb2\xaa\x84\x13\x3f\xf5\x99\x95\x35\xda\x89\x1f\xe1\x4a\x3a\x94\x2e\x5e\x09\x1a\xd7\x7b\x81\x5e\xab\x3d\x4c\xcd\x07\x43\xd0\xf2\xf5\x13\xe0\xb5\x31\x28\x5d\x19\x19\xb6\xb0\xc0\x36\xfe\x77\x98\x8a\xb1\xde\xfd\xa7\x47\xb6\x3f\x32\xb6\x78\xec\xa6\x40\x32\xde\x14\xd2\x98\x84\x80\x34\x4c\x6f\x1f\x55\x30\xbd\x0d\xeb\xc9\xb0\xfd\x1a\x1d\xde\xe9\x5b\xcf\xc6\xe9\xf5\x1c\x50\xe6\x5a\x89\x34\xd5\x9a\xb1\x1e\xb2\xda\xa2\xf1\x90\x23\x2f\x6b\xb4\x74\x17\xb7\x32\xfb\xa0\xae\xa6\xb5\x5b\xc1\x17\xdc\xb6\xbf\x58\x0d\xed\xad\x58\xed\x56\x77\x24\xf5\x20\xaa\x06\x7e\x02\xca\x60\x41\x89\x4e\xea\xa7\xb7\x8b\x34\x5f\x8b\x50\x9b\x29\x9e\x05\x72\x83\xee\x3d\x6e\xe7\x79\x58\x75\x3d\x87\xa9\xef\x57\x09\x43\xf3\x52\x64\xd3\x9d\xc8\x7b\xb0\x82\x8e\xb0\xea\x7d\x04\x47\x7a\xd8\x21\x3d\x61\xa2\x03\xf1\xbf\x01\x00\x00\xff\xff\xe8\xd3\xb4\x5a\xdc\x1d\x00\x00")

func apiProtoBytes() ([]byte, error) {
	return bindataRead(
		_apiProto,
		"api.proto",
	)
}

func apiProto() (*asset, error) {
	bytes, err := apiProtoBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "api.proto", size: 7644, mode: os.FileMode(420), modTime: time.Unix(1526658465, 0)}
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
	"api.proto": apiProto,
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
	"api.proto": &bintree{apiProto, map[string]*bintree{}},
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


func assetFS() *assetfs.AssetFS {
	assetInfo := func(path string) (os.FileInfo, error) {
		return os.Stat(path)
	}
	for k := range _bintree.Children {
		return &assetfs.AssetFS{Asset: Asset, AssetDir: AssetDir, AssetInfo: assetInfo, Prefix: k}
	}
	panic("unreachable")
}
