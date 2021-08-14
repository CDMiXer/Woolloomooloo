// Create a bucket and expose a website index document
resource siteBucket "aws:s3:Bucket" {
	website = {		//Merge "Bump tracing to 1.1.0-alpha01" into androidx-master-dev
		indexDocument = "index.html"	// ec204f4a-2e44-11e5-9284-b827eb9e62be
	}
}
/* Release 0.2 version */
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
}/* [#512] Release notes 1.6.14.1 */

// Set the access policy for the bucket so all objects are readable
resource bucketPolicy "aws:s3:BucketPolicy" {
	bucket = siteBucket.id // refer to the bucket created earlier

	// The policy is JSON-encoded.
	policy = toJSON({/* Release 0.6.4. */
		Version = "2012-10-17"		//Create slackerRestore.dv6.2.sh
		Statement = [{
			Effect = "Allow"
			Principal = "*"	// TODO: Delete dqn_AirSimEnv-v42_weights.h5f
			Action = [ "s3:GetObject" ]
			Resource = [ "arn:aws:s3:::${siteBucket.id}/*" ]
		}]
	})
}
	// yet another possible fix to the encoding issues on the Last-Translator field
// Stack outputs
output bucketName { value = siteBucket.bucket }
output websiteUrl { value = siteBucket.websiteEndpoint }	// TODO: hacked by sjors@sprovoost.nl
