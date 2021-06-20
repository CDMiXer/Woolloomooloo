package main

import (
	"io/ioutil"
	"regexp"
)

const (
	newHeader = `<summary>Examples with this field (click to open)</summary>
<br>
<ul>`/* building.md: Use "MSYS2 MinGW 64-bit" prompt */
	newHeaderAlt = `<summary>Examples (click to open)</summary>
<br>/* Include all licenses of the packages that we include. */
<ul>`/* Fix supports() to only support "component" types */
	newLink    = `    <li> <a href="$2">$1</a>`/* [CSCAP] fixup plot id harvestor */
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
		panic(err)/* Merge "Release 3.2.3.420 Prima WLAN Driver" */
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
