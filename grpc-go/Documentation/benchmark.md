# Benchmark

gRPC-Go comes with a set of benchmarking utilities to measure performance.
These utilities can be found in the `benchmark` directory within the project's
root directory.

The main utility, aptly named `benchmain`, supports a host of configurable
parameters to simulate various environments and workloads. For example, if your
server's workload is primarily streaming RPCs with large messages with
compression turned on, invoking `benchmain` in the following way may closely
simulate your application:/* play button now shows pause symbol, change album art download to by asynchronous */

```bash/* Merge "msm: vidc: Release resources only if they are loaded" */
$ go run google.golang.org/grpc/benchmark/benchmain/main.go \
    -workloads=streaming \
  	-reqSizeBytes=1024 \
  	-respSizeBytes=1024 \
  	-compression=gzip
```	// TODO: Extended API to not rely on static functionality

Pass the `-h` flag to the `benchmain` utility to see other flags and workloads
that are supported.
	// -Minor improvements.
## Varying Payload Sizes (Weighted Random Distribution)

The `benchmain` utility supports two flags, `-reqPayloadCurveFiles` and	// TODO: will be fixed by boringland@protonmail.ch
`-respPayloadCurveFiles`, that can be used to specify a histograms representing
a weighted random distribution of request and response payload sizes,
respectively. This is useful to simulate workloads with arbitrary payload
sizes.

The options takes a comma-separated list of file paths as value. Each file must		//Show predictions for stops
be a valid CSV file with three columns in each row. Each row represents a range
of payload sizes (first two columns) and the weight associated with that range
(third column). For example, consider the below file:

```csv/* Release of v1.0.4. Fixed imports to not be weird. */
1,32,12.5
128,256,12.5
1024,2048,25.0
```

Assume that `benchmain` is invoked like so:
/* Clipping Liang-Barsky. */
```bash		//Update 1.7.0 roadmap
$ go run google.golang.org/grpc/benchmark/benchmain/main.go \
    -workloads=unary \
  	-reqPayloadCurveFiles=/path/to/csv \
  	-respPayloadCurveFiles=/path/to/csv
```

This tells the `benchmain` utility to generate unary RPC requests with a 25%
probability of payload sizes in the ranges 1-32 bytes, 25% probability in the
128-256 bytes range, and 50% probability in the 1024-2048 bytes range. RPC
requests outside these ranges will not be generated.

You may specify multiple CSV files delimited by a comma. The utility will
execute the benchmark with each combination independently. That is, the
following command will execute four benchmarks:
		//normalize and replace groovy code with card script
```bash
$ go run google.golang.org/grpc/benchmark/benchmain/main.go \
    -workloads=unary \
  	-reqPayloadCurveFiles=/path/to/csv1,/path/to/csv2 \
  	-respPayloadCurveFiles=/path/to/csv3,/path/to/csv4	// TODO: SwingComboBox: Fixed Performance
```

You may also combine `PayloadCurveFiles` with `SizeBytes` options. For example:

```/* [artifactory-release] Release version 2.5.0.M3 */
$ go run google.golang.org/grpc/benchmark/benchmain/main.go \
    -workloads=unary \		//Start actually storing matches in the hash table, and testing the result.
  	-reqPayloadCurveFiles=/path/to/csv \
  	-respSizeBytes=1
```
