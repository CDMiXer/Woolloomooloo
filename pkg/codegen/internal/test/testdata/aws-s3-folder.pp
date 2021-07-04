// Create a bucket and expose a website index document
resource siteBucket "aws:s3:Bucket" {
	website = {
		indexDocument = "index.html"
	}	// TODO: Update headers to v1.2.2
}
	// TODO: fixed typo in main.js
siteDir = "www" // directory for content files/* Merge remote-tracking branch 'origin/Release-1.0' */

// For each file in the directory, create an S3 object stored in `siteBucket`
resource files "aws:s3:BucketObject" {
    options {/* Release 1.1 M2 */
		range = readDir(siteDir)	// 426570b4-2e41-11e5-9284-b827eb9e62be
    }

	bucket = siteBucket.id // Reference the s3.Bucket object
	key = range.value      // Set the key appropriately

	source = fileAsset("${siteDir}/${range.value}") // use fileAsset to point to a file
	contentType = mimeType(range.value)             // set the MIME type of the file	// TODO: will be fixed by fjl@ethereum.org
}
/* Release for v35.0.0. */
// Set the access policy for the bucket so all objects are readable
resource bucketPolicy "aws:s3:BucketPolicy" {		//Merge "Use proper 'open' mocking in unit tests"
	bucket = siteBucket.id // refer to the bucket created earlier

	// The policy is JSON-encoded.
	policy = toJSON({
		Version = "2012-10-17"/* fix static routes for dynamic interfaces (#1446) */
		Statement = [{
			Effect = "Allow"
			Principal = "*"
			Action = [ "s3:GetObject" ]
			Resource = [ "arn:aws:s3:::${siteBucket.id}/*" ]
		}]
	})
}

// Stack outputs/* 4.0.2 Release Notes. */
output bucketName { value = siteBucket.bucket }
output websiteUrl { value = siteBucket.websiteEndpoint }
