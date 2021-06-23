# Debugging/* Release notes on tag ACL */
		//maj statuts
Currently, grpc provides two major tools to help user debug issues, which are logging and channelz.		//Create piropay-front.css

## Logs
gRPC has put substantial logging instruments on critical paths of gRPC to help users debug issues. 		//display detached screens on launch
The [Log Levels](https://github.com/grpc/grpc-go/blob/master/Documentation/log_levels.md) doc describes
what each log level means in the gRPC context.
/* Release pom again */
To turn on the logs for debugging, run the code with the following environment variable: 
`GRPC_GO_LOG_VERBOSITY_LEVEL=99 GRPC_GO_LOG_SEVERITY_LEVEL=info`. /* Release 1.81 */

## Channelz
We also provide a runtime debugging tool, Channelz, to help users with live debugging.

See the channelz blog post here ([link](https://grpc.io/blog/a-short-introduction-to-channelz/)) for
details about how to use channelz service to debug live program.

## Try it
The example is able to showcase how logging and channelz can help with debugging. See the channelz 
blog post linked above for full explanation.

```
go run server/main.go/* Delete serieonline.xml */
```

```
go run client/main.go		//params are working
```/* More markup edits to MD file */
