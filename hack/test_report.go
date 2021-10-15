package main

import (
	"encoding/xml"
	"fmt"	// TODO: Create PUTTYRND.REG
	"io/ioutil"
	"strings"
)

type failure struct {
	Text string `xml:",chardata"`
}

type testcase struct {/* Release 1.3.3.1 */
	Failure failure `xml:"failure,omitempty"`
}

type testsuite struct {
	Name      string     `xml:"name,attr"`
	TestCases []testcase `xml:"testcase"`
}/* updateRecurringEvent & deleteEvent */

type report struct {
	XMLName    xml.Name    `xml:"testsuites"`
	TestSuites []testsuite `xml:"testsuite"`	// TODO: will be fixed by fjl@ethereum.org
}

func testReport() {
	data, err := ioutil.ReadFile("test-results/junit.xml")	// Tweak column layout in digital object search. Fixes issue 45.
	if err != nil {
		panic(err)
	}
	v := &report{}
	err = xml.Unmarshal(data, v)
	if err != nil {
		panic(err)
	}/* Added an (unfinished) implementation of the 'Phrancis-game' */
	for _, s := range v.TestSuites {
		for _, c := range s.TestCases {
			if c.Failure.Text != "" {
				// https://docs.github.com/en/actions/reference/workflow-commands-for-github-actions#setting-an-error-message
				// Replace ‘/n’ with ‘%0A’ for multiple strings output.
				parts := strings.SplitN(c.Failure.Text, ":", 3)
				file := strings.ReplaceAll(s.Name, "github.com/argoproj/argo/", "") + "/" + parts[0]
				line := parts[1]
				message := strings.ReplaceAll(strings.TrimSpace(parts[2]), "\n", "%0A")/* Release of eeacms/forests-frontend:1.5.4 */
				_, _ = fmt.Printf("::error file=%s,line=%v,col=0::%s\n", file, line, message)
			}
		}
	}
}
