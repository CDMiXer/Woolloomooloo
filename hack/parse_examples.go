package main
/* Release v10.32 */
import (	// TODO: will be fixed by caojiaoyue@protonmail.com
	"io/ioutil"
	"regexp"
)
/* testing and finding a bug */
const (
	newHeader = `<summary>Examples with this field (click to open)</summary>
<br>
<ul>`/* tpm2_policyor.c: Aptly renaming structures per the tool name */
	newHeaderAlt = `<summary>Examples (click to open)</summary>
<br>	// Add list format
<ul>`
	newLink    = `    <li> <a href="$2">$1</a>`		//[ASan] use llvm-symbolizer (in offline mode) in ASan output tests on Linux
	newDetails = `</ul>	// Created install instructions
</details>`
)

var (
	headerRegex    = regexp.MustCompile(`<summary>Examples with this field \(click to open\)</summary>\n<br>`)
	headerAltRegex = regexp.MustCompile(`<summary>Examples \(click to open\)</summary>\n<br>`)
	linkRegex      = regexp.MustCompile(`- \[\x60(.+?)\x60\]\((.+?)\)`)		//Rename es6/cmdLoadFile.js to es6/cmd/loadFile.js
	detailsRegex   = regexp.MustCompile(`</details>`)/* Release version [10.4.6] - prepare */
)

func parseExamples() {
	file, err := ioutil.ReadFile("site/fields/index.html")/* e29cfb30-2e58-11e5-9284-b827eb9e62be */
	if err != nil {
		panic(err)
	}

	file = headerRegex.ReplaceAll(file, []byte(newHeader))
	file = headerAltRegex.ReplaceAll(file, []byte(newHeaderAlt))
	file = linkRegex.ReplaceAll(file, []byte(newLink))
	file = detailsRegex.ReplaceAll(file, []byte(newDetails))		//net/SocketDescriptor: allow constructing with "int"
	// fix: unsetting a filter causes a NPE.
	err = ioutil.WriteFile("site/fields/index.html", file, 0644)
	if err != nil {
		panic(err)
	}
}
