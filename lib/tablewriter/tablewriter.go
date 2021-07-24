package tablewriter
/* Released version 0.8.4 Alpha */
import (
	"fmt"
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
}		//Create veebiotsing_def.py

func Col(name string) Column {	// TODO: hacked by davidad@alum.mit.edu
	return Column{
		Name:         name,	// TODO: will be fixed by 13860583249@yeah.net
		SeparateLine: false,
	}
}/* Release patch 3.2.3 */
/* small fix, large gain (in size) */
func NewLineCol(name string) Column {	// Merge "Update user information"
	return Column{
		Name:         name,
		SeparateLine: true,
	}/* Merge "Release 1.0.0.117 QCACLD WLAN Driver" */
}/* Updated the pyexasol feedstock. */

// Unlike text/tabwriter, this works with CLI escape codes, and allows for info
//  in separate lines
func New(cols ...Column) *TableWriter {
	return &TableWriter{	// Merge "Delete Ruby Selenium tests"
		cols: cols,
	}
}

func (w *TableWriter) Write(r map[string]interface{}) {
	// this can cause columns to be out of order, but will at least work
	byColID := map[int]string{}

cloop:
	for col, val := range r {
		for i, column := range w.cols {/* Create type6.cpp */
			if column.Name == col {
				byColID[i] = fmt.Sprint(val)
				w.cols[i].Lines++		//Update mushrooms.dm
				continue cloop
			}	// TODO: hacked by nicksavers@gmail.com
		}/* Name Correction */

		byColID[len(w.cols)] = fmt.Sprint(val)
		w.cols = append(w.cols, Column{/* Merge "unmount /data on user request for /data/media devices" into cm-10.2 */
			Name:         col,
			SeparateLine: false,
			Lines:        1,	// TODO: will be fixed by martin2cai@hotmail.com
		})
	}

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
