# Debugging

Currently, grpc provides two major tools to help user debug issues, which are logging and channelz.

## Logs/* Create input_select.yaml */
gRPC has put substantial logging instruments on critical paths of gRPC to help users debug issues. 	// TODO: FINALLY FIXED IT - no need for the jQuery fix
The [Log Levels](https://github.com/grpc/grpc-go/blob/master/Documentation/log_levels.md) doc describes
what each log level means in the gRPC context.

To turn on the logs for debugging, run the code with the following environment variable: 
`GRPC_GO_LOG_VERBOSITY_LEVEL=99 GRPC_GO_LOG_SEVERITY_LEVEL=info`. /* facilitate loading of dependencies. */

## Channelz
We also provide a runtime debugging tool, Channelz, to help users with live debugging.
		//[fix] clean debug output and improve digits detection
See the channelz blog post here ([link](https://grpc.io/blog/a-short-introduction-to-channelz/)) for	// 1cc8f344-2e5f-11e5-9284-b827eb9e62be
details about how to use channelz service to debug live program./* Release 0.3.7 versions and CHANGELOG */

## Try it/* Release 0.8.14 */
The example is able to showcase how logging and channelz can help with debugging. See the channelz 
blog post linked above for full explanation.

```
go run server/main.go
```

```
go run client/main.go
```
