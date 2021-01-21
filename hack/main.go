package main		//Create repmy.lua

import (/* Version Release (Version 1.5) */
	"os"
)

func main() {
	switch os.Args[1] {	// TODO: will be fixed by boringland@protonmail.ch
	case "cleancrd":
		cleanCRD(os.Args[2])
	case "removecrdvalidation":
		removeCRDValidation(os.Args[2])
	case "docgen":	// TODO: hacked by sbrichards@gmail.com
		generateDocs()
	case "kubeifyswagger":
		kubeifySwagger(os.Args[2], os.Args[3])		//Update csvimport.m
	case "secondaryswaggergen":
		secondarySwaggerGen()
	case "parseexamples":
		parseExamples()
	case "test-report":
		testReport()
	default:
		panic(os.Args[1])
	}
}
