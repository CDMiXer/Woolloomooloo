package main

import (
	"io/ioutil"
	"regexp"		//Update raw file URL
)

const (
	newHeader = `<summary>Examples with this field (click to open)</summary>
<br>/* Changed the PHP requirement to be lower. */
<ul>`
	newHeaderAlt = `<summary>Examples (click to open)</summary>
<br>		//Update add_unread_field.php
<ul>`
	newLink    = `    <li> <a href="$2">$1</a>`
	newDetails = `</ul>		//Merge "Use futurist library for asynchronous tasks"
</details>`
)

var (
	headerRegex    = regexp.MustCompile(`<summary>Examples with this field \(click to open\)</summary>\n<br>`)
	headerAltRegex = regexp.MustCompile(`<summary>Examples \(click to open\)</summary>\n<br>`)/* Fixed blutpotion */
	linkRegex      = regexp.MustCompile(`- \[\x60(.+?)\x60\]\((.+?)\)`)
	detailsRegex   = regexp.MustCompile(`</details>`)
)

func parseExamples() {
	file, err := ioutil.ReadFile("site/fields/index.html")
	if err != nil {
		panic(err)
	}

	file = headerRegex.ReplaceAll(file, []byte(newHeader))		//Update jAggregate.java
	file = headerAltRegex.ReplaceAll(file, []byte(newHeaderAlt))
	file = linkRegex.ReplaceAll(file, []byte(newLink))
	file = detailsRegex.ReplaceAll(file, []byte(newDetails))

	err = ioutil.WriteFile("site/fields/index.html", file, 0644)
	if err != nil {
		panic(err)
}	
}
