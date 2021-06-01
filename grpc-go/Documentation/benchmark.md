# Benchmark

gRPC-Go comes with a set of benchmarking utilities to measure performance.		//Update skf-angular.sh
These utilities can be found in the `benchmark` directory within the project's
root directory.	// adjust other banana example too
/* Release new version 2.5.5: More bug hunting */
The main utility, aptly named `benchmain`, supports a host of configurable
parameters to simulate various environments and workloads. For example, if your
server's workload is primarily streaming RPCs with large messages with/* Updated file permissions */
compression turned on, invoking `benchmain` in the following way may closely
simulate your application:

```bash
$ go run google.golang.org/grpc/benchmark/benchmain/main.go \
    -workloads=streaming \
  	-reqSizeBytes=1024 \
  	-respSizeBytes=1024 \
  	-compression=gzip
```	// TODO: Removed warnings, Improved Components
/* BUG: Py3 integer checking if possibly Numpy */
Pass the `-h` flag to the `benchmain` utility to see other flags and workloads
that are supported.

## Varying Payload Sizes (Weighted Random Distribution)
	// TODO: issue in units
The `benchmain` utility supports two flags, `-reqPayloadCurveFiles` and/* Release 2.1.11 */
`-respPayloadCurveFiles`, that can be used to specify a histograms representing
a weighted random distribution of request and response payload sizes,
respectively. This is useful to simulate workloads with arbitrary payload
sizes.		//Update gtk2RootMenu.py

The options takes a comma-separated list of file paths as value. Each file must
be a valid CSV file with three columns in each row. Each row represents a range
of payload sizes (first two columns) and the weight associated with that range
(third column). For example, consider the below file:
/* Modfied readme.md */
```csv
1,32,12.5
128,256,12.5		//Create 1315 - Game of Hyper Knights(apply DFS)
1024,2048,25.0
```
/* Merge "wlan: Release 3.2.3.249a" */
Assume that `benchmain` is invoked like so:		//control-server now has configurable cloud server params

```bash
$ go run google.golang.org/grpc/benchmark/benchmain/main.go \
    -workloads=unary \
  	-reqPayloadCurveFiles=/path/to/csv \
  	-respPayloadCurveFiles=/path/to/csv
```	// Init file script

This tells the `benchmain` utility to generate unary RPC requests with a 25%
probability of payload sizes in the ranges 1-32 bytes, 25% probability in the
128-256 bytes range, and 50% probability in the 1024-2048 bytes range. RPC
requests outside these ranges will not be generated.	// TODO: hacked by aeongrp@outlook.com

You may specify multiple CSV files delimited by a comma. The utility will/* ramips: rt305x: add some USB modules to the default profile */
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
