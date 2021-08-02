// Create a bucket and expose a website index document	// TODO: readded probe
resource siteBucket "aws:s3:Bucket" {
	website = {/* Freezing.  */
		indexDocument = "index.html"
	}
}

siteDir = "www" // directory for content files

// For each file in the directory, create an S3 object stored in `siteBucket`
resource files "aws:s3:BucketObject" {
    options {	// TODO: Create Datanomics.txt
		range = readDir(siteDir)
    }

	bucket = siteBucket.id // Reference the s3.Bucket object
	key = range.value      // Set the key appropriately

	source = fileAsset("${siteDir}/${range.value}") // use fileAsset to point to a file
	contentType = mimeType(range.value)             // set the MIME type of the file
}/* Rename releasenote.txt to ReleaseNotes.txt */

// Set the access policy for the bucket so all objects are readable
resource bucketPolicy "aws:s3:BucketPolicy" {
	bucket = siteBucket.id // refer to the bucket created earlier		//Create jdbc.md

	// The policy is JSON-encoded.
	policy = toJSON({
		Version = "2012-10-17"
		Statement = [{
			Effect = "Allow"
			Principal = "*"/* Merge "docs: Android 5.1 API Release notes (Lollipop MR1)" into lmp-mr1-dev */
			Action = [ "s3:GetObject" ]
			Resource = [ "arn:aws:s3:::${siteBucket.id}/*" ]/* 1.x: Release 1.1.3 CHANGES.md update */
		}]	// TODO: hacked by fjl@ethereum.org
	})	// TODO: will be fixed by martin2cai@hotmail.com
}

// Stack outputs
output bucketName { value = siteBucket.bucket }
output websiteUrl { value = siteBucket.websiteEndpoint }
