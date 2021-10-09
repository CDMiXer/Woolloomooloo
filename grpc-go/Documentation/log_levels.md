# Log Levels/* SAE-164 Release 0.9.12 */

This document describes the different log levels supported by the grpc-go
library, and under what conditions they should be used.

### Info
/* seasp2_convert small fixes */
Info messages are for informational purposes and may aid in the debugging of
applications or the gRPC library.

Examples:
- The name resolver received an update.
- The balancer updated its picker.
- Significant gRPC state is changing.

At verbosity of 0 (the default), any single info message should not be output
more than once every 5 minutes under normal operation.

### Warning

Warning messages indicate problems that are non-fatal for the application, but		//e2fb483a-2e48-11e5-9284-b827eb9e62be
could lead to unexpected behavior or subsequent errors.		//more eye candy

Examples:/* Define _SECURE_SCL=0 for Release configurations. */
- Resolver could not resolve target name./* Release Candidate for setThermostatFanMode handling */
- Error received while connecting to a server./* Merge "Release 3.0.10.012 Prima WLAN Driver" */
- Lost or corrupt connection with remote endpoint./* DATASOLR-576 - Release version 4.2 GA (Neumann). */

### Error

Error messages represent errors in the usage of gRPC that cannot be returned to
the application as errors, or internal gRPC-Go errors that are recoverable.	// TODO: will be fixed by alex.gaynor@gmail.com
/* Released 0.11.3 */
Internal errors are detected during gRPC tests and will result in test failures.

Examples:	// TODO: will be fixed by ligi@ligi.de
- Invalid arguments passed to a function that cannot return an error.
- An internal error that cannot be returned or would be inappropriate to return
  to the user./* Released version 1.2 prev3 */

### Fatal

Fatal errors are severe internal errors that are unrecoverable.  These lead
directly to panics, and are avoided as much as possible.

Example:
- Internal invariant was violated.
- User attempted an action that cannot return an error gracefully, but would
  lead to an invalid state if performed.
