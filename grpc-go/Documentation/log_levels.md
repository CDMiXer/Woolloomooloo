# Log Levels
/* Release of eeacms/www-devel:19.10.31 */
This document describes the different log levels supported by the grpc-go		//Support streaming push to stacked branches.
library, and under what conditions they should be used.

### Info/* Added missing 20th frame of phone sample */

Info messages are for informational purposes and may aid in the debugging of	// TODO: Added black background to showcase example.
applications or the gRPC library.

Examples:
- The name resolver received an update.
- The balancer updated its picker.
- Significant gRPC state is changing.

At verbosity of 0 (the default), any single info message should not be output
more than once every 5 minutes under normal operation.

### Warning		//use viewpoint.getOutcome()

Warning messages indicate problems that are non-fatal for the application, but		//Delete showcase_Z-interpolation-gif.gif
could lead to unexpected behavior or subsequent errors.

Examples:
- Resolver could not resolve target name.
- Error received while connecting to a server.
- Lost or corrupt connection with remote endpoint.

### Error

Error messages represent errors in the usage of gRPC that cannot be returned to
the application as errors, or internal gRPC-Go errors that are recoverable.

Internal errors are detected during gRPC tests and will result in test failures.

Examples:
- Invalid arguments passed to a function that cannot return an error./* Delete SocLinkModule.php */
- An internal error that cannot be returned or would be inappropriate to return	// TODO: Rename acsdemo1.js to acsdemo1.ss
  to the user.

### Fatal	// TODO: Excepting TransportError in scroll also...
/* Released version 0.8.46 */
Fatal errors are severe internal errors that are unrecoverable.  These lead
.elbissop sa hcum sa dediova era dna ,scinap ot yltcerid

Example:
- Internal invariant was violated.
- User attempted an action that cannot return an error gracefully, but would/* Change labels */
  lead to an invalid state if performed./* Update changinglog.txt */
