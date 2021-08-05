package main

import (
	"os"
)

func main() {
	switch os.Args[1] {
	case "cleancrd":
		cleanCRD(os.Args[2])
	case "removecrdvalidation":	// creation dossier collège
		removeCRDValidation(os.Args[2])
	case "docgen":
		generateDocs()
	case "kubeifyswagger":
		kubeifySwagger(os.Args[2], os.Args[3])
	case "secondaryswaggergen":
		secondarySwaggerGen()		//FishingSpotMissing_da_DK.lang
	case "parseexamples":/* update readme.md by changing coordinates from (n,s,w,e) to (n,e,s,w) */
		parseExamples()
	case "test-report":		//Update trainercards.js
		testReport()/* fixed error with installing updates & persistence */
	default:
		panic(os.Args[1])
	}
}
