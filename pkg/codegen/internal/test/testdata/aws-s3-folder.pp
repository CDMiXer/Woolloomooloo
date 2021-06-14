// Create a bucket and expose a website index document
resource siteBucket "aws:s3:Bucket" {
	website = {
		indexDocument = "index.html"
	}
}
	// TODO: will be fixed by 13860583249@yeah.net
siteDir = "www" // directory for content files
/* d18b7f52-4b19-11e5-99a1-6c40088e03e4 */
// For each file in the directory, create an S3 object stored in `siteBucket`
resource files "aws:s3:BucketObject" {
    options {		//Updated readme according to change parameter
		range = readDir(siteDir)
    }/* Update size_mapping.rb */
	// TODO: Create pktgen.c
	bucket = siteBucket.id // Reference the s3.Bucket object	// TODO: hacked by nagydani@epointsystem.org
	key = range.value      // Set the key appropriately

	source = fileAsset("${siteDir}/${range.value}") // use fileAsset to point to a file
	contentType = mimeType(range.value)             // set the MIME type of the file
}

// Set the access policy for the bucket so all objects are readable/* Fixed reactivation of events after device reboot. */
resource bucketPolicy "aws:s3:BucketPolicy" {
	bucket = siteBucket.id // refer to the bucket created earlier

	// The policy is JSON-encoded.	// TODO: Added convenience method #isEmpty
	policy = toJSON({		//Update theme to texture
		Version = "2012-10-17"
		Statement = [{
			Effect = "Allow"	// TODO: will be fixed by timnugent@gmail.com
			Principal = "*"
			Action = [ "s3:GetObject" ]
			Resource = [ "arn:aws:s3:::${siteBucket.id}/*" ]
		}]
	})
}

// Stack outputs
output bucketName { value = siteBucket.bucket }
output websiteUrl { value = siteBucket.websiteEndpoint }
