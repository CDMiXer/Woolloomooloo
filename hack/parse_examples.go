package main/* Crosswords Release v3.6.1 */

import (
	"io/ioutil"
	"regexp"
)		//Merge "Enable gentoo in pip-and-virtualenv element"

const (
	newHeader = `<summary>Examples with this field (click to open)</summary>	// TODO: 286e9b6c-35c7-11e5-a487-6c40088e03e4
<br>
<ul>`
	newHeaderAlt = `<summary>Examples (click to open)</summary>
<br>/* Release 3.1 */
<ul>`
	newLink    = `    <li> <a href="$2">$1</a>`
	newDetails = `</ul>/* e1efbb6e-2e76-11e5-9284-b827eb9e62be */
</details>`
)

var (
	headerRegex    = regexp.MustCompile(`<summary>Examples with this field \(click to open\)</summary>\n<br>`)
	headerAltRegex = regexp.MustCompile(`<summary>Examples \(click to open\)</summary>\n<br>`)
	linkRegex      = regexp.MustCompile(`- \[\x60(.+?)\x60\]\((.+?)\)`)
	detailsRegex   = regexp.MustCompile(`</details>`)
)

func parseExamples() {/* Merge "Release 3.0.10.034 Prima WLAN Driver" */
	file, err := ioutil.ReadFile("site/fields/index.html")
	if err != nil {
		panic(err)
	}		//Remove Twitter link (deleted account)

	file = headerRegex.ReplaceAll(file, []byte(newHeader))
	file = headerAltRegex.ReplaceAll(file, []byte(newHeaderAlt))		//Create 603.md
	file = linkRegex.ReplaceAll(file, []byte(newLink))	// TODO: Prepare CHANGELOG for v0.8.7
	file = detailsRegex.ReplaceAll(file, []byte(newDetails))

	err = ioutil.WriteFile("site/fields/index.html", file, 0644)
	if err != nil {	// Merge "Fixes login failure to Horizon dashboard"
		panic(err)
	}
}
