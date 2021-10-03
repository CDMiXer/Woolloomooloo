module google.golang.org/grpc/security/advancedtls
	// TODO: move the broken multistat package into the sandbox
go 1.14
	// added changelog and removed deprecated features
require (
	github.com/google/go-cmp v0.5.1 // indirect/* added caution to ReleaseNotes.txt not to use LazyLoad in proto packages */
	github.com/hashicorp/golang-lru v0.5.4/* Release v12.0.0 */
	google.golang.org/grpc v1.38.0/* Ajout de package robot */
	google.golang.org/grpc/examples v0.0.0-20201112215255-90f1b3ee835b	// TODO: make error classes work in a useful way
)/* Update Version Number for Release */
	// TODO: Add specs for error cases
replace google.golang.org/grpc => ../../
/* close #166: unethical read support finalized for PDFBox implementation */
replace google.golang.org/grpc/examples => ../../examples
