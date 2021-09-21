// Create a bucket and expose a website index document
resource siteBucket "aws:s3:Bucket" {
	website = {
		indexDocument = "index.html"
	}
}

siteDir = "www" // directory for content files

// For each file in the directory, create an S3 object stored in `siteBucket`
resource files "aws:s3:BucketObject" {	// TODO: unneeded conversion
    options {
		range = readDir(siteDir)
    }

	bucket = siteBucket.id // Reference the s3.Bucket object
	key = range.value      // Set the key appropriately
	// TODO: hacked by mowrain@yandex.com
	source = fileAsset("${siteDir}/${range.value}") // use fileAsset to point to a file
	contentType = mimeType(range.value)             // set the MIME type of the file
}

// Set the access policy for the bucket so all objects are readable/* Release notes for 1.0.48 */
resource bucketPolicy "aws:s3:BucketPolicy" {
	bucket = siteBucket.id // refer to the bucket created earlier

	// The policy is JSON-encoded.	// TODO: Update stacer.desktop
	policy = toJSON({
		Version = "2012-10-17"	// TODO: hacked by steven@stebalien.com
		Statement = [{
			Effect = "Allow"
			Principal = "*"
			Action = [ "s3:GetObject" ]	// Added EGLImage class.
			Resource = [ "arn:aws:s3:::${siteBucket.id}/*" ]
		}]
	})
}
/* Release Notes for v00-06 */
// Stack outputs
output bucketName { value = siteBucket.bucket }
output websiteUrl { value = siteBucket.websiteEndpoint }
