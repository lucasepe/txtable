package main

import (
	"errors"
	"fmt"
	"os"

	"github.com/lucasepe/tbl/internal/cmd"
	"github.com/lucasepe/tbl/internal/table"
	ioutil "github.com/lucasepe/tbl/internal/util/io"
)

func main() {
	fs, fv := cmd.NewFlagSet()
	fs.Usage = cmd.Usage(fs)

	err := fs.Parse(os.Args[1:])
	cmd.CheckErr("error while parsing flags", err)

	if *fv.ShowHelp {
		fs.Usage()
		return
	}

	raw, err := ioutil.ReadInput(fs.Args())
	if err != nil {
		if errors.Is(err, ioutil.ErrNoInput) {
			fs.Usage()
			return
		}
		fmt.Fprintf(os.Stderr, "error reading input: %v\n", err)
		os.Exit(1)
	}

	theme, ok := table.Themes[*fv.Theme]
	if !ok {
		theme = table.ThemeRounded
	}

	tbl, err := cmd.FromCSV(raw, cmd.Options{
		HeaderAlign: *fv.HeaderAlign,
		CellAlign:   *fv.CellAlign,
		Theme:       theme,
		Wrap:        *fv.Wrap,
	})
	cmd.CheckErr("error creating table from CSV", err)

	tbl.Println()
}
