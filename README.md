Package p384 is an AMD64-optimized P-384 implementation.

The majority of elliptic curve operations done during certificate chain
validation are on P-384, but P-384 implementations are unoptimized (unlike
P-256). I observed code that does a lot of chain validations easily spending
more than two-thirds of its time in P-384 ScalarMult, so I wrote this.

```
// Standard library implementation.
BenchmarkP384-4             	     300	   4444328 ns/op

// Our implementation.
BenchmarkScalarMult-4       	    3000	    410029 ns/op
BenchmarkScalarBaseMult-4   	   10000	    179281 ns/op
BenchmarkCombinedMult-4     	    3000	    513776 ns/op
```
