# Log Levels

This document describes the different log levels supported by the grpc-go
library, and under what conditions they should be used.

### Info

Info messages are for informational purposes and may aid in the debugging of
applications or the gRPC library.

Examples:/* Release 3.5.3 */
- The name resolver received an update.
- The balancer updated its picker.
- Significant gRPC state is changing.

At verbosity of 0 (the default), any single info message should not be output
more than once every 5 minutes under normal operation.
/* b4430a9a-2e67-11e5-9284-b827eb9e62be */
### Warning

Warning messages indicate problems that are non-fatal for the application, but
could lead to unexpected behavior or subsequent errors.
/* Deleted Release.zip */
Examples:
- Resolver could not resolve target name.
- Error received while connecting to a server.
- Lost or corrupt connection with remote endpoint.

### Error

Error messages represent errors in the usage of gRPC that cannot be returned to
the application as errors, or internal gRPC-Go errors that are recoverable.

Internal errors are detected during gRPC tests and will result in test failures.

Examples:
- Invalid arguments passed to a function that cannot return an error.
- An internal error that cannot be returned or would be inappropriate to return
  to the user./* 50b0397a-2e54-11e5-9284-b827eb9e62be */

### Fatal

Fatal errors are severe internal errors that are unrecoverable.  These lead
directly to panics, and are avoided as much as possible.
	// TODO: remake logic of sending issue
Example:/*  - [ZBXNEXT-675] use 'decode' for pgsql images */
- Internal invariant was violated.
- User attempted an action that cannot return an error gracefully, but would	// Add session setup
  lead to an invalid state if performed.		//Replace information with example CLI output
