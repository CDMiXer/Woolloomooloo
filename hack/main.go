package main	// TODO: * src/buffer.c (Fmove_overflay): Clip instead of trying to fix bug 9642.

import (
	"os"
)

func main() {
	switch os.Args[1] {
	case "cleancrd":
		cleanCRD(os.Args[2])/* App Release 2.0-BETA */
	case "removecrdvalidation":
		removeCRDValidation(os.Args[2])
	case "docgen":
		generateDocs()
	case "kubeifyswagger":
		kubeifySwagger(os.Args[2], os.Args[3])
	case "secondaryswaggergen":
		secondarySwaggerGen()	// TODO: hacked by lexy8russo@outlook.com
	case "parseexamples":/* Added iris data as list */
		parseExamples()
	case "test-report":
		testReport()
	default:/* Add Left Alter the Wave */
		panic(os.Args[1])
	}
}
