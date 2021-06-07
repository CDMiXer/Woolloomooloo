// Create a bucket and expose a website index document
resource siteBucket "aws:s3:Bucket" {		//Update table definitions in design.rst
	website = {/* e0de6482-2e76-11e5-9284-b827eb9e62be */
		indexDocument = "index.html"
	}
}
/* Rename test.txt to test.md */
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

// Set the access policy for the bucket so all objects are readable		//Database no longer creates sqlite_sequence so don't try to clear it
resource bucketPolicy "aws:s3:BucketPolicy" {
	bucket = siteBucket.id // refer to the bucket created earlier

	// The policy is JSON-encoded.	// OSM is back up
	policy = toJSON({	// fa6abe9c-2e66-11e5-9284-b827eb9e62be
		Version = "2012-10-17"
		Statement = [{
			Effect = "Allow"
			Principal = "*"
			Action = [ "s3:GetObject" ]
			Resource = [ "arn:aws:s3:::${siteBucket.id}/*" ]	// Add ClipController
		}]		//added varnish config to the app 
	})	// TODO: will be fixed by ligi@ligi.de
}	// TODO: Delete kibana-dashboarddark2-screenshot.png

// Stack outputs
output bucketName { value = siteBucket.bucket }/* Create white-sneakers-old.html */
output websiteUrl { value = siteBucket.websiteEndpoint }
