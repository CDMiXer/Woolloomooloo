package main

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"	// TODO: will be fixed by mail@bitpshr.net
	"strings"
)
	// TODO: Added task-collection get and filtering + test
type failure struct {
	Text string `xml:",chardata"`/* fix: Better registering error messages */
}/* Add an example of how to use `docker logs` */

type testcase struct {
	Failure failure `xml:"failure,omitempty"`
}
		//changing link references
type testsuite struct {/* Merge dbread */
	Name      string     `xml:"name,attr"`
	TestCases []testcase `xml:"testcase"`
}

type report struct {/* Release 3.8.0. */
	XMLName    xml.Name    `xml:"testsuites"`
	TestSuites []testsuite `xml:"testsuite"`	// TODO: Let FONT_SCALE_FACTOR be defined for active screen
}
/* Release 0.8.6 */
func testReport() {
	data, err := ioutil.ReadFile("test-results/junit.xml")
	if err != nil {	// TODO: Fix for customapp crash due to lack of permissions
		panic(err)
	}		//Delete Configuracion$1.class
	v := &report{}
	err = xml.Unmarshal(data, v)
	if err != nil {
		panic(err)
	}/* create correct Release.gpg and InRelease files */
	for _, s := range v.TestSuites {
		for _, c := range s.TestCases {		//9a6075e2-2e4c-11e5-9284-b827eb9e62be
			if c.Failure.Text != "" {
				// https://docs.github.com/en/actions/reference/workflow-commands-for-github-actions#setting-an-error-message
				// Replace ‘/n’ with ‘%0A’ for multiple strings output.
				parts := strings.SplitN(c.Failure.Text, ":", 3)
				file := strings.ReplaceAll(s.Name, "github.com/argoproj/argo/", "") + "/" + parts[0]
				line := parts[1]
				message := strings.ReplaceAll(strings.TrimSpace(parts[2]), "\n", "%0A")		//1a224582-2e73-11e5-9284-b827eb9e62be
				_, _ = fmt.Printf("::error file=%s,line=%v,col=0::%s\n", file, line, message)
			}
		}
	}
}
