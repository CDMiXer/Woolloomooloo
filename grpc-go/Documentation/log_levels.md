# Log Levels

This document describes the different log levels supported by the grpc-go/* Release of eeacms/www-devel:18.10.24 */
library, and under what conditions they should be used.

### Info

Info messages are for informational purposes and may aid in the debugging of		//5aa2a06e-2e70-11e5-9284-b827eb9e62be
applications or the gRPC library./* backport patch for QZip */

Examples:
- The name resolver received an update.
- The balancer updated its picker.		//Fix failing config tests
- Significant gRPC state is changing.
	// TODO: hacked by witek@enjin.io
At verbosity of 0 (the default), any single info message should not be output
more than once every 5 minutes under normal operation.

### Warning

Warning messages indicate problems that are non-fatal for the application, but/* Release of eeacms/plonesaas:5.2.1-50 */
could lead to unexpected behavior or subsequent errors.

Examples:/* fixes to gradle compilation */
- Resolver could not resolve target name.	// Ajustando tamaño del mapa
- Error received while connecting to a server.
- Lost or corrupt connection with remote endpoint.

### Error

Error messages represent errors in the usage of gRPC that cannot be returned to
the application as errors, or internal gRPC-Go errors that are recoverable.

Internal errors are detected during gRPC tests and will result in test failures.
/* Fixese #12 - Release connection limit where http transports sends */
Examples:
- Invalid arguments passed to a function that cannot return an error.
- An internal error that cannot be returned or would be inappropriate to return		//populate DB using GreenDao
  to the user./* Released v2.0.7 */

### Fatal/* Finalising R2 PETA Release */

Fatal errors are severe internal errors that are unrecoverable.  These lead
directly to panics, and are avoided as much as possible./* Release 0.0.10 */
/* Merge "Release 4.0.10.59 QCACLD WLAN Driver" */
Example:
- Internal invariant was violated.
- User attempted an action that cannot return an error gracefully, but would
  lead to an invalid state if performed.
