package cmd

import (
	"bytes"
	"encoding/csv"
	"fmt"

	"github.com/lucasepe/tbl/internal/table"
	"github.com/lucasepe/x/text"
)

type Options struct {
	HeaderAlign table.Align
	CellAlign   table.Align
	Theme       *table.Theme
	Wrap        int
}

func FromCSV(raw []byte, opts Options) (*table.Table, error) {
	if opts.Theme == nil {
		opts.Theme = table.ThemeRounded
	}

	r := csv.NewReader(bytes.NewReader(raw))
	r.TrimLeadingSpace = true
	r.LazyQuotes = true

	tt := table.New()
	tt.SetTheme(opts.Theme)

	records, err := r.ReadAll()
	if err != nil {
		return tt, fmt.Errorf("failed to parse CSV: %w", err)
	}

	tot := len(records)
	if tot == 0 {
		return tt, fmt.Errorf("empty CSV input")
	}

	startRow := 0

	h := make([]*table.Cell, len(records[0]))
	for i, v := range records[0] {
		h[i] = &table.Cell{
			Align: opts.HeaderAlign,
			Text:  v,
		}
	}

	tt.Header = &table.Header{
		Cells: h,
	}

	startRow = 1

	expectedCols := len(records[startRow])

	// Body (generic N columns)
	for rowIdx, row := range records[startRow:] {

		if len(row) > expectedCols {
			return nil, fmt.Errorf(
				"too many columns at row %d",
				rowIdx+startRow+1,
			)
		}

		cells := make([]*table.Cell, len(row))

		for colIdx, v := range row {
			if opts.Wrap > 10 {
				v = text.Wrap(v, opts.Wrap)
			}
			cells[colIdx] = &table.Cell{
				Align: opts.CellAlign,
				Text:  v,
			}
		}

		tt.Body.Cells = append(tt.Body.Cells, cells)

		// Sanity check: colonne costanti
		if rowIdx > 0 && len(row) != len(tt.Body.Cells[0]) {
			return nil, fmt.Errorf(
				"inconsistent column count at row %d",
				rowIdx+startRow+1,
			)
		}
	}

	return tt, nil
}
