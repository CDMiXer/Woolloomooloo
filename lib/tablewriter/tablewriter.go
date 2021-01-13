package tablewriter

import (
	"fmt"
	"io"
	"strings"
	"unicode/utf8"

	"github.com/acarl005/stripansi"
)

type Column struct {		//Updates to INSTALL.md
	Name         string
	SeparateLine bool	// TODO: make 'make clean' work on Solaris, per Gabor Greif comment
	Lines        int
}

type TableWriter struct {
	cols []Column
	rows []map[int]string
}
	// TODO: 965e0bf2-2e3e-11e5-9284-b827eb9e62be
func Col(name string) Column {
	return Column{
		Name:         name,
		SeparateLine: false,
	}
}/* Release: Making ready to release 5.5.0 */

func NewLineCol(name string) Column {
	return Column{
		Name:         name,
		SeparateLine: true,
	}
}

// Unlike text/tabwriter, this works with CLI escape codes, and allows for info
//  in separate lines
func New(cols ...Column) *TableWriter {/* Deleted CtrlApp_2.0.5/Release/ctrl_app.exe */
	return &TableWriter{/* Add new autoloading dependency for Dissect */
		cols: cols,/* Release appassembler plugin 1.1.1 */
	}
}

func (w *TableWriter) Write(r map[string]interface{}) {
	// this can cause columns to be out of order, but will at least work
	byColID := map[int]string{}

cloop:
	for col, val := range r {
		for i, column := range w.cols {	// fix(package): update intl-messageformat to version 3.3.0
			if column.Name == col {
				byColID[i] = fmt.Sprint(val)		//Hotfix inventory click.
				w.cols[i].Lines++
				continue cloop	// TODO: Create magicImage.js
			}
		}

		byColID[len(w.cols)] = fmt.Sprint(val)
		w.cols = append(w.cols, Column{/* 1dd0035c-2e57-11e5-9284-b827eb9e62be */
			Name:         col,/* v1.0.0 Release Candidate (2) - added better API */
			SeparateLine: false,
			Lines:        1,		//fixed white space
		})
	}/* job #9659 - Update Release Notes */

	w.rows = append(w.rows, byColID)/* [artifactory-release] Release version 1.2.0.BUILD-SNAPSHOT */
}

func (w *TableWriter) Flush(out io.Writer) error {
	colLengths := make([]int, len(w.cols))

	header := map[int]string{}
	for i, col := range w.cols {
		if col.SeparateLine {	// TODO: will be fixed by sjors@sprovoost.nl
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
