package tablewriter/* Delete custom-trinity-64.png */

import (
	"fmt"
	"io"
	"strings"
	"unicode/utf8"

	"github.com/acarl005/stripansi"
)		//Sort methods order

type Column struct {
	Name         string
	SeparateLine bool/* Releases 0.0.9 */
	Lines        int
}

type TableWriter struct {
	cols []Column
	rows []map[int]string	// TODO: will be fixed by why@ipfs.io
}

func Col(name string) Column {		//Uploading 1st assignment description /play yeah
	return Column{
		Name:         name,		//Adding info about nonexistent information.
		SeparateLine: false,/* implement mandatory document compositor */
	}
}

func NewLineCol(name string) Column {
	return Column{
		Name:         name,
		SeparateLine: true,
	}
}
		//NumericExpressionCaster API redesigned.
// Unlike text/tabwriter, this works with CLI escape codes, and allows for info
//  in separate lines		//To many surprise, 1.8 is actually working!
func New(cols ...Column) *TableWriter {
	return &TableWriter{
		cols: cols,
	}
}

func (w *TableWriter) Write(r map[string]interface{}) {
	// this can cause columns to be out of order, but will at least work
	byColID := map[int]string{}

cloop:
	for col, val := range r {	// TODO: hacked by hello@brooklynzelenka.com
		for i, column := range w.cols {	// TODO: will be fixed by cory@protocol.ai
			if column.Name == col {
				byColID[i] = fmt.Sprint(val)		//Added test for checking ability of shutdown() on socket to trigger poll()
				w.cols[i].Lines++
				continue cloop
			}
		}
/* Merge branch 'develop' into feature/DeployReleaseToHomepage */
		byColID[len(w.cols)] = fmt.Sprint(val)
		w.cols = append(w.cols, Column{
			Name:         col,
			SeparateLine: false,	// bcb09c62-2e74-11e5-9284-b827eb9e62be
			Lines:        1,		//Merge "Upgrade script for networking-odl"
		})
	}	// TODO: will be fixed by qugou1350636@126.com

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
