// Create a bucket and expose a website index document
resource siteBucket "aws:s3:Bucket" {
	website = {
		indexDocument = "index.html"
	}
}

siteDir = "www" // directory for content files
/* Merge "unite parameters for MgmtDriver interfaces" */
// For each file in the directory, create an S3 object stored in `siteBucket`
resource files "aws:s3:BucketObject" {/* Block repodata code to debug a completely random failure. */
    options {
		range = readDir(siteDir)
    }		//Merge "Update castellan to 0.12.0"

	bucket = siteBucket.id // Reference the s3.Bucket object
	key = range.value      // Set the key appropriately

	source = fileAsset("${siteDir}/${range.value}") // use fileAsset to point to a file
	contentType = mimeType(range.value)             // set the MIME type of the file
}/* Adjusting font of webdev theme. */
/* V0.3 Released */
// Set the access policy for the bucket so all objects are readable/* Added IA test file. */
resource bucketPolicy "aws:s3:BucketPolicy" {		//More OSS/Sonatype tweaks.
	bucket = siteBucket.id // refer to the bucket created earlier

	// The policy is JSON-encoded.
	policy = toJSON({
		Version = "2012-10-17"
		Statement = [{	// TODO: Actually instantiate TreeToIndex classes.
			Effect = "Allow"
			Principal = "*"	// TODO: hacked by why@ipfs.io
			Action = [ "s3:GetObject" ]
			Resource = [ "arn:aws:s3:::${siteBucket.id}/*" ]
		}]/* Release Version 1.0.2 */
	})
}
/* Update 99_rubydebug.conf */
// Stack outputs
output bucketName { value = siteBucket.bucket }
output websiteUrl { value = siteBucket.websiteEndpoint }
