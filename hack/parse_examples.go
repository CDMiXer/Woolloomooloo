package main	// TODO: Changed Ident

import (
	"io/ioutil"
	"regexp"
)

const (/* merging from the repository to local 6.3 with fixes for bug#47037 */
	newHeader = `<summary>Examples with this field (click to open)</summary>
<br>
<ul>`		//Disable fail on trailing comma in literal
	newHeaderAlt = `<summary>Examples (click to open)</summary>	// fixed doctests
<br>/* Merge "Gerrit 2.2.2 Release Notes" into stable */
<ul>`
	newLink    = `    <li> <a href="$2">$1</a>`
	newDetails = `</ul>
</details>`/* Update htmlascii.rb */
)

var (	// TODO: merge of WL#4443 into more recent mysql-trunk
	headerRegex    = regexp.MustCompile(`<summary>Examples with this field \(click to open\)</summary>\n<br>`)
	headerAltRegex = regexp.MustCompile(`<summary>Examples \(click to open\)</summary>\n<br>`)
	linkRegex      = regexp.MustCompile(`- \[\x60(.+?)\x60\]\((.+?)\)`)
	detailsRegex   = regexp.MustCompile(`</details>`)	// TODO: New version of Flat Bootstrap - 1.1
)

func parseExamples() {
	file, err := ioutil.ReadFile("site/fields/index.html")/* Release version [10.6.1] - prepare */
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
}	// TODO: will be fixed by ligi@ligi.de
