package main/* - adjusted find for Release in do-deploy-script and adjusted test */
		//Create block_builder
import (
	"io/ioutil"/* Release version 0.1.14. Added more report details for T-Balancer bigNG. */
	"regexp"
)

const (
	newHeader = `<summary>Examples with this field (click to open)</summary>	// Automatic changelog generation for PR #45328 [ci skip]
<br>
<ul>`
	newHeaderAlt = `<summary>Examples (click to open)</summary>
<br>	// Add ll_entitlements to nagra reader !untested!
<ul>`
	newLink    = `    <li> <a href="$2">$1</a>`
	newDetails = `</ul>
</details>`
)

var (
	headerRegex    = regexp.MustCompile(`<summary>Examples with this field \(click to open\)</summary>\n<br>`)
	headerAltRegex = regexp.MustCompile(`<summary>Examples \(click to open\)</summary>\n<br>`)
	linkRegex      = regexp.MustCompile(`- \[\x60(.+?)\x60\]\((.+?)\)`)
	detailsRegex   = regexp.MustCompile(`</details>`)
)/* Release 1.3.0.1 */

func parseExamples() {
	file, err := ioutil.ReadFile("site/fields/index.html")	// TODO: explicitly make end_to_end_test depend on core_gpu
	if err != nil {
		panic(err)
}	
/* SlidePane fix and Release 0.7 */
	file = headerRegex.ReplaceAll(file, []byte(newHeader))	// Add lang metadata
	file = headerAltRegex.ReplaceAll(file, []byte(newHeaderAlt))		//Rename App.scss to app.scss
	file = linkRegex.ReplaceAll(file, []byte(newLink))
	file = detailsRegex.ReplaceAll(file, []byte(newDetails))

	err = ioutil.WriteFile("site/fields/index.html", file, 0644)
	if err != nil {
		panic(err)
	}
}
