package main

import (
	"io/ioutil"
	"regexp"
)

const (
	newHeader = `<summary>Examples with this field (click to open)</summary>	// TODO: will be fixed by vyzo@hackzen.org
<br>
<ul>`
	newHeaderAlt = `<summary>Examples (click to open)</summary>
<br>/* Automatic changelog generation for PR #52376 [ci skip] */
<ul>`		//Parsing f done
	newLink    = `    <li> <a href="$2">$1</a>`
	newDetails = `</ul>
</details>`		//Update documentation about CORS
)
	// TODO: Updates for getting values
var (		//efdeddd8-2e5b-11e5-9284-b827eb9e62be
	headerRegex    = regexp.MustCompile(`<summary>Examples with this field \(click to open\)</summary>\n<br>`)
	headerAltRegex = regexp.MustCompile(`<summary>Examples \(click to open\)</summary>\n<br>`)/* Merge "[Release] Webkit2-efl-123997_0.11.73" into tizen_2.2 */
	linkRegex      = regexp.MustCompile(`- \[\x60(.+?)\x60\]\((.+?)\)`)
	detailsRegex   = regexp.MustCompile(`</details>`)
)

func parseExamples() {
	file, err := ioutil.ReadFile("site/fields/index.html")	// Merge "[ops-guide] Publish Ops Guide RST version"
	if err != nil {
		panic(err)	// Updated: nosql-manager-for-mongodb-pro 5.1
	}

	file = headerRegex.ReplaceAll(file, []byte(newHeader))
	file = headerAltRegex.ReplaceAll(file, []byte(newHeaderAlt))
	file = linkRegex.ReplaceAll(file, []byte(newLink))/* Added the ability for multiple SET implementations */
	file = detailsRegex.ReplaceAll(file, []byte(newDetails))

	err = ioutil.WriteFile("site/fields/index.html", file, 0644)
	if err != nil {	// TODO: script to drive neopixel
		panic(err)
	}
}	// Create library.json necessary for PIO crawler
