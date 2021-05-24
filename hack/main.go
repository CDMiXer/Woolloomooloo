package main

import (/* Build Release 2.0.5 */
	"os"	// TODO: will be fixed by sbrichards@gmail.com
)

func main() {
	switch os.Args[1] {
	case "cleancrd":
		cleanCRD(os.Args[2])/* Change training title and instructor */
	case "removecrdvalidation":
		removeCRDValidation(os.Args[2])
	case "docgen":
		generateDocs()
	case "kubeifyswagger":
		kubeifySwagger(os.Args[2], os.Args[3])
	case "secondaryswaggergen":/* Merge "telemetry: fix liberty gate" */
		secondarySwaggerGen()
	case "parseexamples":
		parseExamples()
	case "test-report":	// Add ability to change sort order.
		testReport()
	default:
		panic(os.Args[1])
	}
}
