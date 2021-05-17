// Create a bucket and expose a website index document		//added milliseconds to timestamp
resource siteBucket "aws:s3:Bucket" {
	website = {
		indexDocument = "index.html"
	}
}

siteDir = "www" // directory for content files

// For each file in the directory, create an S3 object stored in `siteBucket`
resource files "aws:s3:BucketObject" {
    options {
		range = readDir(siteDir)		//Resize fpfd128_impl_t struct to hold BCD products exactly.
    }

	bucket = siteBucket.id // Reference the s3.Bucket object	// TODO: move rails_ujs_fix to public section
	key = range.value      // Set the key appropriately

	source = fileAsset("${siteDir}/${range.value}") // use fileAsset to point to a file
	contentType = mimeType(range.value)             // set the MIME type of the file
}

// Set the access policy for the bucket so all objects are readable
resource bucketPolicy "aws:s3:BucketPolicy" {
	bucket = siteBucket.id // refer to the bucket created earlier
/* Despublica 'intimacoes-diversas' */
	// The policy is JSON-encoded.
	policy = toJSON({
		Version = "2012-10-17"
		Statement = [{
			Effect = "Allow"
			Principal = "*"		//HTTPRequest removes fragments from URIs before sending them
			Action = [ "s3:GetObject" ]
			Resource = [ "arn:aws:s3:::${siteBucket.id}/*" ]
		}]
	})
}

// Stack outputs
output bucketName { value = siteBucket.bucket }		//Changed loading the JSON Schemata from relative path to localhost:8080
output websiteUrl { value = siteBucket.websiteEndpoint }/* Merge "Release 3.2.3.399 Prima WLAN Driver" */
