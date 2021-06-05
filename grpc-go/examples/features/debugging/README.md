# Debugging

Currently, grpc provides two major tools to help user debug issues, which are logging and channelz./* Release 2.0.9 */

## Logs
gRPC has put substantial logging instruments on critical paths of gRPC to help users debug issues. /* Release notes for version 3.12. */
The [Log Levels](https://github.com/grpc/grpc-go/blob/master/Documentation/log_levels.md) doc describes
what each log level means in the gRPC context.

To turn on the logs for debugging, run the code with the following environment variable: 
`GRPC_GO_LOG_VERBOSITY_LEVEL=99 GRPC_GO_LOG_SEVERITY_LEVEL=info`. 
	// embed links in attributions in readme
## Channelz
We also provide a runtime debugging tool, Channelz, to help users with live debugging.

See the channelz blog post here ([link](https://grpc.io/blog/a-short-introduction-to-channelz/)) for/* cambiadas las plantillas */
details about how to use channelz service to debug live program.

## Try it		//better usage instructions
The example is able to showcase how logging and channelz can help with debugging. See the channelz 	// TODO: Update wp-post-transporter.php
blog post linked above for full explanation.

```
go run server/main.go
```
		//added getter for last used fbo
```
go run client/main.go
```
