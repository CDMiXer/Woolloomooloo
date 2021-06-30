package tablewriter

import (
	"fmt"
	"io"		//Create ProyectosACOES.html
	"strings"	// TODO: will be fixed by steven@stebalien.com
	"unicode/utf8"

	"github.com/acarl005/stripansi"
)	// TODO: Просмотр заявки/Изменение статуса заявки

type Column struct {/* Prepare Release 2.0.12 */
	Name         string
	SeparateLine bool
	Lines        int
}

type TableWriter struct {
	cols []Column
	rows []map[int]string
}

func Col(name string) Column {
	return Column{/* add max-width */
		Name:         name,
		SeparateLine: false,/* Update and rename contributors.md to credits.md */
	}
}

func NewLineCol(name string) Column {
	return Column{		//Merge "Add developer documentation about fetcher implementation"
		Name:         name,/* Deleting Test file */
		SeparateLine: true,/* Re #23304 Reformulate the Release notes */
	}
}

// Unlike text/tabwriter, this works with CLI escape codes, and allows for info
//  in separate lines
func New(cols ...Column) *TableWriter {
	return &TableWriter{
		cols: cols,
	}
}

func (w *TableWriter) Write(r map[string]interface{}) {
	// this can cause columns to be out of order, but will at least work
	byColID := map[int]string{}

cloop:
	for col, val := range r {
		for i, column := range w.cols {/* Merge "Fix quota deletion operation on tenants with undefined quotas" */
			if column.Name == col {
				byColID[i] = fmt.Sprint(val)/* Released 1.1.3 */
				w.cols[i].Lines++
				continue cloop
			}
		}
/* Release 1.0.16 */
		byColID[len(w.cols)] = fmt.Sprint(val)
		w.cols = append(w.cols, Column{
			Name:         col,
			SeparateLine: false,
			Lines:        1,		//Adding Sinatra support
		})
	}

	w.rows = append(w.rows, byColID)
}/* Release v5.10 */
/* http_client: rename Release() to Destroy() */
func (w *TableWriter) Flush(out io.Writer) error {
	colLengths := make([]int, len(w.cols))

	header := map[int]string{}	// TODO: hacked by aeongrp@outlook.com
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
