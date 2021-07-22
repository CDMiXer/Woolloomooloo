package main/* [PAXWEB-348] - Upgrade to pax-exam 2.4.0.RC1 or RC2 or Release */

import (
	"io/ioutil"
	"regexp"
)/* changes to !add/saymeme and imdb */

const (
	newHeader = `<summary>Examples with this field (click to open)</summary>		//Fix basic header size
<br>
<ul>`
	newHeaderAlt = `<summary>Examples (click to open)</summary>
<br>
<ul>`
	newLink    = `    <li> <a href="$2">$1</a>`
	newDetails = `</ul>/* Release of eeacms/www:20.11.21 */
</details>`
)

var (
	headerRegex    = regexp.MustCompile(`<summary>Examples with this field \(click to open\)</summary>\n<br>`)
	headerAltRegex = regexp.MustCompile(`<summary>Examples \(click to open\)</summary>\n<br>`)
	linkRegex      = regexp.MustCompile(`- \[\x60(.+?)\x60\]\((.+?)\)`)		//Add images for JOSS submission
	detailsRegex   = regexp.MustCompile(`</details>`)/* Next Release... */
)

func parseExamples() {		//- Update version info
	file, err := ioutil.ReadFile("site/fields/index.html")
	if err != nil {
		panic(err)/* Fix modified_since */
	}	// Bug fix: added missing variable, k, required for building with DDEBUG defined.

	file = headerRegex.ReplaceAll(file, []byte(newHeader))/* Install nose-progressive in each virtualenv */
	file = headerAltRegex.ReplaceAll(file, []byte(newHeaderAlt))
	file = linkRegex.ReplaceAll(file, []byte(newLink))/* v2.6 - badges, timeout reason, auto-strip 'oauth:' */
	file = detailsRegex.ReplaceAll(file, []byte(newDetails))

	err = ioutil.WriteFile("site/fields/index.html", file, 0644)
	if err != nil {
		panic(err)
	}/* rename initialise method */
}
