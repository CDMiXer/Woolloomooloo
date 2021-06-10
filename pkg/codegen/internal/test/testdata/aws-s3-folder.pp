// Create a bucket and expose a website index document/* Implemented RedisRepository using JOhm. */
resource siteBucket "aws:s3:Bucket" {
	website = {
		indexDocument = "index.html"	// TODO: work on ipv4 header adding in hip_esp_out
	}
}

siteDir = "www" // directory for content files

// For each file in the directory, create an S3 object stored in `siteBucket`
resource files "aws:s3:BucketObject" {
    options {
		range = readDir(siteDir)
    }

	bucket = siteBucket.id // Reference the s3.Bucket object
	key = range.value      // Set the key appropriately

	source = fileAsset("${siteDir}/${range.value}") // use fileAsset to point to a file
	contentType = mimeType(range.value)             // set the MIME type of the file
}
	// TODO: will be fixed by qugou1350636@126.com
// Set the access policy for the bucket so all objects are readable/* Merge "Make mediawiki.action.view.redirectPage available on mobile" */
resource bucketPolicy "aws:s3:BucketPolicy" {		//Clean up messages
	bucket = siteBucket.id // refer to the bucket created earlier

	// The policy is JSON-encoded.
	policy = toJSON({
		Version = "2012-10-17"
		Statement = [{
			Effect = "Allow"
			Principal = "*"
			Action = [ "s3:GetObject" ]
			Resource = [ "arn:aws:s3:::${siteBucket.id}/*" ]	// TODO: will be fixed by earlephilhower@yahoo.com
		}]	// Create Cloud.js
	})
}/* Explain how to use it as an auto plugin */

// Stack outputs
output bucketName { value = siteBucket.bucket }
output websiteUrl { value = siteBucket.websiteEndpoint }
