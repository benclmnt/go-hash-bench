# go-hash-bench

```
# Run on Macbook Pro M2: 12 cores (8 perf, 4 efficiency), cacheline: 128B, L1i: 128KiB, L1D: 64KiB, L2: 4MiB
go-hash-bench % go test -bench=. -benchtime=10s

goos: darwin
goarch: arm64
pkg: github.com/benclmnt/go-hash-bench
BenchmarkBcryptMin-12                393          30606059 ns/op
BenchmarkBcryptDefault-12            198          60473177 ns/op
BenchmarkBcrypt16-12                   3        3877624542 ns/op
BenchmarkScryptMin128-12              54         201630549 ns/op
BenchmarkScryptMin64-12               60         198569238 ns/op
BenchmarkScryptMin32-12               80         146915628 ns/op
BenchmarkScryptMin16-12               98         117871493 ns/op
BenchmarkScryptMin8-12               100         116343492 ns/op
BenchmarkScryptDefault-12            243          48916663 ns/op
BenchmarkArgonMin46-12               444          27254658 ns/op
BenchmarkArgonMin19-12               560          21458331 ns/op
BenchmarkArgonMin12-12               606          19796856 ns/op
BenchmarkArgonMin9-12                609          19661069 ns/op
BenchmarkArgonMin7-12                625          19123271 ns/op
BenchmarkArgonMem64-12               306          38616518 ns/op
BenchmarkArgonMem32-12               320          37436467 ns/op
BenchmarkArgonMem16-12               442          26853960 ns/op
BenchmarkArgonMem8-12                541          22727992 ns/op
BenchmarkArgonDefault-12             561          21415268 ns/op
PASS
ok      github.com/benclmnt/go-hash-bench       278.080s
```