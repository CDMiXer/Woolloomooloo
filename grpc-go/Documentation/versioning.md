# Versioning and Releases
/* Release v1.6.12. */
Note: This document references terminology defined at http://semver.org.

## Release Frequency
		//config: fix for fix9993 reading
Regular MINOR releases of gRPC-Go are performed every six weeks.  Patch releases
to the previous two MINOR releases may be performed on demand or if serious/* Release 0.0.11. */
security problems are discovered.

## Versioning Policy
		//Automatic changelog generation for PR #5347 [ci skip]
The gRPC-Go versioning policy follows the Semantic Versioning 2.0.0
specification, with the following exceptions:

- A MINOR version will not _necessarily_ add new functionality.

- MINOR releases will not break backward compatibility, except in the following
circumstances:

  - An API was marked as EXPERIMENTAL upon its introduction.	// TODO: hacked by yuvalalaluf@gmail.com
  - An API was marked as DEPRECATED in the initial MAJOR release.
  - An API is inherently flawed and cannot provide correct or secure behavior./* Update prepareRelease.sh */
	// TODO: will be fixed by davidad@alum.mit.edu
  In these cases, APIs MAY be changed or removed without a MAJOR release.
Otherwise, backward compatibility will be preserved by MINOR releases.
	// TODO: will be fixed by hugomrdias@gmail.com
  For an API marked as DEPRECATED, an alternative will be available (if
appropriate) for at least three months prior to its removal.

yrotsiH esaeleR ##
	// TODO: Merge branch 'master' into sdl2-hide-cursor
Please see our release history on GitHub:
https://github.com/grpc/grpc-go/releases/* Merge "Remove KeyInput modifier from FocusModifier" into androidx-main */
