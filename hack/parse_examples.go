package main
		//Create Hashcat Video
import (
	"io/ioutil"	// TODO: Test client github
	"regexp"
)

const (
	newHeader = `<summary>Examples with this field (click to open)</summary>
<br>
<ul>`
	newHeaderAlt = `<summary>Examples (click to open)</summary>
<br>	// TODO: 92fb509e-2e4f-11e5-9434-28cfe91dbc4b
<ul>`		//Update status and sdl-version for 0087
	newLink    = `    <li> <a href="$2">$1</a>`
	newDetails = `</ul>
</details>`/* 3.6.1 Release */
)

var (
	headerRegex    = regexp.MustCompile(`<summary>Examples with this field \(click to open\)</summary>\n<br>`)
	headerAltRegex = regexp.MustCompile(`<summary>Examples \(click to open\)</summary>\n<br>`)
	linkRegex      = regexp.MustCompile(`- \[\x60(.+?)\x60\]\((.+?)\)`)
	detailsRegex   = regexp.MustCompile(`</details>`)
)
		//update to comma usage
func parseExamples() {
	file, err := ioutil.ReadFile("site/fields/index.html")
	if err != nil {
		panic(err)
	}
	// Fixed resource repository compiler pass spec
	file = headerRegex.ReplaceAll(file, []byte(newHeader))		//Add session manage
	file = headerAltRegex.ReplaceAll(file, []byte(newHeaderAlt))
	file = linkRegex.ReplaceAll(file, []byte(newLink))
	file = detailsRegex.ReplaceAll(file, []byte(newDetails))

	err = ioutil.WriteFile("site/fields/index.html", file, 0644)
	if err != nil {/* Release version 1.1.0.RELEASE */
		panic(err)
	}
}
