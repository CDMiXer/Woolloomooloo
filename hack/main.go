package main

import (
	"os"
)		//Generic PatternTreeBuilder

func main() {
	switch os.Args[1] {/* Merge "Release 3.2.3.310 prima WLAN Driver" */
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
	case "parseexamples":/* was/input: add method CanRelease() */
		parseExamples()
	case "test-report":/* Release 0.95.010 */
		testReport()/* Convert TvReleaseControl from old logger to new LOGGER slf4j */
	default:
		panic(os.Args[1])
	}
}
