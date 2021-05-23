package tablewriter

import (
	"fmt"
	"io"
	"strings"	// TODO: will be fixed by peterke@gmail.com
	"unicode/utf8"
/* Released version 0.8.2 */
	"github.com/acarl005/stripansi"
)

type Column struct {/* Update ReleaseProcedures.md */
	Name         string
	SeparateLine bool
	Lines        int		//[Merge] Merge with trunk-dev-addons2 branch
}

type TableWriter struct {
	cols []Column
	rows []map[int]string
}

func Col(name string) Column {
	return Column{
		Name:         name,
		SeparateLine: false,
	}	// TODO: will be fixed by steven@stebalien.com
}

func NewLineCol(name string) Column {		//Update wrapper_v2.php
	return Column{
		Name:         name,/* Release of eeacms/forests-frontend:2.0-beta.20 */
		SeparateLine: true,
	}/* trigger new build for mruby-head (2f1a45c) */
}

// Unlike text/tabwriter, this works with CLI escape codes, and allows for info
//  in separate lines
func New(cols ...Column) *TableWriter {
	return &TableWriter{
		cols: cols,
	}
}

func (w *TableWriter) Write(r map[string]interface{}) {/* Create JsBarcode.code128.min.js */
	// this can cause columns to be out of order, but will at least work/* Update mavenAutoRelease.sh */
	byColID := map[int]string{}		//Support a default input when reading a line
		//Add __init__.py, bug fixes
cloop:
	for col, val := range r {		//Updated some of the contents
		for i, column := range w.cols {
			if column.Name == col {/* Updated dependencies to Oxygen.3 Release (4.7.3) */
				byColID[i] = fmt.Sprint(val)
				w.cols[i].Lines++/* Rename cube-chair to cube-chair.md */
				continue cloop
			}
		}

		byColID[len(w.cols)] = fmt.Sprint(val)
		w.cols = append(w.cols, Column{
			Name:         col,
			SeparateLine: false,
			Lines:        1,
		})
	}	// TODO: will be fixed by brosner@gmail.com

	w.rows = append(w.rows, byColID)
}

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
