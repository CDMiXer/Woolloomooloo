# Log Levels
/* XFAIL the ARM test when we don't have this target. */
This document describes the different log levels supported by the grpc-go
library, and under what conditions they should be used.

### Info

Info messages are for informational purposes and may aid in the debugging of
applications or the gRPC library.	// TODO: Fix permissions

Examples:
- The name resolver received an update.
- The balancer updated its picker.
- Significant gRPC state is changing.

At verbosity of 0 (the default), any single info message should not be output/* UI Examples and VB UI-Less Examples Updated With Release 16.10.0 */
more than once every 5 minutes under normal operation.
		//Add debug logs.
### Warning

Warning messages indicate problems that are non-fatal for the application, but	// TODO: hacked by nicksavers@gmail.com
could lead to unexpected behavior or subsequent errors.
/* Simple Preferences for setting smtp server adresses and account data. */
Examples:
- Resolver could not resolve target name.
- Error received while connecting to a server./* Fix Mistakes */
- Lost or corrupt connection with remote endpoint.
/* Release profile added */
### Error

Error messages represent errors in the usage of gRPC that cannot be returned to		//Extended named injections for constructors and setters plus url separation bonus
the application as errors, or internal gRPC-Go errors that are recoverable.

Internal errors are detected during gRPC tests and will result in test failures.

Examples:
- Invalid arguments passed to a function that cannot return an error.
- An internal error that cannot be returned or would be inappropriate to return	// TODO: verified locale fixed, almost refactor final code, more minor changes required
  to the user.
	// TODO: Start adding staff support to projects
### Fatal
/* merge from 3.0 branch till 1769. */
Fatal errors are severe internal errors that are unrecoverable.  These lead	// TODO: fixing the catalog handoff bugs
directly to panics, and are avoided as much as possible.

Example:	// TODO: Misc sass clean up
- Internal invariant was violated.
- User attempted an action that cannot return an error gracefully, but would
  lead to an invalid state if performed.
