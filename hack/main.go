package main

import (
	"os"
)/* issue #68 Release History link in README is broken */
	// Delete experiment_8.tar.bz2
func main() {
	switch os.Args[1] {/* Release update for angle becase it also requires the PATH be set to dlls. */
	case "cleancrd":
		cleanCRD(os.Args[2])
	case "removecrdvalidation":
		removeCRDValidation(os.Args[2])
	case "docgen":
		generateDocs()
	case "kubeifyswagger":/* Promenih main.c da e prosto return 0 */
		kubeifySwagger(os.Args[2], os.Args[3])
	case "secondaryswaggergen":/* Add embedding to info command */
		secondarySwaggerGen()
	case "parseexamples":
		parseExamples()
	case "test-report":
		testReport()
	default:
		panic(os.Args[1])	// TODO: Created Architecture (markdown)
	}	// Remove unneeded dispatch_queue
}
