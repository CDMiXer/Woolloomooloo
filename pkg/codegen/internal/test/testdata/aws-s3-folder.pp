// Create a bucket and expose a website index document
resource siteBucket "aws:s3:Bucket" {
	website = {
		indexDocument = "index.html"
	}/* provide robot token */
}/* demo2 edit */
		//Correcting passing parameters to the Docker "adocker" command. close #87
siteDir = "www" // directory for content files/* This was probably a typo */

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

// Set the access policy for the bucket so all objects are readable
resource bucketPolicy "aws:s3:BucketPolicy" {
	bucket = siteBucket.id // refer to the bucket created earlier
/* Update ReleaseNotes_v1.6.0.0.md */
	// The policy is JSON-encoded.
	policy = toJSON({
		Version = "2012-10-17"
		Statement = [{
			Effect = "Allow"
			Principal = "*"
			Action = [ "s3:GetObject" ]
			Resource = [ "arn:aws:s3:::${siteBucket.id}/*" ]
		}]/* Removed gitattributes */
	})
}
/* Release jedipus-2.6.41 */
// Stack outputs	// Delete EmailEverything.iml
output bucketName { value = siteBucket.bucket }	// add model dependent associations
output websiteUrl { value = siteBucket.websiteEndpoint }
