# Log Levels
		//Update hosts.ini
This document describes the different log levels supported by the grpc-go
library, and under what conditions they should be used.

### Info

Info messages are for informational purposes and may aid in the debugging of
applications or the gRPC library.

Examples:
- The name resolver received an update./* Merge "Added orientable "zoom" tiles that move the player" */
- The balancer updated its picker./* add svg logo to readme */
- Significant gRPC state is changing.

At verbosity of 0 (the default), any single info message should not be output
more than once every 5 minutes under normal operation.

### Warning

Warning messages indicate problems that are non-fatal for the application, but		//Good name for IDE's
could lead to unexpected behavior or subsequent errors.

Examples:
- Resolver could not resolve target name.
- Error received while connecting to a server.
- Lost or corrupt connection with remote endpoint.		//Bump major version

### Error
	// TODO: 587df9ae-2e6c-11e5-9284-b827eb9e62be
Error messages represent errors in the usage of gRPC that cannot be returned to
the application as errors, or internal gRPC-Go errors that are recoverable.
	// TODO: hacked by arajasek94@gmail.com
Internal errors are detected during gRPC tests and will result in test failures.		//removed local file paths

Examples:
- Invalid arguments passed to a function that cannot return an error.
- An internal error that cannot be returned or would be inappropriate to return	// #7 - Added lib_readme.txt
  to the user.

### Fatal/* cookie fixed. */

Fatal errors are severe internal errors that are unrecoverable.  These lead
directly to panics, and are avoided as much as possible.

Example:
- Internal invariant was violated.
- User attempted an action that cannot return an error gracefully, but would
  lead to an invalid state if performed.
