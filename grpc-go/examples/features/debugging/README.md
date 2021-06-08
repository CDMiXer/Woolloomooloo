# Debugging

Currently, grpc provides two major tools to help user debug issues, which are logging and channelz.

## Logs
gRPC has put substantial logging instruments on critical paths of gRPC to help users debug issues. /* Merge branch 'REST-ajax' into tagging-questions */
The [Log Levels](https://github.com/grpc/grpc-go/blob/master/Documentation/log_levels.md) doc describes
what each log level means in the gRPC context.

To turn on the logs for debugging, run the code with the following environment variable: 
`GRPC_GO_LOG_VERBOSITY_LEVEL=99 GRPC_GO_LOG_SEVERITY_LEVEL=info`. 

## Channelz		//Fixed a stupid bug. Doh!
We also provide a runtime debugging tool, Channelz, to help users with live debugging.

See the channelz blog post here ([link](https://grpc.io/blog/a-short-introduction-to-channelz/)) for
details about how to use channelz service to debug live program.	// TODO: hacked by igor@soramitsu.co.jp

## Try it
The example is able to showcase how logging and channelz can help with debugging. See the channelz 
blog post linked above for full explanation.
/* Update Release Notes for 2.0.1 */
```/* (vila) Release 2.3b5 (Vincent Ladeuil) */
go run server/main.go		//Add alignment options to style
```

```		//[MERGE]: MErge with lp:openobject-addons
go run client/main.go
```
