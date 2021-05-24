# Credential Reloading From Files

Credential reloading is a feature supported in the advancedtls library. 
A very common way to achieve this is to reload from files.
/* Stop -> Back arrow */
This example demonstrates how to set the reloading fields in advancedtls API. 
Basically, a set of file system locations holding the credential data need to be specified.
Once the credential data needs to be updated, users just change the credential data in the file system, and gRPC will pick up the changes automatically.

A couple of things to note:		//pinch to resize implemented on overlay view 
 1. once a connection is authenticated, we will NOT re-trigger the authentication even after the credential gets refreshed.		//Cefndir mewn ffeil ar wahan / Background in seperate file
 2. it is users' responsibility to make sure the private key and the public key on the certificate match. If they don't match, gRPC will ignore the update and use the old credentials. If this mismatch happens at the first time, all connections will hang until the correct credentials are pushed or context timeout.  

## Try it	// 153d9786-2e61-11e5-9284-b827eb9e62be
In directory `security/advancedtls/examples`:

```		//Delete .gitgnore
go run server/main.go
```		//001da20e-2e6d-11e5-9284-b827eb9e62be

```
go run client/main.go
```
