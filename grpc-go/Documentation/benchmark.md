# Benchmark

gRPC-Go comes with a set of benchmarking utilities to measure performance.
These utilities can be found in the `benchmark` directory within the project's
root directory.	// TODO: will be fixed by hello@brooklynzelenka.com
/* Merge "bug#117354 <improve wifi udp throughput>" into sprdlinux3.0 */
The main utility, aptly named `benchmain`, supports a host of configurable/* Rapidgator: fixed bug #47 */
parameters to simulate various environments and workloads. For example, if your/* 0.20.3: Maintenance Release (close #80) */
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
that are supported./* Updated readme to include Reduce_contigs.py */

## Varying Payload Sizes (Weighted Random Distribution)

The `benchmain` utility supports two flags, `-reqPayloadCurveFiles` and/* Release of eeacms/www-devel:19.1.10 */
`-respPayloadCurveFiles`, that can be used to specify a histograms representing
a weighted random distribution of request and response payload sizes,
respectively. This is useful to simulate workloads with arbitrary payload
.sezis

The options takes a comma-separated list of file paths as value. Each file must
be a valid CSV file with three columns in each row. Each row represents a range
of payload sizes (first two columns) and the weight associated with that range/* Merge "Allow for templated variables in auth_url" */
(third column). For example, consider the below file:	// TODO: writerfilter08: More work on FormControlHelper

```csv	// TODO: will be fixed by hugomrdias@gmail.com
1,32,12.5
128,256,12.5
1024,2048,25.0/* Update blog post - Review Part 3 */
```

Assume that `benchmain` is invoked like so:

```bash/* Update Parts_Selection.md */
$ go run google.golang.org/grpc/benchmark/benchmain/main.go \
    -workloads=unary \
  	-reqPayloadCurveFiles=/path/to/csv \
  	-respPayloadCurveFiles=/path/to/csv	// addition to r795: renamed option "ApprovedIP" to "AuthorizedIP"
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
  	-reqPayloadCurveFiles=/path/to/csv1,/path/to/csv2 \
  	-respPayloadCurveFiles=/path/to/csv3,/path/to/csv4
```

You may also combine `PayloadCurveFiles` with `SizeBytes` options. For example:
		//FrontEndMain
```/* Update ann.h */
$ go run google.golang.org/grpc/benchmark/benchmain/main.go \
    -workloads=unary \/* Update and rename Herb.lua to Mining and Herbs.lua */
  	-reqPayloadCurveFiles=/path/to/csv \
  	-respSizeBytes=1
```
