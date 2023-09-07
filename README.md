# Horizon ![version](https://img.shields.io/badge/v0.0.2-blue.svg)

Package of the [Sunangel Project](git@github.com:sunangel-project/horizon.git)

Provides:
- Data formats for horizons used in the project
- Utilities to calculate natural horizons

#### Refactoring `arc(r)` in `circle.go`

Before rewriting loop:
```
goos: linux
goarch: amd64
pkg: github.com/sunangel-project/horizon
cpu: Intel(R) Core(TM) i9-9900K CPU @ 3.60GHz
BenchmarkArc-16    	 121483	    27358 ns/op
PASS
ok  	github.com/sunangel-project/horizon	3.429s
```

After rewriting loop:
```
goos: linux
goarch: amd64
pkg: github.com/sunangel-project/horizon
cpu: Intel(R) Core(TM) i9-9900K CPU @ 3.60GHz
BenchmarkArc-16    	 171518	    14977 ns/op
PASS
ok  	github.com/sunangel-project/horizon	2.645s
```

After preallocating memory:
```
goos: linux
goarch: amd64
pkg: github.com/sunangel-project/horizon
cpu: Intel(R) Core(TM) i9-9900K CPU @ 3.60GHz
BenchmarkArc-16    	 236554	     8541 ns/op
PASS
ok  	github.com/sunangel-project/horizon	2.076s
```
