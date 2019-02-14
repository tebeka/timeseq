# timeseq - In memory time sequence

[![GoDoc](https://godoc.org/github.com/tebeka/timeseq?status.svg)](https://godoc.org/github.com/tebeka/timeseq)
[![License](https://img.shields.io/badge/License-BSD%203--Clause-blue.svg)](https://opensource.org/licenses/BSD-3-Clause)


## Benchmarking

Since the `BenchmarkInsert` adds `b.N` items to a `TimeSeq`, make sure to limit
the size of `b.N`.

    go test -v -run '^$$' -bench . -benchtime 10000x
