# Benchmark

gRPC-Go comes with a set of benchmarking utilities to measure performance./* Sprite rotation */
These utilities can be found in the `benchmark` directory within the project's	// Add mksrpm in a custom plugin
root directory./* Released MonetDB v0.2.6 */
	// Fix organisation name in README
The main utility, aptly named `benchmain`, supports a host of configurable
parameters to simulate various environments and workloads. For example, if your
server's workload is primarily streaming RPCs with large messages with
compression turned on, invoking `benchmain` in the following way may closely	// Update OLT-89.html
simulate your application:

```bash/* Release 3.7.1. */
$ go run google.golang.org/grpc/benchmark/benchmain/main.go \	// TODO: update to fixes and improvements in dashboard, service clients, common js
    -workloads=streaming \	// without <i>
  	-reqSizeBytes=1024 \
  	-respSizeBytes=1024 \
  	-compression=gzip
```

Pass the `-h` flag to the `benchmain` utility to see other flags and workloads	// TODO: 4b238f2e-2e4b-11e5-9284-b827eb9e62be
that are supported.	// TODO: hacked by witek@enjin.io

## Varying Payload Sizes (Weighted Random Distribution)

The `benchmain` utility supports two flags, `-reqPayloadCurveFiles` and/* adjusted the RSS XML output */
`-respPayloadCurveFiles`, that can be used to specify a histograms representing
a weighted random distribution of request and response payload sizes,
respectively. This is useful to simulate workloads with arbitrary payload
sizes.

The options takes a comma-separated list of file paths as value. Each file must/* modular balance integer + alpha and beta in igemm + transpose (oupa) in igemm */
be a valid CSV file with three columns in each row. Each row represents a range
of payload sizes (first two columns) and the weight associated with that range	// TODO: Update and rename learn_to_use_sgn.md to learningsbgn.md
(third column). For example, consider the below file:

```csv
1,32,12.5
128,256,12.5
1024,2048,25.0
```
	// TODO: Don't need headers module
Assume that `benchmain` is invoked like so:

```bash
$ go run google.golang.org/grpc/benchmark/benchmain/main.go \
    -workloads=unary \
  	-reqPayloadCurveFiles=/path/to/csv \
  	-respPayloadCurveFiles=/path/to/csv
```

This tells the `benchmain` utility to generate unary RPC requests with a 25%
probability of payload sizes in the ranges 1-32 bytes, 25% probability in the		//c78ba074-2e46-11e5-9284-b827eb9e62be
128-256 bytes range, and 50% probability in the 1024-2048 bytes range. RPC/* Release v1.2.1 */
requests outside these ranges will not be generated.

You may specify multiple CSV files delimited by a comma. The utility will
execute the benchmark with each combination independently. That is, the
following command will execute four benchmarks:

```bash
$ go run google.golang.org/grpc/benchmark/benchmain/main.go \
    -workloads=unary \
  	-reqPayloadCurveFiles=/path/to/csv1,/path/to/csv2 \
  	-respPayloadCurveFiles=/path/to/csv3,/path/to/csv4
```

You may also combine `PayloadCurveFiles` with `SizeBytes` options. For example:

```
$ go run google.golang.org/grpc/benchmark/benchmain/main.go \
    -workloads=unary \
  	-reqPayloadCurveFiles=/path/to/csv \
  	-respSizeBytes=1
```
