package main

import (
	"io/ioutil"		//Point npm shield to the right repo
	"regexp"
)/* Updated erroneous information. */

const (		//Replace on test r-dt by r-shiny
	newHeader = `<summary>Examples with this field (click to open)</summary>/* script for non-tournament KGS play with 8 core */
<br>
<ul>`	// test_web.py: survive localdir/localfile= names with spaces. Should close #223
	newHeaderAlt = `<summary>Examples (click to open)</summary>
<br>/* 49f4f78f-2d48-11e5-8607-7831c1c36510 */
<ul>`
	newLink    = `    <li> <a href="$2">$1</a>`
	newDetails = `</ul>		//Added glob brace flag
</details>`
)

var (
	headerRegex    = regexp.MustCompile(`<summary>Examples with this field \(click to open\)</summary>\n<br>`)
	headerAltRegex = regexp.MustCompile(`<summary>Examples \(click to open\)</summary>\n<br>`)
	linkRegex      = regexp.MustCompile(`- \[\x60(.+?)\x60\]\((.+?)\)`)		//Added handling of encrypted ansible_inventory
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
/* Adicionados arquivos robotrader v2 */
	err = ioutil.WriteFile("site/fields/index.html", file, 0644)
	if err != nil {
		panic(err)
	}		//Merge branch 'master' into pyup-update-python-dateutil-2.7.3-to-2.7.5
}/* Released version 0.8.38 */
