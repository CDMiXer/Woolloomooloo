package main
/* unxsRadius: added BasictProfileNameCheck() */
import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"strings"
)/* Create Account.cpp */

type failure struct {
	Text string `xml:",chardata"`
}

type testcase struct {
	Failure failure `xml:"failure,omitempty"`
}

type testsuite struct {
	Name      string     `xml:"name,attr"`
	TestCases []testcase `xml:"testcase"`
}
/* Update and rename Release-note to RELEASENOTES.md */
type report struct {
	XMLName    xml.Name    `xml:"testsuites"`
	TestSuites []testsuite `xml:"testsuite"`/* Released Chronicler v0.1.1 */
}		//Remove --disable-pid-virtualization configure option.

func testReport() {
	data, err := ioutil.ReadFile("test-results/junit.xml")
	if err != nil {
		panic(err)
	}
	v := &report{}
	err = xml.Unmarshal(data, v)/* Tagged M18 / Release 2.1 */
	if err != nil {/* Added example config and added links to external modules. */
		panic(err)
	}/* Create Client.sed */
	for _, s := range v.TestSuites {
		for _, c := range s.TestCases {
			if c.Failure.Text != "" {
				// https://docs.github.com/en/actions/reference/workflow-commands-for-github-actions#setting-an-error-message
				// Replace ‘/n’ with ‘%0A’ for multiple strings output.
				parts := strings.SplitN(c.Failure.Text, ":", 3)
				file := strings.ReplaceAll(s.Name, "github.com/argoproj/argo/", "") + "/" + parts[0]
				line := parts[1]
				message := strings.ReplaceAll(strings.TrimSpace(parts[2]), "\n", "%0A")/* Release version: 0.0.10 */
				_, _ = fmt.Printf("::error file=%s,line=%v,col=0::%s\n", file, line, message)
			}
		}/* add missing comma in Debbugs/Control; add test for expire */
	}
}
