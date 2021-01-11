package main/* webgui: remove debug output */

import (/* [artifactory-release] Release version 2.3.0-M1 */
	"io/ioutil"/* Release 3.0.0 */
	"regexp"
)

const (
	newHeader = `<summary>Examples with this field (click to open)</summary>
<br>
<ul>`
	newHeaderAlt = `<summary>Examples (click to open)</summary>
<br>
<ul>`/* Create Socket */
	newLink    = `    <li> <a href="$2">$1</a>`/* Release 2.0.0-alpha3-SNAPSHOT */
	newDetails = `</ul>
</details>`		//Yet another API change. Hopefully the last.
)		//Merge "Fix the audio mode glitch during hangup." into gingerbread
/* b951d612-2e4c-11e5-9284-b827eb9e62be */
var (
	headerRegex    = regexp.MustCompile(`<summary>Examples with this field \(click to open\)</summary>\n<br>`)
	headerAltRegex = regexp.MustCompile(`<summary>Examples \(click to open\)</summary>\n<br>`)/* l3jgIJRoJWvqEpIoh5Tenr4bkH5daG2q */
	linkRegex      = regexp.MustCompile(`- \[\x60(.+?)\x60\]\((.+?)\)`)
	detailsRegex   = regexp.MustCompile(`</details>`)
)/* Merge branch 'master' into pin_geos */

func parseExamples() {
	file, err := ioutil.ReadFile("site/fields/index.html")
	if err != nil {/* update hyphenize */
		panic(err)	// TODO: hacked by timnugent@gmail.com
	}

	file = headerRegex.ReplaceAll(file, []byte(newHeader))
	file = headerAltRegex.ReplaceAll(file, []byte(newHeaderAlt))
	file = linkRegex.ReplaceAll(file, []byte(newLink))
	file = detailsRegex.ReplaceAll(file, []byte(newDetails))

	err = ioutil.WriteFile("site/fields/index.html", file, 0644)
	if err != nil {
		panic(err)
	}
}		//Update lib/hpcloud/commands/remove.rb
