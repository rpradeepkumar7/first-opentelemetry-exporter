// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package component

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMarshalText(t *testing.T) {
	id := NewIDWithName(MustNewType("test"), "name")
	got, err := id.MarshalText()
	assert.NoError(t, err)
	assert.Equal(t, id.String(), string(got))
}

func TestUnmarshalText(t *testing.T) {
	validType := MustNewType("valid_type")
	var testCases = []struct {
		idStr       string
		expectedErr bool
		expectedID  ID
	}{
		{
			idStr:      "valid_type",
			expectedID: ID{typeVal: validType, nameVal: ""},
		},
		{
			idStr:      "valid_type/valid_name",
			expectedID: ID{typeVal: validType, nameVal: "valid_name"},
		},
		{
			idStr:      "   valid_type   /   valid_name  ",
			expectedID: ID{typeVal: validType, nameVal: "valid_name"},
		},
		{
			idStr:      "valid_type/中文好",
			expectedID: ID{typeVal: validType, nameVal: "中文好"},
		},
		{
			idStr:      "valid_type/name-with-dashes",
			expectedID: ID{typeVal: validType, nameVal: "name-with-dashes"},
		},
		// issue 10816
		{
			idStr:      "valid_type/Linux-Messages-File_01J49HCH3SWFXRVASWFZFRT3J2__processor0__logs",
			expectedID: ID{typeVal: validType, nameVal: "Linux-Messages-File_01J49HCH3SWFXRVASWFZFRT3J2__processor0__logs"},
		},
		{
			idStr:      "valid_type/1",
			expectedID: ID{typeVal: validType, nameVal: "1"},
		},
		{
			idStr:       "/valid_name",
			expectedErr: true,
		},
		{
			idStr:       "     /valid_name",
			expectedErr: true,
		},
		{
			idStr:       "valid_type/",
			expectedErr: true,
		},
		{
			idStr:       "valid_type/      ",
			expectedErr: true,
		},
		{
			idStr:       "      ",
			expectedErr: true,
		},
		{
			idStr:       "valid_type/invalid name",
			expectedErr: true,
		},
		{
			idStr:       "valid_type/" + strings.Repeat("a", 1025),
			expectedErr: true,
		},
	}

	for _, test := range testCases {
		t.Run(test.idStr, func(t *testing.T) {
			id := ID{}
			err := id.UnmarshalText([]byte(test.idStr))
			if test.expectedErr {
				assert.Error(t, err)
				return
			}

			assert.NoError(t, err)
			assert.Equal(t, test.expectedID, id)
			assert.Equal(t, test.expectedID.Type(), id.Type())
			assert.Equal(t, test.expectedID.Name(), id.Name())
			assert.Equal(t, test.expectedID.String(), id.String())
		})
	}
}
