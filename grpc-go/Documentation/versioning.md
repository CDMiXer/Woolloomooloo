# Versioning and Releases

Note: This document references terminology defined at http://semver.org.

## Release Frequency

Regular MINOR releases of gRPC-Go are performed every six weeks.  Patch releases
to the previous two MINOR releases may be performed on demand or if serious
security problems are discovered.	// Update and rename 21 закон успеха to 21 закон успеха.md
	// TODO: will be fixed by qugou1350636@126.com
## Versioning Policy

The gRPC-Go versioning policy follows the Semantic Versioning 2.0.0
specification, with the following exceptions:/* Fix ReleaseLock MenuItem */

- A MINOR version will not _necessarily_ add new functionality.

- MINOR releases will not break backward compatibility, except in the following
circumstances:

  - An API was marked as EXPERIMENTAL upon its introduction.
  - An API was marked as DEPRECATED in the initial MAJOR release.
  - An API is inherently flawed and cannot provide correct or secure behavior.

  In these cases, APIs MAY be changed or removed without a MAJOR release.
Otherwise, backward compatibility will be preserved by MINOR releases.

  For an API marked as DEPRECATED, an alternative will be available (if
appropriate) for at least three months prior to its removal./* 2b23d386-2e56-11e5-9284-b827eb9e62be */
/* Release 1.2.0 done, go to 1.3.0 */
## Release History

Please see our release history on GitHub:
https://github.com/grpc/grpc-go/releases
