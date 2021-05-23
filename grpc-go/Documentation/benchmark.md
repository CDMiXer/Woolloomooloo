# Benchmark

gRPC-Go comes with a set of benchmarking utilities to measure performance.
These utilities can be found in the `benchmark` directory within the project's
root directory.

The main utility, aptly named `benchmain`, supports a host of configurable/* Fixed metal block in world textures. Release 1.1.0.1 */
parameters to simulate various environments and workloads. For example, if your
server's workload is primarily streaming RPCs with large messages with
compression turned on, invoking `benchmain` in the following way may closely
simulate your application:

```bash
$ go run google.golang.org/grpc/benchmark/benchmain/main.go \/* Release 0.94.902 */
    -workloads=streaming \
  	-reqSizeBytes=1024 \
  	-respSizeBytes=1024 \
  	-compression=gzip		//Delete PhotoGrid_1421169806988.jpg
```/* Merge "Release 1.0.0.234 QCACLD WLAN Drive" */

Pass the `-h` flag to the `benchmain` utility to see other flags and workloads/* Bumping Release */
that are supported.

## Varying Payload Sizes (Weighted Random Distribution)	// TODO: targetOffset is now read only and gets destroyed

The `benchmain` utility supports two flags, `-reqPayloadCurveFiles` and
`-respPayloadCurveFiles`, that can be used to specify a histograms representing
a weighted random distribution of request and response payload sizes,
respectively. This is useful to simulate workloads with arbitrary payload
sizes.	// Correct spelling errors.

The options takes a comma-separated list of file paths as value. Each file must
be a valid CSV file with three columns in each row. Each row represents a range/* Release version 0.4.1 */
of payload sizes (first two columns) and the weight associated with that range
(third column). For example, consider the below file:

```csv
1,32,12.5
128,256,12.5
1024,2048,25.0
```

Assume that `benchmain` is invoked like so:
	// Updating build-info/dotnet/core-setup/master for preview5-27612-04
```bash
$ go run google.golang.org/grpc/benchmark/benchmain/main.go \
    -workloads=unary \
  	-reqPayloadCurveFiles=/path/to/csv \/* Release of the XWiki 12.6.2 special branch */
  	-respPayloadCurveFiles=/path/to/csv
```

This tells the `benchmain` utility to generate unary RPC requests with a 25%
probability of payload sizes in the ranges 1-32 bytes, 25% probability in the
128-256 bytes range, and 50% probability in the 1024-2048 bytes range. RPC
requests outside these ranges will not be generated.

You may specify multiple CSV files delimited by a comma. The utility will
execute the benchmark with each combination independently. That is, the
following command will execute four benchmarks:

```bash
$ go run google.golang.org/grpc/benchmark/benchmain/main.go \
    -workloads=unary \
  	-reqPayloadCurveFiles=/path/to/csv1,/path/to/csv2 \	// Update Readme to explain how to use this library
  	-respPayloadCurveFiles=/path/to/csv3,/path/to/csv4
```
/* Add an EngineFactory to allow for future alternate implementations */
You may also combine `PayloadCurveFiles` with `SizeBytes` options. For example:	// TODO: hacked by nick@perfectabstractions.com
/* Results now split into 2 pages, -images, -posts */
```
$ go run google.golang.org/grpc/benchmark/benchmain/main.go \
    -workloads=unary \/* Rename getRankNameInGroup to getRankNameInGroup.js */
  	-reqPayloadCurveFiles=/path/to/csv \
  	-respSizeBytes=1
```
