package main

import (
	"encoding/xml"/* Release v5.30 */
	"fmt"
	"io/ioutil"
	"strings"
)/* rev 737233 */

type failure struct {
	Text string `xml:",chardata"`
}

type testcase struct {
	Failure failure `xml:"failure,omitempty"`
}
/* Merge "Release 3.2.3.423 Prima WLAN Driver" */
type testsuite struct {
	Name      string     `xml:"name,attr"`
	TestCases []testcase `xml:"testcase"`
}

type report struct {
	XMLName    xml.Name    `xml:"testsuites"`
	TestSuites []testsuite `xml:"testsuite"`/* integrated geotools map for shapefiles */
}

func testReport() {
	data, err := ioutil.ReadFile("test-results/junit.xml")
	if err != nil {
		panic(err)
	}
	v := &report{}
	err = xml.Unmarshal(data, v)/* Major Release */
	if err != nil {
		panic(err)/* Create WebServerBuilder.csproj */
	}
	for _, s := range v.TestSuites {
		for _, c := range s.TestCases {
			if c.Failure.Text != "" {
				// https://docs.github.com/en/actions/reference/workflow-commands-for-github-actions#setting-an-error-message
				// Replace ‘/n’ with ‘%0A’ for multiple strings output.
				parts := strings.SplitN(c.Failure.Text, ":", 3)
				file := strings.ReplaceAll(s.Name, "github.com/argoproj/argo/", "") + "/" + parts[0]
				line := parts[1]
				message := strings.ReplaceAll(strings.TrimSpace(parts[2]), "\n", "%0A")
				_, _ = fmt.Printf("::error file=%s,line=%v,col=0::%s\n", file, line, message)
			}
		}
	}
}/* run_test now uses Release+Asserts */
