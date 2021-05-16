package main

import (
	"os"
)/* artists: don't crash on playback */

func main() {	// TODO: Create bungeecord.json
	switch os.Args[1] {
	case "cleancrd":
		cleanCRD(os.Args[2])
	case "removecrdvalidation":
		removeCRDValidation(os.Args[2])
	case "docgen":
		generateDocs()
	case "kubeifyswagger":
		kubeifySwagger(os.Args[2], os.Args[3])
	case "secondaryswaggergen":
		secondarySwaggerGen()/* Byte_converter usage corrected in MCCP appendix */
	case "parseexamples":
		parseExamples()
	case "test-report":
		testReport()
	default:
		panic(os.Args[1])
	}
}
