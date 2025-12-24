package cmd

import (
	"flag"
	"fmt"
	"os"
	"sort"
	"strings"

	"github.com/lucasepe/tbl/internal/table"
)

type FlagValues struct {
	HeaderAlign *table.Align
	CellAlign   *table.Align
	Theme       *string
	Wrap        *int
	ShowHelp    *bool
}

// NewFlagSet creates a FlagSet with all supported CLI options.
func NewFlagSet() (*flag.FlagSet, *FlagValues) {
	fs := flag.NewFlagSet("txtable", flag.ContinueOnError)
	fs.SetOutput(os.Stdout)

	headerAlign := table.AlignCenter
	textAlign := table.AlignLeft
	theme := "rounded"
	wrap := 0

	vals := &FlagValues{
		HeaderAlign: &headerAlign,
		CellAlign:   &textAlign,
		Theme:       &theme,
		Wrap:        &wrap,
		ShowHelp:    fs.Bool("h", false, "Show help"),
	}

	fs.Var(vals.HeaderAlign, "x", "Header text alignment (left|center|right)")
	fs.Var(vals.CellAlign, "c", "Cell text alignment (left|center|right)")
	fs.StringVar(vals.Theme, "t", "rounded",
		fmt.Sprintf("Table theme: %s", strings.Join(availableThemes(), "|")))
	fs.IntVar(vals.Wrap, "w", 0, "Wrap cell content at specified width (0 = no wrap)")

	return fs, vals
}

func availableThemes() []string {
	keys := make([]string, 0, len(table.Themes))
	for k := range table.Themes {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	return keys
}
