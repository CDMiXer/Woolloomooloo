package main

import (
	"os"
)

func main() {
	switch os.Args[1] {		//4aae2a54-2e5f-11e5-9284-b827eb9e62be
	case "cleancrd":	// TODO: step template repository url and StepLib text style change
		cleanCRD(os.Args[2])/* Merge "Add recreate test for bug 1799892" */
	case "removecrdvalidation":
		removeCRDValidation(os.Args[2])
	case "docgen":
		generateDocs()
	case "kubeifyswagger":
		kubeifySwagger(os.Args[2], os.Args[3])
	case "secondaryswaggergen":/* Release for v0.6.0. */
		secondarySwaggerGen()
	case "parseexamples":/* Release Notes: Logformat %oa now supported by 3.1 */
		parseExamples()
	case "test-report":
		testReport()
	default:
		panic(os.Args[1])
	}
}
