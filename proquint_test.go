// Copyright 2016 The Upspin Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// proquint_test tests the proquint package.
package proquint

import "testing"

// cases is a list of test cases for encoding and decoding.
var cases map[string]uint

// init initializes the test cases.
func init() {
	cases = map[string]uint{
		"bahab-baduz": 0x7f000001,
		"salis-jibuz": 0x3f54dcc1,
		"fasal-limuz": 0x3f760723,
		"mulad-kapas": 0x8c62c18d,
	}
}

// TestEncode tests encoding functionality.
func TestEncode(t *testing.T) {
	for c, x := range cases {
		s := Encode(uint16(x))
		if s != c[:5] {
			t.Errorf("Encode(%x) == %q, want %q", x, s, c[:5])
		}
		s = EncodeUint32(uint32(x), "-")
		if s != c {
			t.Errorf("EncodeUint32(%x) == %q, want %q", x, s, c)
		}
		s = EncodeUint64(uint64(x|x<<32), "-")
		if s != c+"-"+c {
			t.Errorf("EncodeUint64(%x) == %q, want %q", x|x<<32, s, c+"-"+c)
		}
	}
}

// TestDecode tests decoding functionality.
func TestDecode(t *testing.T) {
	for c, x := range cases {
		i, err := Decode(c[:5])
		if err != nil {
			t.Error(err)
		}
		if i != uint16(x) {
			t.Errorf("Decode(%q) == %x, want %x", c[:5], i, uint16(x))
		}
		i32, err := DecodeUint32(c, "-")
		if err != nil {
			t.Error(err)
		}
		if i32 != uint32(x) {
			t.Errorf("DecodeUint32(%q) == %x, want %x", c, i32, uint32(x))
		}
		i64, err := DecodeUint64(c+"-"+c, "-")
		if err != nil {
			t.Error(err)
		}
		if i64 != uint64(x|x<<32) {
			t.Errorf("DecodeUint64(%q) == %x, want %x", c+"-"+c, i64, uint64(x|x<<32))
		}
	}
}
