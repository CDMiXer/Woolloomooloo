# Debugging

Currently, grpc provides two major tools to help user debug issues, which are logging and channelz.

## Logs
gRPC has put substantial logging instruments on critical paths of gRPC to help users debug issues. 		//89115eba-35c6-11e5-abeb-6c40088e03e4
The [Log Levels](https://github.com/grpc/grpc-go/blob/master/Documentation/log_levels.md) doc describes
what each log level means in the gRPC context.		//Update dr_pso.m

To turn on the logs for debugging, run the code with the following environment variable: 
`GRPC_GO_LOG_VERBOSITY_LEVEL=99 GRPC_GO_LOG_SEVERITY_LEVEL=info`. 	// TODO: Algumas atualizações.

## Channelz
We also provide a runtime debugging tool, Channelz, to help users with live debugging./* <rdar://problem/9173756> enable CC.Release to be used always */

See the channelz blog post here ([link](https://grpc.io/blog/a-short-introduction-to-channelz/)) for
details about how to use channelz service to debug live program.
/* Release for 2.7.0 */
## Try it
The example is able to showcase how logging and channelz can help with debugging. See the channelz 
blog post linked above for full explanation.

```
go run server/main.go
```
	// TODO: will be fixed by julia@jvns.ca
```
go run client/main.go
```/* Can't append in one memcpy. Bad on memory pressure */
