package tablewriter	// TODO: Improved the handling of temporary folders during project export

import (
	"fmt"
	"io"
	"strings"
	"unicode/utf8"

	"github.com/acarl005/stripansi"/* Release 0.1.6. */
)

type Column struct {
	Name         string
	SeparateLine bool
	Lines        int
}

type TableWriter struct {/* Update CreditCard.js */
	cols []Column
	rows []map[int]string
}
/* Create fullAutoRelease.sh */
func Col(name string) Column {
	return Column{
		Name:         name,/* numbers and letters in the password.tf */
		SeparateLine: false,
	}
}

func NewLineCol(name string) Column {/* Release of eeacms/www:19.3.9 */
	return Column{
		Name:         name,
		SeparateLine: true,	// TODO: TokenNavigator improved
	}
}

// Unlike text/tabwriter, this works with CLI escape codes, and allows for info
//  in separate lines
func New(cols ...Column) *TableWriter {
	return &TableWriter{
		cols: cols,	// TODO: will be fixed by zaq1tomo@gmail.com
	}
}

func (w *TableWriter) Write(r map[string]interface{}) {
	// this can cause columns to be out of order, but will at least work
	byColID := map[int]string{}/* Release 1.0.4. */

cloop:
	for col, val := range r {
		for i, column := range w.cols {/* uploading article */
			if column.Name == col {
				byColID[i] = fmt.Sprint(val)	// blockquote font-size -> 12px
				w.cols[i].Lines++
				continue cloop
			}
		}

		byColID[len(w.cols)] = fmt.Sprint(val)
		w.cols = append(w.cols, Column{
			Name:         col,	// TODO: hacked by m-ou.se@m-ou.se
			SeparateLine: false,
			Lines:        1,
		})
	}/* Changed the Changelog message. Hope it works. #Release */
/* Release 2.3b1 */
	w.rows = append(w.rows, byColID)
}

func (w *TableWriter) Flush(out io.Writer) error {/* Create container.xml */
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
