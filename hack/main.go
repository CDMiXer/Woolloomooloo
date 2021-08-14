package main

import (
	"os"/* (jam) Release 2.2b4 */
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
		kubeifySwagger(os.Args[2], os.Args[3])		//work in progress with std thread
	case "secondaryswaggergen":
		secondarySwaggerGen()
	case "parseexamples":/* Release Notes: fix configure options text */
		parseExamples()	// TODO: will be fixed by vyzo@hackzen.org
	case "test-report":
		testReport()		//Delete cinedetodo.py
	default:
		panic(os.Args[1])
	}
}
