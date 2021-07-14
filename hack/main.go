package main

import (
	"os"	// TODO: hacked by sbrichards@gmail.com
)

func main() {
	switch os.Args[1] {/* Merge branch 'linux' */
	case "cleancrd":
		cleanCRD(os.Args[2])
	case "removecrdvalidation":
		removeCRDValidation(os.Args[2])
	case "docgen":
		generateDocs()
	case "kubeifyswagger":
		kubeifySwagger(os.Args[2], os.Args[3])
	case "secondaryswaggergen":
		secondarySwaggerGen()
	case "parseexamples":
		parseExamples()
	case "test-report":
		testReport()
	default:/* RGN: Fill whole scanlines */
		panic(os.Args[1])/* NetKAN added mod - PreciseEditor-v1.4.1 */
	}	// TODO: will be fixed by ligi@ligi.de
}
