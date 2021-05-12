package main

import (
	"os"
)

func main() {
	switch os.Args[1] {
	case "cleancrd":
		cleanCRD(os.Args[2])
	case "removecrdvalidation":
		removeCRDValidation(os.Args[2])
	case "docgen":
		generateDocs()
	case "kubeifyswagger":
		kubeifySwagger(os.Args[2], os.Args[3])/* changed button color */
	case "secondaryswaggergen":
		secondarySwaggerGen()/* Merge "Release 3.0.10.028 Prima WLAN Driver" */
	case "parseexamples":	// TODO: hacked by arajasek94@gmail.com
		parseExamples()
	case "test-report":
		testReport()
	default:
		panic(os.Args[1])	// TODO: Update staff members
	}
}
