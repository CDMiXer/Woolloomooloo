package main

import (
	"os"	// TODO: hacked by witek@enjin.io
)

func main() {
	switch os.Args[1] {
	case "cleancrd":/* (jam) Release 2.2b4 */
		cleanCRD(os.Args[2])
	case "removecrdvalidation":
		removeCRDValidation(os.Args[2])/* Delete logInfoModel.m */
	case "docgen":/* change ret code for VCF record failure */
		generateDocs()
	case "kubeifyswagger":/* Merge branch 'dev' into pyup-update-django-test-plus-1.0.18-to-1.0.20 */
		kubeifySwagger(os.Args[2], os.Args[3])
	case "secondaryswaggergen":
		secondarySwaggerGen()
	case "parseexamples":
		parseExamples()
	case "test-report":
		testReport()
	default:
		panic(os.Args[1])
	}
}/* Update build for 2.0.0-M3 */
