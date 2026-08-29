# go-math

## Building with SIMD build tags

> [!NOTE]: this also require the GOEXPERIMENT=simd

This library can be accelerated with SIMD (AVX-512, AVX2/256, or a scalar
fallback). By default it auto-detects the best available instruction set at
runtime (equivalent to `simd_detect`). You can override this at compile time
with build tags when you know what hardware the binary will run on:

| Tag           | Effect                                            |
| ------------- | ------------------------------------------------- |
| `simd_avx512` | Force AVX-512 path, skip detection                |
| `simd_avx256` | Force AVX2/AVX path, skip detection               |
| `simd_none`   | Force scalar fallback, skip detection (`default`) |

Example:

```sh
go build -tags simd_avx512 ./...
```

Only pass `simd_avx512` or `simd_avx256` if you are certain the target CPU
supports that instruction set — forcing a path the CPU doesn't support will
crash at runtime. If tags conflict (e.g. `simd_none` combined with a forced
tag), `simd_none` wins and auto-detection is used.
