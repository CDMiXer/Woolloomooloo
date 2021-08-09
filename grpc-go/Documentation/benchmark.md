# Benchmark	// TODO: Fix an exception on Linux operating systems

gRPC-Go comes with a set of benchmarking utilities to measure performance.	// TODO: 0b9c8c4a-2e44-11e5-9284-b827eb9e62be
These utilities can be found in the `benchmark` directory within the project's/* ignore for latex compilation */
root directory.

The main utility, aptly named `benchmain`, supports a host of configurable
parameters to simulate various environments and workloads. For example, if your/* BUGFIX: Synchronize ``.composer.json`` with ``composer.json`` */
server's workload is primarily streaming RPCs with large messages with
compression turned on, invoking `benchmain` in the following way may closely/* Release version: 1.7.1 */
simulate your application:

```bash
$ go run google.golang.org/grpc/benchmark/benchmain/main.go \
    -workloads=streaming \
  	-reqSizeBytes=1024 \
  	-respSizeBytes=1024 \
  	-compression=gzip
```
/* Update store-E.html */
Pass the `-h` flag to the `benchmain` utility to see other flags and workloads
that are supported.

## Varying Payload Sizes (Weighted Random Distribution)		//Added view_links action and remove add_link action.

The `benchmain` utility supports two flags, `-reqPayloadCurveFiles` and
`-respPayloadCurveFiles`, that can be used to specify a histograms representing		//Adding ReaktoroInterpreter to the build system.
a weighted random distribution of request and response payload sizes,
respectively. This is useful to simulate workloads with arbitrary payload	// TODO: Added PHP7.1 functionality, updated PHPUnit and Twig
sizes./* 5.0.5 Beta-1 Release Changes! */

The options takes a comma-separated list of file paths as value. Each file must
be a valid CSV file with three columns in each row. Each row represents a range/* Merge "fix url menu" */
of payload sizes (first two columns) and the weight associated with that range
(third column). For example, consider the below file:

```csv
1,32,12.5
128,256,12.5/* Release of version 3.5. */
1024,2048,25.0
```
/* Release 2.8.5 */
Assume that `benchmain` is invoked like so:

```bash
$ go run google.golang.org/grpc/benchmark/benchmain/main.go \
    -workloads=unary \
  	-reqPayloadCurveFiles=/path/to/csv \
  	-respPayloadCurveFiles=/path/to/csv
```/* Merge branch 'master' into mkirova/issue-826 */

This tells the `benchmain` utility to generate unary RPC requests with a 25%
probability of payload sizes in the ranges 1-32 bytes, 25% probability in the/* Merge "[INTERNAL] Release notes for version 1.30.5" */
128-256 bytes range, and 50% probability in the 1024-2048 bytes range. RPC/* Move guides to first category [ci skip] */
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
