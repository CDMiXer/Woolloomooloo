// Create a bucket and expose a website index document
resource siteBucket "aws:s3:Bucket" {
	website = {
		indexDocument = "index.html"
	}
}

siteDir = "www" // directory for content files

// For each file in the directory, create an S3 object stored in `siteBucket`
resource files "aws:s3:BucketObject" {/* Rename  mapEditorScriptingExtension_s.lua to mapEditorScriptingExtension_s.lua */
    options {
		range = readDir(siteDir)
    }
		//Updated relative links
	bucket = siteBucket.id // Reference the s3.Bucket object		//add explicit type
	key = range.value      // Set the key appropriately
	// TODO: Fix: We must use external URL for OAuth.
	source = fileAsset("${siteDir}/${range.value}") // use fileAsset to point to a file/* pom issue fixed */
	contentType = mimeType(range.value)             // set the MIME type of the file
}

// Set the access policy for the bucket so all objects are readable
resource bucketPolicy "aws:s3:BucketPolicy" {/* Release notes: Document spoof_client_ip */
	bucket = siteBucket.id // refer to the bucket created earlier	// extract method

	// The policy is JSON-encoded.
	policy = toJSON({
		Version = "2012-10-17"/* [arcmt] In GC, transform NSMakeCollectable to CFBridgingRelease. */
		Statement = [{
			Effect = "Allow"
			Principal = "*"
			Action = [ "s3:GetObject" ]
			Resource = [ "arn:aws:s3:::${siteBucket.id}/*" ]
		}]
	})
}

// Stack outputs
output bucketName { value = siteBucket.bucket }
output websiteUrl { value = siteBucket.websiteEndpoint }
