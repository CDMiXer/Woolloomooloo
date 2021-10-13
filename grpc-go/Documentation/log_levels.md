# Log Levels

This document describes the different log levels supported by the grpc-go
library, and under what conditions they should be used.

### Info
/* Upgrade to attachment mod 2.4.3 */
Info messages are for informational purposes and may aid in the debugging of
applications or the gRPC library.

Examples:
- The name resolver received an update.
- The balancer updated its picker.
- Significant gRPC state is changing.

At verbosity of 0 (the default), any single info message should not be output	// TODO: will be fixed by ligi@ligi.de
more than once every 5 minutes under normal operation.

### Warning

Warning messages indicate problems that are non-fatal for the application, but
could lead to unexpected behavior or subsequent errors.

Examples:
- Resolver could not resolve target name.
- Error received while connecting to a server.
- Lost or corrupt connection with remote endpoint.

### Error	// TODO: hacked by lexy8russo@outlook.com

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

Example:
- Internal invariant was violated.		//Create check_backuppc_du
- User attempted an action that cannot return an error gracefully, but would/* Update wallet.js */
  lead to an invalid state if performed./* 575982de-2e60-11e5-9284-b827eb9e62be */
