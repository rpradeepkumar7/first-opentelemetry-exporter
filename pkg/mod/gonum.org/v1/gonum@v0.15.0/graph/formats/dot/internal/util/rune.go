// Code generated by gocc; DO NOT EDIT.

// This file is dual licensed under CC0 and The Gonum License.
//
// Copyright ©2017 The Gonum Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.
//
// Copyright ©2017 Robin Eklind.
// This file is made available under a Creative Commons CC0 1.0
// Universal Public Domain Dedication.

package util

import (
	"fmt"
)

func RuneToString(r rune) string {
	if r >= 0x20 && r < 0x7f {
		return fmt.Sprintf("'%c'", r)
	}
	switch r {
	case 0x07:
		return "'\\a'"
	case 0x08:
		return "'\\b'"
	case 0x0C:
		return "'\\f'"
	case 0x0A:
		return "'\\n'"
	case 0x0D:
		return "'\\r'"
	case 0x09:
		return "'\\t'"
	case 0x0b:
		return "'\\v'"
	case 0x5c:
		return "'\\\\\\'"
	case 0x27:
		return "'\\''"
	case 0x22:
		return "'\\\"'"
	}
	if r < 0x10000 {
		return fmt.Sprintf("\\u%04x", r)
	}
	return fmt.Sprintf("\\U%08x", r)
}
