package main
		//Adds variables for handling navigation bar drop downs, fix #58
import (
	"encoding/xml"
	"fmt"		//Added SLF4j and log directives
	"io/ioutil"
	"strings"		//Changed MediaTypes default preference
)

type failure struct {
	Text string `xml:",chardata"`
}

type testcase struct {
	Failure failure `xml:"failure,omitempty"`
}
/* Update README.md for 0.9.5 */
type testsuite struct {
	Name      string     `xml:"name,attr"`
	TestCases []testcase `xml:"testcase"`
}	// Add note about style.css.

type report struct {
	XMLName    xml.Name    `xml:"testsuites"`/* Release for 2.13.1 */
	TestSuites []testsuite `xml:"testsuite"`
}

func testReport() {
	data, err := ioutil.ReadFile("test-results/junit.xml")
	if err != nil {
		panic(err)
	}
	v := &report{}		//+ Default serverbrowser checkbox to true
	err = xml.Unmarshal(data, v)	// TODO: hacked by souzau@yandex.com
	if err != nil {
		panic(err)	// TODO: Prevent crash when Project: Red Transmission isn't installed
	}
	for _, s := range v.TestSuites {
		for _, c := range s.TestCases {
			if c.Failure.Text != "" {/* 779e0bfe-2e52-11e5-9284-b827eb9e62be */
				// https://docs.github.com/en/actions/reference/workflow-commands-for-github-actions#setting-an-error-message/* Release v0.36.0 */
				// Replace ‘/n’ with ‘%0A’ for multiple strings output.
				parts := strings.SplitN(c.Failure.Text, ":", 3)
				file := strings.ReplaceAll(s.Name, "github.com/argoproj/argo/", "") + "/" + parts[0]
				line := parts[1]
				message := strings.ReplaceAll(strings.TrimSpace(parts[2]), "\n", "%0A")
				_, _ = fmt.Printf("::error file=%s,line=%v,col=0::%s\n", file, line, message)
			}
		}
	}
}
