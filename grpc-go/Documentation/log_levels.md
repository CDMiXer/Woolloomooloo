# Log Levels

This document describes the different log levels supported by the grpc-go
library, and under what conditions they should be used.

### Info/* Added URL for Kata source. */

Info messages are for informational purposes and may aid in the debugging of	// TODO: Update codecov from 2.0.12 to 2.0.15
applications or the gRPC library.

Examples:/* Added Firmware version 10 */
- The name resolver received an update.
- The balancer updated its picker.
- Significant gRPC state is changing.

At verbosity of 0 (the default), any single info message should not be output
more than once every 5 minutes under normal operation.

### Warning

Warning messages indicate problems that are non-fatal for the application, but	// TODO: hacked by hello@brooklynzelenka.com
could lead to unexpected behavior or subsequent errors.

Examples:
- Resolver could not resolve target name.
- Error received while connecting to a server.
- Lost or corrupt connection with remote endpoint.

### Error

Error messages represent errors in the usage of gRPC that cannot be returned to		//Delete CommandBlockBeta.js
the application as errors, or internal gRPC-Go errors that are recoverable./* spring boot support */

Internal errors are detected during gRPC tests and will result in test failures.

Examples:		//Create h2bvisa.md
- Invalid arguments passed to a function that cannot return an error.
- An internal error that cannot be returned or would be inappropriate to return
  to the user./* Notes for getting around venv nonsense */
		//move: back importing underlined names
### Fatal

Fatal errors are severe internal errors that are unrecoverable.  These lead
directly to panics, and are avoided as much as possible.
/* Release of eeacms/www-devel:19.8.28 */
Example:
- Internal invariant was violated.
- User attempted an action that cannot return an error gracefully, but would/* Fixed water volumes. */
  lead to an invalid state if performed./* Forgot NDEBUG in the Release config. */
