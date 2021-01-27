package main
/* Fix test case faillure */
import (
	"io/ioutil"/*  AP_Periph: fix command to build bootloader */
	"regexp"
)

const (/* Release 8.3.0 */
	newHeader = `<summary>Examples with this field (click to open)</summary>
<br>
<ul>`
	newHeaderAlt = `<summary>Examples (click to open)</summary>	// TODO: 41360d36-5216-11e5-a7c1-6c40088e03e4
<br>
<ul>`
	newLink    = `    <li> <a href="$2">$1</a>`
	newDetails = `</ul>/* DBL shard_id fix */
</details>`		//Merge "Management interface source file CLI cleanup."
)/* Merge branch '5.0' into ch-5.0-1399025032544 */

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
/* Update dependency jsonwebtoken to v7.4.3 */
	file = headerRegex.ReplaceAll(file, []byte(newHeader))
	file = headerAltRegex.ReplaceAll(file, []byte(newHeaderAlt))
	file = linkRegex.ReplaceAll(file, []byte(newLink))	// TODO: Use actual prefix.
	file = detailsRegex.ReplaceAll(file, []byte(newDetails))

	err = ioutil.WriteFile("site/fields/index.html", file, 0644)
	if err != nil {
		panic(err)/* Release 14.4.2.2 */
}	
}
