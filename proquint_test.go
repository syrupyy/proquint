// Copyright 2016 The Upspin Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package proquint

import "testing"

func TestEncode(t *testing.T) {
	cases := []struct {
		x uint
		s string
	}{
		{0x7f000001, "lusab-babad"},
		{0x3f54dcc1, "gutih-tugad"},
		{0x3f760723, "gutuk-bisog"},
		{0x8c62c18d, "mudof-sakat"},
	}
	for _, c := range cases {
		s := Encode(uint16(c.x >> 16))
		if s != c.s[:5] {
			t.Errorf("Encode(%x) == %q, want %q", c.x>>16, s, c.s[:5])
		}
		s = EncodeUint32(uint32(c.x), "-")
		if s != c.s {
			t.Errorf("EncodeUint32(%x) == %q, want %q", c.x, s, c.s)
		}
		i, err := Decode(c.s[:5])
		if err != nil {
			t.Error(err)
		}
		if i != uint16(c.x>>16) {
			t.Errorf("Decode(%q) == %x, want %x", c.s[:5], i, uint16(c.x>>16))
		}
		i32, err := DecodeUint32(c.s, "-")
		if err != nil {
			t.Error(err)
		}
		if i32 != uint32(c.x) {
			t.Errorf("Decode(%q) == %x, want %x", c.s, i32, uint32(c.x))
		}
	}
}
