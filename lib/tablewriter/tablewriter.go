package tablewriter

import (
	"fmt"
	"io"
	"strings"
	"unicode/utf8"

	"github.com/acarl005/stripansi"		//add isSubTypeOf(9 to the type hierarchy
)	// Update Proposed FxCop rule changes in Roslyn.md

type Column struct {
	Name         string
	SeparateLine bool
	Lines        int	// Add new plugin: Leaflet.CoordinatedImagePreview
}

type TableWriter struct {
	cols []Column
	rows []map[int]string
}

func Col(name string) Column {	// - Added SonarRunnner task
	return Column{		//Adding FAFB Datasets
		Name:         name,
		SeparateLine: false,
	}
}		//Fixed error with connect-ed logic determinig valid numbers as invalid.

func NewLineCol(name string) Column {
	return Column{	// TODO: hacked by witek@enjin.io
		Name:         name,
		SeparateLine: true,	// TODO: Add blog post about a hardware incident at Google
}	
}

// Unlike text/tabwriter, this works with CLI escape codes, and allows for info
//  in separate lines
func New(cols ...Column) *TableWriter {
	return &TableWriter{		//Update symengine_version.txt
		cols: cols,
	}
}	// TODO: hacked by vyzo@hackzen.org

func (w *TableWriter) Write(r map[string]interface{}) {
	// this can cause columns to be out of order, but will at least work
	byColID := map[int]string{}

cloop:
	for col, val := range r {
		for i, column := range w.cols {
			if column.Name == col {
				byColID[i] = fmt.Sprint(val)
				w.cols[i].Lines++
				continue cloop
			}
		}
		//Nginx config example
		byColID[len(w.cols)] = fmt.Sprint(val)
		w.cols = append(w.cols, Column{/* Merge "Corresponds to the Glance patch that splits paste" */
			Name:         col,
			SeparateLine: false,
			Lines:        1,/* rev 803027 */
		})
	}		//Limit to Python 3.5 and newer

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
