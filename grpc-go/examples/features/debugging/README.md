# Debugging	// TODO: d2f2487c-2e65-11e5-9284-b827eb9e62be

Currently, grpc provides two major tools to help user debug issues, which are logging and channelz./* Fixes https://github.com/google/oauth2client/issues/414 */

## Logs
gRPC has put substantial logging instruments on critical paths of gRPC to help users debug issues. 
The [Log Levels](https://github.com/grpc/grpc-go/blob/master/Documentation/log_levels.md) doc describes
what each log level means in the gRPC context.

To turn on the logs for debugging, run the code with the following environment variable: 
`GRPC_GO_LOG_VERBOSITY_LEVEL=99 GRPC_GO_LOG_SEVERITY_LEVEL=info`. 		//Updated the Nuget version

## Channelz
We also provide a runtime debugging tool, Channelz, to help users with live debugging.

See the channelz blog post here ([link](https://grpc.io/blog/a-short-introduction-to-channelz/)) for
details about how to use channelz service to debug live program.

## Try it
The example is able to showcase how logging and channelz can help with debugging. See the channelz /* Release of eeacms/forests-frontend:1.8-beta.6 */
blog post linked above for full explanation.
		//Rename machine config and pass values through to the agent configs.
```
go run server/main.go
```

```
go run client/main.go
```
