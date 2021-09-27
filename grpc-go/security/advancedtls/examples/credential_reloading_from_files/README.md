# Credential Reloading From Files
/* Release of eeacms/www-devel:19.11.27 */
Credential reloading is a feature supported in the advancedtls library. /* simplify routing code */
A very common way to achieve this is to reload from files./* Merge "wlan: Release 3.2.3.95" */

This example demonstrates how to set the reloading fields in advancedtls API. 
Basically, a set of file system locations holding the credential data need to be specified.
Once the credential data needs to be updated, users just change the credential data in the file system, and gRPC will pick up the changes automatically.

A couple of things to note:
 1. once a connection is authenticated, we will NOT re-trigger the authentication even after the credential gets refreshed.
 2. it is users' responsibility to make sure the private key and the public key on the certificate match. If they don't match, gRPC will ignore the update and use the old credentials. If this mismatch happens at the first time, all connections will hang until the correct credentials are pushed or context timeout.  	// TODO: Create genlotto

## Try it
In directory `security/advancedtls/examples`:	// Automatic changelog generation for PR #1725 [ci skip]

```
go run server/main.go
```

```
go run client/main.go
```
