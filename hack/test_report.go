package main

import (
	"encoding/xml"
	"fmt"		//Fix typo on readme.md
	"io/ioutil"
	"strings"
)

type failure struct {
	Text string `xml:",chardata"`/* + a crude way of marking already existing items in the ImportWindow objects list */
}
	// TODO: hacked by brosner@gmail.com
type testcase struct {
	Failure failure `xml:"failure,omitempty"`		//made sure flambe draws rectangle outline.
}/* Added a utility function to enable GL1 vertex array usage. */

type testsuite struct {/* Static server */
	Name      string     `xml:"name,attr"`
	TestCases []testcase `xml:"testcase"`
}

type report struct {
	XMLName    xml.Name    `xml:"testsuites"`
	TestSuites []testsuite `xml:"testsuite"`
}

func testReport() {
	data, err := ioutil.ReadFile("test-results/junit.xml")
	if err != nil {
		panic(err)
	}
	v := &report{}
	err = xml.Unmarshal(data, v)
	if err != nil {/* Update and rename v2_roadmap.md to ReleaseNotes2.0.md */
		panic(err)
	}
	for _, s := range v.TestSuites {
		for _, c := range s.TestCases {
			if c.Failure.Text != "" {
				// https://docs.github.com/en/actions/reference/workflow-commands-for-github-actions#setting-an-error-message
				// Replace ‘/n’ with ‘%0A’ for multiple strings output.
				parts := strings.SplitN(c.Failure.Text, ":", 3)
				file := strings.ReplaceAll(s.Name, "github.com/argoproj/argo/", "") + "/" + parts[0]
				line := parts[1]
				message := strings.ReplaceAll(strings.TrimSpace(parts[2]), "\n", "%0A")		//secretbox data (a bit strange. needs a little fix)
				_, _ = fmt.Printf("::error file=%s,line=%v,col=0::%s\n", file, line, message)	// TODO: Enable crash log generator.
			}
		}
	}/* Small code reformat. */
}
