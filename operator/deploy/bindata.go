// Package deploy Code generated by go-bindata. (@generated) DO NOT EDIT.
// sources:
// resources.yaml
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

var _resourcesYaml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xc4\x5a\x5f\x73\xe3\xb6\x11\x7f\xd7\xa7\xd8\x71\x1e\xae\x99\x31\xa5\xf8\x3a\x9d\xe9\xe8\xcd\xbd\x4b\x52\xb7\xe9\xdd\x8d\xed\xe6\x25\x97\x87\x15\xb8\x22\xb7\x06\x01\x16\x00\xa5\xd3\x75\xfa\xdd\x3b\x0b\x90\x14\x25\x91\x96\xe2\x4b\x52\xbe\x58\x04\x81\xfd\xfb\xdb\x3f\x00\x8c\x35\xff\x48\xce\xb3\x35\x4b\xd8\xdc\xcc\x9e\xd8\xe4\x4b\x78\x87\x15\xf9\x1a\x15\xcd\x2a\x0a\x98\x63\xc0\xe5\x0c\xc0\x60\x45\x4b\xb0\x35\x19\x5f\xf2\x3a\x64\xf8\xb9\x71\x94\xd9\x9a\x1c\x06\xeb\x66\x59\x96\xcd\x86\xd4\xb0\x66\xfa\x14\xc8\xc8\x9b\x9f\x3f\xfd\xd9\xcf\xd9\x2e\x36\x37\x2b\x0a\xd8\xf1\x79\xd3\xf8\x60\xab\x7b\xf2\xb6\x71\x8a\xde\xd2\x9a\x0d\x07\xb6\xe6\x80\x2d\x1a\x63\x03\xca\xb0\x97\x57\x00\x65\x4d\x70\x56\x6b\x72\x59\x41\x66\xfe\xd4\xac\x68\xd5\xb0\xce\xc9\x45\x0e\xbd\x36\xdf\xcc\x5f\xcf\xff\x34\x03\x50\x8e\xe2\xf2\x47\xae\xc8\x07\xac\xea\x25\x98\x46\xeb\x5e\x23\xa5\x1b\x1f\xc8\xf9\x39\x3a\x3b\xef\xd5\x9b\xb3\x9d\xf9\x9a\x94\xf0\x2c\x9c\x6d\xea\x25\x9c\x7c\x4f\x14\x5a\xb1\x5a\x95\x12\xb1\x38\xa2\xd9\x87\xbf\x0f\x47\x7f\x60\x1f\xe2\x97\x5a\x37\x0e\xf5\x9e\x75\x1c\xf4\x6c\x8a\x46\xa3\xeb\x87\x67\x00\x5e\xd9\x9a\x06\x0e\xc9\x67\x00\x1b\xd4\x9c\x47\x8d\x12\x63\x11\xe9\xf6\xc3\xdd\x8f\x7f\x7c\x50\x25\x55\x98\x06\x01\x72\xf2\xca\x71\x1d\xe7\x75\x02\x00\x7b\x08\x25\x41\x9a\x09\x6b\xeb\xe2\x6b\x27\x06\xdc\x7e\xb8\x6b\x57\xd7\x4e\x1c\x1b\xb8\xd3\x4e\x9e\x81\x77\xfb\xb1\x23\x3e\xaf\x44\x90\x34\x07\x72\xf1\x27\x25\x86\xad\x57\x28\x07\x9f\x58\xdb\x35\x84\x92\x3d\x38\xaa\x1d\x79\x32\xc9\xc3\x03\xb2\x20\x53\xd0\x80\x5d\xfd\x8b\x54\x98\xc3\x03\x39\x21\x02\xbe\xb4\x8d\xce\x05\x04\x1b\x72\x01\x1c\x29\x5b\x18\xfe\xdc\x53\xf6\x10\x6c\x64\xa9\x31\x50\x6b\xee\xee\x61\x13\xc8\x19\xd4\x62\xc2\x86\xae\x01\x4d\x0e\x15\xee\xc0\x91\xf0\x80\xc6\x0c\xa8\xc5\x29\x7e\x0e\xff\xb0\x8e\x80\xcd\xda\x2e\xa1\x0c\xa1\xf6\xcb\xc5\xa2\xe0\xd0\xe1\x59\xd9\xaa\x6a\x0c\x87\xdd\x22\xa2\x92\x57\x4d\xb0\xce\x2f\x72\xda\x90\x5e\x78\x2e\x32\x74\xaa\xe4\x40\x2a\x34\x8e\x16\x58\x73\x16\x05\x37\x11\xce\xf3\x2a\xff\xca\xb5\xe0\xf7\xaf\x06\x92\x86\x9d\x38\xdd\x07\xc7\xa6\xe8\x87\x23\xbe\x26\xed\x2e\x38\x13\xef\x62\xbb\x2c\xc9\xbf\x37\xaf\x0c\x89\x55\xee\xbf\x7d\x78\x84\x8e\x69\x74\xc1\xa1\xcd\xa3\xb5\xf7\xcb\xfc\xde\xf0\x62\x28\x36\x6b\x72\xc9\x71\x6b\x67\xab\x48\x91\x4c\x5e\x5b\x36\xa1\x45\x12\x93\x39\x34\xba\x6f\x56\x15\x07\xf1\xf4\xbf\x1b\xf2\x41\xfc\x33\x87\x37\x31\xaa\x61\x45\xd0\xd4\x39\x06\xca\xe7\x70\x67\xe0\x0d\x56\xa4\xdf\xa0\xa7\xdf\xdc\xec\x62\x61\x9f\x89\x49\xcf\x1b\x7e\x98\x8c\x0e\x27\x26\x6b\xf5\xc3\x5d\xba\x18\xf5\x50\x1b\x81\x0f\x35\xa9\x83\xc8\xc8\xc9\xb3\x13\xf4\x06\x0c\x24\x98\x1f\x66\x90\xe9\x58\x8c\xf1\xa8\x9c\x24\x86\xc3\xc1\x09\x25\xe4\x29\xc8\xd0\x06\x7f\xb0\x45\xc1\xa6\x38\x5e\x35\xc5\x04\x52\xc6\x5d\x73\x31\x12\xfb\x67\x59\x46\xfb\x59\xc3\xc1\xca\xa7\xef\xdf\x3c\x7c\x6b\x36\xec\xac\xa9\xc8\x84\x2f\x23\x74\x4f\xc5\x97\x09\xf3\x48\x06\x5f\x20\x84\xe9\x52\xf1\x2f\x5c\x39\x8a\x18\x79\xba\x60\xbc\xcb\x8f\x29\x1e\xe0\xa7\xab\x92\x77\x6f\xbb\x24\x7e\x2b\xf5\x77\xb0\x3c\xe5\x54\x1a\x54\x8f\x0b\x44\x1b\x07\x72\xc0\xd0\xf8\xb3\x50\x8e\xb3\x0e\xc0\x6c\x57\x5e\xb2\xc5\x8b\xd0\xac\xac\xc9\x79\x50\xe6\xa7\x98\xf7\xd3\xda\x74\x47\x21\xf2\xe9\x86\x81\x8d\x0f\x68\x14\xf9\xf9\x11\x19\x0e\x54\x8d\x80\xfb\x80\xfa\xd5\x9e\xce\x3e\x07\xa6\x32\x24\x9a\xc5\x22\x75\x50\x98\x5e\xf9\xa4\xeb\x31\x33\x79\x06\xa2\xa2\x23\x59\xd3\x37\x44\x50\x91\x2a\xd1\xb0\xaf\x62\x51\x32\x39\xe5\x52\xb7\x24\x1f\x7a\xca\x61\x5b\x92\x11\x83\x8e\x10\xcd\x29\x20\x6b\xdf\x0b\xb1\x17\x4b\x78\x48\x52\x45\xa8\x1d\x5b\xc7\xf0\x64\xec\xd6\x80\x75\xb0\x8d\x25\x33\x7e\xab\x6b\xbd\x1b\x03\xae\x05\xd4\x7a\x6f\xbb\x48\x1e\x0a\xde\x90\x01\x29\x2d\x73\xf8\x68\x86\xfa\xb4\x55\x78\x45\x80\x79\x1e\xbb\x92\x11\x8a\xf4\xa9\xd6\xac\x38\xe8\x5d\x2a\xd7\xbb\x81\xef\x21\x94\x18\x44\x59\xe7\x63\x19\x56\xb6\xaa\xad\x89\xd6\x56\xd1\x58\x2b\xdb\x84\x11\xb2\x0e\x43\x19\x4b\x10\x9a\x58\x51\xd8\xa5\xca\x66\x3d\x1d\x50\x8f\xb6\x8c\xe5\x4a\x92\x6b\x2c\x56\x56\x56\x8e\x90\x1c\xd8\xd0\xcf\xe1\xbd\x51\xd4\x62\x3a\xbf\x8e\xa0\xae\x08\x8d\x30\x89\x26\xd9\xe3\x43\xa1\x81\x54\xc3\x46\x68\x8a\x73\x0b\xca\x01\xdd\x8a\x83\x43\xc7\x7a\x07\x19\xb0\xcc\x56\xb6\x22\x0f\x35\xba\xd0\x05\xec\xed\x87\xbb\xd4\x8b\x94\x98\xc2\xc8\x63\x35\x46\x74\x85\xea\x69\x8b\x2e\xf7\x59\x9c\xbd\xb6\x2e\xbd\x89\xed\x30\xf0\x8a\x35\x87\x68\x6a\x45\xce\xb4\x08\xd9\x25\xb5\x23\xbf\x31\xdd\x7b\x09\xe6\x57\x27\x9f\x9f\x2b\x09\x00\x1a\x7d\x78\x74\x68\x3c\x77\xcd\xf5\xd8\x2c\x10\x31\x2b\x0c\x4b\x90\x32\x9f\x05\x1e\xd5\xec\x4c\xce\x95\x1a\xec\x3d\x16\x13\x1c\xce\xac\x75\x84\x7e\xbc\x52\x4c\xa5\x96\xfb\xb8\x42\xf2\xcb\x51\x70\x22\x58\x43\xd9\xd6\xba\xfc\x7a\xdf\xac\x8c\x12\x86\xa3\xce\xb6\x4f\xce\x18\xa8\xb0\x6e\x27\xef\x0a\x1b\x4f\xfd\x87\xc6\x39\x32\xa1\xcd\xbd\x63\xe9\x44\x9e\xbb\x30\x22\x55\x4c\x19\x6c\xa2\xe7\x59\x28\x36\xa1\x6e\xc2\x35\xf8\x46\x95\x80\x3e\xca\xac\xd9\x4c\x09\x2a\x1b\x28\x15\x34\x14\x92\x49\xdb\xa5\x82\x2f\x36\xe0\x9b\xaa\x42\xc7\x9f\x23\xfc\x55\x12\xb1\xcd\x0e\x51\xf8\x09\x39\xcf\x38\xe4\xb4\xbc\x5c\xbc\x34\x7e\x3e\xef\xc9\x7d\x1a\x7f\xdc\xd5\xd4\x15\x4c\x59\xdc\x9b\xbb\x8f\xe3\xa8\xaa\x9f\xb0\x4d\xd8\xd5\xac\x50\xeb\x9d\x84\x7e\xe7\xf0\x1c\x04\x01\x92\x58\x7d\x69\x5d\x80\xba\x74\xb1\x67\x1d\xa6\xc8\xc8\x6c\x8a\x6a\x9b\x3d\xd9\xe4\x2c\x78\x68\xab\x25\xc7\x94\x0f\x1f\xaf\x70\x65\x24\x66\x74\x16\x5c\x43\x1f\xaf\xa0\xb6\x1a\x1d\x87\xdd\x1c\xbe\xb3\x63\x09\x4c\x1e\xfa\x84\x55\xad\xe9\x1a\xf8\x58\xbf\x8e\x8b\x4f\x55\x05\x85\x1c\xab\x5d\xc2\x51\xdc\x4b\x5e\x4f\x29\x1f\xa5\x61\x9f\x76\x9c\x1f\xaf\x40\xa1\x8f\xc6\xac\x9d\x5d\xe1\x4a\xef\xe2\x0c\x91\xf5\x1a\xbc\x3d\x64\xfb\xbc\xe6\x2b\x09\x04\xad\x29\x87\x8f\x57\x77\xa6\x25\x3f\x92\x81\xe0\x1c\x22\x52\x09\xa0\x93\xd6\x09\x20\x6b\x61\x36\xf2\x41\x28\x9e\x0c\x4f\xf6\x67\xdd\x27\x74\x0e\x0f\x4b\x67\x77\xf4\x31\xd1\x18\x4f\x0a\x2e\xbb\x4c\xe3\x59\xd3\x5f\xd9\x07\xeb\x76\x67\xfa\xbe\xc3\xc9\xa9\xeb\xd1\xec\xfb\xf2\x21\x89\x18\x6e\xbe\x49\x64\x15\x6b\x02\x54\xb1\x98\x9d\xea\x28\x35\x77\x4b\x8e\x20\xe0\x13\x1d\xfb\x68\xa2\x3f\x7a\xbe\x0a\x24\x56\x17\x84\xe4\x63\xd9\xc9\xd5\x8a\x21\xb5\x6e\x44\x8a\x33\xb6\x4b\x8f\x54\x3c\x4d\xbf\x43\xdd\xb9\xb8\x76\x88\x7a\x69\x32\x6c\xcb\x14\x18\xad\xb6\x5f\xa0\xe8\x73\xe8\x4e\xd4\x47\x3e\x1c\xda\x66\x64\x42\x12\xf3\xd7\x08\x00\x47\x5a\xf6\xee\xef\xe3\x9a\x13\x78\x5c\xd2\x6f\xa7\xa5\xf7\xb4\xa6\x58\x59\xe2\xd9\x1e\xb2\xf1\x40\xc6\x36\x45\x19\x4f\x00\xc4\x8d\x11\x36\x16\x34\x05\xd8\xd9\x66\xc4\x92\x6c\x64\xf7\x1d\x24\x29\x57\x36\xe7\x75\x72\x81\xa3\xb6\x01\x6c\x4f\x91\x7e\x29\xb8\x47\x0f\xbc\x26\x54\xb9\xfd\x70\xd7\x1d\x73\x75\x91\xe9\x92\x5e\x23\x7c\xe1\x3c\xf4\xd6\x4c\x3a\xff\x80\xa1\xbc\x80\xf7\xab\xbb\x75\xab\x6b\xec\x86\xad\x64\x79\x26\x45\x07\x5b\x95\xd8\xdd\x13\xca\x3e\x71\x22\x3b\xcb\x0e\xc5\x04\x76\xd4\xae\xb8\x4e\x47\x3d\xed\x89\xd2\xfe\xe4\x4d\x5c\x04\x98\xca\x02\xfc\xed\xe1\xfd\xbb\xc5\xf7\x76\x82\x64\xd4\x02\x50\x29\xf2\xed\x4e\x49\xb6\xfe\xfb\xde\xa4\x3d\xfe\x78\x88\x7b\xa8\x0a\x0d\xaf\xc9\x87\x79\xcb\x83\x9c\xff\xe9\xf5\xcf\x53\xbd\xd0\x77\xd6\x9d\x94\xbd\xfe\xfc\xaa\x03\x14\xfb\x64\x8e\x9e\x22\x6c\x39\x94\x3c\x55\x9f\xa4\x3a\xe6\xad\xda\x69\xd7\x24\xa1\x0b\xb6\x55\xb7\x21\xd0\xfc\x44\x4b\xb8\x12\xb4\x0d\xc4\xfc\x8f\xc1\x8a\xfe\x3b\x5e\xc0\x00\xfe\xb0\x2d\x25\xe9\x5e\xc9\xa4\xab\x24\x5c\x7f\x4c\x29\x63\x83\xa6\xa4\x15\x32\xa6\xc8\xe0\xb8\x28\xc8\x8d\x6e\xaf\xa0\xad\xcf\xb4\x21\x13\xbe\x16\xd8\xf3\x1a\x8c\x1d\x90\x88\x84\xc5\x7b\x35\x29\x5e\x33\xe5\x27\x42\xff\xf4\xfa\xe7\x49\x89\x0f\xed\x25\x4d\x04\x7d\x82\xd7\xa9\x47\x91\x0e\xc0\xe6\x5f\xcf\xe1\x31\xa2\x63\x67\x02\x7e\x12\x4e\x4a\x76\x61\x53\x96\xb5\x46\xfa\x05\x0b\x25\x6e\x08\xbc\xad\x08\xb6\xa4\x75\xd6\xee\xb4\x60\x8b\xb1\x25\xee\x1c\x27\x78\xc3\x6e\x93\x34\x8d\xd6\xee\x70\xf8\xf1\xfd\xdb\xf7\xcb\x24\x99\x00\xaa\x88\xbd\xbb\x6c\xce\xd6\x6c\x50\xc7\x16\x2f\x1d\x5a\x46\x34\x4e\x76\x7b\xbe\x49\xf0\x09\xb6\xdd\xc0\x75\x3d\xd9\xba\x09\x8d\xa3\xf9\xab\x97\xc4\xf1\xf1\xb9\xed\xfe\x19\x39\xc1\x3d\x4e\x1c\xff\xa7\x73\xd0\x8b\x95\x33\x23\xe7\x8f\x63\xca\xbd\x1b\xa0\xfc\x59\xe5\x64\x33\xe2\x0c\x05\x8a\xfa\xe5\x56\xf9\x45\xdc\xd0\xd6\xc1\x2f\xec\x86\xdc\x86\x69\xbb\xd8\x5a\xf7\xc4\xa6\xc8\x04\x9a\x59\xc2\x80\x5f\xc4\x83\xb9\xc5\x57\xf1\xcf\x8b\x75\x99\x3c\xd9\x1b\x53\x28\x4e\xfe\x3d\xb4\x12\x3e\x7e\xf1\x22\xa5\xba\xe3\xc1\xcb\xeb\xd8\xab\x87\x94\x30\xd4\xf1\x5a\x09\x8b\x6d\xc9\xaa\xec\x6e\x6e\xda\x1c\x3b\x11\x4c\xec\xa1\xc2\x3c\xa5\x66\x34\xbb\xdf\x1c\xca\x62\xd0\xb4\x41\xdd\x65\xed\x15\x61\x86\x26\xcf\x62\xfb\xec\x83\x8c\xbf\xc8\x82\x0d\x5f\x14\xbe\xff\xbc\x7b\xfb\xfb\x00\xbc\xe1\x17\xc5\xea\x2f\x6e\xef\x46\x16\x1c\x0d\xed\xaf\x5a\x6f\x50\xd7\x25\xde\xec\xc7\x62\x3b\x95\xb5\x17\xac\x83\xcf\x00\xe9\x80\x78\x09\xb2\xb7\x4d\x03\xc1\x3a\x2c\xa8\x1d\xd9\x9f\x0d\x48\xcf\x50\x07\xca\xdf\x1d\x5f\xb1\x5e\xa5\xa2\xd5\xdd\xa1\xc6\xd7\xc1\xe1\x31\xfc\xf4\xf3\x2c\x51\xa5\xfc\xc7\x4e\x1a\x19\x3c\xbe\xa0\xee\xaf\xbb\x1f\xc4\xea\x8a\x6e\x95\xb2\x8d\x09\x23\x77\xde\xe8\xec\xfe\x9e\x7b\x98\x28\x2e\xbf\x0d\x77\x2b\x54\x73\x6c\x42\x69\x1d\x7f\x8e\xfd\xec\xfe\x4a\x7c\x76\x70\x75\x7c\x6f\xf5\xe1\xbd\xfb\xb9\xfb\xeb\x0a\x0d\x16\xe4\x32\x27\x0b\x5d\xa3\xc5\x5a\x99\xf4\xae\xdf\x3b\xdb\xd4\xad\x27\xa2\x95\xfa\xcb\xc6\x34\x96\xae\x75\x2a\xac\x7d\xef\xac\x94\x6a\xe2\xab\x4f\x56\xc1\x64\x15\x9f\x7c\xbb\xea\x96\x8a\x4c\x14\x7f\x16\x14\xe2\xdf\x74\x91\x77\x19\x67\x4f\xca\xd1\x05\x34\x75\xba\x30\xcf\xa0\xc6\xa0\xca\x21\x1b\xf9\xb9\x8d\x83\x27\x0c\xb1\xae\xfd\x7c\x73\x73\xca\x35\x47\xaa\xac\xf1\x97\x30\x3e\xe1\x77\xca\xe5\xf4\x9f\x01\x8e\xcd\xbb\xbf\xdc\x1f\x67\x96\x93\xa6\x5f\x45\xe1\x8b\x45\x59\xf4\x67\x22\x03\x89\xa6\x75\xfe\x22\x0c\xff\x85\x4d\x2e\x89\xe8\x34\x9c\x86\x80\x5d\xb5\xb3\xe4\xf7\x3d\xad\x63\xe0\xb7\xda\x3d\xc3\x70\x76\xf4\xef\x16\x31\x66\x46\xe3\xa1\x6d\xe9\x62\x48\x8c\x46\xfb\x17\xc4\xf8\xff\x02\x00\x00\xff\xff\x19\xa9\x97\x77\x3c\x23\x00\x00")

func resourcesYamlBytes() ([]byte, error) {
	return bindataRead(
		_resourcesYaml,
		"resources.yaml",
	)
}

func resourcesYaml() (*asset, error) {
	bytes, err := resourcesYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "resources.yaml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
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
	"resources.yaml": resourcesYaml,
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
	"resources.yaml": {resourcesYaml, map[string]*bintree{}},
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
