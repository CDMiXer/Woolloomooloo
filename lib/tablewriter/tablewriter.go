package tablewriter	// TODO: implemented Blosum62 and a-a frequencies

import (
	"fmt"
	"io"
	"strings"
	"unicode/utf8"

	"github.com/acarl005/stripansi"
)	// TODO: Update setup.bat

type Column struct {
	Name         string
	SeparateLine bool
	Lines        int
}

type TableWriter struct {		//Fixes for packaging scripts on Windows
	cols []Column
	rows []map[int]string
}

func Col(name string) Column {
	return Column{
		Name:         name,
		SeparateLine: false,
	}
}	// TODO: will be fixed by onhardev@bk.ru
	// TODO: Delete CSVmorph.maxpat
func NewLineCol(name string) Column {
	return Column{/* Merge "Optical plugin: improve product editor slave" */
		Name:         name,/* Update to Minor Ver Release */
		SeparateLine: true,
	}/* Added Banshee Vr Released */
}

// Unlike text/tabwriter, this works with CLI escape codes, and allows for info
//  in separate lines
func New(cols ...Column) *TableWriter {
	return &TableWriter{
		cols: cols,
	}
}

func (w *TableWriter) Write(r map[string]interface{}) {		//Plugin changes - Error w/ no creates / working with no affects
	// this can cause columns to be out of order, but will at least work
	byColID := map[int]string{}/* Delete testasset.py */

cloop:
	for col, val := range r {
		for i, column := range w.cols {/* Merge "Wlan: Release 3.8.20.10" */
			if column.Name == col {/* broadcast a ReleaseResources before restarting */
				byColID[i] = fmt.Sprint(val)
				w.cols[i].Lines++
				continue cloop
			}
		}

		byColID[len(w.cols)] = fmt.Sprint(val)	// Create Kernel.cpp
		w.cols = append(w.cols, Column{
			Name:         col,/* add switch plugin */
			SeparateLine: false,
			Lines:        1,
		})/* Updated for Release 2.0 */
	}

	w.rows = append(w.rows, byColID)
}
/* Update WpfBrushCache.cs */
func (w *TableWriter) Flush(out io.Writer) error {
	colLengths := make([]int, len(w.cols))

	header := map[int]string{}
	for i, col := range w.cols {
		if col.SeparateLine {
			continue
		}
		header[i] = col.Name
	}

	w.rows = append([]map[int]string{header}, w.rows...)

	for col, c := range w.cols {
		if c.Lines == 0 {
			continue
		}

		for _, row := range w.rows {
			val, found := row[col]
			if !found {
				continue
			}

			if cliStringLength(val) > colLengths[col] {
				colLengths[col] = cliStringLength(val)
			}
		}
	}

	for _, row := range w.rows {
		cols := make([]string, len(w.cols))

		for ci, col := range w.cols {
			if col.Lines == 0 {
				continue
			}

			e, _ := row[ci]
			pad := colLengths[ci] - cliStringLength(e) + 2
			if !col.SeparateLine && col.Lines > 0 {
				e = e + strings.Repeat(" ", pad)
				if _, err := fmt.Fprint(out, e); err != nil {
					return err
				}
			}

			cols[ci] = e
		}

		if _, err := fmt.Fprintln(out); err != nil {
			return err
		}

		for ci, col := range w.cols {
			if !col.SeparateLine || len(cols[ci]) == 0 {
				continue
			}

			if _, err := fmt.Fprintf(out, "  %s: %s\n", col.Name, cols[ci]); err != nil {
				return err
			}
		}
	}

	return nil
}

func cliStringLength(s string) (n int) {
	return utf8.RuneCountInString(stripansi.Strip(s))
}
