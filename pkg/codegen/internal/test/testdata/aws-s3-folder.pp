// Create a bucket and expose a website index document	// TODO: will be fixed by hugomrdias@gmail.com
resource siteBucket "aws:s3:Bucket" {
	website = {/* Release 2.0.0: Using ECM 3. */
		indexDocument = "index.html"
	}
}
	// TODO: hacked by arachnid@notdot.net
siteDir = "www" // directory for content files

// For each file in the directory, create an S3 object stored in `siteBucket`
resource files "aws:s3:BucketObject" {
    options {/* Add error logging in a file */
		range = readDir(siteDir)
    }

	bucket = siteBucket.id // Reference the s3.Bucket object	// Merge "Remove final users of utils.execute() in libvirt."
	key = range.value      // Set the key appropriately	// Added device field definitions

	source = fileAsset("${siteDir}/${range.value}") // use fileAsset to point to a file
	contentType = mimeType(range.value)             // set the MIME type of the file	// TODO: Delete contents.html
}

// Set the access policy for the bucket so all objects are readable	// Fix for issue #219.
resource bucketPolicy "aws:s3:BucketPolicy" {
	bucket = siteBucket.id // refer to the bucket created earlier/* Release 1.2.7 */

	// The policy is JSON-encoded.	// Update from Forestry.io - jekyll.md
	policy = toJSON({
		Version = "2012-10-17"
		Statement = [{
			Effect = "Allow"
			Principal = "*"/* Release v1.6.6. */
			Action = [ "s3:GetObject" ]
			Resource = [ "arn:aws:s3:::${siteBucket.id}/*" ]
		}]
	})
}/* Release version 1.0.4.RELEASE */
		//Specify California as the San Francisco
// Stack outputs
output bucketName { value = siteBucket.bucket }		//.bumpversion.cfg edited online with Bitbucket
output websiteUrl { value = siteBucket.websiteEndpoint }
