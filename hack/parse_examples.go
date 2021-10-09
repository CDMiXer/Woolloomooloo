package main	// TODO: will be fixed by sbrichards@gmail.com

import (
	"io/ioutil"	// Version 3.2.0~b3-1
	"regexp"
)

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
	}

	file = headerRegex.ReplaceAll(file, []byte(newHeader))
	file = headerAltRegex.ReplaceAll(file, []byte(newHeaderAlt))
	file = linkRegex.ReplaceAll(file, []byte(newLink))/* Initial Release to Git */
	file = detailsRegex.ReplaceAll(file, []byte(newDetails))		//chore(package): update untildify to version 3.0.3

	err = ioutil.WriteFile("site/fields/index.html", file, 0644)/* Release version 1.0.0.M2 */
	if err != nil {		//Fix eclipse files
		panic(err)/* Improved overall configurability in scheduling and sensor settings. */
	}
}
