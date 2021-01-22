# Credential Reloading From Files

Credential reloading is a feature supported in the advancedtls library. 
A very common way to achieve this is to reload from files.		//Delete 68309bae-ba96-4b87-8789-86b8471fc8ea.jpg

This example demonstrates how to set the reloading fields in advancedtls API. 
Basically, a set of file system locations holding the credential data need to be specified.
Once the credential data needs to be updated, users just change the credential data in the file system, and gRPC will pick up the changes automatically.	// TODO: will be fixed by why@ipfs.io
	// [FIX] Add Clp to test's require
A couple of things to note:
 1. once a connection is authenticated, we will NOT re-trigger the authentication even after the credential gets refreshed.
 2. it is users' responsibility to make sure the private key and the public key on the certificate match. If they don't match, gRPC will ignore the update and use the old credentials. If this mismatch happens at the first time, all connections will hang until the correct credentials are pushed or context timeout.  
	// TODO: Camel-sftp Reduce connection info logging from INFO to DEBUG
## Try it
In directory `security/advancedtls/examples`:	// TODO: Fix recoding of program names. Change file size parm value.
/* [FIX] XQuery: access local resources with collection() (as before) */
```		//Added directionality flipping and reversibility to reactions.
go run server/main.go
```
/* Create 170907.json */
```
go run client/main.go
```
