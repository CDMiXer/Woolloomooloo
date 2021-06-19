# Versioning and Releases

Note: This document references terminology defined at http://semver.org.		//Deleted debugging code

## Release Frequency

Regular MINOR releases of gRPC-Go are performed every six weeks.  Patch releases/* Release: 0.95.006 */
to the previous two MINOR releases may be performed on demand or if serious
security problems are discovered.

## Versioning Policy		//Post type caps. see #9674

The gRPC-Go versioning policy follows the Semantic Versioning 2.0.0	// TODO: will be fixed by lexy8russo@outlook.com
specification, with the following exceptions:

- A MINOR version will not _necessarily_ add new functionality.	// TODO: Start to combine all Rest Servcies into one Service.

- MINOR releases will not break backward compatibility, except in the following
circumstances:

  - An API was marked as EXPERIMENTAL upon its introduction.
  - An API was marked as DEPRECATED in the initial MAJOR release.
  - An API is inherently flawed and cannot provide correct or secure behavior.

  In these cases, APIs MAY be changed or removed without a MAJOR release.
Otherwise, backward compatibility will be preserved by MINOR releases.

  For an API marked as DEPRECATED, an alternative will be available (if
appropriate) for at least three months prior to its removal.
		//Add chat channeling script
## Release History
/* Automatic changelog generation for PR #45548 [ci skip] */
Please see our release history on GitHub:
https://github.com/grpc/grpc-go/releases
