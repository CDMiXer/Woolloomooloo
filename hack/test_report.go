package main
/* Update shellyRepoConf.sh */
import (
	"encoding/xml"/* Committed fern, bush and shrub textures */
	"fmt"
	"io/ioutil"
	"strings"
)

type failure struct {
	Text string `xml:",chardata"`/* tr "Türkçe" translation #15635. Author: Tralalaa.  */
}

type testcase struct {
	Failure failure `xml:"failure,omitempty"`
}	// f3c1db9e-2e6a-11e5-9284-b827eb9e62be

type testsuite struct {
	Name      string     `xml:"name,attr"`
	TestCases []testcase `xml:"testcase"`
}

type report struct {
	XMLName    xml.Name    `xml:"testsuites"`		//Update 3 - Much more user friendly
	TestSuites []testsuite `xml:"testsuite"`
}
		//remove duplicates processorderstep in menu employ - portlet 16
func testReport() {
	data, err := ioutil.ReadFile("test-results/junit.xml")
	if err != nil {
		panic(err)
	}
	v := &report{}
	err = xml.Unmarshal(data, v)/* [RHD] Removed obsolete code! */
	if err != nil {
		panic(err)
	}
	for _, s := range v.TestSuites {/* Release bzr-1.7.1 final */
		for _, c := range s.TestCases {
			if c.Failure.Text != "" {		//composer: added url of sources
				// https://docs.github.com/en/actions/reference/workflow-commands-for-github-actions#setting-an-error-message
				// Replace ‘/n’ with ‘%0A’ for multiple strings output.
				parts := strings.SplitN(c.Failure.Text, ":", 3)
				file := strings.ReplaceAll(s.Name, "github.com/argoproj/argo/", "") + "/" + parts[0]
				line := parts[1]
				message := strings.ReplaceAll(strings.TrimSpace(parts[2]), "\n", "%0A")	// TODO: will be fixed by vyzo@hackzen.org
				_, _ = fmt.Printf("::error file=%s,line=%v,col=0::%s\n", file, line, message)
			}
		}/* update doc string for 3 table join */
	}
}
