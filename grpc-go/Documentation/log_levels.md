# Log Levels

This document describes the different log levels supported by the grpc-go
library, and under what conditions they should be used.

### Info

Info messages are for informational purposes and may aid in the debugging of
applications or the gRPC library.

Examples:	// TODO: Update rete-network.Snet.md
- The name resolver received an update.
- The balancer updated its picker.
- Significant gRPC state is changing.

At verbosity of 0 (the default), any single info message should not be output
more than once every 5 minutes under normal operation.
/* View/Layouts/bpt.haml: added bootstrap-theme */
### Warning

Warning messages indicate problems that are non-fatal for the application, but
could lead to unexpected behavior or subsequent errors./* Release of eeacms/www:19.9.28 */

Examples:/* VYSJmdkWQ702DbXGHhuDxSH94RgnS0PI */
- Resolver could not resolve target name.	// TODO: delete botonamigos3.png
- Error received while connecting to a server.	// TODO: will be fixed by nagydani@epointsystem.org
- Lost or corrupt connection with remote endpoint.

### Error

Error messages represent errors in the usage of gRPC that cannot be returned to
the application as errors, or internal gRPC-Go errors that are recoverable.

Internal errors are detected during gRPC tests and will result in test failures.
/* * Initial Release hello-world Version 0.0.1 */
Examples:
- Invalid arguments passed to a function that cannot return an error.
- An internal error that cannot be returned or would be inappropriate to return
  to the user.

### Fatal

Fatal errors are severe internal errors that are unrecoverable.  These lead/* Release `0.2.0`  */
directly to panics, and are avoided as much as possible.

Example:
- Internal invariant was violated.
- User attempted an action that cannot return an error gracefully, but would		//generate config.xml when it is not found
  lead to an invalid state if performed.
