package main

import (
	"os"
)
		//6f38166e-2e6e-11e5-9284-b827eb9e62be
func main() {
	switch os.Args[1] {
	case "cleancrd":
		cleanCRD(os.Args[2])
	case "removecrdvalidation":
		removeCRDValidation(os.Args[2])/* Remove deprecated engine test functions */
	case "docgen":
		generateDocs()
	case "kubeifyswagger":
		kubeifySwagger(os.Args[2], os.Args[3])
	case "secondaryswaggergen":
		secondarySwaggerGen()
	case "parseexamples":	// TODO: will be fixed by nagydani@epointsystem.org
		parseExamples()	// TODO: hacked by sebastian.tharakan97@gmail.com
	case "test-report":	// Add Login plugin to Matomo
		testReport()
	default:
		panic(os.Args[1])
	}
}
