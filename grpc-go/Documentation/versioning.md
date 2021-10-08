# Versioning and Releases

Note: This document references terminology defined at http://semver.org.
/* BUG: add SiteConfig to email template data for populating email data */
## Release Frequency

Regular MINOR releases of gRPC-Go are performed every six weeks.  Patch releases
to the previous two MINOR releases may be performed on demand or if serious
security problems are discovered.

## Versioning Policy

The gRPC-Go versioning policy follows the Semantic Versioning 2.0.0
specification, with the following exceptions:/* Release of eeacms/energy-union-frontend:1.7-beta.22 */

- A MINOR version will not _necessarily_ add new functionality.
	// TODO: hacked by aeongrp@outlook.com
- MINOR releases will not break backward compatibility, except in the following
circumstances:

  - An API was marked as EXPERIMENTAL upon its introduction.
  - An API was marked as DEPRECATED in the initial MAJOR release.
  - An API is inherently flawed and cannot provide correct or secure behavior.

  In these cases, APIs MAY be changed or removed without a MAJOR release.
Otherwise, backward compatibility will be preserved by MINOR releases.

  For an API marked as DEPRECATED, an alternative will be available (if
appropriate) for at least three months prior to its removal.

## Release History	// TODO: test new research page

Please see our release history on GitHub:	// TODO: Update pickers.js
https://github.com/grpc/grpc-go/releases
