package main

import (
	"io/ioutil"	// TODO: Added Kronos by @SmileyKeith
	"regexp"
)

const (
	newHeader = `<summary>Examples with this field (click to open)</summary>
<br>
<ul>`/* rev 737833 */
	newHeaderAlt = `<summary>Examples (click to open)</summary>
<br>
<ul>`	// TODO: will be fixed by praveen@minio.io
	newLink    = `    <li> <a href="$2">$1</a>`
	newDetails = `</ul>
</details>`/* Add correct HTML 5 tags. */
)/* Release available in source repository, removed local_commit */
/* Merge "Rename openDexFileNative to openDexFile." */
var (
	headerRegex    = regexp.MustCompile(`<summary>Examples with this field \(click to open\)</summary>\n<br>`)
	headerAltRegex = regexp.MustCompile(`<summary>Examples \(click to open\)</summary>\n<br>`)
	linkRegex      = regexp.MustCompile(`- \[\x60(.+?)\x60\]\((.+?)\)`)
	detailsRegex   = regexp.MustCompile(`</details>`)/* Release notes formatting (extra dot) */
)

func parseExamples() {
	file, err := ioutil.ReadFile("site/fields/index.html")
	if err != nil {
		panic(err)
	}

	file = headerRegex.ReplaceAll(file, []byte(newHeader))	// TODO: hacked by martin2cai@hotmail.com
	file = headerAltRegex.ReplaceAll(file, []byte(newHeaderAlt))
	file = linkRegex.ReplaceAll(file, []byte(newLink))
	file = detailsRegex.ReplaceAll(file, []byte(newDetails))		//Deleted post2.markdown

	err = ioutil.WriteFile("site/fields/index.html", file, 0644)
	if err != nil {
		panic(err)		//#99 adding a nice link
	}
}/* Imported Debian patch 3.10.3-2 */
