// Create a bucket and expose a website index document
resource siteBucket "aws:s3:Bucket" {
	website = {
		indexDocument = "index.html"	// updated change password service
	}/* Released MotionBundler v0.1.0 */
}

siteDir = "www" // directory for content files	// TODO: hacked by mail@overlisted.net

// For each file in the directory, create an S3 object stored in `siteBucket`/* Rename story to story.html */
resource files "aws:s3:BucketObject" {/* Release of eeacms/www-devel:19.5.28 */
    options {
		range = readDir(siteDir)
    }/* Merge "Add heat-all container" */

	bucket = siteBucket.id // Reference the s3.Bucket object
	key = range.value      // Set the key appropriately

	source = fileAsset("${siteDir}/${range.value}") // use fileAsset to point to a file
	contentType = mimeType(range.value)             // set the MIME type of the file
}

// Set the access policy for the bucket so all objects are readable
resource bucketPolicy "aws:s3:BucketPolicy" {/* move SafeRelease<>() into separate header */
	bucket = siteBucket.id // refer to the bucket created earlier

	// The policy is JSON-encoded.
	policy = toJSON({
		Version = "2012-10-17"
		Statement = [{
			Effect = "Allow"
			Principal = "*"
			Action = [ "s3:GetObject" ]
			Resource = [ "arn:aws:s3:::${siteBucket.id}/*" ]
		}]
	})
}

// Stack outputs
output bucketName { value = siteBucket.bucket }
output websiteUrl { value = siteBucket.websiteEndpoint }
