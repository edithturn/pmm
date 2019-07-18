// Code generated by go-bindata. DO NOT EDIT.
// sources:
// migrations/sql/01_init.down.sql (20B)
// migrations/sql/01_init.up.sql (9.984kB)
// migrations/sql/02_postgresql_columns.down.sql (979B)
// migrations/sql/02_postgresql_columns.up.sql (2.04kB)
// migrations/sql/03_add_agent_type.down.sql (233B)
// migrations/sql/03_add_agent_type.up.sql (270B)

package migrations

import (
	"bytes"
	"compress/gzip"
	"crypto/sha256"
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
	bytes  []byte
	info   os.FileInfo
	digest [sha256.Size]byte
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

var __01_initDownSql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x72\x09\xf2\x0f\x50\x28\x49\x4c\xca\x49\x55\xc8\x4d\x2d\x29\xca\x4c\x2e\xb6\xe6\x02\x04\x00\x00\xff\xff\x6b\xf9\xb4\xa3\x14\x00\x00\x00")

func _01_initDownSqlBytes() ([]byte, error) {
	return bindataRead(
		__01_initDownSql,
		"01_init.down.sql",
	)
}

func _01_initDownSql() (*asset, error) {
	bytes, err := _01_initDownSqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "01_init.down.sql", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info, digest: [32]uint8{0xf5, 0xb9, 0x22, 0xf, 0xca, 0x89, 0xa5, 0x5, 0xf2, 0xed, 0x48, 0xd7, 0xe0, 0x9d, 0xec, 0x2, 0xab, 0x87, 0xf3, 0x9d, 0x3d, 0x5a, 0x71, 0x39, 0xe, 0x4a, 0x88, 0xde, 0x1d, 0x0, 0x1d, 0x4e}}
	return a, nil
}

var __01_initUpSql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xa4\x5a\x4f\x73\xdb\x38\xf2\xbd\xfb\x53\xf4\x2d\x4e\x95\x3d\xf5\x9b\x64\xea\x57\xc9\xa4\x72\x90\x6d\x25\xa3\x2d\x5b\xf6\x48\x9a\x9d\xf5\x89\x86\xc8\x96\x88\x35\x09\x30\x00\x18\x59\xf9\xf4\x5b\xf8\x43\x12\x24\x01\x49\xf1\x9c\x12\xab\x5f\x3f\xf4\x6b\x34\x80\x06\xa4\xeb\xc5\x74\xb2\x9a\xc2\x6a\x72\x75\x3b\x85\x12\x95\xa0\xa9\x84\xf3\x33\x80\xcb\x4b\xb8\x23\x94\x41\x46\x4b\x64\x92\x72\x26\xcf\x00\x9e\xbe\xd5\x28\xf6\x34\x7b\x82\x5b\xbe\xbb\x26\x22\xa3\x8c\x14\x54\xed\xcf\x97\x4a\x50\xb6\x7d\x0b\xd7\xf7\x77\x77\xd3\xf9\x0a\xde\xe4\x44\xe6\xc0\x37\x60\x1c\x60\x43\xd9\x16\x45\x25\x28\x53\x6f\x2e\x34\x8f\x44\xf1\x1d\xc5\x71\x9a\xd9\x03\x70\x01\x39\x97\x8a\x91\x12\x35\xe1\xcd\x15\x58\x67\x4b\x94\x11\x45\xd6\x44\xe2\x71\xaa\x07\x2e\xd5\x56\xe0\xf2\xcf\xdb\xdf\xa1\xf1\x72\xc1\xa4\x39\x96\xe4\x38\xc3\xdd\xbe\xe7\xfc\x09\x7c\x4a\x4b\x62\x09\x6b\x89\x42\xc7\x7b\x9c\x32\x2d\x28\x32\x05\xda\x01\xb4\x87\xf5\xb7\x9f\x26\x5a\xf6\xc9\x14\xfd\x4c\x19\x9e\xcb\x4b\x80\xa5\x22\x2c\x23\x22\x83\x82\xac\xb1\x30\x73\x28\xb0\x2a\x68\x4a\x14\xe5\x2c\x91\x78\xc2\x08\x73\x97\x7a\xcf\x11\x24\xaa\x26\xd6\x5a\xaa\x53\xa6\xf2\xda\x02\x3d\x99\x7a\x1e\x69\x8a\x89\xda\x57\x27\xa4\x6a\xb5\xaf\x4c\x14\xce\xcb\x52\x20\xfb\x4e\x05\x67\x25\xb2\x13\x74\x4c\x3b\xb0\x17\x05\xf9\x71\xdc\x73\xf2\x9d\xd0\x82\xac\xa9\xb6\xc2\x0f\xce\x9c\xab\xc0\x2d\xe5\xec\xb8\xfb\xc2\xe0\xbc\x31\x19\xcf\x30\x29\x79\x86\xc5\x09\xd9\xe7\x19\x82\xc1\xba\x84\x73\xa6\x08\x65\x28\x92\xd3\x4a\xec\xba\xc1\x83\x5f\x18\xd7\xb5\x54\xbc\xf4\xaa\xc2\xfe\xef\x97\x67\xdc\x3f\xc1\x44\x08\xb2\x3f\x0f\x13\xfb\xcc\x3e\x87\x61\x97\x36\x44\xc7\xf5\x9d\x14\x35\xbe\x92\xcd\xf8\x3a\x3a\xb2\xd5\xab\xe1\x94\x5d\x67\x96\x21\x53\x74\x43\x51\xe8\x4a\x31\x7e\xa0\x72\xa2\x20\xe5\x45\x81\xa9\x02\xc2\x32\x90\xc8\xb2\x66\xab\xf3\x07\xb0\x65\x38\x65\x75\xf9\x41\x6f\x80\x00\x6f\xba\xcf\x13\xca\xbe\x93\x82\x66\x6f\xe0\x33\xfc\xdf\x85\xb5\x96\x7b\xf9\xad\xb8\xac\x50\x6c\xdc\xda\x87\xcf\xf0\x6b\xcf\x26\x0b\xbe\x2b\xf8\x56\x1b\xde\x35\x06\xce\xb6\x3c\x5b\x5f\x56\x82\x6f\x68\x81\x42\xdb\xde\x9f\x01\xf8\xc5\x66\xc2\x36\xe5\xde\x8b\x9d\x6f\x9a\xa8\x7f\x07\xc7\x7c\x01\x7a\x78\xb7\xf7\x5c\x00\xaa\xf4\x17\xab\xa8\x42\x41\x79\x96\x48\x45\x84\x7a\x82\x1b\xa2\x70\x45\x4b\xf4\x16\x93\xfe\x6b\x97\x23\x6b\xd8\x75\x7d\xf2\x0d\xac\xeb\xf4\x19\x15\x18\x3f\xcc\x7a\x5c\x05\xb2\xad\xca\x9f\xe0\xaf\x19\x53\xef\xdf\x75\x54\x37\xb5\x20\x8d\xbb\x47\x66\x99\x2c\x83\xb7\xfd\x1f\x9f\x43\x93\x3a\xc8\xe8\x16\xa5\x4a\x14\xbe\xa8\x4f\xee\x10\xd9\x51\x95\xf3\x5a\x99\xfd\xd7\xad\xfe\x17\x52\x56\x05\x3e\x81\x25\xe9\x38\xee\x19\x76\x87\x8f\x43\xc1\x46\xf0\x52\xef\x5b\xb0\xe1\x35\xcb\x80\xf6\x63\x74\xa8\x64\xc3\x45\x49\x54\xbf\x0e\xa6\xff\x99\xdc\x3d\xdc\x4e\x93\x2f\xf7\x8b\xbb\xc9\x2a\x99\xcd\xff\x3d\xb9\x9d\xdd\xf8\xb5\xe0\x10\x7e\x09\x7c\x99\xcd\xbf\x4e\x17\x0f\x8b\xd9\x7c\x65\x0a\xa0\x3f\xc9\x33\x96\xe9\xad\x14\x65\x7f\x8e\x05\x92\xa2\x1f\xb6\x04\x2a\xa1\x12\x3c\xa7\x6b\xda\x4e\x09\x95\x89\x12\x35\xd3\x04\x99\x9d\x91\x0f\x21\x6a\xba\x09\x70\x29\xce\xa1\xe0\x6c\x6b\xd6\xc2\x8e\x48\x68\x99\xfa\x89\x18\x2f\x87\x26\x0d\xab\xc7\x87\x69\x28\x09\x8b\xc9\xfc\xe6\xfe\xce\xcf\xc1\xf2\xf6\xfe\xef\xe9\x72\xe5\x2f\x80\x2f\x93\xe5\xca\x7d\xf4\xde\x7d\xf4\xf7\x6c\xf5\x47\x32\x5d\x2c\xee\x17\xfa\xd3\xdf\x62\x89\xda\xe9\x44\xf5\xa7\x54\x87\x5f\xd1\xf4\x19\x33\xa8\xab\x7e\xf8\x6e\xa9\x8c\x6b\xe3\xce\x35\x39\xa3\xfa\xa0\x0c\xfe\xb5\xbc\x9f\x83\xad\x00\xb7\x90\x58\x5d\x26\x1a\x46\x51\x26\xba\x00\x93\x1d\x11\x8c\xb2\xad\x7c\x82\x2f\x05\x27\xbd\x95\xf0\x07\xdf\x41\x49\xd8\x1e\x9c\x83\x09\x4f\x3b\x41\xe3\x34\x2c\xba\xe6\xf3\x5f\x52\x9e\xb5\x3b\xa5\x5d\x60\x5e\x0a\x6e\xa9\x34\xab\xbf\x41\x8f\x7c\x6b\xbd\xae\xac\xb3\x0b\xaa\xb7\xff\xd7\xcc\xb8\x23\x49\xe3\x91\x8c\x74\xa2\x10\x5c\xfc\x9c\x4a\xe3\x32\x5a\x57\x86\x67\x24\xf0\xff\x7f\x0b\x08\xbc\x25\x52\xe9\x81\x19\x1f\xf8\x7a\x02\x47\xce\x7d\x7d\x1d\xc5\x01\x85\x01\x55\x93\xd2\xd0\x34\x9a\x28\x03\x95\x53\xe9\x13\xe8\x86\xd8\x96\x8e\x66\xb3\x5c\xfb\x44\xd1\x12\x93\x54\x87\x37\xa2\x5c\xe5\xa8\x37\x52\x85\xa6\xe1\xc0\x17\x4c\x6b\xb3\x33\x6a\x17\x3d\x80\xc4\x94\xb3\xcc\xe6\xaf\xc4\xa6\xe0\x7a\xc4\xb2\x2e\x5f\x4f\x1c\x22\x2c\x29\x0b\x10\x2e\x4b\x52\x14\x28\x95\x3d\x75\xdb\x95\x91\x34\x84\x7e\x1a\xfb\x74\xe4\x25\x40\x77\x45\xb7\xdb\xd7\xb0\x55\x1f\x3f\x06\xd8\x3e\x7e\xd4\x87\x5c\xaa\x0f\xf6\xc2\xb0\x9d\x44\x5b\xf0\xf4\x79\x3c\x39\x23\x63\x3c\xc1\x86\x56\x71\x20\xe9\xb7\x9a\x0a\x04\xed\x23\x43\xd9\xed\xc8\xfc\xe4\x8e\x8d\x5e\xaa\x46\x46\x5f\xb9\x33\x0a\xbe\x93\x89\xd4\xfd\x47\x40\x40\x67\x8c\x0b\x60\x75\xb9\xb6\x7d\x90\x46\xeb\xae\x47\x69\x3d\x2a\x47\xb0\x57\x86\x56\x41\xc7\x16\x50\xe0\x19\xc7\x0a\x3a\x63\x4c\x81\xde\x5c\x29\xc3\x2c\xaa\xa2\x05\x84\x95\xcc\x07\x2a\x52\xc2\x18\x66\x70\x09\xcb\xe9\xed\xf4\x7a\xd5\xd7\xd0\x72\xc5\x74\x74\x80\x88\x96\x16\x10\xd3\x43\x36\x1b\x4c\xd5\x01\x3d\x2d\xe0\x24\x3d\x69\x4e\xd8\xd6\xe8\xf9\xeb\xe1\x66\xb2\x9a\x5e\xc0\xcd\xf4\x76\xaa\xff\x9d\xcd\x97\xd3\xc5\x40\x5f\xcb\x1d\xd3\xd7\x01\x22\xfa\x5a\x40\x4c\x9f\x40\x12\xd7\x66\x8c\x27\x57\x9c\x46\xdb\xa6\x4b\x91\x75\x81\xb2\xaf\xc5\x70\xc5\x74\x58\x63\x44\x83\x31\x06\xe2\x2f\x51\x6c\x31\xa9\x88\x94\x28\x43\x12\x7a\xf6\x53\x54\x18\x07\xb0\x0e\xb6\x4d\xd3\x8b\x47\x72\xa1\x80\x14\x5b\x2e\xa8\xca\x4b\xc8\x89\x84\x9c\x64\x7a\x69\x65\xbc\x95\xd8\x1b\x2b\xa0\xb2\x6f\x1f\x0b\xed\xd9\x03\x5a\x29\x63\x3c\x5b\x27\x94\x27\x22\xe1\x55\x50\xee\x10\x12\x56\x6c\x4e\x4f\x69\x84\x75\xc2\x2b\xb2\x45\x3b\x7d\xbc\x42\xdb\xdd\x4b\x73\xc9\xc8\xea\x02\xb3\x56\xe4\x70\x84\x80\xce\x11\x64\x2c\x75\x08\x39\xa2\x76\xbd\x57\xe1\xe9\x1d\x83\xc2\x8a\x97\xb4\xa4\x05\x11\x7a\xc2\x9c\xcb\xec\xde\x0e\x7d\x01\xeb\xda\xce\x71\xcd\xa8\xd2\xbd\xb2\xe1\x09\xea\xb5\x23\x1c\x51\xec\x40\x87\x35\x5b\xd0\x11\xd5\x3b\x42\x83\x07\xc1\x08\x13\xd1\x9c\xeb\x05\x99\xf3\x9d\xed\xfc\xcf\xbb\x63\xec\x2d\x50\x7d\x2c\xf0\x67\x98\x31\xc6\x6f\xae\xec\x91\xa7\x6a\x52\x14\x7b\x5b\x02\x3a\x1f\xfa\xca\xe5\xee\x4f\x8a\x0b\xb2\xc5\x60\x4e\x4c\x00\x47\x52\x62\x31\x87\x33\x62\x30\xf1\x84\x08\x4c\xed\xd1\x79\x24\x29\x7d\xdc\xcf\x27\x46\x0b\x77\x37\x4f\xa2\xaf\x5e\xfa\x62\xa0\x77\x36\xdb\x0a\x0c\x33\xd0\x1f\x2d\x9e\x85\x01\x2e\x9a\x89\x3e\x2e\x9e\x8d\x6f\x35\xd6\x78\x2c\x15\x1e\xe8\x9f\xe4\x41\x56\xa6\xe1\xa4\x2a\x47\x61\x92\xa2\x2f\x56\x8a\x03\x32\x85\xc2\xe0\x5c\x11\x99\xf1\xc0\x5c\x06\x24\xcd\xdc\x3b\x86\xfd\xb0\x71\xd3\xc9\x6c\x5b\xd7\x61\x32\xbd\x78\xe3\x99\xf4\x41\xd1\x34\x7a\xa0\x78\x0e\xf5\x7e\x27\x93\x8c\x4a\x45\x59\x7a\x28\x8f\x03\xe0\xc1\x2d\x95\x54\x95\xe0\x2f\xb4\x24\x0a\x8b\xfd\x60\x83\xad\x19\xfd\x56\xa3\xd9\x67\xa5\x97\x5f\x92\xa6\x28\xe5\x78\x87\x1d\x0c\x1b\x4f\xc9\x10\x18\x4d\xcb\x00\x18\x48\x8d\x6d\xae\xed\xa3\x4f\x28\x23\x3d\xfb\x49\x45\xd5\xe9\xa4\xc3\xbb\x89\xa3\x09\x08\xeb\xdb\xc7\x7a\x7a\xf6\x80\x0c\x77\x12\x44\xfa\x68\xcf\x7a\x4a\x43\x60\xe0\x6d\x27\x4d\x8a\xc2\x75\xd2\x9d\x1c\x8f\x30\x20\xc6\xb7\x8e\xa5\x78\xd6\x80\x10\x55\x56\x89\xed\xa3\x42\x42\x3c\xeb\xb1\xbe\x53\x61\x59\x71\x41\xc4\xde\xb5\x65\x90\x0a\x24\x7a\x7b\xe3\x0c\x4a\x2c\xb9\xd8\x9b\xb5\xd9\xce\x56\xab\xcd\x1b\x23\xa0\xcd\xb7\x8e\xb5\x79\xd6\x88\xb6\x8c\xca\xe7\x23\x02\x7d\xc8\x3f\x51\xa9\x79\x0e\x68\xf4\x87\x89\x08\xed\x41\xc2\x6a\x7d\xc8\xa1\xe9\x4c\x24\xfd\x71\x64\x4e\x1d\x24\x52\xa1\x5c\x91\x02\x96\xf4\x87\xbd\x04\x9b\x0a\xd5\xda\x74\x75\x8e\x92\x50\x4b\xcc\xec\xdb\x46\x74\x72\xdd\x60\x87\x66\xb8\x81\x1c\x98\x66\x07\x19\x0a\xbf\xbc\x84\x2b\xce\x0b\x24\xac\x79\xca\x76\xab\x38\x4d\xf2\xf0\x01\xe6\x2c\x61\xed\x7f\x9a\xdd\xe4\x9a\xa4\x39\x42\x4e\xbd\x65\xb8\xa9\x8b\x22\xd1\x77\xc5\x10\x65\x67\x8c\xaf\x79\xbb\x4f\x55\x28\x36\x5c\x94\x98\x01\x01\xed\x65\xd3\x68\x2e\xa1\xfd\xa1\xfe\xcb\x69\x7c\x28\x63\xfc\xe9\xa1\xb4\x17\x9c\x13\xfb\x6f\xf3\xf4\x4d\x59\x86\x2f\x28\xdf\x06\xa6\xed\x70\x01\x1d\x19\xbd\x59\x1c\x84\x01\x2d\xab\x82\xa6\xba\xfd\xd5\x87\x3a\x23\xa3\x2a\x0a\x8c\xcd\x99\x2d\xf7\x83\x31\x34\xa0\x23\xb1\xc8\xe1\x78\xe6\x7d\x4c\xb7\x9d\xdd\xe2\xed\x92\x4f\x0b\xd4\xf7\xb1\x60\xee\x1b\xdb\x31\xf1\x66\x51\x10\x68\xf0\x63\xf2\x03\xf2\x46\x98\xf8\x60\x0d\xd4\x3e\x5a\xb7\xd3\x3d\x94\x24\xb1\xc0\x54\xd9\xca\x11\x84\x6d\x31\x5a\x5c\x11\xe4\x29\x07\x99\x06\xba\x2b\xad\x53\x6f\x08\x40\x22\x11\x69\xae\x63\x22\x20\x70\x83\x02\x59\x8a\x83\x69\x77\xc3\xda\x11\xe3\x61\x59\xfb\x6b\x82\x31\x9e\x52\x07\xa1\x4c\xda\x84\x54\x07\x43\xc8\x31\x0d\xce\x4d\x00\x75\x7a\x38\xcd\x8a\x7b\xc6\x7d\xf3\x15\x8d\x66\x30\x5b\xeb\x33\xea\xa2\xd1\x77\x64\xb2\xd1\x8d\xaf\x79\x78\x16\x7c\xd7\xc5\xa7\x2b\x22\x9e\xa0\xce\x7a\x4a\x3c\x1a\xee\x42\xd8\xa1\x40\xc8\x38\x43\xa8\xa5\xee\xa1\x6d\xa6\x06\xc3\xf2\x5d\xf0\x34\xe9\x8c\xa7\x0e\xaa\xa7\x82\xef\x06\xf4\xb1\x4d\xb5\x33\xbe\x5a\xd3\x7a\x6f\x9f\xf7\xa8\xeb\x15\xfb\x73\xce\x78\x62\x76\xbf\x44\xd7\x48\x28\x82\x3e\xe0\x94\x28\xda\x2f\x30\xfc\xdd\xd5\x1f\x70\xcb\x79\x76\x7c\xd4\x21\xea\x35\x43\x6b\x0e\x7f\xfc\xcb\x4b\x30\xdf\x14\xf7\x8f\xc9\x8c\xa7\x32\x11\xa8\x6a\x11\x79\x51\xed\x03\x4e\x7a\xa9\x73\x60\xc8\x78\x5a\x97\xbd\x5e\xb6\x4f\x16\x68\x08\x06\x80\x71\x3b\xd0\x07\x84\x9e\x1c\x51\x56\x9c\x49\x3c\x70\xcf\x18\x42\xe2\xa2\x1a\x24\x58\xa4\xe9\x02\xdb\x4d\x5e\xa0\xac\x0b\xd5\xb6\x48\xdd\x5b\xe4\x80\x3e\xf4\x22\x39\x84\x04\xde\x25\x07\x90\x80\x54\x93\x0b\xf7\x7c\x1d\x9d\xba\xc6\x7e\xd2\x1a\x72\x4f\xe1\x91\x89\x6b\xa8\x62\xf3\xd6\xda\x23\xd3\xd6\xd8\x7d\x29\x67\x6f\x61\x3a\xff\x3a\x9b\x4f\xe1\x33\xdc\xa1\xd8\xe2\x4a\x20\xc2\xc3\x64\xb1\x9a\xad\x66\xf7\x73\xb8\x7a\x04\xc5\x1f\x1f\x1f\x1f\xef\xee\x6e\x6e\xce\xfd\xdf\x20\xbc\x3d\xbb\x5f\xdc\x4c\x17\x70\xf5\x78\x06\x60\xbf\x4e\x76\x3f\x25\xbb\x30\x7f\xd8\x9f\x74\xd9\xff\x37\xbf\xaf\x72\x16\xfb\xdb\x06\xf3\xff\xe6\x77\x55\xf6\x2f\xef\x57\x52\xf6\x03\x7f\x40\xf3\x25\xf2\x72\xba\x5a\xcd\xe6\x5f\x97\x76\x65\x25\x5b\x41\x58\x5d\x10\x41\xd5\x1e\x3e\xc3\x87\x5f\x3f\xbe\xfb\x74\xf6\xbf\x00\x00\x00\xff\xff\xfd\x1e\xfe\x14\x00\x27\x00\x00")

func _01_initUpSqlBytes() ([]byte, error) {
	return bindataRead(
		__01_initUpSql,
		"01_init.up.sql",
	)
}

func _01_initUpSql() (*asset, error) {
	bytes, err := _01_initUpSqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "01_init.up.sql", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info, digest: [32]uint8{0xc9, 0xbf, 0x86, 0x4c, 0xbb, 0x3d, 0xec, 0x1b, 0x20, 0xd2, 0xce, 0xc4, 0x2e, 0xcb, 0x96, 0x54, 0x41, 0xf7, 0xae, 0xb, 0x5c, 0xc1, 0x4b, 0x7b, 0x8b, 0x61, 0x98, 0x30, 0x35, 0x79, 0x4b, 0x55}}
	return a, nil
}

var __02_postgresql_columnsDownSql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x8c\xd0\x4d\x0e\x82\x30\x10\xc5\xf1\xbd\xa7\xe8\x01\xbc\x81\x2b\x54\x76\x28\x86\xe0\xba\x7c\x4d\x42\x43\x8b\xa6\x1d\xe3\xf5\x0d\x1a\x12\xad\x94\xd7\xfd\x3f\xbf\x69\x5f\x92\x95\x69\x21\xca\x64\x9f\xa5\xc2\x10\x5b\xd5\xba\x8d\x10\xc7\x22\xbf\x88\x43\x9e\x5d\x4f\x67\x51\x19\xe9\xfa\xda\x52\x27\x1b\x3d\x38\xd9\x2b\x96\xed\xc8\xd5\x16\x67\xee\x61\x50\x66\xa9\xee\x62\xb8\x77\x17\xe1\x75\xca\xb2\xa2\x28\x72\x4e\x23\xd4\xa7\x55\xcc\x34\xc6\xa8\x73\x1a\x50\xf5\xad\xad\x35\x1c\xd2\xab\xb0\xb5\x36\xa3\x9f\x61\x0d\x8c\xb8\x50\x62\x13\x4c\xb8\x50\x06\x4c\x26\x73\xc7\x9f\xf6\x2a\x68\x81\xe7\xfd\x87\x01\xb1\xd1\xc3\xe7\x22\x2b\x43\x21\xed\x37\x5a\x91\xa6\x63\x04\xa9\xaf\x6a\xb2\x76\xaf\x00\x00\x00\xff\xff\x92\x15\x3e\x17\xd3\x03\x00\x00")

func _02_postgresql_columnsDownSqlBytes() ([]byte, error) {
	return bindataRead(
		__02_postgresql_columnsDownSql,
		"02_postgresql_columns.down.sql",
	)
}

func _02_postgresql_columnsDownSql() (*asset, error) {
	bytes, err := _02_postgresql_columnsDownSqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "02_postgresql_columns.down.sql", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info, digest: [32]uint8{0xba, 0xff, 0x8a, 0x34, 0x78, 0x29, 0x20, 0x64, 0xd4, 0x9c, 0x10, 0x61, 0x92, 0xf3, 0x20, 0xc, 0x39, 0xb0, 0xf0, 0x75, 0x33, 0x2e, 0x61, 0x44, 0x22, 0x32, 0x16, 0x53, 0x75, 0xb5, 0x32, 0xd4}}
	return a, nil
}

var __02_postgresql_columnsUpSql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xac\x92\x3b\x4f\xc3\x30\x14\x85\xf7\xfe\x8a\xbb\x15\xa4\x8a\x01\x46\xa6\xd2\x96\xa9\x0f\x09\x85\x39\x75\x9c\x5b\x72\x15\x3f\x2a\xfb\x56\x15\xfc\x7a\xe4\x3e\x20\x7d\xd0\xd8\x11\x4b\xa6\xf3\xe9\x7c\xbe\x39\xc3\x69\x36\x79\x83\x6c\xf8\x32\x9d\x80\x46\x76\x24\x7d\x0f\x60\x38\x1e\xc3\x68\x31\x7d\x9f\xcd\x61\xa9\x73\x5f\x09\x87\x65\x5e\xa8\xda\xe7\x15\x71\x2e\x0d\x2f\xe1\x55\x59\xc1\x4f\x8f\x83\xd6\xb4\xdf\xe8\x9f\x34\x8c\x16\xb3\xd9\x64\x9e\x41\x3f\xb3\x2c\x14\x98\x8d\x2e\xd0\x81\x5d\xc1\x9e\x82\x42\x59\x59\x7b\x90\x42\x56\x08\x15\xb1\x87\xe2\x13\xb8\x42\xf0\x2c\x18\x35\x1a\xee\xdf\x6e\x74\x28\xca\x04\xc1\x5d\xbc\x8b\x61\x00\x2f\xdc\x1e\x5a\xe4\x4a\x72\x4c\x98\xe2\x77\x24\xba\x28\x1e\xd8\x64\xcb\xad\x23\x66\x34\x09\x96\x47\xa2\x8b\xe5\x81\x8d\xb2\x54\x56\x0a\x15\xbb\xc3\xb3\x70\x9c\xdb\x0e\xda\xab\x25\x8e\xb0\x51\x17\xb1\xc1\xf3\x74\xb2\x5d\xc2\x02\x1b\x5d\x71\x03\xbc\x02\x74\xf0\x4b\x99\x5f\xa3\x31\x6e\x7d\x57\x80\x0e\x8a\x29\xdb\x63\xd4\xeb\xe8\xdf\x7b\x16\x8e\x53\x0b\x50\xf2\xcf\xfd\x6d\x8a\x3b\xdc\x65\x3e\x5d\x2e\xe5\x6c\x85\xaa\xf7\x37\x60\xd2\xd8\xe2\x76\x9a\xbd\xe5\x15\x02\xa7\xdd\xe0\xd7\xe1\x1b\x70\x32\x1f\x07\xd3\x01\x90\x01\x4d\x4a\x91\x47\x69\x4d\xe9\xe1\x8e\x56\xc0\x4e\xc8\x3a\x27\x1b\x6a\x42\x96\x3c\xa0\x11\x85\xc2\x72\x00\x96\x2b\x74\x5b\xf2\x08\x5f\xe8\xec\xfd\x5f\x2f\x0a\x17\xc0\xd8\x27\x35\xc2\x9d\xde\x14\xf8\x7f\x7d\x53\x0f\xe0\xf9\x3b\x00\x00\xff\xff\xb5\x67\xf1\x34\xf8\x07\x00\x00")

func _02_postgresql_columnsUpSqlBytes() ([]byte, error) {
	return bindataRead(
		__02_postgresql_columnsUpSql,
		"02_postgresql_columns.up.sql",
	)
}

func _02_postgresql_columnsUpSql() (*asset, error) {
	bytes, err := _02_postgresql_columnsUpSqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "02_postgresql_columns.up.sql", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info, digest: [32]uint8{0xe2, 0x1d, 0xed, 0x3f, 0xb, 0xb3, 0x8d, 0x0, 0x6e, 0x61, 0xd9, 0x9f, 0x1, 0x3e, 0xeb, 0xe5, 0x5c, 0x54, 0x66, 0xbf, 0xcf, 0xd5, 0xc0, 0x15, 0xcf, 0xf, 0x81, 0xc6, 0x35, 0x31, 0x5a, 0x93}}
	return a, nil
}

var __03_add_agent_typeDownSql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x4c\x8e\xcd\x4a\xc4\x30\x18\x45\xf7\xf3\x14\x77\x17\x85\x8e\xf8\xb3\x11\xc5\x45\x1d\x23\x08\xcd\x14\x24\x2e\x5c\xcd\xc4\xf4\x6b\x1b\xc8\x9f\x49\x54\xfa\xf6\x12\xaa\xe8\xf6\x1c\xb8\xe7\xb6\x9d\xe4\xcf\x90\xed\x7d\xc7\xe1\xa8\x24\xa3\x33\x44\xff\xf0\xf4\xf8\x8a\x5d\xdf\xbd\x88\x3d\x8e\x6a\x22\x5f\x0e\x65\x89\x74\x04\xf7\x1f\xee\xfa\x64\x03\xb0\x3f\x7a\x30\xfe\x53\x59\x33\x30\xdc\xe1\xbc\xa9\xce\x2d\xf9\xdd\x6e\x23\xa5\x31\xeb\x99\x9c\xaa\xe6\xe2\x9f\xc9\x36\x7c\xd9\x30\x55\x7c\xb9\xe2\xe0\xa7\x30\xbc\x6d\x63\x0a\xa3\xb1\x94\xaa\xb9\xda\x00\xa7\xd8\xf5\x42\xf0\xbd\x04\x6b\x6b\x0f\x72\x89\x84\x32\xab\x02\x1d\xac\x25\x5d\x10\xc6\xdf\xdf\x37\xf8\xd9\x6d\x50\xd3\x58\xdb\x0d\xa8\xe8\x33\x76\xfb\x1d\x00\x00\xff\xff\xf5\x7c\x8f\xa8\xe9\x00\x00\x00")

func _03_add_agent_typeDownSqlBytes() ([]byte, error) {
	return bindataRead(
		__03_add_agent_typeDownSql,
		"03_add_agent_type.down.sql",
	)
}

func _03_add_agent_typeDownSql() (*asset, error) {
	bytes, err := _03_add_agent_typeDownSqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "03_add_agent_type.down.sql", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info, digest: [32]uint8{0x22, 0xa4, 0x18, 0x3c, 0x12, 0x1b, 0x61, 0xa3, 0x28, 0x96, 0x96, 0x58, 0xc5, 0x18, 0xa0, 0x95, 0xdb, 0xf8, 0xfc, 0xb5, 0x86, 0xad, 0x3d, 0xdc, 0xb4, 0xf9, 0xac, 0x96, 0x9e, 0xdf, 0xbb, 0x8f}}
	return a, nil
}

var __03_add_agent_typeUpSql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x4c\x8f\x4f\x4b\xc3\x40\x10\x47\xef\xfd\x14\xbf\x5b\x14\x52\xf1\xdf\x41\x14\x0f\xb1\x46\x10\x92\x06\x24\x1e\x3c\xb5\xeb\x76\xb2\x0d\xec\x3f\x77\x47\x25\xdf\x5e\xa6\x51\xec\x61\x2f\xef\x2d\x33\x6f\xaa\xa6\xaf\x5f\xd0\x57\x0f\x4d\x0d\x47\x9c\x46\x9d\xd1\x76\x8f\xcf\x4f\x6f\x58\x75\xcd\x6b\xbb\xc6\x56\x19\xf2\xbc\xe1\x29\xd2\x16\xb5\xff\x74\x37\x27\x0b\xa0\xf8\xa7\x9b\xd1\x7f\x29\x3b\xee\x0a\xdc\xe3\xbc\x14\xe7\xa6\xfc\x61\x97\x91\xd2\x90\xf5\x9e\x9c\x12\x73\x71\x64\xb2\x0d\xdf\x36\x18\xc1\x97\x33\x0e\xde\x84\xdd\xfb\x32\xa6\x30\x8c\x96\x92\x98\xab\x83\x89\x21\xb3\x49\x74\x98\x67\x32\x2b\x96\x47\x8e\x3c\x67\xf9\x74\xbd\x00\x4e\xb1\xea\xda\xb6\x5e\xf7\x28\x2a\x89\x42\x3f\x45\x02\xef\x15\x43\x07\x6b\x49\x33\xc2\xf0\x77\xdc\x2d\x7e\x97\x97\x90\x3e\xcc\x81\x25\x88\xf5\x59\x71\xf7\x13\x00\x00\xff\xff\x6e\xa0\x3b\x29\x0e\x01\x00\x00")

func _03_add_agent_typeUpSqlBytes() ([]byte, error) {
	return bindataRead(
		__03_add_agent_typeUpSql,
		"03_add_agent_type.up.sql",
	)
}

func _03_add_agent_typeUpSql() (*asset, error) {
	bytes, err := _03_add_agent_typeUpSqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "03_add_agent_type.up.sql", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info, digest: [32]uint8{0x60, 0xfb, 0x74, 0x75, 0xaa, 0xed, 0x12, 0xc4, 0xfd, 0xf4, 0x7, 0x2a, 0x33, 0xa, 0xd0, 0x6b, 0xea, 0x23, 0x4e, 0x77, 0x74, 0x40, 0xec, 0xc1, 0x17, 0xcc, 0x65, 0x30, 0xf, 0x7b, 0x2b, 0xa5}}
	return a, nil
}

// Asset loads and returns the asset for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
func Asset(name string) ([]byte, error) {
	canonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[canonicalName]; ok {
		a, err := f()
		if err != nil {
			return nil, fmt.Errorf("Asset %s can't read by error: %v", name, err)
		}
		return a.bytes, nil
	}
	return nil, fmt.Errorf("Asset %s not found", name)
}

// AssetString returns the asset contents as a string (instead of a []byte).
func AssetString(name string) (string, error) {
	data, err := Asset(name)
	return string(data), err
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

// MustAssetString is like AssetString but panics when Asset would return an
// error. It simplifies safe initialization of global variables.
func MustAssetString(name string) string {
	return string(MustAsset(name))
}

// AssetInfo loads and returns the asset info for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
func AssetInfo(name string) (os.FileInfo, error) {
	canonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[canonicalName]; ok {
		a, err := f()
		if err != nil {
			return nil, fmt.Errorf("AssetInfo %s can't read by error: %v", name, err)
		}
		return a.info, nil
	}
	return nil, fmt.Errorf("AssetInfo %s not found", name)
}

// AssetDigest returns the digest of the file with the given name. It returns an
// error if the asset could not be found or the digest could not be loaded.
func AssetDigest(name string) ([sha256.Size]byte, error) {
	canonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[canonicalName]; ok {
		a, err := f()
		if err != nil {
			return [sha256.Size]byte{}, fmt.Errorf("AssetDigest %s can't read by error: %v", name, err)
		}
		return a.digest, nil
	}
	return [sha256.Size]byte{}, fmt.Errorf("AssetDigest %s not found", name)
}

// Digests returns a map of all known files and their checksums.
func Digests() (map[string][sha256.Size]byte, error) {
	mp := make(map[string][sha256.Size]byte, len(_bindata))
	for name := range _bindata {
		a, err := _bindata[name]()
		if err != nil {
			return nil, err
		}
		mp[name] = a.digest
	}
	return mp, nil
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
	"01_init.down.sql": _01_initDownSql,

	"01_init.up.sql": _01_initUpSql,

	"02_postgresql_columns.down.sql": _02_postgresql_columnsDownSql,

	"02_postgresql_columns.up.sql": _02_postgresql_columnsUpSql,

	"03_add_agent_type.down.sql": _03_add_agent_typeDownSql,

	"03_add_agent_type.up.sql": _03_add_agent_typeUpSql,
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
// then AssetDir("data") would return []string{"foo.txt", "img"},
// AssetDir("data/img") would return []string{"a.png", "b.png"},
// AssetDir("foo.txt") and AssetDir("notexist") would return an error, and
// AssetDir("") will return []string{"data"}.
func AssetDir(name string) ([]string, error) {
	node := _bintree
	if len(name) != 0 {
		canonicalName := strings.Replace(name, "\\", "/", -1)
		pathList := strings.Split(canonicalName, "/")
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
	"01_init.down.sql":               &bintree{_01_initDownSql, map[string]*bintree{}},
	"01_init.up.sql":                 &bintree{_01_initUpSql, map[string]*bintree{}},
	"02_postgresql_columns.down.sql": &bintree{_02_postgresql_columnsDownSql, map[string]*bintree{}},
	"02_postgresql_columns.up.sql":   &bintree{_02_postgresql_columnsUpSql, map[string]*bintree{}},
	"03_add_agent_type.down.sql":     &bintree{_03_add_agent_typeDownSql, map[string]*bintree{}},
	"03_add_agent_type.up.sql":       &bintree{_03_add_agent_typeUpSql, map[string]*bintree{}},
}}

// RestoreAsset restores an asset under the given directory.
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
	return os.Chtimes(_filePath(dir, name), info.ModTime(), info.ModTime())
}

// RestoreAssets restores an asset under the given directory recursively.
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
	canonicalName := strings.Replace(name, "\\", "/", -1)
	return filepath.Join(append([]string{dir}, strings.Split(canonicalName, "/")...)...)
}
