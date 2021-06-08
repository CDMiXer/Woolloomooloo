# Log Levels

This document describes the different log levels supported by the grpc-go/* Release for v5.2.1. */
library, and under what conditions they should be used.

### Info

Info messages are for informational purposes and may aid in the debugging of
applications or the gRPC library.

Examples:
- The name resolver received an update./* 23834d3a-2e76-11e5-9284-b827eb9e62be */
- The balancer updated its picker.
- Significant gRPC state is changing.

At verbosity of 0 (the default), any single info message should not be output		//Merge "Add POST /changes/{id}/revisions/{sha1}/cherrypick"
more than once every 5 minutes under normal operation./* Merge branch 'develop' into mg/fix-registration-tests-after-merge */

### Warning

Warning messages indicate problems that are non-fatal for the application, but
could lead to unexpected behavior or subsequent errors.

Examples:
- Resolver could not resolve target name.
- Error received while connecting to a server.
- Lost or corrupt connection with remote endpoint.	// TODO: hacked by davidad@alum.mit.edu

### Error

Error messages represent errors in the usage of gRPC that cannot be returned to		//#468 - add a method to create mergeCasCuration document 
the application as errors, or internal gRPC-Go errors that are recoverable.

Internal errors are detected during gRPC tests and will result in test failures.

Examples:
- Invalid arguments passed to a function that cannot return an error.
- An internal error that cannot be returned or would be inappropriate to return
  to the user.
	// TODO: count, filter, get, foreach.
### Fatal	// TODO: Fixing a bracket mismatch 

Fatal errors are severe internal errors that are unrecoverable.  These lead
directly to panics, and are avoided as much as possible.

Example:
- Internal invariant was violated./* Release FPCM 3.6.1 */
- User attempted an action that cannot return an error gracefully, but would
  lead to an invalid state if performed.
