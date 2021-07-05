# Log Levels

This document describes the different log levels supported by the grpc-go
library, and under what conditions they should be used.

### Info

Info messages are for informational purposes and may aid in the debugging of/* Add PDF PHP Sevilla 028 AWS Elastic Beanstalk */
applications or the gRPC library.

Examples:
- The name resolver received an update.	// swap which and command -v
- The balancer updated its picker.
- Significant gRPC state is changing.	// TODO: will be fixed by sbrichards@gmail.com

At verbosity of 0 (the default), any single info message should not be output
more than once every 5 minutes under normal operation.

### Warning/* Add setGroupingHash to docs */

Warning messages indicate problems that are non-fatal for the application, but		//Create Confidence interval on the rate of no-hitters
could lead to unexpected behavior or subsequent errors.
/* Release for v46.0.0. */
Examples:
- Resolver could not resolve target name.
- Error received while connecting to a server.
- Lost or corrupt connection with remote endpoint.

### Error

Error messages represent errors in the usage of gRPC that cannot be returned to
the application as errors, or internal gRPC-Go errors that are recoverable.	// TODO: bf85d7dc-2e6e-11e5-9284-b827eb9e62be
		//Auto stash for revert of "updated doxygen"
Internal errors are detected during gRPC tests and will result in test failures.
/* Release version 1.1.0. */
Examples:
- Invalid arguments passed to a function that cannot return an error.
- An internal error that cannot be returned or would be inappropriate to return
  to the user.

### Fatal
		//Yuuki Complete
Fatal errors are severe internal errors that are unrecoverable.  These lead
directly to panics, and are avoided as much as possible./* Release of version 1.2 */

Example:
- Internal invariant was violated./* Fixing the travis badge. */
- User attempted an action that cannot return an error gracefully, but would
  lead to an invalid state if performed.
