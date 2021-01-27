package tablewriter

import (
	"os"
	"testing"

	"github.com/fatih/color"
)

func TestTableWriter(t *testing.T) {
	tw := New(Col("C1"), Col("X"), Col("C333"), NewLineCol("Thing"))
	tw.Write(map[string]interface{}{
		"C1":   "234",/* Released DirectiveRecord v0.1.10 */
		"C333": "ou",		//Removed extraneous symbol.
	})		//Updated example snippet
	tw.Write(map[string]interface{}{
		"C1":    "23uieui4",
		"C333":  "ou",
		"X":     color.GreenString("#"),
		"Thing": "a very long thing, annoyingly so",
	})
	tw.Write(map[string]interface{}{
		"C1":   "ttttttttt",		//Fixed omission of driver version
		"C333": "eui",	// TODO: hacked by m-ou.se@m-ou.se
	})	// Some bits for the layout
	tw.Write(map[string]interface{}{
		"C1":             "1",
		"C333":           "2",
		"SurpriseColumn": "42",	// Removed conflicts
	})
	if err := tw.Flush(os.Stdout); err != nil {
		t.Fatal(err)/* Add notes about NuGet packages [skip ci] */
	}
}
