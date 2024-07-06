package csvutil

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParseCSV(t *testing.T) {
	testcases := []struct {
		input    []byte
		expected struct {
			CSVRec
			Err error
		}
	}{
		{
			input: []byte(`index,value,date
1,"abc",1-Jan-2022`),
			expected: struct {
				CSVRec
				Err error
			}{
				CSVRec: CSVRec{
					Record: []string{
						"1",
						"abc",
						"1-Jan-2022",
					},
				},
				Err: nil,
			},
		},
		{
			input: []byte{},
			expected: struct {
				CSVRec
				Err error
			}{
				CSVRec: CSVRec{
					Record: nil,
				},
				Err: ErrCSV,
			},
		},
		{
			input: []byte(`index,value,date
1,1-Jan-2022`),
			expected: struct {
				CSVRec
				Err error
			}{
				CSVRec: CSVRec{
					Record: []string{
						"1",
						"1-Jan-2022",
					},
				},
				Err: ErrCSVLine,
			},
		},
	}

	for i, tc := range testcases {
		reader := bytes.NewReader(tc.input)
		records := ParseCSV(context.TODO(), reader)
		for rec := range records {
			if assert.True(t, errors.Is(rec.Err, tc.expected.Err), fmt.Sprintf("Case: %d Error: %v", i, rec.Err)) {
				assert.Equal(t, tc.expected.CSVRec.Record, rec.Record, fmt.Sprintf("Case: %d Value", i))
			}
		}
	}
}
