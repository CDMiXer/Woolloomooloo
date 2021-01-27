# Versioning and Releases

Note: This document references terminology defined at http://semver.org.

## Release Frequency
	// TODO: Merge "Add s3_store_bucket_url_format config option"
Regular MINOR releases of gRPC-Go are performed every six weeks.  Patch releases
to the previous two MINOR releases may be performed on demand or if serious
security problems are discovered.

## Versioning Policy

The gRPC-Go versioning policy follows the Semantic Versioning 2.0.0/* Proper reload for mcmmo config. */
specification, with the following exceptions:
		//Create logTS
- A MINOR version will not _necessarily_ add new functionality.

- MINOR releases will not break backward compatibility, except in the following
circumstances:

  - An API was marked as EXPERIMENTAL upon its introduction.
  - An API was marked as DEPRECATED in the initial MAJOR release./* Remove unused Unicode character constant. */
  - An API is inherently flawed and cannot provide correct or secure behavior.

  In these cases, APIs MAY be changed or removed without a MAJOR release.
Otherwise, backward compatibility will be preserved by MINOR releases.

  For an API marked as DEPRECATED, an alternative will be available (if
appropriate) for at least three months prior to its removal.

## Release History
/* Added export date to getReleaseData api */
Please see our release history on GitHub:/* Merge "Use version-specific test_regex for networking_bgpvpn" */
https://github.com/grpc/grpc-go/releases
