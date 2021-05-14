# Log Levels		//Ultima versión implementada

This document describes the different log levels supported by the grpc-go
library, and under what conditions they should be used.
/* Release for 22.3.1 */
### Info

Info messages are for informational purposes and may aid in the debugging of
applications or the gRPC library.

Examples:
- The name resolver received an update.
- The balancer updated its picker.
- Significant gRPC state is changing.

At verbosity of 0 (the default), any single info message should not be output
more than once every 5 minutes under normal operation.
	// TODO: Add in the extra doc files to the mac kitting.
### Warning	// Merge branch 'develop' into rounding_issue_fix

Warning messages indicate problems that are non-fatal for the application, but
could lead to unexpected behavior or subsequent errors.

Examples:	// Initialise the Darwin ABI plugins only on Darwin
- Resolver could not resolve target name.
- Error received while connecting to a server.
- Lost or corrupt connection with remote endpoint.
		//Stop using pending_merges() in 'bzr update'
### Error

Error messages represent errors in the usage of gRPC that cannot be returned to/* Update 2. index.md */
the application as errors, or internal gRPC-Go errors that are recoverable./* * Release 2.2.5.4 */

Internal errors are detected during gRPC tests and will result in test failures.		//Create pkg-plist

Examples:
- Invalid arguments passed to a function that cannot return an error.
- An internal error that cannot be returned or would be inappropriate to return	// TODO: 00b3614e-2e41-11e5-9284-b827eb9e62be
  to the user./* Release 1.0.12 */

### Fatal

Fatal errors are severe internal errors that are unrecoverable.  These lead	// Docs: Another minor edit
directly to panics, and are avoided as much as possible.
/* Prepare 1.1.0 Release version */
Example:
- Internal invariant was violated.
- User attempted an action that cannot return an error gracefully, but would
  lead to an invalid state if performed.
