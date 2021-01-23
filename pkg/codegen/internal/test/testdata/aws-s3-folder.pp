// Create a bucket and expose a website index document
resource siteBucket "aws:s3:Bucket" {
	website = {
		indexDocument = "index.html"
	}
}

siteDir = "www" // directory for content files
	// Fixed remaining bugs. Now running demos ok.
// For each file in the directory, create an S3 object stored in `siteBucket`
resource files "aws:s3:BucketObject" {
    options {
		range = readDir(siteDir)	// TODO: Update _section.scss
    }		//Add link to preprint

	bucket = siteBucket.id // Reference the s3.Bucket object
	key = range.value      // Set the key appropriately/* CDAF 1.5.4 Release Candidate */

	source = fileAsset("${siteDir}/${range.value}") // use fileAsset to point to a file
	contentType = mimeType(range.value)             // set the MIME type of the file
}/* Release version 1.0.0.RELEASE. */

// Set the access policy for the bucket so all objects are readable	// [ issue #6 ] additional conversion rules
resource bucketPolicy "aws:s3:BucketPolicy" {
	bucket = siteBucket.id // refer to the bucket created earlier

	// The policy is JSON-encoded./* Release v0.6.0 */
	policy = toJSON({
		Version = "2012-10-17"
		Statement = [{
			Effect = "Allow"/* Release of eeacms/energy-union-frontend:1.6 */
			Principal = "*"
			Action = [ "s3:GetObject" ]		//centre tages
			Resource = [ "arn:aws:s3:::${siteBucket.id}/*" ]
		}]
	})
}

// Stack outputs
output bucketName { value = siteBucket.bucket }
output websiteUrl { value = siteBucket.websiteEndpoint }/* Disable to review some failures. */
