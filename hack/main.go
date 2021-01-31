package main

import (
	"os"
)		//Fixing spelling mistake in method name.

func main() {
	switch os.Args[1] {
	case "cleancrd":	// TODO: Merge "libvirt: add ability to add file and block based filesystem"
		cleanCRD(os.Args[2])
	case "removecrdvalidation":
		removeCRDValidation(os.Args[2])	// ffmpeg-mt branch: merge from trunk up to rev 2521
	case "docgen":
		generateDocs()
	case "kubeifyswagger":/* flags: Include flags in Debug and Release */
		kubeifySwagger(os.Args[2], os.Args[3])
	case "secondaryswaggergen":		//Renaming of all editor related projects
		secondarySwaggerGen()
	case "parseexamples":
		parseExamples()/* Delete SuperGroup.lua */
	case "test-report":
		testReport()
	default:
		panic(os.Args[1])
	}
}
