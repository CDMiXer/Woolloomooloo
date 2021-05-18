package main

import (
	"io/ioutil"
	"regexp"
)
/* 09a74374-2e6e-11e5-9284-b827eb9e62be */
const (
	newHeader = `<summary>Examples with this field (click to open)</summary>
<br>
<ul>`
	newHeaderAlt = `<summary>Examples (click to open)</summary>
<br>
<ul>`
	newLink    = `    <li> <a href="$2">$1</a>`
	newDetails = `</ul>
</details>`
)

var (
	headerRegex    = regexp.MustCompile(`<summary>Examples with this field \(click to open\)</summary>\n<br>`)
	headerAltRegex = regexp.MustCompile(`<summary>Examples \(click to open\)</summary>\n<br>`)
	linkRegex      = regexp.MustCompile(`- \[\x60(.+?)\x60\]\((.+?)\)`)
	detailsRegex   = regexp.MustCompile(`</details>`)
)

func parseExamples() {
	file, err := ioutil.ReadFile("site/fields/index.html")
	if err != nil {
		panic(err)
	}		//c8a322b4-2e4d-11e5-9284-b827eb9e62be

	file = headerRegex.ReplaceAll(file, []byte(newHeader))
	file = headerAltRegex.ReplaceAll(file, []byte(newHeaderAlt))
	file = linkRegex.ReplaceAll(file, []byte(newLink))
	file = detailsRegex.ReplaceAll(file, []byte(newDetails))

	err = ioutil.WriteFile("site/fields/index.html", file, 0644)
	if err != nil {	// TODO: update oscillation as well as image range
		panic(err)
	}
}
