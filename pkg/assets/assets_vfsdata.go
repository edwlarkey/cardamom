// Code generated by vfsgen; DO NOT EDIT.

// +build !dev

package assets

import (
	"bytes"
	"compress/gzip"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	pathpkg "path"
	"time"
)

// Assets statically implements the virtual filesystem provided to vfsgen.
var Assets = func() http.FileSystem {
	fs := vfsgen۰FS{
		"/": &vfsgen۰DirInfo{
			name:    "/",
			modTime: time.Date(2020, 8, 26, 19, 7, 11, 446667427, time.UTC),
		},
		"/static": &vfsgen۰DirInfo{
			name:    "static",
			modTime: time.Date(2020, 8, 26, 19, 7, 11, 446667427, time.UTC),
		},
		"/static/css": &vfsgen۰DirInfo{
			name:    "css",
			modTime: time.Date(2020, 10, 30, 18, 59, 22, 273129031, time.UTC),
		},
		"/static/css/main.css": &vfsgen۰CompressedFileInfo{
			name:             "main.css",
			modTime:          time.Date(2020, 10, 30, 18, 59, 22, 273129031, time.UTC),
			uncompressedSize: 1569,

			compressedContent: []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x02\xff\x9c\x54\xd1\x8e\xda\x3a\x10\x7d\xcf\x57\x8c\x76\x75\xa5\x5b\xd4\x80\xc3\xc2\x2e\x6b\x5e\xda\x1f\xe8\x43\xdb\x1f\x70\xb0\x09\x16\x8e\x27\x72\x86\x25\xb4\xe2\xdf\x2b\x3b\x4e\xc8\x66\x43\x1f\x2a\x4b\xa0\xcc\x8c\xcf\x1c\xcf\x99\x99\xc5\x0c\x7e\x1c\xf0\x0c\x9a\x40\xd7\xb0\xd7\x8d\x92\x40\x08\x74\x50\x40\x58\xc1\x6c\x91\xe4\x28\x2f\xf0\x3b\x01\x00\x28\xb5\x4d\x0f\x4a\x17\x07\xe2\xb0\x64\x4e\x95\xdb\xe4\x9a\x24\x8f\x56\xbc\xc5\x00\xa9\xeb\xca\x88\x0b\x87\xbd\x51\xcd\x36\x98\x8c\xae\x29\xad\xe9\x62\x54\x4a\x97\x4a\x71\xb0\x68\x55\xeb\xaa\x84\x94\xda\x16\xa9\x51\x7b\xe2\xb0\xae\x9a\x00\x37\xf7\x9f\x5d\x42\xe1\x0a\x6d\x53\xd7\xa6\x14\x27\xc2\x10\xf2\xa5\x54\x52\x0b\xf8\xdf\xf3\x39\x6b\x49\x07\x0e\x2f\xcf\xaf\x55\xf3\x29\x5c\x1b\xf0\xf1\x34\x52\xa9\x9d\xda\x91\x46\xcb\xc1\xe1\xb9\x63\x15\x23\x00\x22\x40\xb6\x0c\x04\xbc\xe5\x1a\x7e\x87\x3c\xa6\x99\xdc\x62\xc5\xe7\x51\x78\x7c\x1a\x87\xf9\x2a\x94\xa9\xb5\xee\xd1\x52\x5a\xeb\x5f\x8a\x43\x36\xcf\x96\xeb\xce\xe3\x41\xae\xfe\x61\xc9\x62\x16\xe9\xfb\xca\x2f\x66\xe1\x16\xbe\x29\xb7\x37\x78\xe6\x70\xd0\x52\x2a\xbb\x1d\xf8\xfa\x3c\xeb\xaa\x81\x8c\x55\xcd\xd0\x19\x39\x13\x56\x1c\x32\x55\x4e\xb8\x72\x24\xc2\xf2\x83\xd7\x68\xab\x7a\xa1\x9f\x46\xa8\xc3\x47\x2c\x2b\xea\x5c\x57\xff\xdf\xf6\x82\x88\x65\x20\xd5\x50\x2a\xd5\x0e\x9d\x68\xcb\x7f\x93\x7e\x87\x06\x1d\x87\xc7\x27\xb6\xfe\xfa\xc4\x5a\xdb\x18\x78\xa2\x47\x56\x9d\x44\x9d\x39\xea\xb1\x8a\xbd\x13\xd2\xcf\x73\x27\xac\xec\x3a\xe0\x06\xba\x64\xd3\xa0\xec\x0e\x28\x8b\xa0\x51\x93\xb4\x15\x78\x58\x09\x83\x82\x38\x78\xfb\xb0\x40\xfd\x10\xe4\x06\x77\xc7\xf7\xf5\xe9\xb0\x42\x8a\x29\xb0\xe0\x18\xa2\x85\x22\x0a\xa3\x0b\x3b\xe1\xfc\x5b\xaa\xb9\xc1\x02\x4f\x94\xe6\x27\x22\xb4\xb1\x1c\xb9\xd8\x1d\x0b\x87\x27\x2b\x87\x6a\xe4\xe8\xa4\x72\xff\xa4\x4f\x30\xee\x45\xa9\xcd\x85\xc3\xc3\x4f\x5d\xaa\x1a\xbe\xa9\x33\x7c\xc7\x52\xd8\x87\x08\x76\x72\xb5\x47\xab\x50\x5b\x52\x6e\x9b\xc0\x35\x49\x4a\xa1\xed\x70\x4a\x0b\xe7\x3b\x3c\xdb\xde\xdb\x23\x93\xa3\xec\x25\xdf\xa1\x25\x65\xe9\x2e\xd6\x6d\x7c\xfa\xfd\xf1\x48\xa2\x48\xdb\xc5\xa4\xa5\xca\x85\x8b\x97\xe3\x2a\x58\xb2\xff\x42\xd8\x90\x63\x3f\x68\xd9\xca\x4f\xda\x73\xec\x8d\xae\xd5\xbb\x7a\x31\xc6\xd8\x86\xc5\x3d\xe6\x13\xe4\x88\xc7\x52\xb8\x63\xfd\x7e\xa3\xf5\x83\xc7\xfa\xad\xe7\xa3\x4f\xce\x8c\x00\x9f\x5f\xfd\xf9\x28\x40\xdf\xcb\x77\x86\xac\x87\x24\x51\x8c\x20\x37\x9e\xe4\x90\xa3\x14\xa4\x46\x31\x2f\x1b\x7f\x7c\xcc\x9f\x00\x00\x00\xff\xff\xe8\xd1\xa2\xbc\x21\x06\x00\x00"),
		},
		"/static/img": &vfsgen۰DirInfo{
			name:    "img",
			modTime: time.Date(2020, 8, 26, 19, 7, 11, 446667427, time.UTC),
		},
		"/static/img/favicon.ico": &vfsgen۰CompressedFileInfo{
			name:             "favicon.ico",
			modTime:          time.Date(2020, 8, 26, 19, 7, 11, 446667427, time.UTC),
			uncompressedSize: 1406,

			compressedContent: []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x02\xff\xec\xc9\xa1\x0d\x80\x40\x10\x44\xd1\xb7\x82\x04\x79\x0a\x8d\xa4\x34\xfa\x6f\x60\xc9\x01\x09\x1a\x3f\x4f\x4d\xe6\x53\xca\x18\x6e\xe7\xc2\x86\x03\xf3\xda\xcd\xbe\x3e\xb1\x7c\xde\xdd\xdd\x22\x22\xe2\xbf\x2b\x00\x00\xff\xff\xe1\xe9\x47\x67\x7e\x05\x00\x00"),
		},
		"/static/js": &vfsgen۰DirInfo{
			name:    "js",
			modTime: time.Date(2020, 11, 7, 20, 56, 6, 217262186, time.UTC),
		},
		"/static/js/main.js": &vfsgen۰CompressedFileInfo{
			name:             "main.js",
			modTime:          time.Date(2020, 11, 7, 20, 56, 6, 217262186, time.UTC),
			uncompressedSize: 1554,

			compressedContent: []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x02\xff\xcc\x94\xdd\x6a\xdc\x30\x10\x85\xef\xfd\x14\xc3\x36\x60\x1b\xb2\x4a\x7a\xbd\x24\x90\x3f\x4a\xa0\xa5\x17\xcd\x0b\xcc\x4a\xb3\xb6\x89\x2c\xb9\xd2\x68\xc1\x84\xbc\x7b\x91\xe4\xba\xbb\xce\x92\xb4\x77\x35\x18\x27\xd2\x99\x33\x47\x9f\xc7\xab\xac\x0c\x3d\x19\x16\xa8\xd4\xc3\x9e\x0c\x7f\xed\x3c\x93\x21\x57\x95\xf7\xdf\xbf\xdd\x59\xc3\x71\xcd\xa2\x22\x55\x9e\xc3\x2e\x18\xc9\x9d\x35\x55\x0d\x2f\x05\x00\xc0\x6d\xd0\x3d\x3e\x61\xe3\x1f\xcd\x10\x58\x20\x33\xca\xb6\xaa\x37\xc5\x6b\xbd\x29\x8a\x7f\x33\xaf\x6a\xb8\xba\x86\x97\xa2\x00\xb8\xb8\x80\x2f\xc4\x80\x5a\xc3\xca\xe0\x7e\x8b\x6e\xbd\x0d\xae\x21\xb7\x02\xd2\x14\x2d\x7d\x01\x20\xad\xf1\x0c\x67\x59\x70\x9b\xf6\x3d\x5c\xc1\x8d\x73\x38\x8a\xc1\x59\xb6\x3c\x0e\x24\xbc\xee\x24\x09\x89\x5a\x57\x73\xa0\x9f\x81\xdc\xf8\x83\x34\x49\xb6\xee\x46\xeb\xaa\x14\x47\x7d\xca\xfa\x1c\x2e\xe3\x11\x52\x96\xbb\x96\xe4\x33\x74\x3b\xe0\x96\x1c\x01\xc6\xdb\x8c\x90\x2b\x20\x57\xc4\x40\xdd\x0e\xaa\xe3\x38\x42\x93\x69\xb8\x85\x6b\xb8\xac\xf3\xd1\x92\xe1\x8d\x52\x80\x20\x75\x27\x9f\x81\x22\x18\xb0\x06\x08\x65\x0b\x36\x35\xe9\x93\x70\x61\xb5\xb3\xee\x21\xd2\x05\xd2\x19\x14\xa4\x8b\xf4\x09\xbc\xc9\xfa\x98\x69\xbe\x26\xb2\xdc\x12\x30\xba\x86\x18\x76\xce\xf6\xe9\xff\x95\x42\xc6\x75\x5e\x5d\x01\x32\xbb\x6e\x1b\x98\xe6\xd2\x8c\x7b\xaa\xba\x8a\x7d\x63\x81\x27\x16\x79\x6d\xb3\x50\x9e\xcd\xd2\x19\x7b\x43\xfc\x90\xdf\xdf\xed\xf8\xa8\xaa\x2c\xc8\x98\xe7\x78\x4f\xb6\x69\x34\xe5\x44\x9d\x5f\xa3\xe4\x6e\x4f\x2b\x90\x1a\xbd\x8f\x98\xb6\x96\xdb\xbc\xbb\x98\x0c\x34\xea\x68\xbd\x27\x13\x56\xb3\x33\x69\x91\x2c\x22\x23\xc1\xa9\x47\x55\xce\xfe\x65\xfd\x27\xfd\x94\xfb\x23\xf9\xa4\x7f\x9d\x2a\xf3\xf3\xb5\x58\x4c\xfe\xd1\xa0\x55\xe5\x36\x30\x5b\xf3\xc9\x0e\x64\xd6\xbd\x55\xa8\xcb\xfa\x9d\xb7\x37\x7f\x6e\x69\x46\xf2\x37\x97\xfe\x14\x83\x4b\xcf\x7b\xda\x61\xd0\x5c\xa5\xde\x7b\x74\x90\x3c\x0f\x89\x2f\xfa\x8b\xa9\xe9\x26\xa1\x46\xef\x43\xdf\x99\x06\x46\x1b\xa0\xc5\x3d\x81\x35\x7a\x84\xcf\x93\x59\xcb\xfd\x7b\x5e\x71\x3b\x73\x4b\xa6\x07\xbc\x50\xa9\x37\x6c\xa3\xfa\x84\x44\xea\x6e\x18\x48\x4d\x40\xb3\xd1\xc9\xcc\xeb\x2d\xca\xe7\xc6\xd9\x60\xd4\x5f\x32\xfb\xfd\x1b\x45\x27\x69\xbd\x4d\xed\xa8\xb7\xfb\x53\x43\xb1\x88\x7e\xa0\x3b\x48\x9f\x27\xe0\x83\x33\x48\x6d\x3d\xfd\xc7\xf1\xe3\xfd\x2b\x00\x00\xff\xff\x18\xb7\xfe\x21\x12\x06\x00\x00"),
		},
		"/templates": &vfsgen۰DirInfo{
			name:    "templates",
			modTime: time.Date(2020, 12, 21, 15, 38, 38, 202270424, time.UTC),
		},
		"/templates/base.layout.tmpl": &vfsgen۰CompressedFileInfo{
			name:             "base.layout.tmpl",
			modTime:          time.Date(2020, 11, 7, 20, 56, 6, 217262186, time.UTC),
			uncompressedSize: 2804,

			compressedContent: []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x02\xff\xb4\x56\xcd\x8e\xdc\x36\x0c\xbe\xef\x53\xb0\xba\x4c\x0b\xc4\x56\xd2\x53\x5b\xd8\x46\xda\x00\x41\x0f\xe9\x0f\xba\xe9\x03\x70\x24\x8e\xad\x5d\x59\x32\x24\x7a\x36\x0b\x63\xde\xbd\xf0\xdf\x8c\xc7\xe3\x6d\x8a\xb4\xf5\xc5\xfa\xa1\x3e\x91\x1f\x29\x92\x5d\xa7\xe9\x60\x1c\x81\xd8\x63\x24\x71\x3a\xdd\x65\x5f\x69\xaf\xf8\xb9\x21\xa8\xb8\xb6\xc5\x5d\xd6\xff\xc0\xa2\x2b\x73\x41\x4e\x14\x77\x00\x59\x45\xa8\xfb\x01\x40\x56\x13\x23\xa8\x0a\x43\x24\xce\x45\xcb\x87\xe4\x3b\xb1\xdc\x72\x58\x53\x2e\x8e\x86\x9e\x1a\x1f\x58\x80\xf2\x8e\xc9\x71\x2e\x9e\x8c\xe6\x2a\xd7\x74\x34\x8a\x92\x61\xf2\x0a\x8c\x33\x6c\xd0\x26\x51\xa1\xa5\xfc\xcd\x2b\x88\x55\x30\xee\x31\x61\x9f\x1c\x0c\xe7\xce\xcf\xd0\x6c\xd8\x52\xd1\x75\x4c\x75\x63\x91\x09\xc4\xb0\x22\x20\x3d\x9d\x20\x81\x77\x18\x34\xd6\xbe\xce\xe4\x28\x38\x1e\xb2\xc6\x3d\x42\x20\x9b\xef\x62\xe5\x03\xab\x96\xc1\x28\xef\x76\x50\x05\x3a\xe4\x3b\x19\x19\xd9\x28\x69\xea\x52\x1e\xf0\xd8\x6f\xa5\x46\xf9\x1d\xf4\x64\xe4\x3b\x53\x63\x49\xf2\x53\x32\x1c\x59\x23\x8a\xc8\xcf\x96\x62\x45\xc4\x62\x84\x13\x15\x73\x13\x7f\x90\x52\x69\x97\x3e\x44\x4d\xd6\x1c\x43\xea\x88\xa5\x6b\x6a\xb9\x6f\x6d\x8d\x6f\x5f\xa7\xdf\xa7\xaf\xa5\x8a\x71\x9c\xa7\xb5\x71\xa9\x8a\x51\xfc\x5b\xf4\xb7\x2a\x10\xb2\x39\xd2\x00\x3b\x82\x27\x8c\x65\x34\xae\x69\xf9\xed\x9b\xf4\x75\xfa\xad\xd4\x26\xf2\xe5\xee\xcb\xf6\x8b\x5a\xec\x2e\x5a\xac\x29\xeb\x61\x6a\x1c\xcf\x0d\xdc\x64\x72\x0e\x91\x6c\xef\xf5\xf3\x04\xe5\xf0\x08\xca\x62\x8c\xb9\x70\x78\xdc\x63\x00\x13\x93\xd8\xa0\x22\x2d\x20\x78\x4b\xc3\xba\x29\x91\x8d\x77\x02\x30\x18\x4c\x2c\xee\x7b\x06\x7a\x74\x58\x6c\x8e\x88\x00\x99\x36\x2b\xcc\x64\x1f\xd0\xe9\xb3\x00\x40\x86\x2b\x01\xc3\x54\xcf\x3c\x4a\x51\x64\x91\x83\x77\x65\x71\x09\x9a\x69\x21\x93\x58\xdc\x2d\x61\x46\x15\xf7\x2d\x73\xaf\xde\xea\xd6\x36\x94\x14\x60\xfc\xad\x74\x27\xd7\x4e\x2b\xf4\xa9\x41\xa7\x49\xe7\xe2\x80\x36\x92\x00\x8d\xdc\x73\x1f\xca\xfe\xf5\x38\x3c\xfe\xd2\xcb\x5e\x74\x07\xc8\x62\x83\x6e\x3c\x5c\x19\xad\xc9\xe5\x82\x43\x4b\xa2\xc8\x64\xbf\xf3\x7f\x8b\x0e\x1c\xcc\x43\x6d\x8e\x67\x42\x36\x88\x1f\xed\x34\x7a\xcb\x90\x0d\xf1\xc8\x18\xf8\xda\xd6\xcf\x78\xea\x67\x5f\xd3\x42\xa1\xfe\xeb\x3a\x73\x80\xf4\xc7\x96\x2b\x72\x6c\x14\x32\xe9\x3f\x23\x85\xd3\x69\x21\xf3\x39\x60\x6c\x1a\xb9\xf7\xfe\xb1\xc6\xf0\x28\x8a\x9f\xa6\x51\x5c\xdd\xf4\x4f\x50\x18\x4b\x51\x7c\xc4\x32\xde\x68\x49\x4e\x2f\x74\xba\xa6\x72\x93\x1d\xba\x8a\xe1\x4d\x91\xe1\xfa\x2f\x20\x63\x81\x34\xc6\x72\x14\xd7\x86\x02\x64\x07\x1f\xea\x59\xa8\x1f\x27\xc6\x59\xe3\x48\x00\xaa\xfe\xfd\xe5\xbb\xc1\xde\x36\x52\x90\xd6\x97\xbe\xe5\x1d\xd4\xc4\x95\xd7\xf9\xee\xf7\xdf\xee\x3f\xee\xd6\x80\xa3\x7a\x90\xbe\xbb\xff\xe3\xfd\x7b\x43\x56\xc3\x4a\xab\xe9\xde\x21\x05\x5d\x6b\xd7\x27\x89\x26\x98\x1a\xc3\xb3\x18\x73\xb1\x88\xed\xbe\x36\x2c\xe0\x88\xb6\xa5\x5c\x7c\x18\x34\x80\xaf\xbb\xee\xd6\xf8\xf4\x57\xac\xe9\x74\xfa\xe6\xd6\x44\xd9\xdb\xb5\xf2\xf0\xe8\x97\x2b\xc7\xd9\x48\x5f\x42\x20\xfe\x8d\x11\x53\xc4\x44\x53\xba\xb6\x11\x05\xcc\x19\xe8\xde\x94\x0e\xda\xe6\x9c\x80\xe0\x26\x04\xb7\x81\xad\x29\xab\x73\x65\xe8\xfd\x61\x9c\x28\xe0\x83\x2f\xc1\xb8\x5b\x8c\x2d\x23\xaf\xa2\xf3\x46\xe4\x6a\xba\x98\x64\xd2\xe1\x3c\x8c\x34\x04\xc6\xac\xdc\x34\xdd\xcc\xd3\x7d\x0b\x80\xc6\x51\x00\xb1\x78\x03\x5d\xf7\x64\xb8\x82\xf4\xbd\xc5\x58\x2d\xdf\x0a\x06\x36\xca\xd2\x7c\xba\xa6\x18\xb1\xa4\x17\x1f\xc7\xb4\x9f\xf4\x75\x47\x14\x5d\x97\x9e\x4e\x37\xe6\x4c\x90\xc5\xdd\x9a\x82\xc5\xc2\xa5\xb5\x98\x5a\x96\xa1\xb9\xb8\xdb\x62\x61\x32\x76\x9a\x1e\xbc\x67\x0a\x97\xc7\xd3\xcf\x5e\xe4\x81\x1c\x43\x85\x31\x61\xfa\xc4\x89\x22\xc7\x14\xe8\xaa\x78\x35\xe7\xfa\x94\xe1\xaa\xf6\x97\x86\xab\x76\x9f\x2a\x5f\x4b\xd2\x4f\x16\xc3\x23\x3d\x4b\x35\x55\x31\x71\xae\x67\x70\xef\xdb\xa0\x86\xc4\xb9\x28\x6d\xcd\xa6\x3f\x47\x6d\x67\x97\xaa\x60\x1a\x86\x18\xd4\x7f\xd6\x6f\x3c\x6c\xb7\x1b\x0f\x71\xa8\x3c\xc3\x7d\x1b\x97\xcf\x6d\xc6\xc3\xd4\x65\x3c\xc4\x39\x11\xf4\xb4\xc9\x07\x3c\xe2\x28\x7e\x8d\x92\xc9\xb1\xf7\xc8\xe4\xd8\xc9\xce\x4e\xfe\x2b\x00\x00\xff\xff\xd1\xb5\x88\x1e\xf4\x0a\x00\x00"),
		},
		"/templates/bookmarks.page.tmpl": &vfsgen۰CompressedFileInfo{
			name:             "bookmarks.page.tmpl",
			modTime:          time.Date(2020, 11, 7, 20, 56, 6, 220595570, time.UTC),
			uncompressedSize: 3056,

			compressedContent: []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x02\xff\xc4\x56\xc9\x6e\xe4\x36\x10\xbd\xf7\x57\x14\x88\x20\x3e\x49\x1d\xf8\x1a\x49\x40\xe2\x89\x93\x00\x46\x26\xf0\x92\x6b\x40\x89\xa5\x16\x61\x4a\x14\xc8\x6a\xcf\x18\xb4\xfe\x3d\x20\xb5\xb5\x16\x67\x19\x0c\x90\xbe\x34\xa9\xda\x5e\xbd\x2a\x16\xe9\x9c\xc0\x52\x36\x08\x8c\x24\x29\x64\x5d\xe7\x5c\xfc\x20\x09\x1f\xfd\xb6\xeb\x20\x82\x5f\x74\x8d\xce\x61\x23\xba\xee\x70\x98\xf5\x0b\xdd\x10\x36\xc4\xba\xee\x90\x08\xf9\x02\x85\xe2\xd6\xa6\x4c\xe1\x0b\x2a\x96\x1d\x00\x36\x5f\x23\x85\x25\x05\xd1\x9e\x50\x12\xd6\x83\x10\x20\xa9\xae\x47\x69\xc0\x05\xd2\x46\xd7\x2c\xfb\x51\xeb\xe7\x9a\x9b\x67\x9b\x1c\xab\xeb\xc1\xd1\x51\xc8\x97\x10\x6e\x5a\x6c\x5c\x1b\x79\xaa\xfe\x6d\xe0\xfc\x4c\xa4\x1b\x90\x22\x65\xba\xc5\x26\xaa\xb5\xe0\x8a\x8d\x16\xa3\xd4\x46\x4a\x36\xcf\xfe\xdf\xd6\x5c\x29\x96\xdd\x18\xe4\x84\x30\x02\x4c\x8e\xbd\xe6\x3e\xc6\xe1\x6f\xc1\x5b\x1f\x67\x8d\x3f\x7c\x8d\x72\x5e\x3c\x9f\x8c\x3e\x37\x82\x65\xbb\x69\xf6\x6a\x05\x37\x62\xcc\xb2\x42\x2e\xd0\x6c\xe5\x91\x17\xcc\xd9\xb6\x3b\x1a\x7d\x23\x6c\x33\x6a\xd7\x1c\x2d\x4d\x95\xb6\x08\x02\x15\x12\x32\xe0\x46\xf2\x48\xf1\x1c\x55\xca\x82\xc4\x03\x5f\x52\xd2\x23\x1c\x76\x16\x0b\x92\x1b\x97\x1e\x4d\xae\xc5\xeb\x8c\xb7\xd4\xa6\x0e\xb5\xc9\x07\x58\x51\x11\x60\x32\xe0\xc1\x41\x7a\x75\xe4\x6d\x7b\x1c\xa5\xc7\x5e\x7a\x05\x35\x52\xa5\x45\x7a\xf5\xfb\xc7\x87\xc7\xab\xd1\x9b\xff\x39\x07\xf1\xcd\xc3\xfd\xed\xad\x44\x25\xa0\xeb\x16\xa2\x4f\x92\x2a\x88\x6f\xb5\xa9\x17\x82\x25\xf7\xa5\xb7\x64\xd9\x42\x0e\x90\x84\xe4\xa7\x46\xf3\x1b\x96\x3d\xdd\xdf\x25\xc7\xb0\x5e\xab\x8f\xb1\x7e\x32\x46\x1b\x1b\xff\x8c\x04\xec\x6c\x14\x5b\xc5\xdd\x46\xaf\xb8\x8d\x08\x3f\x53\x24\x78\x73\x42\xc3\x32\xe7\xe2\xae\x9b\x9a\x64\x19\xa3\x3f\xc3\x2b\x67\xb2\x69\xcf\x34\xba\x0b\x1b\x06\xf4\xda\x62\xca\xbc\x5f\x06\x0d\xaf\x31\x0d\x60\xe0\x85\xab\x33\xa6\xcc\xb9\x4b\x80\x0c\xf8\x99\x74\xa9\x8b\xb3\x4d\xd9\xb4\x5c\x11\xb2\x03\xe8\x0b\x39\x0c\x73\xe9\x3f\xb0\x38\x8e\xb5\xff\x9d\xc7\x56\xf1\x02\x2b\xad\x04\x9a\x94\x85\x2c\xe0\x93\x54\x0a\x72\x84\x12\xa9\xa8\x50\x40\x69\x74\x0d\x4f\xf7\x77\x20\x4b\x68\x34\x81\xc5\x89\xff\x3e\x8d\x75\x05\xc6\xe4\xbe\x3e\xdb\x45\x85\xc5\x73\xae\x3f\x6f\xd4\xa6\x4c\xfb\xdc\x66\x3d\xb8\x47\x2e\xe0\x8e\x13\x9a\xb5\xe7\xbd\x6a\x7d\xc5\x8e\xe0\x27\xfb\x4e\x43\x2c\x7d\xfa\x1b\xcb\x68\xb5\x93\x52\x98\x40\x0a\x0b\x82\xfa\xac\x48\xb6\x0a\x47\xda\xf9\xc9\x32\x10\x9c\x78\x34\xd4\x72\xfe\xb0\x28\xe8\x4d\xa5\xfd\xfc\xf3\x58\x76\xfd\xfb\xb6\x31\xbe\xb5\xe0\x9b\xf8\x63\xeb\x47\x95\xdd\xed\x49\x80\x44\x07\x31\x38\x27\x4b\x88\x1f\x02\x2c\x14\x5d\x67\x87\xd5\xd0\x7e\xbe\x41\xff\xf0\xdd\xe0\xbb\xb4\x37\x79\x2f\xee\x5e\xbb\xf6\x25\xe8\x7d\xee\xb0\xb6\x53\x9c\x2f\xaa\xd7\xe2\x50\x2c\x2f\xcf\xf1\x74\xd8\x73\x5e\x4b\x9a\x5a\xbb\xbf\x76\xfe\xa9\xa1\x93\xa3\xbf\x09\x96\x83\x7c\x4e\xd3\x27\x56\x8c\x8c\xac\xee\xdc\x03\x0c\xcc\x4e\x6f\x89\xc1\x68\xac\xcf\x46\x70\x99\xe6\x78\xb1\xcc\x37\x12\x87\xca\x60\x99\xb2\xe5\xbd\xe3\x5c\xfc\xeb\x87\x30\x1b\x15\x85\xe3\x3a\x3c\xa7\xc2\x5c\x19\xd6\xc9\x91\x67\xdf\x36\xb9\x6d\xbf\x7f\xeb\xff\x12\xbe\x99\x46\x27\x83\xaf\x6c\x88\xe0\x5c\xfc\x74\x7f\xe7\x9d\x12\x37\x27\xa4\x94\xfd\x99\x2b\xee\x89\x1c\x83\xf4\xe2\xac\x5f\xc1\x1b\x08\x5d\x73\xd9\xfc\xc6\xeb\x21\x5a\x92\x9b\x99\xb2\xf5\xc9\xf0\x6f\xb9\x05\xc5\x24\x6b\x7c\x07\x90\xe0\x84\x5e\x1c\xa2\xf6\x15\x13\x3f\xd0\x10\x7b\xda\xc3\x1b\xd4\xba\xa1\xea\x03\xa7\x00\xc0\x5b\x2c\x4b\xe6\xeb\xe0\x8f\xcc\xa2\x3d\x2f\x81\xd1\xf6\x3c\x4d\x75\xda\x18\x0e\xf5\x18\x4c\x5b\x15\x7d\x07\xad\xf1\xcf\x46\x7f\xc9\x7b\x16\xa0\xeb\xc0\xf3\x70\xf8\xfb\xf3\xb1\xe9\xb6\xb5\xca\x85\xc2\xe5\x72\x66\x77\x36\x70\x0e\x95\xc5\xc9\x36\xe1\x86\x64\xa1\x70\x7e\xdc\xf6\x6f\xdb\xa2\x92\x4a\xf8\x99\x2f\x4b\x59\x70\xdf\xba\x17\x59\x27\x6d\xf6\x58\xa1\xc1\x2b\xeb\x35\x2a\xd9\x9c\x80\x34\x58\x44\xf0\x5f\xe3\x38\x7e\x45\xba\x7c\x9e\x1d\x87\x20\xd9\x61\x46\x32\xfe\xff\x15\x00\x00\xff\xff\x2d\x59\xbc\x02\xf0\x0b\x00\x00"),
		},
		"/templates/create.page.tmpl": &vfsgen۰CompressedFileInfo{
			name:             "create.page.tmpl",
			modTime:          time.Date(2020, 11, 7, 20, 56, 6, 220595570, time.UTC),
			uncompressedSize: 1413,

			compressedContent: []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x02\xff\xc4\x94\xcf\x8e\x9b\x30\x10\xc6\xef\x79\x8a\x91\x55\x29\x27\x40\xb9\x03\x52\x1b\x35\xbd\x54\xda\x6a\xb3\xed\xdd\xc1\x93\x60\xad\xb1\x91\x3d\x64\x5b\x59\x7e\xf7\xca\xfc\x49\x02\x9b\x56\xdd\x5e\x36\x97\x60\x7f\xc3\xcc\xf8\xf7\x0d\xf6\x5e\xe0\x51\x6a\x04\x46\x92\x14\xb2\x10\x3e\x0a\x01\x1c\x34\xbe\xc0\xc1\x98\xe7\x86\xdb\x67\xef\x51\x8b\x10\x56\xab\x6b\x70\x65\x34\xa1\x26\x16\xc2\x2a\x17\xf2\x0c\x95\xe2\xce\x15\xac\x32\xaa\x6b\xb4\x63\xe5\x9d\x5d\x90\x2e\x31\x1a\x13\xaa\xa5\x15\x31\xa2\xde\x4c\x01\x7d\xe9\xa8\x6f\x58\x19\xcb\x7f\x1a\x0b\xe7\x59\xbd\x29\x57\x00\xf9\xd1\xd8\x06\xa4\x28\xd8\xd4\x52\x52\x59\xe4\x84\x0c\x78\x45\xd2\xe8\x62\x9d\xf1\xb6\xcd\x26\x35\x1b\xd4\x35\x34\x48\xb5\x11\xc5\xfa\xdb\xc3\xfe\x69\x1d\x33\xc5\x9f\xf7\x90\x6e\xf7\x8f\xbb\x9d\x44\x25\x20\x84\xcb\xf6\x8b\xa4\x1a\xd2\x9d\xb1\xcd\x65\x13\xe0\xf6\x20\xc7\xf8\x06\x2b\x2f\x1a\x40\xae\xf8\x01\xd5\xa4\xf7\x0b\x56\x3e\xc5\xe3\xe4\x59\xbf\xba\x0d\x9e\x2a\x7c\xb6\xd6\x58\x97\x7e\x41\xba\x52\xbf\x09\x5b\x56\xad\xb9\x4b\x08\x7f\x52\x22\xb8\x3e\xa1\x65\xa5\xf7\x69\x08\x79\x26\xe4\x79\x9e\x7d\x70\xe9\x26\x89\xd4\x6d\x47\x53\x9a\x7e\xc1\x80\x7e\xb5\x58\xb0\x98\x8f\x81\xe6\x0d\x8e\xf4\x19\x9c\xb9\xea\xb0\x60\xde\xcf\x1b\x63\xc0\x3b\x32\x47\x53\x75\xae\x60\x97\xc7\x1b\x08\x8b\x46\xfe\x83\xd7\xf7\xc7\xaf\xff\x48\xab\xb3\xea\x7d\x59\xc5\x06\x96\xa4\x86\xa6\xde\xc6\xe9\x6f\x94\xee\xcf\x14\x3f\xb9\x3b\x90\xe6\x9f\x99\x26\x6b\xd4\x8c\x37\x40\xee\x50\x61\x45\xd0\x74\x8a\x64\xab\x70\x32\x9d\x9f\x1c\x03\xc1\x89\x27\xe3\x29\xaf\x1b\xad\xe2\x15\xd6\x46\x09\xb4\x05\xdb\xd6\xc6\x38\x84\x58\x7f\x91\x39\x62\xb4\x11\x33\x7c\x48\x1f\xda\xf8\x19\xba\xd7\xde\x98\x5e\x00\xef\xe5\x11\xd2\x7d\xdf\x0a\x8a\x10\xdc\xf8\x34\x1a\x11\x6d\xfa\x11\xa9\x46\xaf\x86\x57\x5e\xd7\x5a\x5a\x16\x91\x0e\x79\x66\x44\x66\x98\xdf\x34\x9b\xb3\x11\x38\x74\x44\xa6\xbf\xb4\x5c\x57\x55\xe8\xdc\x34\x0e\xae\x3b\x34\x92\x2e\x53\xb0\x1d\x2e\xa2\x3f\xf8\x3c\x75\x9d\x67\xf1\x0a\x2b\x57\xa3\x38\xfe\x4d\xea\xef\x00\x00\x00\xff\xff\xfb\xb1\x75\x5c\x85\x05\x00\x00"),
		},
		"/templates/create_tag.page.tmpl": &vfsgen۰CompressedFileInfo{
			name:             "create_tag.page.tmpl",
			modTime:          time.Date(2020, 8, 26, 19, 7, 11, 446667427, time.UTC),
			uncompressedSize: 747,

			compressedContent: []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x02\xff\x6c\x52\xc1\x8e\x9b\x30\x10\xbd\xf3\x15\xa3\xb9\xe4\x94\xa0\xdc\xc1\x52\xb5\x6a\x7a\x6b\xab\xee\xfe\x80\x63\x0f\x60\x09\xec\xc8\x1e\x67\x5b\x59\xfe\xf7\xca\x10\x76\x59\x14\x0e\x98\x79\x7e\xe6\xbd\x37\x9e\x94\x34\x75\xc6\x12\x20\x1b\x1e\x09\x73\xfe\xa6\x35\x48\xb0\xf4\x0e\x2c\xfb\x94\xc8\xea\x9c\xab\xea\x93\xa7\x9c\x65\xb2\x8c\x39\x57\x8d\x36\x77\x50\xa3\x0c\xa1\x45\xe5\xc6\x38\xd9\x80\xe2\x09\x0a\x26\x1c\x9d\xa5\x23\x0f\xc6\xeb\xc2\x18\xce\x2b\x61\x56\x2d\xfb\x67\x14\x45\xf9\x4d\xf6\x4d\x3d\x9c\x45\xd5\x74\xce\x4f\x60\x74\x8b\x2c\xfb\xa3\xf2\x24\x99\x10\xa4\x62\xe3\x6c\x7b\xa8\xe5\xed\x56\xb3\xec\xeb\x65\xe3\x00\x13\xf1\xe0\x74\x7b\xf8\xfd\xeb\xf5\xed\x20\x2a\x00\x80\x94\xe0\xf4\xf2\xfa\xe7\x72\x31\x34\x6a\xc8\xf9\x01\xbe\x1b\x1e\xe0\x74\x71\x7e\x7a\x40\x5b\xf8\xbb\xf7\xce\x87\xd3\x0f\x62\xc0\x9e\x2c\x79\xa3\x70\x43\x2b\xcf\x36\x9e\x75\x6c\x3a\xa3\x64\x31\x55\x42\x68\x69\x7b\xf2\x28\xbe\x1c\x00\x68\xae\x91\xd9\xd9\xf5\x94\xa6\x91\x98\x50\x34\xf5\x82\xef\xe9\x29\x9d\xf6\x9a\xb5\x36\x77\xb1\x71\xbb\xdc\xca\x33\x4b\x5d\x49\xbb\x73\xd0\x8c\xf2\x4a\xe3\xca\x98\x0b\x14\x3f\xe5\x44\x4d\x3d\x17\x3b\xb6\xb1\xb7\xc8\x2b\x7b\x2e\x10\xf8\xdf\x8d\x5a\x64\xfa\xcb\x08\x56\x4e\xd4\x62\x79\x23\xdc\xe5\x18\xa9\xc5\x94\x96\x9e\xcd\x60\xce\x08\x32\xb2\xeb\x9c\x8a\xa1\xc5\x8f\xcf\x8d\xab\x4d\xa0\x6d\x98\xaf\xd2\x8f\xae\x99\x70\x0c\x51\x29\x0a\x61\xb5\x11\xe2\x75\x32\xfc\x21\xfe\xb2\x0c\x87\xa8\x9a\xba\x0c\x4d\x59\xe7\xbf\x3f\x96\x55\xe0\x7f\x00\x00\x00\xff\xff\xee\x06\xc7\xd9\xeb\x02\x00\x00"),
		},
		"/templates/edit.page.tmpl": &vfsgen۰CompressedFileInfo{
			name:             "edit.page.tmpl",
			modTime:          time.Date(2020, 8, 26, 19, 7, 11, 446667427, time.UTC),
			uncompressedSize: 1550,

			compressedContent: []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x02\xff\xc4\x54\x4d\x8f\x9b\x30\x10\xbd\xf3\x2b\x46\x56\x24\x4e\x21\xf7\x15\x70\xe8\xb6\x5b\x55\x5a\x69\xab\x4d\xd2\xbb\x83\x87\x60\xad\xb1\x91\x6d\xd2\x56\x96\xff\x7b\x65\x30\x04\x92\x3d\xb4\x95\xaa\x9e\xc0\xe3\xf9\x7a\xef\x8d\xc7\x39\x86\x35\x97\x08\xc4\x72\x2b\x90\x78\xef\x5c\xb6\xe7\x16\x0f\xe1\xe8\xfd\x03\x38\x97\x7d\x50\xea\xad\xa5\xfa\x2d\x8b\x46\xe7\x50\x32\xef\x93\xe4\x1a\x5d\x29\x69\x51\x5a\xe2\x7d\xe2\xdc\x77\x6e\x1b\xd8\x9c\x62\x14\x3c\x14\x30\xa7\xf0\x3e\x01\x00\xc8\x6b\xa5\x5b\xa8\x04\x35\xa6\x20\x04\x38\x2b\xc8\xe4\xbe\x45\xc6\x2d\x81\x16\x6d\xa3\x58\x91\x7e\x7d\xd9\x1f\xd2\x72\x08\x72\x0e\x36\xd9\xe3\xfe\xf5\xe9\x89\xa3\x60\x10\x53\x4d\xe5\xb2\x27\xa5\xdb\x68\x9b\xad\xd9\x27\xad\x95\x36\xd9\x67\xb4\x40\xce\x28\x51\xf3\x8a\xcc\x4e\x00\x39\xe3\x97\xa9\x0d\xa9\x2c\xaf\x79\x45\x2d\x57\x12\xb8\xd9\x32\x2a\xcf\xa8\x49\x39\x3b\x03\xe4\xa7\xde\x5a\x25\xa7\x08\x86\x02\x2d\x92\x32\xdf\x8d\xf6\xa5\xab\x73\xd9\xb2\xce\x8e\xf1\x4b\x39\xf7\x36\xd2\x17\xaf\xb8\xec\x7a\x0b\xf6\x67\x87\x45\xda\x70\xc6\x50\xa6\x20\x69\x8b\x45\xca\x59\x0a\x17\x2a\x7a\x2c\x52\xe7\x36\x57\x19\xbe\x7c\xf4\x3e\x2d\x93\xe4\x1e\x43\x1d\x98\x59\xb4\x9c\x0b\x7a\x42\x31\xdd\x0e\x07\x52\x0e\x2a\xe6\xbb\xe1\x74\x75\x7d\x8f\xb1\x69\x26\x16\xb8\xd6\xf5\x1a\x6a\xb6\x16\x7f\xd8\x99\xac\x01\xf6\x0a\xed\x2d\xde\x19\x71\x4c\x31\x1c\xc8\x88\x9f\x84\x5c\x64\x44\x1f\x8b\x47\x02\xc8\x8a\x80\x38\x87\x04\x68\x6f\x55\xad\xaa\xde\x14\x64\xfe\x9d\xe1\xc7\x36\xfe\x96\xa6\xe3\xeb\xf3\x6f\x91\xd4\x6b\xf1\xbf\x28\x0a\xa5\xdf\x25\xe8\xf8\xfa\xfc\xaf\xe9\x39\xd0\xb3\xb9\xe3\x67\x9d\x27\x2c\x05\xad\xc4\xea\x09\x01\xe4\x06\x05\x56\x16\xda\x5e\x58\xde\x09\x9c\xe4\xa6\x67\x43\x80\x51\x4b\xb7\x11\xe8\xd5\xd0\x09\x5a\x61\xa3\x04\x43\x5d\x90\xc7\x46\x29\x83\x10\xea\xdf\x64\x0e\x3c\xea\xc0\x32\x6c\xb2\x97\x2e\xbc\x63\x73\x23\x0c\x40\xae\x86\x0b\x70\x8e\xd7\x90\xed\x87\x56\x90\x79\x6f\xe2\x5f\x54\x22\xa8\xf4\x2d\x10\x1b\xa4\x1a\x43\xee\x6b\xad\x35\x1b\x39\x1d\xf3\xac\x18\x59\x4a\xfd\x07\xa4\xaf\x46\x20\x2e\x1e\x6e\xb6\xa6\xaf\x2a\x34\x66\x1a\x07\xd3\x9f\xda\xb0\x2f\xe3\x14\x1c\x3b\x46\xc3\x4a\x4a\xee\x6a\x2f\xdb\xcd\x77\x61\xff\x96\xc9\x64\x9b\xbe\xbf\x02\x00\x00\xff\xff\x69\x5e\xd4\x04\x0e\x06\x00\x00"),
		},
		"/templates/home.page.tmpl": &vfsgen۰CompressedFileInfo{
			name:             "home.page.tmpl",
			modTime:          time.Date(2020, 12, 21, 15, 38, 38, 202270424, time.UTC),
			uncompressedSize: 231,

			compressedContent: []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x02\xff\x6c\x8f\x31\x0e\x83\x30\x0c\x45\xf7\x9c\xc2\xca\x4e\x11\x7b\x9a\x13\x74\xe8\xd6\x99\x26\xa6\x58\x84\x18\x11\x43\x55\x45\xb9\x7b\x05\x12\x65\xe9\x64\xf9\xdb\x4f\x7a\x3f\x67\x8f\x1d\x45\x04\x2d\x24\x01\x75\x29\x37\x7e\x51\xcc\x19\xa3\x2f\x45\xa9\xf3\xee\x38\x0a\x46\xd1\xa5\x28\xe3\x69\x05\x17\xda\x94\xae\xda\x71\x58\xc6\x98\xb4\xfd\x93\x02\xa5\x8a\x23\x56\xd2\xd3\xec\xf7\xa5\xeb\x12\xca\x99\x6d\x54\xdf\x1c\xd0\x6e\xb0\xbd\x35\xda\x3e\x30\x38\x1e\xd1\xd4\x7d\x63\x95\x99\xec\x3d\x60\x9b\x10\xc2\x26\x07\xc2\xb0\x12\xbe\xe1\xc3\xcb\x0c\x4f\xe6\x61\x6c\xe7\x21\x5d\x4c\x3d\x59\x65\x6a\x4f\xeb\x6f\x1c\x3d\xbe\x01\x00\x00\xff\xff\x46\xed\x4a\x08\xe7\x00\x00\x00"),
		},
		"/templates/login.page.tmpl": &vfsgen۰CompressedFileInfo{
			name:             "login.page.tmpl",
			modTime:          time.Date(2020, 8, 26, 19, 7, 11, 446667427, time.UTC),
			uncompressedSize: 951,

			compressedContent: []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x02\xff\x9c\x52\xcb\x8e\xdb\x30\x0c\xbc\xfb\x2b\x08\xde\x13\x23\x77\xdb\x97\x62\xd3\x4b\x81\x2e\xba\xfd\x01\xc5\xa2\x63\x02\xb2\x14\x58\x54\x16\x85\xa0\x7f\x2f\xfc\xd0\xc6\x2e\x72\x69\x4e\xf6\x0c\x87\xe0\x8c\xc8\x18\x35\x75\x6c\x09\x50\x58\x0c\x61\x4a\x3f\xdc\x95\x6d\x8c\x64\x75\x4a\x45\xf1\xa8\xb7\xce\x0a\x59\xc1\x94\x8a\x4a\xf3\x1d\x5a\xa3\xbc\xaf\xb1\x75\x26\x0c\xd6\x63\xf3\x84\x05\xf6\x07\x67\xe9\x20\x3d\x8f\x7a\x06\x5d\xe7\x49\x1e\xdc\xd4\xd5\x9f\x72\xd3\xec\x60\x92\x9d\xb0\x99\x5d\x54\x65\x7f\x6a\x8a\xaa\x73\xe3\x90\x35\x2a\x48\x8f\xa0\x5a\x61\x67\x6b\x2c\xcd\x24\x43\x18\x48\x7a\xa7\x6b\x7c\xff\xf9\xf1\x1b\xc1\xba\xbb\x32\xac\x95\x50\x53\x00\x00\xc4\x08\xc7\x6f\x1f\xbf\xce\x67\x26\xa3\x21\xa5\x95\xfc\x64\xe9\xe1\x78\x76\xe3\xb0\x52\x5b\xfa\x6d\x1c\xdd\xe8\x8f\xdf\x49\x00\xaf\x64\x69\xe4\x16\x37\x32\x80\x6d\x5a\xeb\x84\x3b\x6e\xd5\x64\x6a\xf2\xaf\x95\xbd\xd2\x88\xcd\x46\x0e\x50\x5d\x82\x88\xb3\xb9\x47\x93\x21\x21\x6c\xaa\x72\xe1\xf7\xe2\x18\x8f\xfb\x69\xa5\xe6\x7b\xb3\x71\xb9\xac\xe7\x99\x99\x6e\x4a\xb9\x9b\x5d\x19\x75\x21\x93\xeb\x33\xc0\xe6\x6d\x50\x6c\xaa\x72\x46\x3b\x31\xdb\x5b\x90\x2c\x9e\x01\x82\xfc\xb9\x51\x8d\x34\xb5\x20\x58\x35\x3c\xc0\x5d\x99\x40\x35\xc6\xb8\x3c\xd5\xc2\xa6\x84\xa0\x82\xb8\xce\xb5\x61\xd9\xd9\xf2\xbb\x71\xf5\x4f\xa0\x17\x02\xbc\x2b\xef\x3f\xdd\xa8\xff\x2f\xc3\x6d\xed\xca\x31\xbe\xf0\xab\xd6\x76\xa3\xd6\x15\xb3\x3f\xf8\xd0\xb6\xe4\x7d\x1e\xeb\xc3\x65\x60\xf9\x7a\xae\xf9\xba\x9f\x8f\xcc\xab\xad\xca\xe9\xec\x9b\x62\x2d\xad\x9f\x5c\xfd\x1b\x00\x00\xff\xff\xbf\x41\xe7\x84\xb7\x03\x00\x00"),
		},
		"/templates/settings.page.tmpl": &vfsgen۰CompressedFileInfo{
			name:             "settings.page.tmpl",
			modTime:          time.Date(2020, 12, 21, 15, 21, 20, 955807984, time.UTC),
			uncompressedSize: 1033,

			compressedContent: []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x02\xff\x8c\x53\x4d\x8f\xd3\x30\x10\xbd\xf7\x57\x8c\x7c\xa0\xac\x50\x6a\xa9\xc7\x55\x12\x04\x05\x69\xb9\xb2\xcb\x0f\x98\xd8\xd3\xc6\xc4\xb1\x83\x3d\x29\xaa\xa2\xfc\x77\x94\xaf\xad\xba\x0d\x82\x5e\x5c\xcf\x7b\xf3\xe6\xbd\x89\xdc\x75\x9a\x8e\xc6\x11\x08\x36\x6c\x49\xf4\x7d\xd7\xed\x9e\x0d\xd3\xcb\x70\xed\x7b\x48\xe0\xc9\xd7\xd4\x75\xe4\x74\xdf\x6f\x36\x57\xbe\xf2\x8e\xc9\xb1\xe8\xfb\x4d\xaa\xcd\x19\x94\xc5\x18\x33\x61\xe9\x4c\x56\xe4\x1b\x80\xbb\x6a\x62\xe9\xc8\x23\xb4\x06\x1a\xa6\x7a\x06\x01\xd2\x72\xbf\xa0\xa3\x2f\x30\x31\xd9\x8b\xfc\x99\x98\x8d\x3b\xc5\x54\x96\xfb\x59\x47\x6a\x73\x1e\xa7\xbd\xfe\xb9\x53\x0e\xe6\x54\xfe\xef\xdc\xa2\x65\xf6\x0e\x8c\xce\x84\x6f\xc8\x25\xb5\xd7\x68\xc5\xd2\xb1\xa0\x31\xb1\xc6\x55\xc3\x19\x6b\xb4\x56\xe4\x5f\xb5\xe1\x54\x4e\xf0\xba\xb1\xf9\xb8\xd9\xd5\x24\xfe\xd6\xf4\x58\x4d\x0a\x54\xd5\x29\xf8\xd6\x69\x91\xaf\x66\x9b\x68\x0a\x83\x5e\xa2\x95\x84\x9a\xc2\x3d\x9e\x0c\xc0\x35\x62\xb3\xc2\x98\x3e\x7e\x7e\x08\x84\x4c\xf0\xd9\xfb\xaa\xc6\x50\xa5\xb2\x79\xbb\x98\xdb\x56\xeb\x23\x81\x26\x4b\x4c\x02\x30\x18\x4c\x2c\x16\x64\x33\x31\x22\x83\xf1\xdb\x95\x4c\x0e\xe7\x5b\x24\xc5\xe6\x4e\x72\x70\x53\x78\x7d\x59\x42\xc9\x99\xb6\xb2\xca\x2f\x01\x4f\xc0\xa5\x89\x90\x22\x94\x81\x8e\x99\xf8\x89\x67\x8c\x2a\x98\x86\x1f\xad\x57\x38\x34\xee\x46\x64\x5b\x32\x37\x8f\x52\x76\x1d\xec\x9e\x7c\x64\x87\x35\x41\xdf\x4b\x6c\x1a\x59\xcc\x71\xa5\x1a\xe3\x7f\x6c\x83\xcd\xb6\x73\xee\xbf\xff\x3e\x90\x53\x5e\xd3\x8f\xef\xdf\x0e\xbe\x6e\xbc\x23\xc7\xef\x6f\x66\x3e\xfc\x5b\x62\xfb\x6e\x5c\x7c\xb6\x5d\x13\xd3\x5e\xb5\x35\x39\xde\x8d\x9c\x07\x91\x7f\xd2\x1a\x5e\x3c\x1c\x30\x68\xac\x7d\x9d\x4a\xcc\x61\xf1\x6e\x89\x81\x3d\x5c\x7c\x1b\x5e\x6b\x11\x0a\x0c\x43\xf5\x57\x6b\x54\x65\x2f\x80\x5a\x83\xa3\xdf\x57\xc2\x6e\x78\xd0\xd3\xcb\xfe\x13\x00\x00\xff\xff\x0f\x8a\xd6\x49\x09\x04\x00\x00"),
		},
		"/templates/show.page.tmpl": &vfsgen۰CompressedFileInfo{
			name:             "show.page.tmpl",
			modTime:          time.Date(2020, 11, 7, 20, 56, 6, 220595570, time.UTC),
			uncompressedSize: 790,

			compressedContent: []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x02\xff\x7c\x52\x4d\xab\x14\x31\x10\xbc\xcf\xaf\x68\xe2\x79\x26\xb0\x78\x92\x6c\x40\xdf\x7a\x10\xc4\x83\x3e\xcf\xd2\xbb\xe9\x9d\x09\x3b\x93\x19\x92\x7e\x4f\x25\xe6\xbf\x4b\xe6\x23\xfb\x81\x78\x4a\x52\xdd\x55\xd5\x45\x27\x46\x43\x67\xeb\x08\x04\x5b\xee\x49\xa4\x14\x63\xf3\xcd\x32\x3d\xe7\x67\x4a\xef\x20\xc6\xe6\xc3\x38\x5e\x06\xf4\x97\x66\x05\x63\x24\x67\x52\xaa\xaa\x2b\xfb\x34\x3a\x26\xc7\x22\xa5\x4a\x05\x3a\xb1\x1d\x1d\x58\xb3\x2f\xb8\xae\x00\x00\x62\xfc\x69\xb9\x83\x22\x98\xd2\x0c\xab\x6e\x07\xa7\x1e\x43\xd8\x2f\x53\x80\x0d\xf5\x4e\x68\x85\xd0\x79\x3a\xef\x45\x8c\xcd\xf7\xaf\x9f\x53\x12\xc0\xe8\x5b\xe2\xbd\xf8\x71\xec\xd1\x5d\x04\x60\xcf\x73\x79\x1d\x4c\xe8\xeb\x5d\x49\xd4\x4a\x76\xbb\xc5\x59\x4d\x9b\x43\x78\x39\x16\x93\xb7\x42\x2b\xb6\x03\x6d\xb5\x0e\x43\xcd\xf4\x8b\xeb\xd6\xd3\x6f\x01\x06\x99\x72\x79\xb6\x78\xf2\x84\x4c\xe6\x3d\xaf\x36\xe5\x0d\x7f\xa0\x7b\x19\xd0\x1d\x90\x67\xdb\xcc\xd0\x6a\x5a\x7c\x73\x66\x7b\x86\xe6\x19\xdb\xb0\xa6\x05\x50\xc6\xbe\x96\xc0\xd8\x06\xb1\xf5\xe6\x6e\x8f\xae\xa5\x07\xc2\x4c\x0a\x13\xba\x1b\x56\x9e\xbf\xb7\xee\x92\x4f\x37\xfa\x01\xfb\x05\x69\x3b\xbe\xd1\x5b\x34\xa1\xf9\x82\x03\xc1\xbd\x9e\xcc\x82\xb7\xce\xcb\x52\xb7\xaa\xb1\xaf\xd7\x08\x77\xa5\x9b\xe9\xef\xd7\x9b\x3b\x9b\xa7\x05\xfa\xb7\x92\x3a\xfa\x72\xdd\xb6\x2b\x71\x9a\xe4\x71\xfd\x11\x32\xc6\xe6\xd3\x21\x25\x49\xc6\xb2\x28\x36\xe8\x4d\x7d\x1e\x47\x26\x5f\x5b\xa6\x41\xe8\x8f\xc6\x72\xde\xf0\xa3\xd8\x9b\xff\x70\x0e\xd4\x13\x53\x61\x6d\xa9\x94\x5c\x3f\xac\xae\x36\xe8\x6f\x00\x00\x00\xff\xff\x86\x8f\x40\x99\x16\x03\x00\x00"),
		},
		"/templates/signup.page.tmpl": &vfsgen۰CompressedFileInfo{
			name:             "signup.page.tmpl",
			modTime:          time.Date(2020, 9, 4, 1, 0, 37, 73625636, time.UTC),
			uncompressedSize: 1399,

			compressedContent: []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x02\xff\xc4\x94\xc1\x6e\xdb\x30\x0c\x86\xef\x7e\x0a\x42\x97\x9c\x12\x23\x77\x5b\x97\xa1\xd9\x65\xd8\x8a\x65\x2f\xa0\x58\x74\x2c\x40\x96\x32\x89\x4a\x37\x08\x7a\xf7\xc1\xb2\x9d\x3a\x85\x5b\x74\xc3\x80\x9e\x6c\x92\x3f\xc9\x5f\x9f\x00\xc5\x28\xb1\x55\x06\x81\x91\x22\x8d\x2c\xa5\xa3\x3a\x9b\x70\x89\x11\x8d\x4c\xa9\x28\x9e\x05\x8d\x35\x84\x86\x58\x4a\x45\x25\xd5\x15\x1a\x2d\xbc\xaf\x59\x63\x75\xe8\x8d\x67\x7c\x25\x0b\xca\x6f\xad\xc1\xed\xcf\x20\x1c\xa1\xcb\x61\xdb\x7a\xa4\x9c\xa5\x4e\x39\x39\xf4\x75\xfb\xb9\x2d\x9b\x18\x64\x7b\xc6\x47\x23\x55\xd9\xed\x79\x51\xb5\xd6\xf5\xb3\x48\x04\xea\x18\x88\x86\x94\x35\xf5\xa6\xf4\x59\xb7\x81\x1e\xa9\xb3\xb2\xde\x3c\x7e\x3b\xfe\xd8\x80\xb1\x57\xa1\x95\x14\x84\xbc\x00\x00\x88\x11\x76\x9f\x8e\xdf\x0f\x07\x85\x5a\x42\x4a\x53\xf2\x49\x51\x07\xbb\x83\x75\xfd\x94\x5a\xa6\x1f\x9c\xb3\xce\xef\x3e\x23\x01\x3b\xa3\x41\xa7\x1a\xb6\x90\x2d\xcf\x8b\x83\x94\xf1\x18\x77\x29\x55\xa5\x54\x57\xbe\x98\x36\xa2\x5c\x6b\x6b\x07\x37\xec\x59\x9b\xeb\x5a\x9c\x50\xcf\x8a\x1c\x30\xfe\x55\xf4\x58\x95\x39\xb8\x57\xaf\x79\x35\xa2\xc7\xa5\xd1\xb5\xcd\x9d\xf0\x5b\xc2\x5f\xb4\x95\xc2\x9c\xf1\x15\xeb\x6b\xf6\xf3\x20\x65\x2e\x81\xe6\x51\x39\x60\x40\xbf\x2f\x58\xb3\x61\x26\x83\xc1\x41\x3d\xfa\x80\xab\xd0\x01\x6b\x16\xe3\x9d\x39\x06\x22\x90\x6d\x6d\x13\xc6\x1b\x1d\x7f\x17\x28\x5e\x58\xf9\x27\x6a\x0f\xbd\x50\xfa\xdd\xd8\x70\x50\xaf\x72\xbb\x9b\xfd\x1a\xb9\xd5\x2d\x7f\xc5\x6e\x34\x30\xc1\x9b\x82\x17\xf4\x66\x8f\xff\x19\xd4\xa3\xf0\xfe\xc9\x3a\xf9\x6e\x56\x97\xa9\xe1\x23\x71\xdd\x3c\x4c\xc4\x6e\xf1\x9b\x6c\xf8\x1b\xe3\x4f\x81\xc8\xe6\x57\xcb\x87\xa6\x41\xef\xe7\x55\x3e\x9c\x7a\x45\xb7\xdb\xf8\x62\xcf\xca\xac\xaf\x99\xcf\x50\x95\xc3\x93\xc5\x8b\xa9\x34\x7d\xe6\xea\x9f\x00\x00\x00\xff\xff\x88\x4d\x88\xbb\x77\x05\x00\x00"),
		},
		"/templates/tags.page.tmpl": &vfsgen۰CompressedFileInfo{
			name:             "tags.page.tmpl",
			modTime:          time.Date(2020, 10, 26, 13, 37, 19, 918548888, time.UTC),
			uncompressedSize: 1059,

			compressedContent: []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x02\xff\x84\x53\xc1\x6e\xdb\x30\x0c\xbd\xfb\x2b\x08\xee\x90\x53\x6c\xa0\x67\x5b\x97\x62\xdd\x4e\xdb\xb0\xe6\x3e\xa8\x16\x6d\x0b\x95\x25\x43\xa2\xdb\x15\x82\xfe\x7d\x90\x9d\xd4\x4e\x5a\x60\x39\x44\xd2\xd3\xa3\xf8\xf8\x48\xc7\xa8\xa8\xd3\x96\x00\x59\xb3\x21\x4c\x29\xc6\xf2\x51\x33\x9d\xf2\x31\x25\x38\xc2\x77\x37\x52\x8c\x64\x55\x4a\x45\xb1\xf1\x5b\x67\x99\x2c\x63\x4a\x45\x3d\xdc\x41\x6b\x64\x08\xcd\xfa\x0a\xe8\x70\xbc\x43\x71\x92\x7d\xa8\xab\xe1\x4e\x14\xb5\xd2\x2f\x17\x46\xeb\xcc\x3c\xda\x80\x9f\xa1\x39\xd0\x59\x3a\xf2\xa0\xbd\xca\x8c\xce\xf9\x11\xb4\x6a\x90\x65\x7f\x6c\x3d\x49\x26\x04\xd9\xb2\x76\xb6\x39\x54\x72\x9a\x2a\x96\x7d\xb5\x5e\x1c\x60\x24\x1e\x9c\x6a\x0e\xbf\x7e\x3e\x9e\x0e\xa2\x00\x00\x88\x11\xca\xfb\xc7\xdf\x0f\x0f\x9a\x8c\x82\x94\xce\xe0\xab\xe6\x01\xca\x07\xe7\xc7\x33\xb4\x87\xbf\x7a\xef\x7c\x28\xbf\x11\x03\xf6\x64\xc9\xeb\x16\x77\xb4\xfc\xdb\x4b\xb7\x8e\x75\xa7\x5b\x99\x45\xe5\x02\x94\xb4\x3d\x79\x14\x57\x01\x00\xf5\xd3\xcc\xec\xec\x25\x4a\x91\x21\x26\x14\x75\xb5\xe2\xb7\xf4\x18\xcb\xdb\x9c\x95\xd2\x2f\x62\xa7\x76\x6d\xc9\x47\x41\xdd\x52\xeb\x20\xc3\x51\x2a\xe5\x16\xab\x3f\xd3\x9d\x1b\xe8\x9d\xb9\x12\x5a\x6b\x3b\xcd\x7c\x61\x2c\x07\x04\x7e\x9b\xa8\x41\xa6\xbf\x8c\x30\x19\xd9\xd2\xe0\x8c\x22\xdf\xe0\xfd\xe2\x3b\x58\x7a\x05\x96\x3d\x82\x95\x23\x35\x98\xff\x11\x5e\xa4\x99\xa9\xc1\x18\x57\x1f\x17\x30\x25\x04\x39\xb3\xeb\x5c\x3b\x87\x06\xdf\xb7\x7b\x81\xd7\x45\xfe\x57\xef\xd9\xd4\x55\x62\x98\x9f\x46\xcd\x78\x09\x38\xdf\xe9\x70\x34\xda\x3e\xa3\x58\xe5\x7e\x34\xfc\x2a\xe7\xee\x70\x71\xb8\xae\xf2\x18\x8a\xe2\x7c\x75\x5e\x8a\x4c\xd0\x1d\x94\x79\xcc\xdf\xfb\x10\xa3\xcf\xed\xbf\x41\xeb\xd9\xec\xd2\x19\x2d\x6a\x09\x83\xa7\xae\xc1\x2f\x08\x2c\x7d\x4f\xdc\xe0\x9f\x27\x23\xed\x33\x82\x34\xbc\xf8\xf6\x43\x8e\x94\x12\x8a\xf7\x6d\x5d\x49\x51\x57\x46\x6f\x52\xb7\x67\xb7\x69\x88\x91\x4c\xa0\x2d\xf5\x24\x4e\x03\x79\x02\xe9\x09\xac\xcb\x8d\x0a\xf0\x46\x5c\xc2\xc9\xbf\xc1\xf2\xe5\x68\xdb\x83\xb3\x54\xd6\xd5\x24\x8a\xed\xa9\xcb\xfa\x2f\x00\x00\xff\xff\xd6\x08\xb8\x65\x23\x04\x00\x00"),
		},
	}
	fs["/"].(*vfsgen۰DirInfo).entries = []os.FileInfo{
		fs["/static"].(os.FileInfo),
		fs["/templates"].(os.FileInfo),
	}
	fs["/static"].(*vfsgen۰DirInfo).entries = []os.FileInfo{
		fs["/static/css"].(os.FileInfo),
		fs["/static/img"].(os.FileInfo),
		fs["/static/js"].(os.FileInfo),
	}
	fs["/static/css"].(*vfsgen۰DirInfo).entries = []os.FileInfo{
		fs["/static/css/main.css"].(os.FileInfo),
	}
	fs["/static/img"].(*vfsgen۰DirInfo).entries = []os.FileInfo{
		fs["/static/img/favicon.ico"].(os.FileInfo),
	}
	fs["/static/js"].(*vfsgen۰DirInfo).entries = []os.FileInfo{
		fs["/static/js/main.js"].(os.FileInfo),
	}
	fs["/templates"].(*vfsgen۰DirInfo).entries = []os.FileInfo{
		fs["/templates/base.layout.tmpl"].(os.FileInfo),
		fs["/templates/bookmarks.page.tmpl"].(os.FileInfo),
		fs["/templates/create.page.tmpl"].(os.FileInfo),
		fs["/templates/create_tag.page.tmpl"].(os.FileInfo),
		fs["/templates/edit.page.tmpl"].(os.FileInfo),
		fs["/templates/home.page.tmpl"].(os.FileInfo),
		fs["/templates/login.page.tmpl"].(os.FileInfo),
		fs["/templates/settings.page.tmpl"].(os.FileInfo),
		fs["/templates/show.page.tmpl"].(os.FileInfo),
		fs["/templates/signup.page.tmpl"].(os.FileInfo),
		fs["/templates/tags.page.tmpl"].(os.FileInfo),
	}

	return fs
}()

type vfsgen۰FS map[string]interface{}

func (fs vfsgen۰FS) Open(path string) (http.File, error) {
	path = pathpkg.Clean("/" + path)
	f, ok := fs[path]
	if !ok {
		return nil, &os.PathError{Op: "open", Path: path, Err: os.ErrNotExist}
	}

	switch f := f.(type) {
	case *vfsgen۰CompressedFileInfo:
		gr, err := gzip.NewReader(bytes.NewReader(f.compressedContent))
		if err != nil {
			// This should never happen because we generate the gzip bytes such that they are always valid.
			panic("unexpected error reading own gzip compressed bytes: " + err.Error())
		}
		return &vfsgen۰CompressedFile{
			vfsgen۰CompressedFileInfo: f,
			gr:                        gr,
		}, nil
	case *vfsgen۰DirInfo:
		return &vfsgen۰Dir{
			vfsgen۰DirInfo: f,
		}, nil
	default:
		// This should never happen because we generate only the above types.
		panic(fmt.Sprintf("unexpected type %T", f))
	}
}

// vfsgen۰CompressedFileInfo is a static definition of a gzip compressed file.
type vfsgen۰CompressedFileInfo struct {
	name              string
	modTime           time.Time
	compressedContent []byte
	uncompressedSize  int64
}

func (f *vfsgen۰CompressedFileInfo) Readdir(count int) ([]os.FileInfo, error) {
	return nil, fmt.Errorf("cannot Readdir from file %s", f.name)
}
func (f *vfsgen۰CompressedFileInfo) Stat() (os.FileInfo, error) { return f, nil }

func (f *vfsgen۰CompressedFileInfo) GzipBytes() []byte {
	return f.compressedContent
}

func (f *vfsgen۰CompressedFileInfo) Name() string       { return f.name }
func (f *vfsgen۰CompressedFileInfo) Size() int64        { return f.uncompressedSize }
func (f *vfsgen۰CompressedFileInfo) Mode() os.FileMode  { return 0444 }
func (f *vfsgen۰CompressedFileInfo) ModTime() time.Time { return f.modTime }
func (f *vfsgen۰CompressedFileInfo) IsDir() bool        { return false }
func (f *vfsgen۰CompressedFileInfo) Sys() interface{}   { return nil }

// vfsgen۰CompressedFile is an opened compressedFile instance.
type vfsgen۰CompressedFile struct {
	*vfsgen۰CompressedFileInfo
	gr      *gzip.Reader
	grPos   int64 // Actual gr uncompressed position.
	seekPos int64 // Seek uncompressed position.
}

func (f *vfsgen۰CompressedFile) Read(p []byte) (n int, err error) {
	if f.grPos > f.seekPos {
		// Rewind to beginning.
		err = f.gr.Reset(bytes.NewReader(f.compressedContent))
		if err != nil {
			return 0, err
		}
		f.grPos = 0
	}
	if f.grPos < f.seekPos {
		// Fast-forward.
		_, err = io.CopyN(ioutil.Discard, f.gr, f.seekPos-f.grPos)
		if err != nil {
			return 0, err
		}
		f.grPos = f.seekPos
	}
	n, err = f.gr.Read(p)
	f.grPos += int64(n)
	f.seekPos = f.grPos
	return n, err
}
func (f *vfsgen۰CompressedFile) Seek(offset int64, whence int) (int64, error) {
	switch whence {
	case io.SeekStart:
		f.seekPos = 0 + offset
	case io.SeekCurrent:
		f.seekPos += offset
	case io.SeekEnd:
		f.seekPos = f.uncompressedSize + offset
	default:
		panic(fmt.Errorf("invalid whence value: %v", whence))
	}
	return f.seekPos, nil
}
func (f *vfsgen۰CompressedFile) Close() error {
	return f.gr.Close()
}

// vfsgen۰DirInfo is a static definition of a directory.
type vfsgen۰DirInfo struct {
	name    string
	modTime time.Time
	entries []os.FileInfo
}

func (d *vfsgen۰DirInfo) Read([]byte) (int, error) {
	return 0, fmt.Errorf("cannot Read from directory %s", d.name)
}
func (d *vfsgen۰DirInfo) Close() error               { return nil }
func (d *vfsgen۰DirInfo) Stat() (os.FileInfo, error) { return d, nil }

func (d *vfsgen۰DirInfo) Name() string       { return d.name }
func (d *vfsgen۰DirInfo) Size() int64        { return 0 }
func (d *vfsgen۰DirInfo) Mode() os.FileMode  { return 0755 | os.ModeDir }
func (d *vfsgen۰DirInfo) ModTime() time.Time { return d.modTime }
func (d *vfsgen۰DirInfo) IsDir() bool        { return true }
func (d *vfsgen۰DirInfo) Sys() interface{}   { return nil }

// vfsgen۰Dir is an opened dir instance.
type vfsgen۰Dir struct {
	*vfsgen۰DirInfo
	pos int // Position within entries for Seek and Readdir.
}

func (d *vfsgen۰Dir) Seek(offset int64, whence int) (int64, error) {
	if offset == 0 && whence == io.SeekStart {
		d.pos = 0
		return 0, nil
	}
	return 0, fmt.Errorf("unsupported Seek in directory %s", d.name)
}

func (d *vfsgen۰Dir) Readdir(count int) ([]os.FileInfo, error) {
	if d.pos >= len(d.entries) && count > 0 {
		return nil, io.EOF
	}
	if count <= 0 || count > len(d.entries)-d.pos {
		count = len(d.entries) - d.pos
	}
	e := d.entries[d.pos : d.pos+count]
	d.pos += count
	return e, nil
}
