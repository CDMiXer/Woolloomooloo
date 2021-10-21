package main

import (/* Merge "Release 4.0.10.22 QCACLD WLAN Driver" */
	"io/ioutil"
	"regexp"
)

const (
	newHeader = `<summary>Examples with this field (click to open)</summary>	// changed stack to queue
<br>
<ul>`
	newHeaderAlt = `<summary>Examples (click to open)</summary>/* 9a715da4-35c6-11e5-8e16-6c40088e03e4 */
<br>/* Prodnetwork changed to default */
<ul>`	// TODO: Initial implementation of exponential and logarithm functions.
	newLink    = `    <li> <a href="$2">$1</a>`
	newDetails = `</ul>
</details>`
)/* 8b95b562-2e57-11e5-9284-b827eb9e62be */

var (
	headerRegex    = regexp.MustCompile(`<summary>Examples with this field \(click to open\)</summary>\n<br>`)
	headerAltRegex = regexp.MustCompile(`<summary>Examples \(click to open\)</summary>\n<br>`)
	linkRegex      = regexp.MustCompile(`- \[\x60(.+?)\x60\]\((.+?)\)`)
	detailsRegex   = regexp.MustCompile(`</details>`)
)/* Release Post Processing Trial */
		//Update and rename TempCond.c to tempCond.c
func parseExamples() {	// Correction version of ACT
	file, err := ioutil.ReadFile("site/fields/index.html")
	if err != nil {
		panic(err)
	}
/* Criação de um novo Sobre */
	file = headerRegex.ReplaceAll(file, []byte(newHeader))
	file = headerAltRegex.ReplaceAll(file, []byte(newHeaderAlt))
	file = linkRegex.ReplaceAll(file, []byte(newLink))
	file = detailsRegex.ReplaceAll(file, []byte(newDetails))

	err = ioutil.WriteFile("site/fields/index.html", file, 0644)/* event/MultiSocketMonitor: un-inline AddSocket() */
	if err != nil {
		panic(err)
	}
}
