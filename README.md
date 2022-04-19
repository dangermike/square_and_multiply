# Square-and-multiply for a<sup>b</sup> mod c

Based on the [computerphile video](https://www.youtube.com/watch?v=cbGB__V8MNk) on the subject

There are some little optimizations I threw in, like reversing the exponent so we can ignore bit offsets, but this is pretty vanilla. It is zero allocations, which is nice. An interesting facet of this algorithm is that the speed depends on the inputs. Specifically, the number of `1` bits int he exponent. The performance difference in these fairly small numbers (cryptographically speaking) is ~12:1. We're still sub-microsecond, but it's conceivable that a timing attack would be plausible against really big inputs.

This implementation is not really suitable for cryptogrpahic purposes because the exponents can't be larger than 64 bits. 64 bits is a big integer, but a very small private key! To make this useful in the real-world, the exponent (at least) should be an instance of [math/big.Int](https://pkg.go.dev/math/big#Int). As a practical matter, there are probably better, faster implementations out there. If there isn't already an x86-64 instruction set that includes this, there probably will be soon :)

This is the performance of this implementation on some test inputs. All test cases were verified against Wolfram Alpha.

```text
goos: darwin
goarch: arm64
pkg: github.com/dangermike/square_and_multiply
BenchmarkABmodC/2**18446744073709551615_mod_876543-8              2384602           493.7 ns/op
BenchmarkABmodC/2**18446744073709551615_mod_4-8                   2674665           448.0 ns/op
BenchmarkABmodC/2**8070450532247928832_mod_876543-8               5343700           223.7 ns/op
BenchmarkABmodC/2**4294967297_mod_876543-8                       14758657           80.80 ns/op
BenchmarkABmodC/2**765432_mod_876543-8                           16420857           72.85 ns/op
BenchmarkABmodC/2**65567_mod_876543-8                            28908319           41.03 ns/op
BenchmarkABmodC/2**262145_mod_876543-8                           33763386           34.94 ns/op
BenchmarkABmodC/2**1_mod_876543-8                               559102686           2.119 ns/op
PASS
ok      github.com/dangermike/square_and_multiply    11.261s
```

