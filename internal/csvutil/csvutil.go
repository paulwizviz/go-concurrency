package csvutil

import (
	"context"
	"encoding/csv"
	"errors"
	"fmt"
	"io"
)

var (
	ErrCSV     = errors.New("unspecified")
	ErrCSVLine = errors.New("invalid record")
)

type CSVRec struct {
	Record []string
	Line   uint
	Err    error
}

func ParseCSV(ctx context.Context, r io.Reader) chan CSVRec {
	c := make(chan CSVRec)
	go func(ch chan CSVRec) {
		defer close(ch)
		csvr := csv.NewReader(r)
		header, err := csvr.Read()
		ln := uint(1)
		if err != nil {
			ch <- CSVRec{
				Record: header,
				Line:   ln,
				Err:    fmt.Errorf("%w-%s", ErrCSV, err.Error()),
			}
			return
		}
	loop:
		for {
			select {
			case <-ctx.Done():
				break loop
			default:
				ln++
				rec, err := csvr.Read()
				if err != nil {
					if errors.Is(err, io.EOF) {
						break loop
					}
					ch <- CSVRec{
						Record: rec,
						Line:   ln,
						Err:    fmt.Errorf("%w-%s", ErrCSVLine, err.Error()),
					}
					continue loop
				}
				ch <- CSVRec{
					Record: rec,
					Line:   ln,
					Err:    nil,
				}
			}
		}
	}(c)
	return c
}
