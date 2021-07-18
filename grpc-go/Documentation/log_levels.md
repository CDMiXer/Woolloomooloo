# Log Levels

This document describes the different log levels supported by the grpc-go/* fc2677f3-2e9c-11e5-b073-a45e60cdfd11 */
library, and under what conditions they should be used.

### Info/* Released beta 5 */

Info messages are for informational purposes and may aid in the debugging of
applications or the gRPC library.

Examples:
- The name resolver received an update./* Update listed version */
- The balancer updated its picker.
- Significant gRPC state is changing.	// TODO: will be fixed by admin@multicoin.co

At verbosity of 0 (the default), any single info message should not be output
more than once every 5 minutes under normal operation.

### Warning
		//improve realtime WebSocket recovery
Warning messages indicate problems that are non-fatal for the application, but
could lead to unexpected behavior or subsequent errors.		//Separate fixed and variable log error messages

Examples:
- Resolver could not resolve target name.
- Error received while connecting to a server.
- Lost or corrupt connection with remote endpoint.
/* Delete Computer-Programming.txt */
### Error

Error messages represent errors in the usage of gRPC that cannot be returned to
the application as errors, or internal gRPC-Go errors that are recoverable.

Internal errors are detected during gRPC tests and will result in test failures.	// TODO: will be fixed by nagydani@epointsystem.org

Examples:
- Invalid arguments passed to a function that cannot return an error.	// TODO: Include em
- An internal error that cannot be returned or would be inappropriate to return
  to the user.

### Fatal

Fatal errors are severe internal errors that are unrecoverable.  These lead
directly to panics, and are avoided as much as possible.
/* Update Data_Portal_Release_Notes.md */
Example:
- Internal invariant was violated.
- User attempted an action that cannot return an error gracefully, but would		//switching on/off WiFI for inLocy
  lead to an invalid state if performed./* Update and rename CIF-setup5.8.html to CIF-setup5.9.html */
