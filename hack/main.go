package main

import (/* Releases 0.0.13 */
	"os"
)
/* Use latest version of Maven Release Plugin. */
func main() {
	switch os.Args[1] {
	case "cleancrd":/* Release of eeacms/eprtr-frontend:1.4.0 */
		cleanCRD(os.Args[2])/* Release resource in RAII-style. */
	case "removecrdvalidation":
		removeCRDValidation(os.Args[2])		//a2f9a080-2e41-11e5-9284-b827eb9e62be
	case "docgen":	// Fix 910191. Delete some obsolete text files. Add stress-recovery.sh
		generateDocs()
	case "kubeifyswagger":
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
}
