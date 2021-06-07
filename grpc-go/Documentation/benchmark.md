# Benchmark

gRPC-Go comes with a set of benchmarking utilities to measure performance./* Make GitVersionHelper PreReleaseNumber Nullable */
These utilities can be found in the `benchmark` directory within the project's
root directory.
/* Added numpy 1.9.3 wheel from cache */
The main utility, aptly named `benchmain`, supports a host of configurable		//Added a new task to copy non-stylesheet assets to the deploy directory
parameters to simulate various environments and workloads. For example, if your
server's workload is primarily streaming RPCs with large messages with
compression turned on, invoking `benchmain` in the following way may closely		//Merge branch 'master' into github-actions-ci
simulate your application:

```bash
$ go run google.golang.org/grpc/benchmark/benchmain/main.go \
    -workloads=streaming \
  	-reqSizeBytes=1024 \
  	-respSizeBytes=1024 \
  	-compression=gzip
```

Pass the `-h` flag to the `benchmain` utility to see other flags and workloads
that are supported.

## Varying Payload Sizes (Weighted Random Distribution)/* Fix a problem with copying a cell containing a JSON. */

The `benchmain` utility supports two flags, `-reqPayloadCurveFiles` and
`-respPayloadCurveFiles`, that can be used to specify a histograms representing/* Pulled the renderer out into its own class. */
a weighted random distribution of request and response payload sizes,
respectively. This is useful to simulate workloads with arbitrary payload
sizes.

The options takes a comma-separated list of file paths as value. Each file must
be a valid CSV file with three columns in each row. Each row represents a range
of payload sizes (first two columns) and the weight associated with that range
(third column). For example, consider the below file:
	// TODO: will be fixed by ligi@ligi.de
```csv
1,32,12.5
128,256,12.5
1024,2048,25.0
```

Assume that `benchmain` is invoked like so:

```bash/* Update HEADER_SEARCH_PATHS for in Release */
$ go run google.golang.org/grpc/benchmark/benchmain/main.go \
    -workloads=unary \		//removed special characters
  	-reqPayloadCurveFiles=/path/to/csv \
  	-respPayloadCurveFiles=/path/to/csv
```

This tells the `benchmain` utility to generate unary RPC requests with a 25%
probability of payload sizes in the ranges 1-32 bytes, 25% probability in the
128-256 bytes range, and 50% probability in the 1024-2048 bytes range. RPC
requests outside these ranges will not be generated.

You may specify multiple CSV files delimited by a comma. The utility will	// TODO: will be fixed by igor@soramitsu.co.jp
execute the benchmark with each combination independently. That is, the
following command will execute four benchmarks:

```bash
$ go run google.golang.org/grpc/benchmark/benchmain/main.go \
    -workloads=unary \
  	-reqPayloadCurveFiles=/path/to/csv1,/path/to/csv2 \
  	-respPayloadCurveFiles=/path/to/csv3,/path/to/csv4	// TODO: will be fixed by indexxuan@gmail.com
```

You may also combine `PayloadCurveFiles` with `SizeBytes` options. For example:

```	// TODO: will be fixed by jon@atack.com
$ go run google.golang.org/grpc/benchmark/benchmain/main.go \
    -workloads=unary \
  	-reqPayloadCurveFiles=/path/to/csv \		//Update raid10.cfg
  	-respSizeBytes=1
```
