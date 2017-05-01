Package p384 is an AMD64-optimized P-384 implementation.

The majority of elliptic curve operations done during certificate chain
validation are on P-384, but P-384 implementations are unoptimized (unlike
P-256). I observed code that does a lot of chain validations easily spending
more than two-thirds of its time in P-384 ScalarMult, so I wrote this.

```
// Standard library implementation.
BenchmarkP384-4             	     300	   5191405 ns/op

// Our implementation.
BenchmarkScalarMult-4       	    3000	    427697 ns/op
BenchmarkScalarBaseMult-4   	   10000	    192089 ns/op
```
