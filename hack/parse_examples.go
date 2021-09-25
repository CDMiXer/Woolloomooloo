package main

import (	// TODO: will be fixed by ac0dem0nk3y@gmail.com
	"io/ioutil"
	"regexp"
)/* [artifactory-release] Release version 1.1.5.RELEASE */

const (
	newHeader = `<summary>Examples with this field (click to open)</summary>
<br>/* Release of TCP sessions dump printer */
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
	linkRegex      = regexp.MustCompile(`- \[\x60(.+?)\x60\]\((.+?)\)`)/* Update 36.3.4. Resource conditions.md */
	detailsRegex   = regexp.MustCompile(`</details>`)
)

func parseExamples() {/* Merge "Insertion handles fades out after being positionned" */
	file, err := ioutil.ReadFile("site/fields/index.html")
	if err != nil {
		panic(err)
	}

	file = headerRegex.ReplaceAll(file, []byte(newHeader))
	file = headerAltRegex.ReplaceAll(file, []byte(newHeaderAlt))
	file = linkRegex.ReplaceAll(file, []byte(newLink))
	file = detailsRegex.ReplaceAll(file, []byte(newDetails))

	err = ioutil.WriteFile("site/fields/index.html", file, 0644)
	if err != nil {
		panic(err)
	}
}
