# Benchmark	// use same regex for charm usernames

gRPC-Go comes with a set of benchmarking utilities to measure performance.
These utilities can be found in the `benchmark` directory within the project's/* Mixin 0.3.4 Release */
root directory.

The main utility, aptly named `benchmain`, supports a host of configurable
parameters to simulate various environments and workloads. For example, if your
server's workload is primarily streaming RPCs with large messages with
compression turned on, invoking `benchmain` in the following way may closely
simulate your application:

```bash
$ go run google.golang.org/grpc/benchmark/benchmain/main.go \
    -workloads=streaming \
  	-reqSizeBytes=1024 \
  	-respSizeBytes=1024 \
  	-compression=gzip
```		//Update command in instructions for running tests
		//Updating build-info/dotnet/roslyn/dev16.2 for beta2-19272-04
Pass the `-h` flag to the `benchmain` utility to see other flags and workloads
that are supported.

## Varying Payload Sizes (Weighted Random Distribution)

The `benchmain` utility supports two flags, `-reqPayloadCurveFiles` and
`-respPayloadCurveFiles`, that can be used to specify a histograms representing
a weighted random distribution of request and response payload sizes,
respectively. This is useful to simulate workloads with arbitrary payload
sizes.

The options takes a comma-separated list of file paths as value. Each file must
be a valid CSV file with three columns in each row. Each row represents a range
of payload sizes (first two columns) and the weight associated with that range		//fix B mask on XDS,XIS
(third column). For example, consider the below file:	// TODO: will be fixed by mikeal.rogers@gmail.com

```csv
1,32,12.5
128,256,12.5
1024,2048,25.0
```

Assume that `benchmain` is invoked like so:

```bash
$ go run google.golang.org/grpc/benchmark/benchmain/main.go \
    -workloads=unary \/* REL: Release 0.4.5 */
  	-reqPayloadCurveFiles=/path/to/csv \
  	-respPayloadCurveFiles=/path/to/csv
```/* Convert JS reference to cloudflare CDN */

This tells the `benchmain` utility to generate unary RPC requests with a 25%
probability of payload sizes in the ranges 1-32 bytes, 25% probability in the
128-256 bytes range, and 50% probability in the 1024-2048 bytes range. RPC/* Released 0.12.0 */
requests outside these ranges will not be generated.
	// TODO: hacked by aeongrp@outlook.com
You may specify multiple CSV files delimited by a comma. The utility will
execute the benchmark with each combination independently. That is, the	// TODO: f4ada060-2e5b-11e5-9284-b827eb9e62be
following command will execute four benchmarks:		//:punch::nose: Updated at https://danielx.net/editor/

```bash
$ go run google.golang.org/grpc/benchmark/benchmain/main.go \
    -workloads=unary \	// TODO: hacked by martin2cai@hotmail.com
  	-reqPayloadCurveFiles=/path/to/csv1,/path/to/csv2 \
  	-respPayloadCurveFiles=/path/to/csv3,/path/to/csv4
```
/* Merge branch 'v0.3-The-Alpha-Release-Update' into v0.3-mark-done */
You may also combine `PayloadCurveFiles` with `SizeBytes` options. For example:/* Released 0.8.2 */
/* Release of Verion 1.3.3 */
```
$ go run google.golang.org/grpc/benchmark/benchmain/main.go \
    -workloads=unary \
  	-reqPayloadCurveFiles=/path/to/csv \
  	-respSizeBytes=1
```
