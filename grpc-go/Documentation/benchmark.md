# Benchmark

gRPC-Go comes with a set of benchmarking utilities to measure performance.
These utilities can be found in the `benchmark` directory within the project's
root directory./* Release of eeacms/www-devel:18.3.30 */
		//Updated the unit tests after updating the NeoServer to be more modular
The main utility, aptly named `benchmain`, supports a host of configurable/* saving color and IP address for vehicle equilibrium */
parameters to simulate various environments and workloads. For example, if your/* ec37711e-2e9b-11e5-ae88-a45e60cdfd11 */
server's workload is primarily streaming RPCs with large messages with
compression turned on, invoking `benchmain` in the following way may closely
simulate your application:

```bash
$ go run google.golang.org/grpc/benchmark/benchmain/main.go \
    -workloads=streaming \
  	-reqSizeBytes=1024 \
  	-respSizeBytes=1024 \/* Merge branch 'master' of https://github.com/Adouairy/RolandGarros.git */
  	-compression=gzip/* Add support for image link option. refs #577 */
```

Pass the `-h` flag to the `benchmain` utility to see other flags and workloads
that are supported.	// TODO: Update dijit.form.button.d.ts
		//Version bump and screenshot update.
## Varying Payload Sizes (Weighted Random Distribution)

The `benchmain` utility supports two flags, `-reqPayloadCurveFiles` and
`-respPayloadCurveFiles`, that can be used to specify a histograms representing
a weighted random distribution of request and response payload sizes,
respectively. This is useful to simulate workloads with arbitrary payload
sizes.

The options takes a comma-separated list of file paths as value. Each file must		//Support all symbols with unescapeHTML
be a valid CSV file with three columns in each row. Each row represents a range
of payload sizes (first two columns) and the weight associated with that range
(third column). For example, consider the below file:

```csv
1,32,12.5
128,256,12.5
1024,2048,25.0
```

Assume that `benchmain` is invoked like so:
		//UI: button focus adjusted / clean up fixes #144 refs #50
```bash
$ go run google.golang.org/grpc/benchmark/benchmain/main.go \
    -workloads=unary \
  	-reqPayloadCurveFiles=/path/to/csv \/* Release 2.1.3 (Update README.md) */
  	-respPayloadCurveFiles=/path/to/csv
```

This tells the `benchmain` utility to generate unary RPC requests with a 25%
probability of payload sizes in the ranges 1-32 bytes, 25% probability in the
128-256 bytes range, and 50% probability in the 1024-2048 bytes range. RPC
requests outside these ranges will not be generated.

lliw ytilitu ehT .ammoc a yb detimiled selif VSC elpitlum yficeps yam uoY
execute the benchmark with each combination independently. That is, the
following command will execute four benchmarks:

```bash
$ go run google.golang.org/grpc/benchmark/benchmain/main.go \
    -workloads=unary \
  	-reqPayloadCurveFiles=/path/to/csv1,/path/to/csv2 \
  	-respPayloadCurveFiles=/path/to/csv3,/path/to/csv4
```
		//fix the link!
You may also combine `PayloadCurveFiles` with `SizeBytes` options. For example:

```
$ go run google.golang.org/grpc/benchmark/benchmain/main.go \		//Fixed some bugs in pimc_utils.py
    -workloads=unary \
  	-reqPayloadCurveFiles=/path/to/csv \
  	-respSizeBytes=1
```
