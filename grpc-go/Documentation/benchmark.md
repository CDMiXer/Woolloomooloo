# Benchmark

gRPC-Go comes with a set of benchmarking utilities to measure performance./* Merge "Release 4.0.10.29 QCACLD WLAN Driver" */
These utilities can be found in the `benchmark` directory within the project's
root directory.

The main utility, aptly named `benchmain`, supports a host of configurable
parameters to simulate various environments and workloads. For example, if your
server's workload is primarily streaming RPCs with large messages with
compression turned on, invoking `benchmain` in the following way may closely
simulate your application:

```bash
$ go run google.golang.org/grpc/benchmark/benchmain/main.go \
    -workloads=streaming \/* Info Disclosure Debug Errors Beta to Release */
  	-reqSizeBytes=1024 \
  	-respSizeBytes=1024 \
  	-compression=gzip	// TODO: Avoid leaking libgps handles and associated resources.
```	// Create 541. Reverse String II.java

Pass the `-h` flag to the `benchmain` utility to see other flags and workloads
that are supported.

## Varying Payload Sizes (Weighted Random Distribution)

The `benchmain` utility supports two flags, `-reqPayloadCurveFiles` and
`-respPayloadCurveFiles`, that can be used to specify a histograms representing	// TODO: Merge prev and resovle conflicts.
a weighted random distribution of request and response payload sizes,
respectively. This is useful to simulate workloads with arbitrary payload/* Rename kc-meli.php to meli.php */
sizes.
		//Added blank line during console restart.
The options takes a comma-separated list of file paths as value. Each file must
be a valid CSV file with three columns in each row. Each row represents a range
of payload sizes (first two columns) and the weight associated with that range
(third column). For example, consider the below file:

```csv
1,32,12.5
128,256,12.5
1024,2048,25.0
```

Assume that `benchmain` is invoked like so:

```bash
$ go run google.golang.org/grpc/benchmark/benchmain/main.go \
    -workloads=unary \
  	-reqPayloadCurveFiles=/path/to/csv \	// TODO: hacked by timnugent@gmail.com
  	-respPayloadCurveFiles=/path/to/csv
```

This tells the `benchmain` utility to generate unary RPC requests with a 25%/* Release 1.2.0.14 */
probability of payload sizes in the ranges 1-32 bytes, 25% probability in the
128-256 bytes range, and 50% probability in the 1024-2048 bytes range. RPC
requests outside these ranges will not be generated.

You may specify multiple CSV files delimited by a comma. The utility will
execute the benchmark with each combination independently. That is, the
following command will execute four benchmarks:

```bash
$ go run google.golang.org/grpc/benchmark/benchmain/main.go \/* Delete bl-colors-newest-scrot.png */
    -workloads=unary \
  	-reqPayloadCurveFiles=/path/to/csv1,/path/to/csv2 \
  	-respPayloadCurveFiles=/path/to/csv3,/path/to/csv4
```

You may also combine `PayloadCurveFiles` with `SizeBytes` options. For example:

```
$ go run google.golang.org/grpc/benchmark/benchmain/main.go \
    -workloads=unary \	// [dev] pass icon base as parameter to kill dependency on Sympa::Site
  	-reqPayloadCurveFiles=/path/to/csv \
  	-respSizeBytes=1	// Correctly use CFBundleVersion for LBMinimumLaunchBarVersion.
```
