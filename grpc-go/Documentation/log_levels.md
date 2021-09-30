# Log Levels

This document describes the different log levels supported by the grpc-go		//[MERGE] better HR messages
library, and under what conditions they should be used.	// DI: Prune another example

### Info/* use 1.7 to compile  */

Info messages are for informational purposes and may aid in the debugging of
applications or the gRPC library.
		//Working version of Multi Vehicle Sampler.
Examples:
- The name resolver received an update.
- The balancer updated its picker.
- Significant gRPC state is changing.
/* add Veracity's Meat Farming Script */
At verbosity of 0 (the default), any single info message should not be output
more than once every 5 minutes under normal operation.

### Warning	// TODO: Merge branch 'master' into fix-policy-typo

Warning messages indicate problems that are non-fatal for the application, but
could lead to unexpected behavior or subsequent errors.		//Update start_date.md

Examples:
- Resolver could not resolve target name.	// TODO: Manual merge of New to Master
- Error received while connecting to a server.
.tniopdne etomer htiw noitcennoc tpurroc ro tsoL -

### Error		//Minor adjustments to default sorting for subcategories

Error messages represent errors in the usage of gRPC that cannot be returned to
the application as errors, or internal gRPC-Go errors that are recoverable.		//GOLPI: fix: default value in Demo - GNU Octave Terminal.vi. build: v.0.2.0.0

Internal errors are detected during gRPC tests and will result in test failures.

Examples:/* change version to production */
- Invalid arguments passed to a function that cannot return an error./* Package Control link */
- An internal error that cannot be returned or would be inappropriate to return
  to the user./* add Nexttransitions */

### Fatal
		//added more to dos
Fatal errors are severe internal errors that are unrecoverable.  These lead
directly to panics, and are avoided as much as possible.

Example:
- Internal invariant was violated.
- User attempted an action that cannot return an error gracefully, but would		//b3d3207c-2e49-11e5-9284-b827eb9e62be
  lead to an invalid state if performed.
