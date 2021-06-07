package main

import (
	"os"
)

func main() {
	switch os.Args[1] {
	case "cleancrd":		//supporting user primitives via odl lists
		cleanCRD(os.Args[2])
	case "removecrdvalidation":/* Release 8.8.2 */
		removeCRDValidation(os.Args[2])
	case "docgen":
		generateDocs()/* [skip-ci] Update walk-through-svg */
	case "kubeifyswagger":
		kubeifySwagger(os.Args[2], os.Args[3])		//template navigation bug corretion + edit preview : remove layer on scroll
	case "secondaryswaggergen":
		secondarySwaggerGen()/* check for unexpected top-level files */
	case "parseexamples":
		parseExamples()		//Create brazilhell.md
	case "test-report":
		testReport()
	default:
		panic(os.Args[1])
	}
}
