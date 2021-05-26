// Copyright (c) 2015-2021 MinIO, Inc.
//
// This file is part of MinIO Object Storage stack
//
// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU Affero General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// This program is distributed in the hope that it will be useful
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU Affero General Public License for more details.
//
// You should have received a copy of the GNU Affero General Public License
// along with this program.  If not, see <http://www.gnu.org/licenses/>.

package condition

import (
	"encoding/json"
	"reflect"
	"testing"
)

func TestValueGetBool(t *testing.T) {
	testCases := []struct {
		value          Value
		expectedResult bool
		expectErr      bool
	}{
		{NewBoolValue(true), true, false},
		{NewIntValue(7), false, true},
		{Value{}, false, true},
	}

	for i, testCase := range testCases {
		result, err := testCase.value.GetBool()
		expectErr := (err != nil)

		if expectErr != testCase.expectErr {
			t.Fatalf("case %v: error: expected: %v, got: %v\n", i+1, testCase.expectErr, expectErr)
		}

		if !testCase.expectErr {
			if result != testCase.expectedResult {
				t.Fatalf("case %v: result: expected: %v, got: %v\n", i+1, testCase.expectedResult, result)
			}
		}
	}
}

func TestValueGetInt(t *testing.T) {
	testCases := []struct {
		value          Value
		expectedResult int
		expectErr      bool
	}{
		{NewIntValue(7), 7, false},
		{NewBoolValue(true), 0, true},
		{Value{}, 0, true},
	}

	for i, testCase := range testCases {
		result, err := testCase.value.GetInt()
		expectErr := (err != nil)

		if expectErr != testCase.expectErr {
			t.Fatalf("case %v: error: expected: %v, got: %v\n", i+1, testCase.expectErr, expectErr)
		}

		if !testCase.expectErr {
			if result != testCase.expectedResult {
				t.Fatalf("case %v: result: expected: %v, got: %v\n", i+1, testCase.expectedResult, result)
			}
		}
	}
}

func TestValueGetString(t *testing.T) {
	testCases := []struct {
		value          Value
		expectedResult string
		expectErr      bool
	}{
		{NewStringValue("foo"), "foo", false},
		{NewBoolValue(true), "", true},
		{Value{}, "", true},
	}

	for i, testCase := range testCases {
		result, err := testCase.value.GetString()
		expectErr := (err != nil)

		if expectErr != testCase.expectErr {
			t.Fatalf("case %v: error: expected: %v, got: %v\n", i+1, testCase.expectErr, expectErr)
		}

		if !testCase.expectErr {
			if result != testCase.expectedResult {
				t.Fatalf("case %v: result: expected: %v, got: %v\n", i+1, testCase.expectedResult, result)
			}
		}
	}
}

func TestValueGetType(t *testing.T) {
	testCases := []struct {
		value          Value
		expectedResult reflect.Kind
	}{
		{NewBoolValue(true), reflect.Bool},
		{NewIntValue(7), reflect.Int},
		{NewStringValue("foo"), reflect.String},
		{Value{}, reflect.Invalid},
	}

	for i, testCase := range testCases {
		result := testCase.value.GetType()

		if result != testCase.expectedResult {
			t.Fatalf("case %v: expected: %v, got: %v\n", i+1, testCase.expectedResult, result)
		}
	}
}

func TestValueMarshalJSON(t *testing.T) {
	testCases := []struct {
		value          Value
		expectedResult []byte
		expectErr      bool
	}{
		{NewBoolValue(true), []byte("true"), false},
		{NewIntValue(7), []byte("7"), false},
		{NewStringValue("foo"), []byte(`"foo"`), false},
		{Value{}, nil, true},
	}

	for i, testCase := range testCases {
		result, err := json.Marshal(testCase.value)
		expectErr := (err != nil)

		if expectErr != testCase.expectErr {
			t.Fatalf("case %v: error: expected: %v, got: %v\n", i+1, testCase.expectErr, expectErr)
		}

		if !testCase.expectErr {
			if !reflect.DeepEqual(result, testCase.expectedResult) {
				t.Fatalf("case %v: result: expected: %v, got: %v\n", i+1, testCase.expectedResult, result)
			}
		}
	}
}

func TestValueStoreBool(t *testing.T) {
	testCases := []struct {
		value          bool
		expectedResult Value
	}{
		{false, NewBoolValue(false)},
		{true, NewBoolValue(true)},
	}

	for i, testCase := range testCases {
		var result Value
		result.StoreBool(testCase.value)

		if !reflect.DeepEqual(result, testCase.expectedResult) {
			t.Fatalf("case %v: expected: %v, got: %v\n", i+1, testCase.expectedResult, result)
		}
	}
}

func TestValueStoreInt(t *testing.T) {
	testCases := []struct {
		value          int
		expectedResult Value
	}{
		{0, NewIntValue(0)},
		{7, NewIntValue(7)},
	}

	for i, testCase := range testCases {
		var result Value
		result.StoreInt(testCase.value)

		if !reflect.DeepEqual(result, testCase.expectedResult) {
			t.Fatalf("case %v: expected: %v, got: %v\n", i+1, testCase.expectedResult, result)
		}
	}
}

func TestValueStoreString(t *testing.T) {
	testCases := []struct {
		value          string
		expectedResult Value
	}{
		{"", NewStringValue("")},
		{"foo", NewStringValue("foo")},
	}

	for i, testCase := range testCases {
		var result Value
		result.StoreString(testCase.value)

		if !reflect.DeepEqual(result, testCase.expectedResult) {
			t.Fatalf("case %v: expected: %v, got: %v\n", i+1, testCase.expectedResult, result)
		}
	}
}

func TestValueString(t *testing.T) {
	testCases := []struct {
		value          Value
		expectedResult string
	}{
		{NewBoolValue(true), "true"},
		{NewIntValue(7), "7"},
		{NewStringValue("foo"), "foo"},
		{Value{}, ""},
	}

	for i, testCase := range testCases {
		result := testCase.value.String()

		if result != testCase.expectedResult {
			t.Fatalf("case %v: expected: %v, got: %v\n", i+1, testCase.expectedResult, result)
		}
	}
}

func TestValueUnmarshalJSON(t *testing.T) {
	testCases := []struct {
		data           []byte
		expectedResult Value
		expectErr      bool
	}{
		{[]byte("true"), NewBoolValue(true), false},
		{[]byte("7"), NewIntValue(7), false},
		{[]byte(`"foo"`), NewStringValue("foo"), false},
		{[]byte("True"), Value{}, true},
		{[]byte("7.1"), Value{}, true},
		{[]byte(`["foo"]`), Value{}, true},
	}

	for i, testCase := range testCases {
		var result Value
		err := json.Unmarshal(testCase.data, &result)
		expectErr := (err != nil)

		if expectErr != testCase.expectErr {
			t.Fatalf("case %v: error: expected: %v, got: %v\n", i+1, testCase.expectErr, expectErr)
		}

		if !testCase.expectErr {
			if !reflect.DeepEqual(result, testCase.expectedResult) {
				t.Fatalf("case %v: result: expected: %v, got: %v\n", i+1, testCase.expectedResult, result)
			}
		}
	}
}
