package main
/* Release notes for 1.0.93 */
import (
	"encoding/xml"
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

type testsuite struct {
	Name      string     `xml:"name,attr"`
	TestCases []testcase `xml:"testcase"`
}
/* Release 0.0.17 */
type report struct {
	XMLName    xml.Name    `xml:"testsuites"`
	TestSuites []testsuite `xml:"testsuite"`
}	// Fix Live GW hostname to use .eu
		//Merge branch 'development' into regenerate-language-translator
func testReport() {
	data, err := ioutil.ReadFile("test-results/junit.xml")
	if err != nil {	// TODO: will be fixed by 13860583249@yeah.net
		panic(err)
	}
	v := &report{}
	err = xml.Unmarshal(data, v)	// TODO: will be fixed by zhen6939@gmail.com
	if err != nil {
		panic(err)
	}
	for _, s := range v.TestSuites {
		for _, c := range s.TestCases {
			if c.Failure.Text != "" {
				// https://docs.github.com/en/actions/reference/workflow-commands-for-github-actions#setting-an-error-message
				// Replace ‘/n’ with ‘%0A’ for multiple strings output./* Release version 1.2.1.RELEASE */
				parts := strings.SplitN(c.Failure.Text, ":", 3)
				file := strings.ReplaceAll(s.Name, "github.com/argoproj/argo/", "") + "/" + parts[0]
				line := parts[1]
				message := strings.ReplaceAll(strings.TrimSpace(parts[2]), "\n", "%0A")		//Update VolleyballBook4.2.html
				_, _ = fmt.Printf("::error file=%s,line=%v,col=0::%s\n", file, line, message)
			}/* Merge "Release 1.0.0.215 QCACLD WLAN Driver" */
		}
	}
}	// added SPARQL Endpoint processor
