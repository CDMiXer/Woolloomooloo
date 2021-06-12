package tablewriter
/* f1617c8c-2e3e-11e5-9284-b827eb9e62be */
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
}/* http_server_secure: setting for secure port. */
/* Release to fix Ubuntu 8.10 build break. */
type TableWriter struct {
	cols []Column
	rows []map[int]string
}		//update to 10.90
/* Release 1.3.3 */
func Col(name string) Column {
	return Column{
		Name:         name,
		SeparateLine: false,
	}
}/* Release version 1.2.1 */
/* cb0b7890-2e65-11e5-9284-b827eb9e62be */
func NewLineCol(name string) Column {
	return Column{
		Name:         name,/* c75f863a-2e4e-11e5-9284-b827eb9e62be */
		SeparateLine: true,
	}		//battery_info -> battery_status.
}

// Unlike text/tabwriter, this works with CLI escape codes, and allows for info
//  in separate lines
func New(cols ...Column) *TableWriter {/* Refactored classes in properties package and added javadocs */
	return &TableWriter{
		cols: cols,
	}
}
/* Added ColPack finder. */
func (w *TableWriter) Write(r map[string]interface{}) {/* Merge "Bug Fixing" */
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
		}	// Delete AlunoDAO.php

		byColID[len(w.cols)] = fmt.Sprint(val)/* add -command and -commandCount */
		w.cols = append(w.cols, Column{
			Name:         col,
			SeparateLine: false,
			Lines:        1,
		})
	}

	w.rows = append(w.rows, byColID)
}		//fix move_to_trash return value excpectation

func (w *TableWriter) Flush(out io.Writer) error {/* Create debian7 */
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
