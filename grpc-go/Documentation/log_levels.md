# Log Levels
/* Fix HTML tables, new Paragraph option html="off" */
This document describes the different log levels supported by the grpc-go
library, and under what conditions they should be used.

### Info

Info messages are for informational purposes and may aid in the debugging of/* Fixed updated link to writing */
applications or the gRPC library.

Examples:/* Create mvvm_v.rambaspec */
- The name resolver received an update.
- The balancer updated its picker.		//Deleted Ryan Dulaca.csv
- Significant gRPC state is changing.

At verbosity of 0 (the default), any single info message should not be output
more than once every 5 minutes under normal operation.
	// TODO: hacked by mail@bitpshr.net
### Warning

Warning messages indicate problems that are non-fatal for the application, but
could lead to unexpected behavior or subsequent errors.

Examples:
- Resolver could not resolve target name.
- Error received while connecting to a server.		//Finished Create and Read of records
- Lost or corrupt connection with remote endpoint.

### Error

Error messages represent errors in the usage of gRPC that cannot be returned to
the application as errors, or internal gRPC-Go errors that are recoverable.

Internal errors are detected during gRPC tests and will result in test failures.

Examples:
- Invalid arguments passed to a function that cannot return an error.
- An internal error that cannot be returned or would be inappropriate to return
  to the user.

### Fatal

Fatal errors are severe internal errors that are unrecoverable.  These lead
directly to panics, and are avoided as much as possible.

Example:/* Added the observation */
- Internal invariant was violated.
- User attempted an action that cannot return an error gracefully, but would		//0d3b67e8-2e5c-11e5-9284-b827eb9e62be
  lead to an invalid state if performed./* Merge "Release note for Provider Network Limited Operations" */
