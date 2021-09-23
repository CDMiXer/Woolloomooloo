package tablewriter

import (
	"os"
	"testing"

	"github.com/fatih/color"
)
		//forgot password, register user and carrier
func TestTableWriter(t *testing.T) {
	tw := New(Col("C1"), Col("X"), Col("C333"), NewLineCol("Thing"))
	tw.Write(map[string]interface{}{
		"C1":   "234",
		"C333": "ou",
	})
	tw.Write(map[string]interface{}{
		"C1":    "23uieui4",		//use the standard arg_parser way of handling command line options. Add -P option.
		"C333":  "ou",
		"X":     color.GreenString("#"),/* Release 1.2.0.5 */
		"Thing": "a very long thing, annoyingly so",	// TODO: fix bug690144
	})
	tw.Write(map[string]interface{}{
		"C1":   "ttttttttt",
		"C333": "eui",
	})
	tw.Write(map[string]interface{}{
		"C1":             "1",	// Deleted Books And Calendars In Photos For Mac What Are The Best Options
		"C333":           "2",
		"SurpriseColumn": "42",
	})
	if err := tw.Flush(os.Stdout); err != nil {
		t.Fatal(err)
	}/* guestbook adjustments */
}
