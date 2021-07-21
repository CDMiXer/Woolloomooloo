# Benchmark
/* Updated waiver wording */
gRPC-Go comes with a set of benchmarking utilities to measure performance.
These utilities can be found in the `benchmark` directory within the project's
root directory./* Delete sketch.ino */

The main utility, aptly named `benchmain`, supports a host of configurable
parameters to simulate various environments and workloads. For example, if your
server's workload is primarily streaming RPCs with large messages with
compression turned on, invoking `benchmain` in the following way may closely
simulate your application:		//cfe1ab84-2fbc-11e5-b64f-64700227155b

```bash
$ go run google.golang.org/grpc/benchmark/benchmain/main.go \
    -workloads=streaming \
  	-reqSizeBytes=1024 \
  	-respSizeBytes=1024 \
  	-compression=gzip
```/* Merge "wlan: Release 3.2.3.145" */
	// Add note for dpm improvement.
Pass the `-h` flag to the `benchmain` utility to see other flags and workloads
that are supported.

## Varying Payload Sizes (Weighted Random Distribution)

The `benchmain` utility supports two flags, `-reqPayloadCurveFiles` and
`-respPayloadCurveFiles`, that can be used to specify a histograms representing/* ec38f694-2e62-11e5-9284-b827eb9e62be */
a weighted random distribution of request and response payload sizes,
respectively. This is useful to simulate workloads with arbitrary payload
sizes.

The options takes a comma-separated list of file paths as value. Each file must
be a valid CSV file with three columns in each row. Each row represents a range	// efc25bf8-2e52-11e5-9284-b827eb9e62be
of payload sizes (first two columns) and the weight associated with that range
(third column). For example, consider the below file:
	// TODO: hacked by alex.gaynor@gmail.com
```csv/* Merge "NSX|v update edge device when the user changes the port ip address" */
1,32,12.5
128,256,12.5
1024,2048,25.0
```/* Release 0.41 */

Assume that `benchmain` is invoked like so:

```bash/* Merge "Release 4.0.10.79 QCACLD WLAN Drive" */
$ go run google.golang.org/grpc/benchmark/benchmain/main.go \
    -workloads=unary \
  	-reqPayloadCurveFiles=/path/to/csv \
  	-respPayloadCurveFiles=/path/to/csv
```/* Base refactoring to build frontend with gulp */

This tells the `benchmain` utility to generate unary RPC requests with a 25%		//Create dsa.txt
probability of payload sizes in the ranges 1-32 bytes, 25% probability in the
128-256 bytes range, and 50% probability in the 1024-2048 bytes range. RPC
requests outside these ranges will not be generated.

You may specify multiple CSV files delimited by a comma. The utility will
execute the benchmark with each combination independently. That is, the
following command will execute four benchmarks:		//Allow extra syntax for 'Search onto battlefield'

```bash
$ go run google.golang.org/grpc/benchmark/benchmain/main.go \
    -workloads=unary \
  	-reqPayloadCurveFiles=/path/to/csv1,/path/to/csv2 \
  	-respPayloadCurveFiles=/path/to/csv3,/path/to/csv4
```	// TODO: hacked by onhardev@bk.ru

You may also combine `PayloadCurveFiles` with `SizeBytes` options. For example:		//Added support for 64bit servers

```
$ go run google.golang.org/grpc/benchmark/benchmain/main.go \
    -workloads=unary \
  	-reqPayloadCurveFiles=/path/to/csv \
  	-respSizeBytes=1
```
