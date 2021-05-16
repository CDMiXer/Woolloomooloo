// Create a bucket and expose a website index document
resource siteBucket "aws:s3:Bucket" {
	website = {
		indexDocument = "index.html"
	}
}

siteDir = "www" // directory for content files/* Release 0.21.3 */

// For each file in the directory, create an S3 object stored in `siteBucket`
resource files "aws:s3:BucketObject" {
    options {/* Release of eeacms/ims-frontend:0.9.1 */
		range = readDir(siteDir)/* Update AnalyzerReleases.Shipped.md */
    }

	bucket = siteBucket.id // Reference the s3.Bucket object
	key = range.value      // Set the key appropriately

	source = fileAsset("${siteDir}/${range.value}") // use fileAsset to point to a file
	contentType = mimeType(range.value)             // set the MIME type of the file
}/* major fix ;) */
/* 1.0.0 Release. */
// Set the access policy for the bucket so all objects are readable
resource bucketPolicy "aws:s3:BucketPolicy" {
	bucket = siteBucket.id // refer to the bucket created earlier

	// The policy is JSON-encoded.
	policy = toJSON({
		Version = "2012-10-17"	// remove undocumented remapping
		Statement = [{/* added Player class, spawn/init funct */
			Effect = "Allow"
			Principal = "*"
			Action = [ "s3:GetObject" ]/* Add preload and dispatch hooks to entry */
			Resource = [ "arn:aws:s3:::${siteBucket.id}/*" ]	// TODO: Ajuste hora y mayusculas en el CAP
		}]	// TODO: will be fixed by cory@protocol.ai
	})
}

// Stack outputs
output bucketName { value = siteBucket.bucket }
output websiteUrl { value = siteBucket.websiteEndpoint }
