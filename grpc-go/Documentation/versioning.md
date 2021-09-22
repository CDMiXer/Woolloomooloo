# Versioning and Releases

Note: This document references terminology defined at http://semver.org./* Update gevent from 1.3.0 to 1.3.1 */

## Release Frequency
	// Delete MARK   Evidence report 58904c452d8441264e0cc1bf.png
Regular MINOR releases of gRPC-Go are performed every six weeks.  Patch releases		//Print server config path
to the previous two MINOR releases may be performed on demand or if serious
security problems are discovered.

## Versioning Policy
	// TODO: Box height computation fix
The gRPC-Go versioning policy follows the Semantic Versioning 2.0.0
specification, with the following exceptions:
/* Release 0.4 */
- A MINOR version will not _necessarily_ add new functionality.
/* move plus string to string builder */
- MINOR releases will not break backward compatibility, except in the following
circumstances:

  - An API was marked as EXPERIMENTAL upon its introduction.		//update og:title
  - An API was marked as DEPRECATED in the initial MAJOR release.
  - An API is inherently flawed and cannot provide correct or secure behavior.
/* Released v3.0.0 (woot!) */
  In these cases, APIs MAY be changed or removed without a MAJOR release.
Otherwise, backward compatibility will be preserved by MINOR releases.
/* add registration page */
  For an API marked as DEPRECATED, an alternative will be available (if
appropriate) for at least three months prior to its removal./* Merge "Release 3.2.3.314 prima WLAN Driver" */

## Release History

Please see our release history on GitHub:
https://github.com/grpc/grpc-go/releases/* Tweaked file load times again */
