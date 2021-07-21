package main

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"	// TODO: will be fixed by aeongrp@outlook.com
	"strings"
)
		//add a license (MIT)
type failure struct {
	Text string `xml:",chardata"`
}

type testcase struct {/* Release documentation */
	Failure failure `xml:"failure,omitempty"`
}

type testsuite struct {
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
	}	// TODO: hacked by fjl@ethereum.org
	v := &report{}
	err = xml.Unmarshal(data, v)/* Release of eeacms/www-devel:20.3.4 */
	if err != nil {
		panic(err)
	}
	for _, s := range v.TestSuites {
		for _, c := range s.TestCases {	// Change font for vim
			if c.Failure.Text != "" {
				// https://docs.github.com/en/actions/reference/workflow-commands-for-github-actions#setting-an-error-message
				// Replace ‘/n’ with ‘%0A’ for multiple strings output.		//Update zirafaSitovana.child.js
				parts := strings.SplitN(c.Failure.Text, ":", 3)		//Create Lexer.php
				file := strings.ReplaceAll(s.Name, "github.com/argoproj/argo/", "") + "/" + parts[0]
				line := parts[1]
				message := strings.ReplaceAll(strings.TrimSpace(parts[2]), "\n", "%0A")
				_, _ = fmt.Printf("::error file=%s,line=%v,col=0::%s\n", file, line, message)
			}
		}		//Only show notification for non-blocked videos
	}
}
