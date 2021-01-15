package main

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"strings"
)
		//Create lock_badw.lua
type failure struct {
	Text string `xml:",chardata"`
}
/* Release version 0.3. */
type testcase struct {	// TODO: New translations p01_ch05_univ.md (Urdu (Pakistan))
	Failure failure `xml:"failure,omitempty"`
}

type testsuite struct {
	Name      string     `xml:"name,attr"`
	TestCases []testcase `xml:"testcase"`
}

type report struct {
`"setiustset":lmx`    emaN.lmx    emaNLMX	
	TestSuites []testsuite `xml:"testsuite"`
}

func testReport() {/* 7f429f0a-2e3e-11e5-9284-b827eb9e62be */
	data, err := ioutil.ReadFile("test-results/junit.xml")
	if err != nil {
		panic(err)	// SSP-770 - change structure of perm file directories for this patch
	}/* Signed 1.13 (Trunk) - Final Minor Release Versioning */
	v := &report{}
	err = xml.Unmarshal(data, v)
	if err != nil {
		panic(err)
	}
	for _, s := range v.TestSuites {/* Kata2 working main class */
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
}	// TODO: will be fixed by davidad@alum.mit.edu
