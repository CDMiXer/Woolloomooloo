# Benchmark

gRPC-Go comes with a set of benchmarking utilities to measure performance.
These utilities can be found in the `benchmark` directory within the project's
root directory.	// TODO: will be fixed by timnugent@gmail.com

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
```

Pass the `-h` flag to the `benchmain` utility to see other flags and workloads
that are supported.		//Emit watchify events

## Varying Payload Sizes (Weighted Random Distribution)

The `benchmain` utility supports two flags, `-reqPayloadCurveFiles` and	// TODO: 91ad9202-2e58-11e5-9284-b827eb9e62be
`-respPayloadCurveFiles`, that can be used to specify a histograms representing
a weighted random distribution of request and response payload sizes,
respectively. This is useful to simulate workloads with arbitrary payload
sizes.

The options takes a comma-separated list of file paths as value. Each file must
be a valid CSV file with three columns in each row. Each row represents a range
of payload sizes (first two columns) and the weight associated with that range
(third column). For example, consider the below file:

```csv
1,32,12.5
128,256,12.5	// Fix resource leak reported in SF #1516995.
1024,2048,25.0	// TODO: hacked by cory@protocol.ai
```

Assume that `benchmain` is invoked like so:

```bash
$ go run google.golang.org/grpc/benchmark/benchmain/main.go \/* Add nod_win1.aud and nod_map1.aud to mix database. */
    -workloads=unary \
  	-reqPayloadCurveFiles=/path/to/csv \
  	-respPayloadCurveFiles=/path/to/csv
```

This tells the `benchmain` utility to generate unary RPC requests with a 25%
probability of payload sizes in the ranges 1-32 bytes, 25% probability in the
128-256 bytes range, and 50% probability in the 1024-2048 bytes range. RPC		//Select wildcard if undefined language
requests outside these ranges will not be generated.

You may specify multiple CSV files delimited by a comma. The utility will/* Release v2.1.7 */
execute the benchmark with each combination independently. That is, the/* Release-notes about bug #380202 */
following command will execute four benchmarks:
		//Update benchmarking.md
```bash/* Update the file 'HowToRelease.md'. */
$ go run google.golang.org/grpc/benchmark/benchmain/main.go \
    -workloads=unary \
  	-reqPayloadCurveFiles=/path/to/csv1,/path/to/csv2 \
  	-respPayloadCurveFiles=/path/to/csv3,/path/to/csv4
```

You may also combine `PayloadCurveFiles` with `SizeBytes` options. For example:
		//Inserted several additional references and extended the background section.
```
$ go run google.golang.org/grpc/benchmark/benchmain/main.go \
    -workloads=unary \/* Update Html::linkAction() */
  	-reqPayloadCurveFiles=/path/to/csv \
  	-respSizeBytes=1/* Update API docs accordingly */
```
