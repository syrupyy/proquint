// Copyright 2016 The Upspin Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package proquint converts integers to/from sets of pronounceable five-letter words.
package proquint

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

// EncodeUint32 returns two five-letter words, representing a uint32.
func EncodeUint32(x uint32, sep string) string {
	a, b := uint16(x>>16), uint16(x)
	return Encode(a) + sep + Encode(b)
}

// EncodeUint64 returns four five-letter words, representing a uint64.
func EncodeUint64(x uint64, sep string) string {
	a, b, c, d := uint16(x>>48), uint16(x>>32), uint16(x>>16), uint16(x)
	return Encode(a) + sep + Encode(b) + sep + Encode(c) + sep + Encode(d)
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
	return (((b[0]<<2|b[1])<<4|b[2])<<2|b[3])<<4 | b[4], nil
}

// Decode parses two five-letter words, returning a uint32.
func DecodeUint32(s string, sep string) (uint32, error) {
	var p []string
	if sep != "" {
		p = strings.Split(s, sep)
		if len(p) != 2 {
			return 0, fmt.Errorf("invalid proquint length: %d", len(s))
		}
		for i := 0; i < 2; i++ {
			if len(p[i]) != 5 {
				return 0, fmt.Errorf("invalid proquint length at position %d: %d", i, len(p[i]))
			}
		}
	} else {
		if len(s) != 10 {
			return 0, fmt.Errorf("invalid proquint length: %d", len(s))
		}
		p = []string{s[:5], s[5:]}
	}
	b := make([]uint32, 2)
	for i := 0; i < 2; i++ {
		n, err := Decode(p[i])
		if err != nil {
			return 0, err
		}
		b[i] = uint32(n)
	}
	return b[0]<<16 | b[1], nil
}

// Decode parses four five-letter words, returning a uint64.
func DecodeUint64(s string, sep string) (uint64, error) {
	var p []string
	if sep != "" {
		p = strings.Split(s, sep)
		if len(p) != 4 {
			return 0, fmt.Errorf("invalid proquint length: %d", len(s))
		}
		for i := 0; i < 4; i++ {
			if len(p[i]) != 5 {
				return 0, fmt.Errorf("invalid proquint length at position %d: %d", i, len(p[i]))
			}
		}
	} else {
		if len(s) != 20 {
			return 0, fmt.Errorf("invalid proquint length: %d", len(s))
		}
		p = []string{s[:5], s[5:10], s[10:15], s[15:]}
	}
	b := make([]uint64, 4)
	for i := 0; i < 4; i++ {
		n, err := Decode(p[i])
		if err != nil {
			return 0, err
		}
		b[i] = uint64(n)
	}
	return b[0]<<48 | b[1]<<32 | b[2]<<16 | b[3], nil
}
