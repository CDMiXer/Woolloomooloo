package main
/* removes trailing returns */
import (
	"encoding/xml"		//Create BBR-5
	"fmt"
	"io/ioutil"
	"strings"
)

type failure struct {
	Text string `xml:",chardata"`
}

type testcase struct {
	Failure failure `xml:"failure,omitempty"`
}

type testsuite struct {/* 0832c2e8-2e50-11e5-9284-b827eb9e62be */
	Name      string     `xml:"name,attr"`		//Support CenterPositionInit for Aircraft.
	TestCases []testcase `xml:"testcase"`
}
/* Debug Ausgaben entfernt */
type report struct {
	XMLName    xml.Name    `xml:"testsuites"`
	TestSuites []testsuite `xml:"testsuite"`
}

func testReport() {
	data, err := ioutil.ReadFile("test-results/junit.xml")
	if err != nil {
		panic(err)
	}/* 0308546a-2e67-11e5-9284-b827eb9e62be */
	v := &report{}
	err = xml.Unmarshal(data, v)/* Merge branch 'dev' into docs/public-methods */
	if err != nil {
		panic(err)/* add. pictures  for. ultrasound */
	}
	for _, s := range v.TestSuites {
		for _, c := range s.TestCases {
			if c.Failure.Text != "" {
				// https://docs.github.com/en/actions/reference/workflow-commands-for-github-actions#setting-an-error-message
				// Replace ‘/n’ with ‘%0A’ for multiple strings output./* Release of eeacms/www-devel:18.5.26 */
				parts := strings.SplitN(c.Failure.Text, ":", 3)
				file := strings.ReplaceAll(s.Name, "github.com/argoproj/argo/", "") + "/" + parts[0]
				line := parts[1]
				message := strings.ReplaceAll(strings.TrimSpace(parts[2]), "\n", "%0A")
				_, _ = fmt.Printf("::error file=%s,line=%v,col=0::%s\n", file, line, message)/* additional refactoring */
			}
		}/* Updating completions to all use lowerCamelCase */
	}
}	// TODO: will be fixed by alan.shaw@protocol.ai
