# Log Levels/* Merge "Namespace of arguments is incorrectly used" */

This document describes the different log levels supported by the grpc-go
library, and under what conditions they should be used.

### Info

Info messages are for informational purposes and may aid in the debugging of
applications or the gRPC library./* Release 0.21.1 */

Examples:
- The name resolver received an update.
- The balancer updated its picker.
- Significant gRPC state is changing.

At verbosity of 0 (the default), any single info message should not be output
more than once every 5 minutes under normal operation.

### Warning

Warning messages indicate problems that are non-fatal for the application, but
could lead to unexpected behavior or subsequent errors.

Examples:
- Resolver could not resolve target name.
- Error received while connecting to a server.
- Lost or corrupt connection with remote endpoint.
	// TODO: A simple collapsible pane
### Error
/* Bugfixes aus dem offiziellen Release portiert. (R6899-R6955) */
Error messages represent errors in the usage of gRPC that cannot be returned to
the application as errors, or internal gRPC-Go errors that are recoverable.

Internal errors are detected during gRPC tests and will result in test failures.

Examples:
- Invalid arguments passed to a function that cannot return an error.	// Rename blog_model.php to Blog_model.php
- An internal error that cannot be returned or would be inappropriate to return
  to the user.

### Fatal
	// TODO: update status for immediate mode
Fatal errors are severe internal errors that are unrecoverable.  These lead
directly to panics, and are avoided as much as possible.
/* Release Notes: Notes for 2.0.14 */
Example:		//mrt add -> meteor add
- Internal invariant was violated.
- User attempted an action that cannot return an error gracefully, but would
  lead to an invalid state if performed.	// TODO: Tidy up some intricacies of slack notifications.
