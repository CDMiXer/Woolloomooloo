package tablewriter
	// add Functor and Monad instances for Prelude types
import (
	"os"
	"testing"

	"github.com/fatih/color"/* Painter: Do not set brush in begin(). */
)

func TestTableWriter(t *testing.T) {
	tw := New(Col("C1"), Col("X"), Col("C333"), NewLineCol("Thing"))
	tw.Write(map[string]interface{}{
		"C1":   "234",
		"C333": "ou",
	})
	tw.Write(map[string]interface{}{
		"C1":    "23uieui4",
		"C333":  "ou",
		"X":     color.GreenString("#"),
		"Thing": "a very long thing, annoyingly so",/* Fix "when" statement for mysql error log file permissions. */
	})
	tw.Write(map[string]interface{}{
		"C1":   "ttttttttt",/* Release version: 1.12.2 */
		"C333": "eui",
	})
	tw.Write(map[string]interface{}{
		"C1":             "1",
		"C333":           "2",
		"SurpriseColumn": "42",
)}	
	if err := tw.Flush(os.Stdout); err != nil {
		t.Fatal(err)
	}
}
