package main

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

func static_index_html() ([]byte, error) {
	return bindata_read([]byte{
		0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x00, 0xff, 0xe4, 0x59,
		0xeb, 0x6f, 0xdc, 0xb8, 0x11, 0xff, 0x9e, 0xbf, 0x82, 0xa7, 0xa4, 0xf0,
		0x1a, 0x67, 0x49, 0x6b, 0xd7, 0x06, 0x9a, 0xf5, 0xae, 0x8b, 0x3e, 0x0e,
		0x28, 0xfa, 0xa1, 0x39, 0x20, 0xe9, 0x87, 0xe2, 0x70, 0x30, 0xb8, 0xe2,
		0xac, 0x97, 0x31, 0x25, 0xaa, 0x22, 0xb5, 0xeb, 0xed, 0x9d, 0xff, 0xf7,
		0x0e, 0x49, 0x3d, 0xa8, 0xc7, 0xbe, 0x90, 0xe6, 0x10, 0xa0, 0x08, 0x90,
		0x95, 0x38, 0xc3, 0x79, 0xfc, 0xe6, 0xc1, 0xa1, 0x3c, 0xff, 0xee, 0xaf,
		0x1f, 0xfe, 0xf2, 0xe9, 0x5f, 0x3f, 0xfe, 0x40, 0xd6, 0x3a, 0x15, 0x0f,
		0x6f, 0xe6, 0xee, 0x87, 0x90, 0xf9, 0x1a, 0x28, 0x33, 0x0f, 0xf8, 0x98,
		0x82, 0xa6, 0x24, 0xa3, 0x29, 0x2c, 0x82, 0x0d, 0x87, 0x6d, 0x2e, 0x0b,
		0x1d, 0x90, 0x44, 0x66, 0x1a, 0x32, 0xbd, 0x08, 0xb6, 0x9c, 0xe9, 0xf5,
		0x82, 0xc1, 0x86, 0x27, 0x10, 0xda, 0x97, 0x2b, 0xc2, 0x33, 0xae, 0x39,
		0x15, 0xa1, 0x4a, 0xa8, 0x80, 0xc5, 0x75, 0x34, 0xbd, 0x22, 0x29, 0x7d,
		0xe1, 0x69, 0x99, 0xfa, 0x4b, 0xa5, 0x82, 0xc2, 0xbe, 0xd3, 0x25, 0x2e,
		0x65, 0x32, 0x20, 0xf1, 0x50, 0x23, 0xcd, 0x73, 0x01, 0x61, 0x2a, 0x97,
		0x1c, 0x7f, 0xb6, 0xb0, 0x0c, 0x71, 0x21, 0x4c, 0x68, 0x6e, 0xf6, 0x78,
		0x56, 0xec, 0x40, 0xb5, 0xdb, 0x05, 0xcf, 0x9e, 0xc9, 0xba, 0x80, 0xd5,
		0x22, 0x58, 0x6b, 0x9d, 0xcf, 0xe2, 0x18, 0xb5, 0x27, 0x2c, 0x8b, 0x96,
		0x52, 0x6a, 0xa5, 0x0b, 0x9a, 0x9b, 0x97, 0x44, 0xa6, 0xf1, 0x0a, 0xf7,
		0x87, 0x74, 0x0b, 0x4a, 0xa6, 0x10, 0xdf, 0x46, 0x37, 0xd1, 0x34, 0x4e,
		0x94, 0xea, 0x2c, 0x47, 0x29, 0x47, 0x5e, 0x85, 0xe2, 0x0b, 0x10, 0x8b,
		0x40, 0xe9, 0x9d, 0x00, 0xb5, 0x06, 0xd0, 0xad, 0x3e, 0x95, 0x14, 0x3c,
		0xd7, 0x44, 0xef, 0x72, 0x34, 0x58, 0xc3, 0x8b, 0x8e, 0x3f, 0xd3, 0x0d,
		0x75, 0xab, 0x01, 0x51, 0x45, 0xe2, 0xec, 0x50, 0x68, 0x48, 0x22, 0x19,
		0x44, 0x9f, 0xff, 0x5d, 0x42, 0xb1, 0xb3, 0x06, 0xb8, 0xc7, 0xf0, 0x26,
		0xba, 0xc6, 0x7f, 0x46, 0xd3, 0x67, 0x15, 0x3c, 0xcc, 0x63, 0xb7, 0xb7,
		0x12, 0xaf, 0xb9, 0x16, 0xf0, 0x20, 0xd3, 0x97, 0x5c, 0xd0, 0x1d, 0x14,
		0xf3, 0xd8, 0x2d, 0x98, 0x28, 0xc5, 0x75, 0x98, 0xe6, 0xd6, 0x2e, 0xb7,
		0xc1, 0xc4, 0x90, 0xfc, 0x62, 0x1f, 0xf1, 0x05, 0xf8, 0xd3, 0x5a, 0xcf,
		0xc8, 0xf5, 0x74, 0xfa, 0xbb, 0x7b, 0xbb, 0xf6, 0xfa, 0xc6, 0xfe, 0x2c,
		0x25, 0xdb, 0x35, 0x5c, 0x39, 0x65, 0x8c, 0x67, 0x4f, 0x33, 0x32, 0xcd,
		0x5f, 0xee, 0xab, 0xb5, 0x94, 0x16, 0x4f, 0x3c, 0xeb, 0x2c, 0x0d, 0x85,
		0x11, 0x62, 0xb1, 0x5a, 0xd1, 0x94, 0x8b, 0xdd, 0x8c, 0xfc, 0x0d, 0xc4,
		0x06, 0x34, 0x4f, 0x68, 0x47, 0x55, 0xb4, 0x2c, 0xb5, 0x96, 0x59, 0xa3,
		0x2d, 0x29, 0x0b, 0x25, 0x8b, 0x19, 0xc9, 0x25, 0xc7, 0xf0, 0x15, 0x3d,
		0xde, 0x42, 0x6e, 0x15, 0x34, 0xbc, 0x8c, 0x2b, 0xe3, 0xf6, 0x8c, 0x64,
		0x32, 0x83, 0x51, 0xce, 0x52, 0x9c, 0xe7, 0xc6, 0x70, 0xbb, 0xe0, 0xe7,
		0x02, 0x21, 0xb8, 0xd2, 0xa1, 0x85, 0x3c, 0x34, 0x51, 0xf7, 0xad, 0xf3,
		0x89, 0xfd, 0xf5, 0x0c, 0xc2, 0x06, 0xc2, 0xe8, 0x0e, 0xd2, 0x9a, 0xb2,
		0x94, 0x05, 0xc3, 0x5a, 0x58, 0x4a, 0x84, 0x29, 0x45, 0x5a, 0xfe, 0x42,
		0x94, 0x14, 0x9c, 0x91, 0xb7, 0x49, 0x92, 0xdc, 0xf7, 0x4d, 0xfb, 0x43,
		0x6b, 0x87, 0x05, 0x5f, 0xf1, 0xff, 0x80, 0x11, 0x78, 0x5d, 0x0b, 0x1c,
		0xf5, 0xb0, 0xf5, 0x31, 0x91, 0xc2, 0xa0, 0xff, 0xf6, 0xfd, 0xfb, 0xf7,
		0x5d, 0x17, 0xc3, 0xc2, 0xd9, 0x76, 0x7b, 0x10, 0x2a, 0xda, 0x08, 0x32,
		0xa9, 0x1e, 0x32, 0x48, 0x64, 0x41, 0x35, 0x97, 0x59, 0xd7, 0xdb, 0x5a,
		0xcb, 0x74, 0x3a, 0xad, 0x97, 0xb6, 0x6b, 0xae, 0x21, 0x54, 0x39, 0x4d,
		0x2c, 0x32, 0x5b, 0x2c, 0xc4, 0x9a, 0x24, 0x37, 0x50, 0xac, 0x84, 0xdc,
		0xce, 0xc8, 0x9a, 0x33, 0x06, 0xd9, 0xbd, 0xaf, 0xa2, 0x25, 0x82, 0x10,
		0x3c, 0x57, 0x5c, 0xdd, 0xf7, 0xb3, 0x63, 0x29, 0x64, 0xf2, 0xdc, 0x35,
		0xda, 0xf4, 0x86, 0x42, 0x0a, 0x75, 0x28, 0x95, 0x10, 0x55, 0xa9, 0xb8,
		0x33, 0x1e, 0xcb, 0x1b, 0xdd, 0xd8, 0x34, 0x14, 0xdb, 0xcb, 0xba, 0xb9,
		0x3e, 0x56, 0x4c, 0x5d, 0x5d, 0x5e, 0x26, 0x35, 0x82, 0xe9, 0x12, 0x83,
		0x59, 0xea, 0x46, 0xb0, 0x96, 0xf9, 0x8c, 0xdc, 0xa1, 0x04, 0x22, 0x60,
		0xa5, 0xdd, 0x63, 0x45, 0x32, 0x1d, 0xee, 0x99, 0xeb, 0x10, 0x7b, 0x54,
		0xa6, 0x56, 0xb2, 0xc0, 0x64, 0xb0, 0x8f, 0x68, 0x19, 0x4c, 0x42, 0x64,
		0xbc, 0x22, 0xe6, 0xff, 0xcb, 0x3d, 0x9e, 0x46, 0x05, 0x6c, 0x79, 0xc6,
		0x46, 0x6c, 0x58, 0xf1, 0x17, 0x60, 0x1d, 0x03, 0xa6, 0x7d, 0xa7, 0x6e,
		0x5a, 0x33, 0x06, 0xae, 0xef, 0xd5, 0xd4, 0x2f, 0xed, 0xa3, 0x70, 0xde,
		0x1d, 0x42, 0x13, 0x13, 0x5a, 0x48, 0x8a, 0x4b, 0x06, 0x98, 0x4e, 0x0a,
		0x50, 0xc1, 0x9f, 0x50, 0x68, 0x02, 0x6d, 0xc3, 0xe8, 0x64, 0xff, 0x8d,
		0x5f, 0x4c, 0x2f, 0x66, 0xd1, 0x96, 0x4a, 0x53, 0x58, 0x2f, 0x7b, 0x4a,
		0xed, 0xb6, 0x2d, 0x35, 0x00, 0x38, 0xea, 0x2c, 0x4d, 0x9e, 0xb7, 0xb4,
		0x60, 0xd8, 0xef, 0x94, 0x6e, 0x7c, 0xae, 0x44, 0x56, 0xc5, 0x73, 0x73,
		0x96, 0x44, 0x8c, 0xf1, 0x5e, 0x81, 0x2e, 0x3b, 0x4e, 0x93, 0xb7, 0xc1,
		0x0c, 0x4b, 0xe1, 0x58, 0xe0, 0x6b, 0xaf, 0xbf, 0x30, 0xf6, 0x95, 0xb2,
		0x6f, 0x25, 0xf6, 0xbf, 0x3f, 0x33, 0xf6, 0x36, 0xfd, 0x4f, 0x0b, 0x7c,
		0xed, 0xa9, 0xfb, 0x0d, 0xcb, 0xfc, 0x7f, 0x10, 0xa3, 0x5a, 0x18, 0x93,
		0xdb, 0xec, 0x8b, 0x72, 0x28, 0xa5, 0x7c, 0x0c, 0xfc, 0x61, 0xa9, 0x7b,
		0xd1, 0xad, 0x33, 0xe0, 0xbc, 0x80, 0x5b, 0x4d, 0x91, 0x69, 0x9e, 0xbf,
		0x65, 0x7b, 0x6b, 0xb3, 0x84, 0x96, 0x5a, 0xf6, 0xcc, 0xbd, 0xf5, 0x7d,
		0x72, 0xd8, 0x51, 0xc6, 0x4b, 0xd5, 0x4d, 0xa9, 0xe6, 0xe8, 0xee, 0x88,
		0x18, 0xc9, 0x2a, 0x32, 0x4c, 0xab, 0xdb, 0x5e, 0x5a, 0xad, 0x29, 0x33,
		0x47, 0x10, 0xcf, 0x14, 0x68, 0x14, 0x68, 0xfe, 0x5d, 0xe3, 0x48, 0x40,
		0xde, 0x32, 0xc6, 0x8e, 0x03, 0x37, 0x5b, 0x02, 0x3a, 0x3c, 0x32, 0xd7,
		0x78, 0x27, 0x17, 0xa9, 0xe7, 0xd9, 0x19, 0xb9, 0x20, 0x17, 0xbd, 0x73,
		0xdf, 0xa5, 0xed, 0xf1, 0x28, 0xd5, 0x3d, 0xea, 0x6b, 0x46, 0xea, 0xe6,
		0xfa, 0xee, 0x8c, 0x50, 0xf9, 0xe9, 0x87, 0xc6, 0x3d, 0x15, 0xb2, 0xcc,
		0x98, 0x71, 0xbc, 0x84, 0xaf, 0x11, 0x42, 0x5f, 0x47, 0x5b, 0x3e, 0xe3,
		0xa7, 0xc5, 0x3e, 0x10, 0xab, 0xb6, 0xfc, 0x15, 0x31, 0xbc, 0xfe, 0x3f,
		0x81, 0xb0, 0xca, 0xfb, 0xab, 0xc3, 0xc9, 0xfa, 0x9b, 0x55, 0x87, 0x2a,
		0x97, 0xf6, 0xf6, 0xa4, 0x4e, 0x0d, 0xed, 0x8d, 0x99, 0xfa, 0x49, 0xdd,
		0x97, 0x47, 0xee, 0x42, 0x87, 0xa3, 0xf5, 0x25, 0x0d, 0xa8, 0x73, 0x45,
		0xf0, 0xb7, 0x58, 0x6b, 0xfa, 0x73, 0x36, 0xa5, 0xf4, 0x88, 0xeb, 0xe8,
		0xcf, 0x79, 0x5e, 0x57, 0x87, 0xdb, 0xb7, 0xe1, 0xb4, 0x3f, 0x16, 0x8c,
		0xfb, 0x8c, 0xd7, 0xe6, 0xfa, 0x12, 0x3c, 0x37, 0xb7, 0xdb, 0xea, 0xfa,
		0xcc, 0xf8, 0x86, 0x24, 0x82, 0x2a, 0xb5, 0x08, 0xdc, 0x3d, 0x26, 0x78,
		0xa8, 0xc4, 0xcc, 0x4b, 0x81, 0x77, 0xed, 0x52, 0x54, 0x8c, 0x31, 0x72,
		0x3e, 0xbc, 0x19, 0x6c, 0xaa, 0x71, 0x6c, 0xb7, 0x79, 0x44, 0x37, 0xc7,
		0x35, 0xa4, 0x9e, 0x3a, 0x37, 0x1f, 0xd1, 0xc4, 0x20, 0x4d, 0x3a, 0xc3,
		0x63, 0x40, 0x18, 0xd5, 0x74, 0x11, 0x28, 0x80, 0xe7, 0x47, 0x43, 0x79,
		0xb4, 0xab, 0x0f, 0x73, 0x5e, 0x6f, 0x5e, 0x51, 0xb2, 0xa2, 0x96, 0x37,
		0xac, 0x77, 0x9a, 0x2f, 0x03, 0xfc, 0xa1, 0xb2, 0xf3, 0xb8, 0x42, 0x7f,
		0xb6, 0xec, 0xe8, 0xab, 0x08, 0x87, 0x54, 0x56, 0x2c, 0xbe, 0xc6, 0x4a,
		0x5f, 0x0d, 0x82, 0x6f, 0x85, 0x6f, 0x83, 0x49, 0xb6, 0x3d, 0x78, 0xb4,
		0xc5, 0xe7, 0x2c, 0x6c, 0x8c, 0xaa, 0xd7, 0x83, 0x87, 0x8f, 0xff, 0xfc,
		0xf3, 0xc7, 0x03, 0x0e, 0x9a, 0x1c, 0x46, 0xa6, 0x4f, 0x1f, 0x7e, 0x3c,
		0x19, 0x85, 0x06, 0xbc, 0x3e, 0xe2, 0x43, 0xcf, 0xc7, 0x70, 0x6e, 0x1c,
		0x3e, 0x01, 0xe9, 0x31, 0x90, 0x47, 0xf0, 0x1d, 0x42, 0x7b, 0x82, 0x0a,
		0xd3, 0x0e, 0x6b, 0xf9, 0x39, 0x2d, 0x31, 0x8b, 0x09, 0x67, 0x8b, 0xa0,
		0x02, 0x72, 0xa0, 0xc4, 0xb2, 0x0f, 0x35, 0xec, 0x8d, 0x9a, 0x9b, 0x49,
		0x4f, 0xca, 0x63, 0x3b, 0xb7, 0x7a, 0x33, 0x6c, 0x6d, 0x95, 0x5b, 0x7a,
		0xb4, 0x4b, 0x03, 0x7b, 0x7c, 0xfe, 0xf3, 0x1c, 0xc7, 0x99, 0xbb, 0x99,
		0xbe, 0x7b, 0xaa, 0x70, 0x61, 0x9f, 0x22, 0x4b, 0x3a, 0xe0, 0x7d, 0xa7,
		0xd8, 0x0f, 0x7e, 0xbf, 0xab, 0x37, 0xaf, 0xca, 0xcc, 0x19, 0x94, 0xc8,
		0x34, 0xa5, 0x19, 0x9b, 0x98, 0xef, 0x93, 0x97, 0x4d, 0x43, 0x25, 0xe4,
		0x5d, 0xf4, 0x04, 0xfa, 0xef, 0x1f, 0x3f, 0xfc, 0x63, 0x12, 0xc4, 0x15,
		0x4f, 0x1c, 0x90, 0xef, 0xed, 0x77, 0xcc, 0x2b, 0xf2, 0xcb, 0xeb, 0x55,
		0x23, 0x62, 0x52, 0x80, 0xca, 0xfd, 0xad, 0xf6, 0x68, 0xc3, 0x36, 0x0c,
		0x91, 0x90, 0x4f, 0x8e, 0x7a, 0xdf, 0x10, 0x5f, 0x9b, 0xe7, 0xaa, 0xb3,
		0x7b, 0xb6, 0x98, 0x30, 0x4f, 0x72, 0xaa, 0xd7, 0xfb, 0x0c, 0xb1, 0x79,
		0x80, 0xda, 0xf1, 0x56, 0x60, 0x3e, 0x47, 0x19, 0x56, 0x72, 0xd8, 0x12,
		0xbe, 0x22, 0x76, 0x15, 0x8f, 0xca, 0x24, 0x01, 0xa5, 0xba, 0x54, 0x94,
		0x3e, 0x09, 0xaa, 0xef, 0x41, 0xc1, 0x65, 0xb4, 0xe6, 0x0c, 0x26, 0x9e,
		0xad, 0x35, 0x43, 0xd3, 0x33, 0x2f, 0x23, 0xb5, 0x96, 0xdb, 0x2e, 0xcb,
		0xeb, 0x49, 0xae, 0x39, 0x1d, 0x43, 0xe7, 0x26, 0x41, 0x29, 0x8c, 0x66,
		0x9d, 0x8a, 0x49, 0x10, 0xa0, 0x80, 0x86, 0x64, 0x9d, 0x5b, 0xb8, 0x9f,
		0x5f, 0x7f, 0x25, 0x41, 0xe0, 0xd1, 0x7c, 0x4c, 0x2a, 0xeb, 0x0d, 0x2a,
		0x86, 0xf7, 0x24, 0x54, 0xe2, 0x98, 0xfc, 0x89, 0x31, 0xbc, 0xd8, 0xea,
		0xb2, 0x70, 0x3d, 0x85, 0x80, 0x80, 0x14, 0x8f, 0x31, 0x03, 0xd8, 0x16,
		0x08, 0xce, 0x4a, 0x50, 0x10, 0x8a, 0xb2, 0x0a, 0xb3, 0xc8, 0x78, 0x01,
		0x89, 0x96, 0xc5, 0xae, 0x87, 0xac, 0x55, 0xf5, 0xdd, 0x02, 0x8d, 0xeb,
		0xe3, 0xca, 0x35, 0xa4, 0xaa, 0x32, 0x3f, 0xc2, 0xe9, 0x87, 0x6b, 0x34,
		0x15, 0xfd, 0x73, 0x84, 0x28, 0x97, 0x79, 0x1f, 0x68, 0x7b, 0x6c, 0x54,
		0x4e, 0x5b, 0xa6, 0x9f, 0xa6, 0x3f, 0xd7, 0x8e, 0xf7, 0x02, 0x32, 0x17,
		0x58, 0x0a, 0xd4, 0x7d, 0x13, 0xbf, 0x78, 0x7b, 0xe1, 0xca, 0xe8, 0xc2,
		0x64, 0x67, 0x2b, 0xe4, 0x7b, 0x12, 0x5c, 0xb4, 0xe5, 0x74, 0xe1, 0xca,
		0x49, 0xc0, 0x06, 0x04, 0x56, 0xd3, 0x85, 0xad, 0x26, 0xcb, 0x3d, 0x8f,
		0x29, 0xbe, 0xa0, 0x40, 0x8c, 0x02, 0xcd, 0x73, 0xc8, 0xd8, 0x27, 0xd9,
		0xe4, 0x04, 0x31, 0xc1, 0xe9, 0x06, 0xdb, 0x7b, 0xc1, 0xd6, 0x47, 0x26,
		0x1c, 0xaf, 0x50, 0x64, 0x88, 0x30, 0xb1, 0x29, 0x8a, 0xae, 0x18, 0xd2,
		0x4f, 0xfc, 0xe7, 0xae, 0x0f, 0xa6, 0x88, 0x90, 0x66, 0x58, 0x22, 0xf3,
		0x9f, 0x79, 0xbf, 0x7f, 0xd3, 0x05, 0xf0, 0x20, 0xbc, 0x8d, 0x88, 0xda,
		0xd7, 0xa6, 0x36, 0xbb, 0x8a, 0x5e, 0xbb, 0x42, 0x31, 0xee, 0x3f, 0x70,
		0xbd, 0xb6, 0xa1, 0xb5, 0xf6, 0x49, 0xf3, 0x34, 0x16, 0x5d, 0x62, 0x41,
		0x7d, 0x34, 0x4d, 0xa4, 0x36, 0xb4, 0x61, 0x23, 0x7f, 0x24, 0x01, 0xbe,
		0x04, 0x64, 0x46, 0x02, 0x43, 0xe9, 0x05, 0x28, 0x51, 0xea, 0xd1, 0xa2,
		0x3e, 0xba, 0xd1, 0x1e, 0x19, 0x02, 0xb3, 0xcb, 0x6d, 0xc7, 0x37, 0x2e,
		0xd2, 0xa0, 0xe7, 0xfc, 0x9e, 0x18, 0xdb, 0xaf, 0xd3, 0x2e, 0xd0, 0xad,
		0x75, 0x26, 0xd0, 0x5e, 0x06, 0x58, 0x5c, 0x86, 0xc1, 0x37, 0xa4, 0xd6,
		0x30, 0x47, 0x37, 0x29, 0x60, 0xd6, 0x3b, 0x61, 0x30, 0xb4, 0x73, 0x73,
		0xa2, 0x79, 0x1a, 0x36, 0x80, 0x77, 0x93, 0xa6, 0x0c, 0xf7, 0x35, 0x35,
		0xa5, 0xa9, 0x2e, 0x55, 0x70, 0xb4, 0xa9, 0x36, 0xad, 0xac, 0x28, 0xb3,
		0x0c, 0x2f, 0x09, 0x63, 0xad, 0xec, 0xc4, 0x4e, 0x45, 0xb0, 0xdc, 0xbd,
		0xbf, 0x4a, 0xb4, 0xfb, 0x9b, 0x56, 0x38, 0xdc, 0x4d, 0xea, 0x1e, 0x76,
		0xc0, 0x79, 0xbf, 0xaf, 0xd9, 0x09, 0x1d, 0x25, 0xa1, 0x2f, 0x41, 0x22,
		0x38, 0x4e, 0x28, 0x9e, 0x77, 0xfd, 0xe3, 0xc2, 0x9d, 0x41, 0x6e, 0x20,
		0xea, 0xc8, 0x37, 0xf3, 0xa8, 0xdc, 0xe2, 0x41, 0x92, 0xd8, 0x8f, 0xf2,
		0x51, 0x01, 0x42, 0x52, 0x36, 0xe9, 0x1d, 0x28, 0x1d, 0xbd, 0xd5, 0x10,
		0x71, 0x96, 0xe6, 0x77, 0x13, 0xbd, 0xe6, 0x0a, 0xa3, 0xad, 0x75, 0x31,
		0x09, 0x4c, 0x36, 0x05, 0x97, 0x07, 0x74, 0x98, 0x79, 0xbc, 0xa7, 0x20,
		0xa0, 0xfb, 0xb5, 0x6c, 0x68, 0x51, 0xb7, 0xf3, 0x31, 0x45, 0xf7, 0x3d,
		0xd6, 0xaa, 0xe8, 0x86, 0xac, 0x36, 0xff, 0x3b, 0x87, 0x84, 0xcb, 0x0a,
		0xb7, 0x61, 0x51, 0x55, 0x63, 0x3f, 0x2f, 0xda, 0x53, 0xf5, 0x8c, 0x64,
		0xf0, 0xcf, 0xab, 0xa3, 0xc9, 0x5e, 0x3d, 0xf9, 0x7f, 0xe8, 0x9b, 0xc7,
		0xee, 0xce, 0x32, 0x8f, 0xed, 0xdf, 0x60, 0xff, 0x1b, 0x00, 0x00, 0xff,
		0xff, 0x5f, 0xa6, 0xac, 0x72, 0x9a, 0x1d, 0x00, 0x00,
	},
		"static/index.html",
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
	"static/index.html": static_index_html,
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
	"static": &_bintree_t{nil, map[string]*_bintree_t{
		"index.html": &_bintree_t{static_index_html, map[string]*_bintree_t{
		}},
	}},
}}
