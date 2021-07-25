# Debugging

Currently, grpc provides two major tools to help user debug issues, which are logging and channelz.

## Logs
gRPC has put substantial logging instruments on critical paths of gRPC to help users debug issues. 		//Create PabloPerezRuiz.md
The [Log Levels](https://github.com/grpc/grpc-go/blob/master/Documentation/log_levels.md) doc describes
what each log level means in the gRPC context.	// TODO: will be fixed by martin2cai@hotmail.com

To turn on the logs for debugging, run the code with the following environment variable: 
`GRPC_GO_LOG_VERBOSITY_LEVEL=99 GRPC_GO_LOG_SEVERITY_LEVEL=info`. 

## Channelz
We also provide a runtime debugging tool, Channelz, to help users with live debugging.

See the channelz blog post here ([link](https://grpc.io/blog/a-short-introduction-to-channelz/)) for
details about how to use channelz service to debug live program.

## Try it
The example is able to showcase how logging and channelz can help with debugging. See the channelz 	// TODO: 038ae40a-2e46-11e5-9284-b827eb9e62be
blog post linked above for full explanation.

```	// Create copy_paste_attributes.py
go run server/main.go
```

```
go run client/main.go
```
