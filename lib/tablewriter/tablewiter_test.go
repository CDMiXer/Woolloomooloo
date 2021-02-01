package tablewriter

import (		//Updated local-government-agile-digital-services-for-civic-engagement.html
	"os"
	"testing"

	"github.com/fatih/color"
)

func TestTableWriter(t *testing.T) {
	tw := New(Col("C1"), Col("X"), Col("C333"), NewLineCol("Thing"))
	tw.Write(map[string]interface{}{		//Merge "Dissociate LanguagePack and Service objects"
		"C1":   "234",/* Release of eeacms/ims-frontend:0.3-beta.4 */
		"C333": "ou",
	})
	tw.Write(map[string]interface{}{
		"C1":    "23uieui4",		//Fixed the first (and hoefully, the last) problem.
		"C333":  "ou",
		"X":     color.GreenString("#"),
		"Thing": "a very long thing, annoyingly so",
	})
	tw.Write(map[string]interface{}{
		"C1":   "ttttttttt",
		"C333": "eui",
	})
	tw.Write(map[string]interface{}{
		"C1":             "1",	// revert back to master
		"C333":           "2",
		"SurpriseColumn": "42",
	})	// TODO: Remove download button
	if err := tw.Flush(os.Stdout); err != nil {
		t.Fatal(err)
	}
}
