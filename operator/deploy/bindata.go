// Package deploy Code generated by go-bindata. (@generated) DO NOT EDIT.
// sources:
// staticresources/aro.openshift.io_clusters.yaml
// staticresources/namespace.yaml
// staticresources/role.yaml
// staticresources/serviceaccount.yaml
package deploy

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
		return nil, fmt.Errorf("read %q: %v", name, err)
	}

	var buf bytes.Buffer
	_, err = io.Copy(&buf, gz)
	clErr := gz.Close()

	if err != nil {
		return nil, fmt.Errorf("read %q: %v", name, err)
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

// Name return file name
func (fi bindataFileInfo) Name() string {
	return fi.name
}

// Size return file size
func (fi bindataFileInfo) Size() int64 {
	return fi.size
}

// Mode return file mode
func (fi bindataFileInfo) Mode() os.FileMode {
	return fi.mode
}

// ModTime return file modify time
func (fi bindataFileInfo) ModTime() time.Time {
	return fi.modTime
}

// IsDir return file whether a directory
func (fi bindataFileInfo) IsDir() bool {
	return fi.mode&os.ModeDir != 0
}

// Sys return file is sys mode
func (fi bindataFileInfo) Sys() interface{} {
	return nil
}

var _staticresourcesAroOpenshiftIo_clustersYaml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xc4\x59\xcf\x72\xdb\x38\xd2\xbf\xfb\x29\xba\x3c\x07\x4f\xea\xb3\xa4\x2f\xd9\xcb\xae\x6e\xa9\x64\x66\x4a\xbb\xb3\x89\x2b\xf6\xe6\x92\xe4\xd0\x02\x5b\x64\xaf\x41\x80\x0b\x34\xa5\x28\x53\xf3\xee\x5b\x0d\x90\x12\x25\x93\xb1\x93\xaa\x99\xe5\xc5\x66\x13\xe8\xff\xfd\xeb\x06\x74\x31\x9b\xcd\x2e\xb0\xe1\xf7\x14\x22\x7b\xb7\x04\x6c\x98\x3e\x0b\x39\x7d\x8b\xf3\xfb\xbf\xc6\x39\xfb\xc5\xf6\xf9\x9a\x04\x9f\x5f\xdc\xb3\x2b\x96\xf0\xaa\x8d\xe2\xeb\x77\x14\x7d\x1b\x0c\xbd\xa6\x0d\x3b\x16\xf6\xee\xa2\x26\xc1\x02\x05\x97\x17\x00\xe8\x9c\x17\x54\x72\xd4\x57\x00\xe3\x9d\x04\x6f\x2d\x85\x59\x49\x6e\x7e\xdf\xae\x69\xdd\xb2\x2d\x28\x24\x09\xbd\xfc\x1f\x5b\x77\xef\xfc\xce\x3d\xbb\x00\x30\x81\x12\x87\x3b\xae\x29\x0a\xd6\xcd\x12\x5c\x6b\xed\x05\x80\xc3\x9a\x96\x60\x6c\x1b\x85\x42\x9c\x63\xf0\x73\xdf\x90\x8b\x15\x6f\x64\xce\xfe\x22\x36\x64\x54\x6c\x19\x7c\xdb\x2c\xe1\xc1\xf7\xcc\xa1\xd3\xac\xb3\x2a\x33\x4b\x14\xcb\x51\xfe\x31\xa4\xfe\xca\x51\xd2\x97\xc6\xb6\x01\xed\x51\x74\x22\x46\x76\x65\x6b\x31\x1c\xc8\x17\x00\xd1\xf8\x86\x86\x5c\x63\xbb\x0e\x9d\xcb\x3a\xb9\x51\x50\xda\xb8\x84\xdf\x7e\xbf\x00\xd8\xa2\xe5\x22\x59\x9b\x3f\xaa\xba\x2f\x6f\x56\xef\xff\x72\x6b\x2a\xaa\x31\x13\x01\x0a\x8a\x26\x70\x93\xd6\xf5\xcc\x81\x23\x48\x45\x90\x57\xc2\xc6\x87\xf4\xda\xab\x08\x2f\x6f\x56\xdd\xee\x26\xf8\x86\x82\x70\xaf\x81\x3e\x83\xe0\x1f\x68\x67\x72\xae\x54\x91\xbc\x06\x0a\x0d\x37\x65\x81\x5d\xd0\xa8\x80\x98\x45\xfb\x0d\x48\xc5\x11\x02\x35\x81\x22\xb9\x9c\x00\x03\xb6\xa0\x4b\xd0\x81\x5f\xff\x9b\x8c\xcc\xe1\x96\x82\x32\x81\x58\xf9\xd6\x16\x9a\x23\x5b\x0a\x02\x81\x8c\x2f\x1d\x7f\x39\x70\x8e\x20\x3e\x89\xb4\x28\xd4\x85\xa2\x7f\xd8\x09\x05\x87\x56\x5d\xd8\xd2\x35\xa0\x2b\xa0\xc6\x3d\x04\x52\x19\xd0\xba\x01\xb7\xb4\x24\xce\xe1\x9f\x3e\x10\xb0\xdb\xf8\x25\x54\x22\x4d\x5c\x2e\x16\x25\x4b\x9f\xee\xc6\xd7\x75\xeb\x58\xf6\x8b\x94\xb4\xbc\x6e\xc5\x87\xb8\x28\x68\x4b\x76\x11\xb9\x9c\x61\x30\x15\x0b\x19\x69\x03\x2d\xb0\xe1\x59\x52\xdc\xa5\x6c\x9f\xd7\xc5\x0f\x87\x40\x5f\x0d\x34\x95\xbd\x26\x44\x94\xc0\xae\x3c\x90\x53\xee\x4d\xfa\x5d\x73\x50\xa3\x8b\xdd\xb6\xac\xff\xd1\xbd\x4a\x52\xaf\xbc\xfb\xe9\xf6\x0e\x7a\xa1\x29\x04\xa7\x3e\x4f\xde\x3e\x6e\x8b\x47\xc7\xab\xa3\xd8\x6d\x28\xe4\xc0\x6d\x82\xaf\x13\x47\x72\x45\xe3\xd9\x49\x97\x49\x4c\xee\xd4\xe9\xb1\x5d\xd7\x2c\x1a\xe9\xff\xb4\x14\x45\xe3\x33\x87\x57\xa9\xe8\x61\x4d\xd0\x36\x05\x0a\x15\x73\x58\x39\x78\x85\x35\xd9\x57\x18\xe9\x0f\x77\xbb\x7a\x38\xce\xd4\xa5\x8f\x3b\x7e\x88\x55\xa7\x0b\xb3\xb7\x0e\xe4\x1e\x4a\x46\x23\xd4\x55\xe0\x6d\x43\xe6\xa4\x32\x0a\x8a\x1c\x34\x7b\x05\x85\x34\xe7\x87\xe8\x32\x5d\x8b\xa9\x1e\x4d\x78\xa3\xe8\x76\x42\x9c\x30\x42\x9f\x92\x1c\x6d\xf1\x57\x5f\x96\xec\xca\xf3\x5d\x53\x42\x20\x03\xf2\x86\xcb\x91\xda\x3f\x6c\x46\xd1\xca\x5a\xc2\xd5\x87\xff\x9f\xfd\xed\xd3\xff\xcd\xf3\x9f\xab\x91\xa5\x93\xda\x25\x57\x7b\xc7\xe2\xf5\xd3\x2f\xaf\x6e\x7f\x72\x5b\x0e\xde\xd5\xe4\x64\x4c\x26\xb9\xb6\x1e\xa3\xcf\xe0\x35\x63\xe9\x7c\x14\x36\xf1\x26\xf8\x62\x74\xcd\xdd\x39\x36\x7c\xab\x76\xef\xa8\x9c\x70\xc6\x13\x79\xdc\x91\xc3\x71\xcb\xbe\xca\x20\x75\xa3\x06\xcd\x83\xb0\x3f\xb2\x73\x34\x63\xf5\xe9\xc1\x60\x55\x9c\x73\x3c\xc9\xdf\xbe\x89\xaf\x5e\xf7\x4d\xe4\xe5\x97\x36\xd0\x60\x7b\xc6\x74\x1a\x74\xb6\x27\xa8\x36\x5e\x48\xb9\xdd\x3d\x56\x4a\x69\xd5\x49\x31\xf9\x75\x54\xb4\xfa\xae\x6a\x32\xde\x15\x3c\x98\x42\xa6\x84\x1f\x96\x75\x70\x4b\x92\xe4\xf4\x64\x60\x17\x05\x9d\xa1\x38\x3f\x63\xc3\x42\xf5\x48\x71\x9d\x70\xbf\x3c\xf2\x39\x62\x70\x6e\x83\x6a\x59\x6a\x92\x27\x8d\xf1\x2a\x66\x5b\xcf\x85\xe9\x33\x50\x15\x03\xe9\x9e\xc3\xbc\x06\x35\x99\x0a\x1d\xc7\x3a\x35\x45\x57\x50\xa1\x7d\x53\xf1\x38\x52\x01\xbb\x8a\x9c\x3a\x74\x84\x69\x41\x82\x6c\xe3\x41\x89\xa3\x5a\x2a\x43\x41\x1d\xa1\x09\xec\x03\x43\x9a\xce\xc0\x07\xd8\xa5\x96\x9d\xbe\x35\x8d\xdd\x8f\x25\xae\x07\xb4\xf6\xe8\xbb\xc4\x1e\x4a\xde\x92\x03\x6d\x6d\x73\xf8\xe8\x86\xf6\x74\x53\xc0\x9a\x00\x8b\x82\xc6\x4a\x5c\x3c\xd0\xe7\xc6\xb2\x61\xb1\xfb\x3c\x2e\xec\x07\xb1\x07\xa9\x50\xd4\xd8\x10\xd3\x18\x60\x7c\xdd\x78\x97\xbc\x6d\x92\xb3\xd6\xbe\x1d\x03\x88\x80\x52\xa5\x16\x88\x2e\x75\x34\x0e\xb9\xb3\xfa\x48\x27\xdc\x93\x2f\x53\xbb\x54\x70\x4f\xcd\xd2\xeb\xce\x11\x96\x03\x1f\xc6\x39\xbc\x75\x86\xba\x9c\x2e\xae\x53\x52\xd7\x84\x4e\x85\x24\x97\x1c\xf3\xc3\xa0\x83\xdc\x43\x47\x78\x6a\x70\x4b\x2a\x00\xc3\x9a\x25\x60\x60\xbb\x87\x19\xb0\xae\x36\xbe\xa6\x08\x0d\x06\xe9\x0b\xf6\xe5\xcd\x2a\xcf\x42\x15\xe6\x32\x8a\x58\x8f\x31\x5d\xa3\xb9\xdf\x61\x28\xe2\x2c\xad\xde\xf8\x90\xdf\xd4\x77\x28\xbc\x66\xcb\x92\x5c\x6d\x28\xb8\x2e\x43\xf6\xd9\xec\x24\x6f\xcc\xf6\x83\x06\xf3\xcb\x07\x9f\xbf\xd6\x92\x00\x2c\x46\xb9\x0b\xe8\x22\xf7\x83\xff\xd8\x2a\x50\x35\x6b\x94\x25\xe8\x98\x31\x13\x1e\xb5\xec\x11\xcc\xd5\x19\x20\x46\x2c\x27\x24\x3c\xb2\x37\x10\xc6\xf1\x4e\x31\x05\x2d\xef\xd2\x0e\xc5\x97\xb3\xe2\x44\xf0\x8e\x66\x3b\x1f\x8a\xeb\xe3\xb0\x34\xca\x18\xce\x26\xeb\x03\x38\xa3\x50\xe9\xc3\x5e\xdf\x0d\xb6\x91\x0e\x1f\xda\x10\xc8\x49\x87\xbd\x63\x70\xa2\xcf\x4a\x46\xb4\x4a\x90\xc1\x2e\x45\x9e\x95\x63\x2b\x4d\x2b\xd7\x10\x5b\x53\x01\xc6\xa4\xb3\x65\x37\xa5\xa8\x9e\xef\x8c\x58\x28\x15\x49\xbb\xad\x9a\x5f\xec\x20\xb6\x75\x8d\x81\xbf\xa4\xf4\x37\x59\xc5\x0e\x1d\x92\xf2\x13\x7a\x3e\x12\x90\x87\xed\xe5\xc9\x5b\xd3\xe7\xc7\x23\x79\x84\xf1\xbb\x7d\x43\x7d\xc3\xd4\xcd\x07\x77\x1f\xea\x38\x99\x1a\x27\x7c\x23\xfb\x86\x0d\x5a\xbb\xd7\xd2\xef\x03\x5e\x80\x66\x80\x02\x6b\xac\x7c\x10\x68\xaa\x90\x66\xe6\x21\x44\x26\x61\x53\x5c\x3b\xf4\x64\x57\xb0\xe6\x43\xd7\x2d\x39\x41\x3e\x7c\xbc\xc4\xb5\xd3\x9a\xb1\x33\x09\x2d\x7d\xbc\x84\xc6\x5b\x0c\x2c\xfb\x39\xfc\xec\xc7\x00\x4c\x1f\xfa\x8c\x75\x63\xe9\x1a\xf8\xdc\xbe\x5e\x4a\xcc\x5d\x05\x95\x1d\x9b\x7d\xce\xa3\x74\x96\xbd\x9e\x32\x3e\x69\xc3\x31\x9f\x78\x3f\x5e\x82\xc1\x98\x9c\xd9\x04\xbf\xc6\xb5\xdd\xa7\x15\xaa\xeb\x35\x44\x7f\x2a\xf6\xeb\x96\xaf\xb5\x10\xac\xa5\x02\x3e\x5e\xae\x5c\xc7\x7e\x04\x81\xe0\xb1\x8c\xc8\x2d\x80\x1e\x8c\x4e\x3a\x62\xe6\x34\x1b\xf9\xa0\x1c\x1f\x90\x27\xe7\xb3\xfe\x13\x86\x80\xa7\xad\x53\xd1\x11\xc5\x87\x89\xc1\x7c\x52\xf1\x40\x7a\x36\x2e\xde\x26\x59\x0f\xea\xe0\x29\x63\x4a\xde\xfa\x8e\x36\x94\x0a\x32\xdd\xd8\x20\xbb\x08\xe4\x7c\x5b\x56\xe9\xe0\xa6\xa8\x9b\x12\xd1\x83\x25\x81\xbd\x6f\x47\xdc\xcb\x4e\x0f\x4d\xa2\xb9\x5c\xfb\x82\x37\x39\xa4\x81\xba\xbe\xd9\x1d\xfe\xbf\xb1\x33\x8c\xdf\x53\x4c\x98\xf2\xf2\x66\xd5\xdf\x4e\xf4\xb5\x19\xb2\x5d\x23\x72\xbf\xea\xd6\xfc\x6c\x98\x6c\x71\x83\x52\x3d\x41\xf6\xd5\x6a\xd3\xd9\x9a\x86\x08\xaf\xc5\xc1\x64\xe8\x64\xc2\x4b\x43\x11\xa1\x8e\xd7\x13\x49\xad\x83\x9d\x13\x0e\xd4\xed\xb8\xce\x27\xf4\xee\x22\xe0\x78\x61\xa2\x21\x02\xcc\xd5\x04\x7f\xbf\x7d\xfb\x66\xf1\x8b\x9f\x60\x99\xac\x00\x34\x86\x62\x37\x60\xea\x31\xec\x08\xe9\xdd\xa9\xf5\x36\x8d\x9e\x35\x3a\xde\x50\x94\x79\x27\x83\x42\xfc\xf0\xe2\xd3\x54\x0b\xf9\xd9\x87\x07\x68\x71\xb8\x76\xe8\x13\x8a\x63\x76\xc7\x81\x23\xec\x58\x2a\x9e\x2a\x6b\x05\x95\xa2\x33\x3b\x0f\x9b\x82\xf7\x04\xbe\x33\xb7\x25\xb0\x7c\x4f\x4b\xb8\xd4\x6c\x1b\xa8\xf9\x9b\x9e\xa6\x7e\x1f\xaf\x7b\x80\x1f\x77\x15\x05\x82\x4b\x5d\x74\x99\x95\x3b\xdc\x2e\x29\x6d\x80\xe5\x9d\x92\x69\xa8\x94\xc0\x65\x49\x61\x74\x2a\x85\x0e\xd6\x68\x4b\x4e\x9e\x69\xda\xf3\x06\x9c\x1f\xb0\x48\x8c\x35\x7a\x0d\x19\xde\x30\x15\x0f\x94\xfe\xf0\xe2\xd3\xa4\xc6\xa7\xfe\x52\xec\xa5\xcf\xf0\x22\x43\xbb\x02\xa7\x2f\x9e\xcd\xe1\x2e\x65\xc7\xde\x09\x7e\x56\x49\x46\x87\xd7\x29\xcf\x7a\xa7\x30\xeb\xa1\xc2\x2d\x41\xf4\x35\xc1\x8e\xac\x9d\x75\x03\x2a\xec\x30\x4d\x12\x7d\xe0\x34\xdf\xb0\x9f\x2d\xa7\xb3\xb5\xbf\xd3\xbb\x7b\xfb\xfa\xed\x32\x6b\xa6\x09\x55\xa6\x91\x47\x67\xda\x0d\x3b\xb4\xa9\x33\xe6\xbb\xa6\x94\x8d\x93\x4d\x32\xb6\x39\x7d\xc4\x77\x73\x6f\xdf\xca\x36\xad\xb4\x81\xe6\x63\xd7\x10\x8f\xd6\xf1\xf9\x75\xdb\xf1\x19\xb9\x78\x3b\x07\x8e\xff\xd1\xf5\xd5\x93\x8d\x73\x23\xd7\x46\x63\xc6\xbd\x19\x64\xf9\x57\x8d\xd3\x19\x2e\x38\x12\x4a\xf6\x15\xde\xc4\x45\x3a\x07\x34\x12\x17\x7e\x4b\x61\xcb\xb4\x5b\xec\x7c\xb8\x67\x57\xce\x34\x35\x67\x39\x07\xe2\x22\xdd\x67\x2c\x7e\x48\x7f\xbe\xdb\x96\xc9\x0b\x91\x31\x83\xd2\xe2\x3f\xc3\x2a\x95\x13\x17\xdf\x65\x54\x7f\xab\xf2\xf4\x3e\x76\x75\x9b\x01\xc3\x9c\xef\xd5\xb2\xd8\x55\x6c\xaa\xfe\xc2\xbd\xc3\xd8\x89\x62\xe2\x08\x35\x16\x19\x9a\xd1\xed\xff\xf0\x54\x56\x87\xe6\xb9\x7e\x3f\xeb\x7e\xf8\x99\xa1\x2b\xf4\xff\xc8\x51\x94\xfe\x5d\x1e\x6c\xf9\x49\xe5\xfb\xaf\xd5\xeb\x3f\x27\xc1\x5b\xfe\xae\x5a\xfd\xe6\xb1\x70\x64\xc3\x19\xe9\xf0\x03\xda\xf6\x39\xda\xa6\xc2\xe7\x47\x5a\x1a\xa7\x66\xdd\x6f\x66\x83\xcf\x00\xf9\x5e\x6d\x09\x7a\x24\xc8\x04\xf1\x41\x4f\xc4\x99\x72\x3c\x52\xe9\xcc\xd0\x08\x15\x6f\xce\x7f\x35\xbb\xcc\x4d\xab\xff\x59\x2c\xbd\x0e\xee\xdc\xe0\xc3\xa7\x8b\xcc\x95\x8a\xf7\xbd\x36\x4a\xfc\x6f\x00\x00\x00\xff\xff\x8b\xf6\x6a\x10\x78\x1c\x00\x00")

func staticresourcesAroOpenshiftIo_clustersYamlBytes() ([]byte, error) {
	return bindataRead(
		_staticresourcesAroOpenshiftIo_clustersYaml,
		"staticresources/aro.openshift.io_clusters.yaml",
	)
}

func staticresourcesAroOpenshiftIo_clustersYaml() (*asset, error) {
	bytes, err := staticresourcesAroOpenshiftIo_clustersYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "staticresources/aro.openshift.io_clusters.yaml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _staticresourcesNamespaceYaml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x14\xca\xb1\x0d\xc2\x40\x0c\x05\xd0\xde\x53\x78\x01\x17\xb4\x37\x04\x25\xfd\x27\xf9\x11\x56\x88\xef\x64\x5f\x28\x32\x3d\x4a\xf9\xa4\x27\x66\x26\x18\xfe\x62\x96\xf7\x68\xfa\x7b\xc8\xee\xb1\x36\x7d\xe2\x60\x0d\x2c\x94\x83\x13\x2b\x26\x9a\xa8\x06\x0e\x36\xed\x83\x51\x1f\xdf\xa6\xe1\x3a\x93\xd6\x07\x13\xb3\xa7\xd4\xe0\x72\xb7\xcd\x03\x5f\xbf\x98\x75\xcb\x74\x3f\xdf\xcc\xe0\x64\xc9\x3f\x00\x00\xff\xff\x44\x6f\xf6\xda\x72\x00\x00\x00")

func staticresourcesNamespaceYamlBytes() ([]byte, error) {
	return bindataRead(
		_staticresourcesNamespaceYaml,
		"staticresources/namespace.yaml",
	)
}

func staticresourcesNamespaceYaml() (*asset, error) {
	bytes, err := staticresourcesNamespaceYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "staticresources/namespace.yaml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _staticresourcesRoleYaml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xac\x91\xb1\x6e\xf3\x30\x0c\x84\x77\x3d\x85\x90\xdd\x0e\xb2\xfd\xd0\xfa\x0f\xdd\x8b\xa2\x3b\x23\x33\x09\x11\x49\x14\x48\xca\x05\xfa\xf4\x85\xec\x16\x28\xea\x06\xf0\xd0\xc9\xe4\x01\x77\xdf\xd1\x72\xc3\x30\x38\xa8\xf4\x8a\xa2\xc4\x25\x78\x39\x43\x1c\xa1\xd9\x8d\x85\xde\xc1\x88\xcb\x78\xff\xa7\x23\xf1\x71\x3e\xb9\x3b\x95\x29\xf8\xff\xa9\xa9\xa1\x3c\x73\x42\x97\xd1\x60\x02\x83\xe0\xbc\x8f\x82\x8b\xe1\x85\x32\xaa\x41\xae\xc1\x97\x96\x92\xf3\xbe\x40\xc6\xe0\x33\x14\xb8\xa2\x0c\xd2\x8d\xd2\x12\x6a\x70\x83\x87\x4a\x4f\xc2\xad\x6a\x8f\x18\xfc\xe1\xe0\xbc\x17\x54\x6e\x12\xf1\x53\x8b\x5c\x2e\x74\xcd\x50\x75\x59\x7b\x9a\x56\x88\xb8\xae\x8a\x32\x53\x44\x88\x91\x5b\xb1\xae\xcd\x28\xe7\x2f\x6b\xef\x84\xcb\x78\x45\x5b\xbe\xad\x4e\x5d\xda\x45\x56\x8c\x82\x3b\x32\x13\xe9\x3a\x54\xb0\x78\xfb\x8e\xe9\xe3\xdb\x22\x6e\x80\x50\xab\x8e\xf3\x69\x4b\x9d\x00\x33\x17\xdd\x03\xde\xf0\xb6\x14\xe1\x91\x2b\x16\xbd\xd1\xc5\x46\xe2\x5f\x7e\xef\xfa\x9e\x8f\x61\x13\x26\xfc\x93\x83\x77\x57\x39\xaa\x81\xb5\x1f\x8d\x1e\xde\xfc\x11\x00\x00\xff\xff\x8d\xa9\xfa\xde\xc5\x02\x00\x00")

func staticresourcesRoleYamlBytes() ([]byte, error) {
	return bindataRead(
		_staticresourcesRoleYaml,
		"staticresources/role.yaml",
	)
}

func staticresourcesRoleYaml() (*asset, error) {
	bytes, err := staticresourcesRoleYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "staticresources/role.yaml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _staticresourcesServiceaccountYaml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x44\xca\x31\x0e\x02\x41\x08\x05\xd0\x9e\x53\x70\x01\x0a\x5b\x3a\xcf\x60\x62\x4f\x66\xbf\x91\x98\x85\x09\xc3\x6e\xe1\xe9\x6d\x4c\xb6\x7d\x79\x24\x22\x64\xd3\x9f\xa8\xe5\x19\xca\xe7\x8d\x3e\x1e\x9b\xf2\x03\x75\xfa\xc0\x7d\x8c\x3c\xa2\x69\x47\xdb\x66\x6d\x4a\xcc\x61\x3b\x94\xad\x52\x72\xa2\xac\xb3\xfe\xb8\xa6\x0d\x28\xe7\x44\xac\xb7\xbf\x5a\xec\x7b\x14\xae\xf5\x0b\x00\x00\xff\xff\xf6\x4a\x0e\x0f\x6e\x00\x00\x00")

func staticresourcesServiceaccountYamlBytes() ([]byte, error) {
	return bindataRead(
		_staticresourcesServiceaccountYaml,
		"staticresources/serviceaccount.yaml",
	)
}

func staticresourcesServiceaccountYaml() (*asset, error) {
	bytes, err := staticresourcesServiceaccountYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "staticresources/serviceaccount.yaml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
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
	"staticresources/aro.openshift.io_clusters.yaml": staticresourcesAroOpenshiftIo_clustersYaml,
	"staticresources/namespace.yaml":                 staticresourcesNamespaceYaml,
	"staticresources/role.yaml":                      staticresourcesRoleYaml,
	"staticresources/serviceaccount.yaml":            staticresourcesServiceaccountYaml,
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
	"staticresources": {nil, map[string]*bintree{
		"aro.openshift.io_clusters.yaml": {staticresourcesAroOpenshiftIo_clustersYaml, map[string]*bintree{}},
		"namespace.yaml":                 {staticresourcesNamespaceYaml, map[string]*bintree{}},
		"role.yaml":                      {staticresourcesRoleYaml, map[string]*bintree{}},
		"serviceaccount.yaml":            {staticresourcesServiceaccountYaml, map[string]*bintree{}},
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
