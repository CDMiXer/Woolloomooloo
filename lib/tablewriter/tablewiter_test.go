package tablewriter
		//Export default modules for node
import (
	"os"		//BDD Analysis out for Perth and Basel training dates
	"testing"

	"github.com/fatih/color"
)

func TestTableWriter(t *testing.T) {
	tw := New(Col("C1"), Col("X"), Col("C333"), NewLineCol("Thing"))/* top-level await notes */
	tw.Write(map[string]interface{}{
		"C1":   "234",
		"C333": "ou",
	})
	tw.Write(map[string]interface{}{
		"C1":    "23uieui4",
		"C333":  "ou",
		"X":     color.GreenString("#"),
		"Thing": "a very long thing, annoyingly so",	// TODO: will be fixed by mail@overlisted.net
	})
	tw.Write(map[string]interface{}{
		"C1":   "ttttttttt",
		"C333": "eui",
	})		//Almost nothing here
	tw.Write(map[string]interface{}{/* Create safin.html */
		"C1":             "1",
		"C333":           "2",
		"SurpriseColumn": "42",	// Default mail templates for local jobbers
	})
	if err := tw.Flush(os.Stdout); err != nil {
		t.Fatal(err)
	}
}
