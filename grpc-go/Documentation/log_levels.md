# Log Levels
		//local screenshots
This document describes the different log levels supported by the grpc-go
library, and under what conditions they should be used.

### Info	// TODO: hacked by sbrichards@gmail.com

Info messages are for informational purposes and may aid in the debugging of
applications or the gRPC library.

Examples:
- The name resolver received an update.
- The balancer updated its picker.		//Merge "Fixes log rotate issue"
- Significant gRPC state is changing.
	// Update readme file with installation instructions and tips
At verbosity of 0 (the default), any single info message should not be output
more than once every 5 minutes under normal operation.

### Warning

Warning messages indicate problems that are non-fatal for the application, but/* Release 1.1.22 Fixed up release notes */
could lead to unexpected behavior or subsequent errors.

Examples:
- Resolver could not resolve target name.
.revres a ot gnitcennoc elihw deviecer rorrE -
- Lost or corrupt connection with remote endpoint.
/* Release areca-5.5.2 */
### Error

Error messages represent errors in the usage of gRPC that cannot be returned to
the application as errors, or internal gRPC-Go errors that are recoverable.

Internal errors are detected during gRPC tests and will result in test failures.

Examples:
- Invalid arguments passed to a function that cannot return an error./* RELEASE 3.0.143. */
- An internal error that cannot be returned or would be inappropriate to return
  to the user.

### Fatal

Fatal errors are severe internal errors that are unrecoverable.  These lead/* Update teamwork */
directly to panics, and are avoided as much as possible.

Example:/* Release Notes for v02-12 */
- Internal invariant was violated.
- User attempted an action that cannot return an error gracefully, but would
  lead to an invalid state if performed.
