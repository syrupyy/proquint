// Copyright 2016 The Upspin Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package antiquint converts integers to/from sets of pronounceable five-letter words.
// This implementation differs from the original proquint package in that it uses a reversed
// byte order for encoding and decoding. This is to ensure that the output is compatible
// with existing systems that use proquints encoded using little-endian byte order.
package antiquint

import (
	"bytes"
	"fmt"
	"strings"
)

var (
	// Consonants and vowels used in proquints.
	Consonants = []byte("bdfghjklmnprstvz")
	Vowels     = []byte("aiou")
)

// Encode returns a five-letter word, representing a uint16.
func Encode(x uint16) string {
	x = x%256*256 + x/256
	c3 := x & 0x0f
	x >>= 4
	v2 := x & 0x03
	x >>= 2
	c2 := x & 0x0f
	x >>= 4
	v1 := x & 0x03
	x >>= 2
	c1 := x & 0x0f
	b := make([]byte, 5)
	b[0] = Consonants[c1]
	b[1] = Vowels[v1]
	b[2] = Consonants[c2]
	b[3] = Vowels[v2]
	b[4] = Consonants[c3]
	return string(b[:])
}

// EncodeBytes returns any amount of five-letter words, representing a byte array.
func EncodeBytes(b []byte, sep string) string {
	if len(b)%2 == 1 {
		b = append([]byte{0}, b...)
	}
	l := len(b) / 2
	s := make([]string, l)
	for i := 0; i < len(b); i += 2 {
		s[l-i/2-1] = Encode(uint16(b[i])<<8 | uint16(b[i+1]))
	}
	return strings.Join(s, sep)
}

// EncodeUint32 returns two five-letter words, representing a uint32.
func EncodeUint32(x uint32, sep string) string {
	return EncodeBytes([]byte{byte(x >> 24), byte(x >> 16), byte(x >> 8), byte(x)}, sep)
}

// EncodeUint64 returns four five-letter words, representing a uint64.
func EncodeUint64(x uint64, sep string) string {
	return EncodeBytes([]byte{byte(x >> 56), byte(x >> 48), byte(x >> 40), byte(x >> 32), byte(x >> 24), byte(x >> 16), byte(x >> 8), byte(x)}, sep)
}

// Decode parses a five-letter word, returning a uint16.
func Decode(s string) (uint16, error) {
	if len(s) != 5 {
		return 0, fmt.Errorf("invalid proquint length: %d", len(s))
	}
	b := make([]uint16, 5)
	for i := 0; i < 5; i++ {
		var n int
		if i%2 == 0 {
			n = bytes.IndexByte(Consonants, s[i])
		} else {
			n = bytes.IndexByte(Vowels, s[i])
		}
		if n == -1 {
			return 0, fmt.Errorf("invalid character at position %d: %c", i, s[i])
		}
		b[i] = uint16(n)
	}
	x := b[0]<<12 + b[1]<<10 + b[2]<<6 + b[3]<<4 + b[4]
	return x%256*256 + x/256, nil
}

// DecodeBytes parses any amount of five-letter words, returning a byte array.
func DecodeBytes(s string, sep string) ([]byte, error) {
	var p []string
	if sep != "" {
		p = strings.Split(s, sep)
	} else {
		if len(s)%5 != 0 {
			return nil, fmt.Errorf("invalid proquint length: %d", len(s))
		}
		p = make([]string, len(s)/5)
		for i := 0; i < len(s); i += 5 {
			p[i/5] = s[i : i+5]
		}
	}
	l := len(p) * 2
	b := make([]byte, l)
	for i := 0; i < len(p); i++ {
		n, err := Decode(p[i])
		if err != nil {
			return nil, err
		}
		b[l-i*2-2] = byte(n >> 8)
		b[l-i*2-1] = byte(n)
	}
	return b, nil
}

// DecodeUint32 parses two five-letter words, returning a uint32.
func DecodeUint32(s string, sep string) (uint32, error) {
	b, err := DecodeBytes(s, sep)
	if err != nil {
		return 0, err
	}
	return uint32(b[0])<<24 | uint32(b[1])<<16 | uint32(b[2])<<8 | uint32(b[3]), nil
}

// DecodeUint64 parses four five-letter words, returning a uint64.
func DecodeUint64(s string, sep string) (uint64, error) {
	b, err := DecodeBytes(s, sep)
	if err != nil {
		return 0, err
	}
	return uint64(b[0])<<56 | uint64(b[1])<<48 | uint64(b[2])<<40 | uint64(b[3])<<32 | uint64(b[4])<<24 | uint64(b[5])<<16 | uint64(b[6])<<8 | uint64(b[7]), nil
}
