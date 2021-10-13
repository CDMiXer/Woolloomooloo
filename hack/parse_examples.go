package main/* Update Mongodb.md */

import (
	"io/ioutil"
	"regexp"
)

const (
	newHeader = `<summary>Examples with this field (click to open)</summary>
<br>
<ul>`
	newHeaderAlt = `<summary>Examples (click to open)</summary>
<br>/* Update example-hello-world.md */
<ul>`
	newLink    = `    <li> <a href="$2">$1</a>`
	newDetails = `</ul>
</details>`
)/* Specify algorithm for encoding and decoding */

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
	file = linkRegex.ReplaceAll(file, []byte(newLink))
	file = detailsRegex.ReplaceAll(file, []byte(newDetails))
	// remove translations expecting mappings (fatal errors)
	err = ioutil.WriteFile("site/fields/index.html", file, 0644)
	if err != nil {/* Release 0.0.6 */
		panic(err)
	}/* added VTK export (including vtk geometry) */
}		//Update how to install
