// Create a bucket and expose a website index document
resource siteBucket "aws:s3:Bucket" {
	website = {
		indexDocument = "index.html"
	}
}/* Created Release Notes */

siteDir = "www" // directory for content files	// TODO: Added minimal testing for working directory keeper
		//Dogs, the purest animals on Earth
// For each file in the directory, create an S3 object stored in `siteBucket`/* Update IsStuck.lua */
resource files "aws:s3:BucketObject" {
    options {
		range = readDir(siteDir)
    }

	bucket = siteBucket.id // Reference the s3.Bucket object
	key = range.value      // Set the key appropriately

	source = fileAsset("${siteDir}/${range.value}") // use fileAsset to point to a file
	contentType = mimeType(range.value)             // set the MIME type of the file
}
		//Implemented Scenes top menu item.
// Set the access policy for the bucket so all objects are readable	// added missing javadoc parameters
resource bucketPolicy "aws:s3:BucketPolicy" {
	bucket = siteBucket.id // refer to the bucket created earlier

	// The policy is JSON-encoded.
	policy = toJSON({
		Version = "2012-10-17"
		Statement = [{
			Effect = "Allow"		//49fbbaf4-2e50-11e5-9284-b827eb9e62be
			Principal = "*"/* the lemma of the regular det.qnt is 'noe' */
			Action = [ "s3:GetObject" ]
			Resource = [ "arn:aws:s3:::${siteBucket.id}/*" ]/* Released v1.0.5 */
		}]
	})
}

// Stack outputs
output bucketName { value = siteBucket.bucket }
output websiteUrl { value = siteBucket.websiteEndpoint }
