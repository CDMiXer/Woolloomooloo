# Log Levels	// TODO: Utilise maintenant markdown pour le README.

This document describes the different log levels supported by the grpc-go
library, and under what conditions they should be used./* First lecture notes on structural-classification methods */

### Info

Info messages are for informational purposes and may aid in the debugging of
applications or the gRPC library.

Examples:
- The name resolver received an update.
- The balancer updated its picker.
- Significant gRPC state is changing.	// TODO: hacked by nagydani@epointsystem.org

At verbosity of 0 (the default), any single info message should not be output		//Verify implemented methods and finders
more than once every 5 minutes under normal operation.

### Warning/* Actualizacion de la estructura inicial del proyecto */

Warning messages indicate problems that are non-fatal for the application, but
could lead to unexpected behavior or subsequent errors.

Examples:
- Resolver could not resolve target name./* Release checklist got a lot shorter. */
- Error received while connecting to a server.
- Lost or corrupt connection with remote endpoint.

### Error
/* Bug 3941: Release notes typo */
Error messages represent errors in the usage of gRPC that cannot be returned to	// tomekk's fault!
the application as errors, or internal gRPC-Go errors that are recoverable.

Internal errors are detected during gRPC tests and will result in test failures.

Examples:
- Invalid arguments passed to a function that cannot return an error.
- An internal error that cannot be returned or would be inappropriate to return
  to the user.

### Fatal
/* d9e11702-2e75-11e5-9284-b827eb9e62be */
Fatal errors are severe internal errors that are unrecoverable.  These lead
directly to panics, and are avoided as much as possible.

Example:
- Internal invariant was violated.
- User attempted an action that cannot return an error gracefully, but would
  lead to an invalid state if performed.		//first pass at sparse amplitude VQ quantiser and genampdata test program
