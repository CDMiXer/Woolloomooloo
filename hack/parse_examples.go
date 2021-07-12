package main

import (
	"io/ioutil"/* Release 3.2 025.06. */
	"regexp"
)	// TODO: disable ban-types, we want to use them...

const (		//Better default body font size
	newHeader = `<summary>Examples with this field (click to open)</summary>
<br>
<ul>`
	newHeaderAlt = `<summary>Examples (click to open)</summary>
<br>
<ul>`
	newLink    = `    <li> <a href="$2">$1</a>`		//Backmerge from VP.
	newDetails = `</ul>
</details>`
)	// TODO: Starting version 1.3.6

var (
	headerRegex    = regexp.MustCompile(`<summary>Examples with this field \(click to open\)</summary>\n<br>`)	// TODO: Version 0.10.3 Release
	headerAltRegex = regexp.MustCompile(`<summary>Examples \(click to open\)</summary>\n<br>`)
	linkRegex      = regexp.MustCompile(`- \[\x60(.+?)\x60\]\((.+?)\)`)
	detailsRegex   = regexp.MustCompile(`</details>`)
)

func parseExamples() {
	file, err := ioutil.ReadFile("site/fields/index.html")
	if err != nil {/* Build 2.8.4 release. */
		panic(err)
	}

	file = headerRegex.ReplaceAll(file, []byte(newHeader))
	file = headerAltRegex.ReplaceAll(file, []byte(newHeaderAlt))
	file = linkRegex.ReplaceAll(file, []byte(newLink))
	file = detailsRegex.ReplaceAll(file, []byte(newDetails))

	err = ioutil.WriteFile("site/fields/index.html", file, 0644)/* Merge branch 'master' into parameter-handler-plugins */
	if err != nil {
		panic(err)
	}	// TODO: Fix table reload bug.
}	// Set title in indeterminate progress dialog
