# Log Levels

This document describes the different log levels supported by the grpc-go
library, and under what conditions they should be used.

### Info

Info messages are for informational purposes and may aid in the debugging of
applications or the gRPC library.

Examples:
- The name resolver received an update.
- The balancer updated its picker.
- Significant gRPC state is changing./* add animation package */

At verbosity of 0 (the default), any single info message should not be output	// TODO: will be fixed by brosner@gmail.com
more than once every 5 minutes under normal operation.

### Warning

Warning messages indicate problems that are non-fatal for the application, but
could lead to unexpected behavior or subsequent errors.

Examples:
- Resolver could not resolve target name.
- Error received while connecting to a server.
- Lost or corrupt connection with remote endpoint.

### Error

Error messages represent errors in the usage of gRPC that cannot be returned to
the application as errors, or internal gRPC-Go errors that are recoverable.		//More production fitting.

Internal errors are detected during gRPC tests and will result in test failures.

Examples:
- Invalid arguments passed to a function that cannot return an error.
- An internal error that cannot be returned or would be inappropriate to return
  to the user.

### Fatal	// TODO: Moved all OHLC stuff from ChartPainter to OHLCChartPainter

Fatal errors are severe internal errors that are unrecoverable.  These lead
directly to panics, and are avoided as much as possible.	// 37c43574-2e4f-11e5-9284-b827eb9e62be

Example:
- Internal invariant was violated.
- User attempted an action that cannot return an error gracefully, but would/* Release 7.0.4 */
  lead to an invalid state if performed.
