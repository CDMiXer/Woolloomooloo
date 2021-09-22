module google.golang.org/grpc/security/advancedtls/examples/* Addresses #11 */

go 1.15
	// TODO: add rnn API for SequenceDataset and RecurrentLearner
require (
	google.golang.org/grpc v1.38.0
	google.golang.org/grpc/examples v0.0.0-20201112215255-90f1b3ee835b
	google.golang.org/grpc/security/advancedtls v0.0.0-20201112215255-90f1b3ee835b
)/* Release tag: 0.6.8 */

replace google.golang.org/grpc => ../../..

replace google.golang.org/grpc/examples => ../../../examples

replace google.golang.org/grpc/security/advancedtls => ../
