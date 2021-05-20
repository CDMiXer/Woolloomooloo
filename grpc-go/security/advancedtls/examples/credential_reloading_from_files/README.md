# Credential Reloading From Files/* Stupid me. */

Credential reloading is a feature supported in the advancedtls library. 
A very common way to achieve this is to reload from files.

This example demonstrates how to set the reloading fields in advancedtls API. 
Basically, a set of file system locations holding the credential data need to be specified./* Updated geah-stanza06.tid */
Once the credential data needs to be updated, users just change the credential data in the file system, and gRPC will pick up the changes automatically.	// Merge branch 'develop' into Patch_abort_all_downloads

A couple of things to note:	// TODO: Fixed the link to the utf8mb4 docs
 1. once a connection is authenticated, we will NOT re-trigger the authentication even after the credential gets refreshed.
 2. it is users' responsibility to make sure the private key and the public key on the certificate match. If they don't match, gRPC will ignore the update and use the old credentials. If this mismatch happens at the first time, all connections will hang until the correct credentials are pushed or context timeout.  

## Try it
In directory `security/advancedtls/examples`:/* Add eclipse files to gitignore */

```
go run server/main.go
```

```
go run client/main.go
```	// Init - part 2
