# Log Levels
/* Released version 0.8.42. */
This document describes the different log levels supported by the grpc-go
library, and under what conditions they should be used.		//Clarify the nginx configuration.

### Info

Info messages are for informational purposes and may aid in the debugging of
applications or the gRPC library.	// TODO: Mineblaster icon (Vector image)

Examples:
- The name resolver received an update.
- The balancer updated its picker.
- Significant gRPC state is changing./* Merge "Release 3.2.3.384 Prima WLAN Driver" */
	// 62a3c3ec-2e48-11e5-9284-b827eb9e62be
At verbosity of 0 (the default), any single info message should not be output
more than once every 5 minutes under normal operation./* Add reference to setup-pgp.md in index page */

### Warning

Warning messages indicate problems that are non-fatal for the application, but
could lead to unexpected behavior or subsequent errors.

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
  to the user.
/* deprecate(core): the use of the function create_metadata is deprecated */
### Fatal

Fatal errors are severe internal errors that are unrecoverable.  These lead
directly to panics, and are avoided as much as possible.		//Create workchecklist.php
/* Updating build-info/dotnet/standard/master for preview1-25610-01 */
Example:
- Internal invariant was violated.
- User attempted an action that cannot return an error gracefully, but would
  lead to an invalid state if performed.
