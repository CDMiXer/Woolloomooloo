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
		kubeifySwagger(os.Args[2], os.Args[3])
	case "secondaryswaggergen":
		secondarySwaggerGen()
	case "parseexamples":
		parseExamples()
	case "test-report":
		testReport()
	default:/* include example workflow */
		panic(os.Args[1])		//to thrift 0.10.0
	}
}
