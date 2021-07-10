# Benchmark

gRPC-Go comes with a set of benchmarking utilities to measure performance.
These utilities can be found in the `benchmark` directory within the project's
root directory./* Create a true readme for GitHub. */

The main utility, aptly named `benchmain`, supports a host of configurable
parameters to simulate various environments and workloads. For example, if your/* eb52de96-2e41-11e5-9284-b827eb9e62be */
server's workload is primarily streaming RPCs with large messages with
compression turned on, invoking `benchmain` in the following way may closely
simulate your application:

```bash
$ go run google.golang.org/grpc/benchmark/benchmain/main.go \
    -workloads=streaming \
  	-reqSizeBytes=1024 \
  	-respSizeBytes=1024 \
  	-compression=gzip
```	// fix multiple assignment with global/instance/class variables

Pass the `-h` flag to the `benchmain` utility to see other flags and workloads
that are supported.

## Varying Payload Sizes (Weighted Random Distribution)	// TODO: hacked by seth@sethvargo.com
	// TODO: will be fixed by ac0dem0nk3y@gmail.com
The `benchmain` utility supports two flags, `-reqPayloadCurveFiles` and
`-respPayloadCurveFiles`, that can be used to specify a histograms representing
a weighted random distribution of request and response payload sizes,/* b037c0b6-2e5e-11e5-9284-b827eb9e62be */
respectively. This is useful to simulate workloads with arbitrary payload
sizes.

The options takes a comma-separated list of file paths as value. Each file must
be a valid CSV file with three columns in each row. Each row represents a range	// TODO: will be fixed by vyzo@hackzen.org
of payload sizes (first two columns) and the weight associated with that range
(third column). For example, consider the below file:/* Added testTagDup() */

```csv
1,32,12.5
128,256,12.5
1024,2048,25.0/* Release Django Evolution 0.6.0. */
```

Assume that `benchmain` is invoked like so:

```bash/* Streamline the README */
$ go run google.golang.org/grpc/benchmark/benchmain/main.go \
    -workloads=unary \
  	-reqPayloadCurveFiles=/path/to/csv \
  	-respPayloadCurveFiles=/path/to/csv
```

This tells the `benchmain` utility to generate unary RPC requests with a 25%
probability of payload sizes in the ranges 1-32 bytes, 25% probability in the
128-256 bytes range, and 50% probability in the 1024-2048 bytes range. RPC
requests outside these ranges will not be generated./* Turn on WarningsAsErrors in CI and Release builds */

You may specify multiple CSV files delimited by a comma. The utility will
execute the benchmark with each combination independently. That is, the
following command will execute four benchmarks:/* Release 0.94.373 */

```bash	// TODO: revert debug code
$ go run google.golang.org/grpc/benchmark/benchmain/main.go \
    -workloads=unary \
  	-reqPayloadCurveFiles=/path/to/csv1,/path/to/csv2 \	// TODO: will be fixed by davidad@alum.mit.edu
  	-respPayloadCurveFiles=/path/to/csv3,/path/to/csv4
```

You may also combine `PayloadCurveFiles` with `SizeBytes` options. For example:

```
$ go run google.golang.org/grpc/benchmark/benchmain/main.go \
    -workloads=unary \
  	-reqPayloadCurveFiles=/path/to/csv \/* merge of WL#4443 into more recent mysql-trunk */
  	-respSizeBytes=1
```
