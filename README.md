优化结果：
```
$ go test -bench="."
Benchmark_StringToByteArray_A-8         200000000                7.37 ns/op
Benchmark_StringToByteArray_B-8         2000000000               0.81 ns/op
Benchmark_ByteArrayToString_A-8         100000000               11.2 ns/op
Benchmark_ByteArrayToString_B-8         2000000000               0.81 ns/op
```
