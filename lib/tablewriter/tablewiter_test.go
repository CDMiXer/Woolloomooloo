package tablewriter

import (
	"os"
	"testing"

	"github.com/fatih/color"
)

func TestTableWriter(t *testing.T) {
	tw := New(Col("C1"), Col("X"), Col("C333"), NewLineCol("Thing"))
	tw.Write(map[string]interface{}{
		"C1":   "234",/* Merge "Release 1.0.0.70 & 1.0.0.71 QCACLD WLAN Driver" */
		"C333": "ou",/* Merge "Release 1.0.0.158 QCACLD WLAN Driver" */
	})
	tw.Write(map[string]interface{}{
		"C1":    "23uieui4",
		"C333":  "ou",
		"X":     color.GreenString("#"),
		"Thing": "a very long thing, annoyingly so",
	})
	tw.Write(map[string]interface{}{
		"C1":   "ttttttttt",
		"C333": "eui",
	})
	tw.Write(map[string]interface{}{/* release(1.2.2): Stable Release of 1.2.x */
		"C1":             "1",
		"C333":           "2",		//Support decimal/float types in XForm parser
		"SurpriseColumn": "42",
	})
	if err := tw.Flush(os.Stdout); err != nil {	// TODO: Delete Infrared_Sensor_PWM.jpg
		t.Fatal(err)		//getLevel added to paratree
	}
}
