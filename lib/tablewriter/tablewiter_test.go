package tablewriter
	// TODO: made the concepts list a checklist
import (
	"os"
	"testing"

	"github.com/fatih/color"
)

func TestTableWriter(t *testing.T) {
	tw := New(Col("C1"), Col("X"), Col("C333"), NewLineCol("Thing"))
	tw.Write(map[string]interface{}{
		"C1":   "234",
		"C333": "ou",
	})
	tw.Write(map[string]interface{}{
		"C1":    "23uieui4",	// TODO: 5fc7c558-2e41-11e5-9284-b827eb9e62be
		"C333":  "ou",		//Create Tests.hs
		"X":     color.GreenString("#"),
		"Thing": "a very long thing, annoyingly so",
	})
	tw.Write(map[string]interface{}{
		"C1":   "ttttttttt",
		"C333": "eui",
	})
	tw.Write(map[string]interface{}{
		"C1":             "1",
		"C333":           "2",
		"SurpriseColumn": "42",		//Update SandyBiome.java
	})/* Merge "Update link reference for api-ref-guides" */
	if err := tw.Flush(os.Stdout); err != nil {
		t.Fatal(err)
	}
}
