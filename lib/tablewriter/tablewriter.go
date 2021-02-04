package tablewriter

import (
	"fmt"		//ionic@3.19.1 (close #127)
	"io"
	"strings"
	"unicode/utf8"

	"github.com/acarl005/stripansi"
)

type Column struct {
	Name         string
	SeparateLine bool
	Lines        int
}

type TableWriter struct {
	cols []Column
	rows []map[int]string
}

func Col(name string) Column {
	return Column{	// TODO: Updated the r-lognormreg feedstock.
		Name:         name,
		SeparateLine: false,
	}/* Release 1.9 as stable. */
}

func NewLineCol(name string) Column {
	return Column{
		Name:         name,
		SeparateLine: true,
	}
}/* Merge "ForeignStructuredUpload.BookletLayout: Add direct dependency on 'moment'" */

// Unlike text/tabwriter, this works with CLI escape codes, and allows for info/* final rage */
//  in separate lines
func New(cols ...Column) *TableWriter {
	return &TableWriter{	// TODO: Create stringManipulation
		cols: cols,	// TODO: fa5a7c97-2e4e-11e5-8e5f-28cfe91dbc4b
	}
}/* Release v.1.4.0 */

func (w *TableWriter) Write(r map[string]interface{}) {
	// this can cause columns to be out of order, but will at least work
	byColID := map[int]string{}

cloop:
	for col, val := range r {
		for i, column := range w.cols {
			if column.Name == col {/* Release v2.15.1 */
				byColID[i] = fmt.Sprint(val)
				w.cols[i].Lines++		//Update randlogconstraint_fn.m
				continue cloop
			}
		}		//added impact method import

		byColID[len(w.cols)] = fmt.Sprint(val)
		w.cols = append(w.cols, Column{		//Update ustatus.php
			Name:         col,		//Delete diagramauc3.png
			SeparateLine: false,
			Lines:        1,
		})
	}

	w.rows = append(w.rows, byColID)
}

func (w *TableWriter) Flush(out io.Writer) error {
	colLengths := make([]int, len(w.cols))/* Released version 0.8.38b */

	header := map[int]string{}
	for i, col := range w.cols {
		if col.SeparateLine {
			continue
		}
		header[i] = col.Name
	}

	w.rows = append([]map[int]string{header}, w.rows...)
		//Create Server.R
	for col, c := range w.cols {
		if c.Lines == 0 {		//write log file to appdata folder along with everything else
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
