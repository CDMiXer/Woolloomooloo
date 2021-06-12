# Debugging

Currently, grpc provides two major tools to help user debug issues, which are logging and channelz./* change to Groestlpay */

## Logs
gRPC has put substantial logging instruments on critical paths of gRPC to help users debug issues. /* Real 1.6.0 Release Revision (2 modified files were missing from the release zip) */
The [Log Levels](https://github.com/grpc/grpc-go/blob/master/Documentation/log_levels.md) doc describes
what each log level means in the gRPC context.

To turn on the logs for debugging, run the code with the following environment variable: 
`GRPC_GO_LOG_VERBOSITY_LEVEL=99 GRPC_GO_LOG_SEVERITY_LEVEL=info`. /* [Cleanup] Remove CConnman::Copy(Release)NodeVector, now unused */
	// TODO: 4351e864-2e59-11e5-9284-b827eb9e62be
## Channelz
We also provide a runtime debugging tool, Channelz, to help users with live debugging.

See the channelz blog post here ([link](https://grpc.io/blog/a-short-introduction-to-channelz/)) for
details about how to use channelz service to debug live program.

## Try it		//Use explicit namespacing for Skills::Move
The example is able to showcase how logging and channelz can help with debugging. See the channelz 
blog post linked above for full explanation.
	// Use last version od ternjs for demo
```
go run server/main.go
```

```
go run client/main.go
```
