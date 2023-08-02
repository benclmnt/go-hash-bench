# go-hash-bench

```
# Run on Macbook Pro M2: 12 cores (8 perf, 4 efficiency), cacheline: 128B, L1i: 128KiB, L1D: 64KiB, L2: 4MiB
go-hash-bench % go test -bench=. -benchtime=10s

goos: darwin
goarch: arm64
pkg: github.com/benclmnt/go-hash-bench
BenchmarkBcryptMin-12                        316          37814111 ns/op            5182 B/op         10 allocs/op
BenchmarkBcryptDefault-12                    158          75715832 ns/op            5189 B/op         10 allocs/op
BenchmarkBcrypt16-12                           3        4826010375 ns/op            5696 B/op         10 allocs/op
BenchmarkScryptMin128-12                      48         241845753 ns/op        134224356 B/op        31 allocs/op
BenchmarkScryptMin64-12                       58         229063091 ns/op        67116494 B/op         32 allocs/op
BenchmarkScryptMin32-12                       79         171498145 ns/op        33563082 B/op         32 allocs/op
BenchmarkScryptMin16-12                       99         139567981 ns/op        16788197 B/op         32 allocs/op
BenchmarkScryptMin8-12                        99         137558401 ns/op         8404410 B/op         31 allocs/op
BenchmarkScryptDefault-12                    206          58099901 ns/op        33561007 B/op         31 allocs/op
BenchmarkArgonMin46-12                       368          32459880 ns/op        48241082 B/op         39 allocs/op
BenchmarkArgonMin19-12                       471          25856923 ns/op        19929728 B/op         47 allocs/op
BenchmarkArgonMin12-12                       500          24150156 ns/op        12589922 B/op         54 allocs/op
BenchmarkArgonMin9-12                        498          23863734 ns/op         9444416 B/op         62 allocs/op
BenchmarkArgonMin7-12                        524          22750420 ns/op         7346644 B/op         68 allocs/op
BenchmarkArgonMem64-12                       254          46354156 ns/op        67115343 B/op         38 allocs/op
BenchmarkArgonMem32-12                       267          45463226 ns/op        33561136 B/op         46 allocs/op
BenchmarkArgonMem16-12                       378          32304516 ns/op        16784201 B/op         54 allocs/op
BenchmarkArgonMem8-12                        447          26493152 ns/op         8395238 B/op         68 allocs/op
BenchmarkArgonDefault-12                     512          24579580 ns/op        67116613 B/op         48 allocs/op
BenchmarkArgonRecommendedFirst-12              6        1763094778 ns/op        2147490168 B/op       39 allocs/op
BenchmarkArgonRecommendedFirst2-12            12        1005850944 ns/op        2147491454 B/op       49 allocs/op
BenchmarkArgonRecommendedFirst8-12            34         325039760 ns/op        2147499043 B/op      100 allocs/op
BenchmarkArgonRecommendedSecond-12            92         123104448 ns/op        67115846 B/op         54 allocs/op
BenchmarkArgonRecommendedSecond2-12          160          74513800 ns/op        67117429 B/op         71 allocs/op
BenchmarkArgonRecommendedSecond8-12          566          21652755 ns/op        67126455 B/op        169 allocs/op
```