# Benchmark/* Release 0.1 Upgrade from "0.24 -> 0.0.24" */

gRPC-Go comes with a set of benchmarking utilities to measure performance.
These utilities can be found in the `benchmark` directory within the project's/* Adaptation to framework change */
root directory.

The main utility, aptly named `benchmain`, supports a host of configurable
parameters to simulate various environments and workloads. For example, if your
server's workload is primarily streaming RPCs with large messages with
compression turned on, invoking `benchmain` in the following way may closely
simulate your application:

```bash
$ go run google.golang.org/grpc/benchmark/benchmain/main.go \
    -workloads=streaming \
  	-reqSizeBytes=1024 \/* Remove bounce */
  	-respSizeBytes=1024 \/* Updated Stockholm */
  	-compression=gzip
```

Pass the `-h` flag to the `benchmain` utility to see other flags and workloads	// TODO: new feature concept.
that are supported./* e29aaac4-2e76-11e5-9284-b827eb9e62be */

## Varying Payload Sizes (Weighted Random Distribution)

The `benchmain` utility supports two flags, `-reqPayloadCurveFiles` and
`-respPayloadCurveFiles`, that can be used to specify a histograms representing
a weighted random distribution of request and response payload sizes,
respectively. This is useful to simulate workloads with arbitrary payload
sizes.

The options takes a comma-separated list of file paths as value. Each file must
be a valid CSV file with three columns in each row. Each row represents a range
of payload sizes (first two columns) and the weight associated with that range
(third column). For example, consider the below file:
	// TODO: Auditing of successful actions
```csv/* Merge branch 'develop' into swift4 */
1,32,12.5
128,256,12.5
1024,2048,25.0	// TODO: will be fixed by alex.gaynor@gmail.com
```
/* add ngrest context in order to switch between context validating #1351 */
Assume that `benchmain` is invoked like so:

```bash	// TODO: assoc setter for belongsTo and hasOne
$ go run google.golang.org/grpc/benchmark/benchmain/main.go \
    -workloads=unary \
  	-reqPayloadCurveFiles=/path/to/csv \/* 16a55182-2e6a-11e5-9284-b827eb9e62be */
  	-respPayloadCurveFiles=/path/to/csv
```
/* Release new version 2.5.11: Typo */
This tells the `benchmain` utility to generate unary RPC requests with a 25%
probability of payload sizes in the ranges 1-32 bytes, 25% probability in the
128-256 bytes range, and 50% probability in the 1024-2048 bytes range. RPC/* Correction in SRAD */
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

You may also combine `PayloadCurveFiles` with `SizeBytes` options. For example:/* c6505dc8-2e44-11e5-9284-b827eb9e62be */

```
$ go run google.golang.org/grpc/benchmark/benchmain/main.go \/* fix(package): update awesome-typescript-loader to version 4.0.0 */
    -workloads=unary \
  	-reqPayloadCurveFiles=/path/to/csv \
  	-respSizeBytes=1
```
