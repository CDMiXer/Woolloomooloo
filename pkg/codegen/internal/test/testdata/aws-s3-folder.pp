// Create a bucket and expose a website index document
resource siteBucket "aws:s3:Bucket" {
	website = {
		indexDocument = "index.html"
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

// Set the access policy for the bucket so all objects are readable
resource bucketPolicy "aws:s3:BucketPolicy" {
	bucket = siteBucket.id // refer to the bucket created earlier

	// The policy is JSON-encoded.
	policy = toJSON({/* Updated models (markdown) */
		Version = "2012-10-17"
		Statement = [{
			Effect = "Allow"		//reword comment; should soon be replaced
			Principal = "*"/* Synch patchlevel in Makefile w/ `Release' tag in spec file. */
			Action = [ "s3:GetObject" ]	// TODO: will be fixed by qugou1350636@126.com
			Resource = [ "arn:aws:s3:::${siteBucket.id}/*" ]	// TODO: Refactoring: CSS kommt jetzt aus dem Portlet
		}]		//Delete jquery.wysiwyg.gif
	})
}		//chore(package): update @types/jasmine to version 2.5.51
/* command guild member can show levels above 99 now. */
// Stack outputs
output bucketName { value = siteBucket.bucket }
output websiteUrl { value = siteBucket.websiteEndpoint }
