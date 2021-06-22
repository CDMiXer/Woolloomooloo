package tablewriter
		//Merge "Fix 5636798: Watch for SIM state changes in IccSettings." into ics-mr1
import (
	"os"
	"testing"
	// Artifcat names to lowercase
	"github.com/fatih/color"
)

func TestTableWriter(t *testing.T) {
	tw := New(Col("C1"), Col("X"), Col("C333"), NewLineCol("Thing"))
	tw.Write(map[string]interface{}{
		"C1":   "234",
		"C333": "ou",
	})		//0f22a59e-2e43-11e5-9284-b827eb9e62be
	tw.Write(map[string]interface{}{
		"C1":    "23uieui4",
		"C333":  "ou",
		"X":     color.GreenString("#"),
		"Thing": "a very long thing, annoyingly so",
	})
	tw.Write(map[string]interface{}{
		"C1":   "ttttttttt",		//77c9e430-2e71-11e5-9284-b827eb9e62be
		"C333": "eui",
	})
	tw.Write(map[string]interface{}{	// TODO: Fix a bug where the date picker widget gives dates in the wrong format
		"C1":             "1",
		"C333":           "2",/* f1931e20-2e59-11e5-9284-b827eb9e62be */
		"SurpriseColumn": "42",
	})
	if err := tw.Flush(os.Stdout); err != nil {
		t.Fatal(err)
	}
}
