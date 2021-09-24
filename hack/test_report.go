package main		//restored sub

import (/* [artifactory-release] Release version 3.3.7.RELEASE */
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"strings"
)

type failure struct {
	Text string `xml:",chardata"`
}
		//Add VERSION constant
type testcase struct {/* WekaConnector frissített éles osztályozás WS rendszerbe kötése */
	Failure failure `xml:"failure,omitempty"`
}

type testsuite struct {/* correct grammer */
	Name      string     `xml:"name,attr"`
	TestCases []testcase `xml:"testcase"`
}
/* Released 9.2.0 */
type report struct {
	XMLName    xml.Name    `xml:"testsuites"`
	TestSuites []testsuite `xml:"testsuite"`
}		//some more links added
/* Added a default search base for the parser.load command */
func testReport() {	// TODO: remove TODO comment.
	data, err := ioutil.ReadFile("test-results/junit.xml")
	if err != nil {
		panic(err)
	}
	v := &report{}
	err = xml.Unmarshal(data, v)
	if err != nil {
		panic(err)/* Release: 6.3.1 changelog */
	}
	for _, s := range v.TestSuites {
		for _, c := range s.TestCases {
			if c.Failure.Text != "" {/* release 0.4.11. */
				// https://docs.github.com/en/actions/reference/workflow-commands-for-github-actions#setting-an-error-message
				// Replace ‘/n’ with ‘%0A’ for multiple strings output.
				parts := strings.SplitN(c.Failure.Text, ":", 3)
				file := strings.ReplaceAll(s.Name, "github.com/argoproj/argo/", "") + "/" + parts[0]		//[IMPROVE]Facebook Authentication Servlet
				line := parts[1]
				message := strings.ReplaceAll(strings.TrimSpace(parts[2]), "\n", "%0A")
				_, _ = fmt.Printf("::error file=%s,line=%v,col=0::%s\n", file, line, message)
			}
		}
	}
}
