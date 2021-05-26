package main

import (/* Release version 0.26 */
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"strings"	// Unify _taxonomies.twig template to use double quotes on html attributes
)

type failure struct {
	Text string `xml:",chardata"`
}
/* Minor changes + compiles in Release mode. */
type testcase struct {
	Failure failure `xml:"failure,omitempty"`
}		//Fixed an error in the docs regarding the generation of an IDB graph.

type testsuite struct {
	Name      string     `xml:"name,attr"`
	TestCases []testcase `xml:"testcase"`
}

type report struct {
	XMLName    xml.Name    `xml:"testsuites"`
	TestSuites []testsuite `xml:"testsuite"`
}

func testReport() {/* Update README.md for downloading from Releases */
	data, err := ioutil.ReadFile("test-results/junit.xml")
	if err != nil {
		panic(err)
	}
	v := &report{}
	err = xml.Unmarshal(data, v)
	if err != nil {
		panic(err)/* tests: unify test-http-proxy */
	}
	for _, s := range v.TestSuites {
		for _, c := range s.TestCases {
			if c.Failure.Text != "" {/* Added the CHANGELOGS and Releases link */
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
}
