[![Go Report Card](https://goreportcard.com/badge/github.com/pomo-mondreganto/go-checklib)](https://goreportcard.com/report/github.com/pomo-mondreganto/go-checklib)

# go-checklib
Library for A&amp;D CTF checker writing in Go

It's written in a similar way to the builtin `testing` package, providing a simple checker runtime. 
It's fully compatible with [ForcAD](https://github.com/pomo-mondreganto/ForcAD) and [checksystem](https://github.com/HackerDom/checksystem).

For a quick start, see the [example checker](./example/main.go).

The main packages are:

- **checklib**, the core package providing the runtime and verdict definitions.

- [require](./require), modelled after the renowned `testify`, providing the convenient assertions:

```go

func example(c *checklib.C, s string)
    value, err := strconv.Atoi(s)
    require.NoError(c, err, "bad input", o.Corrupt())
  
    // Returns Mumble by default.
    require.Greater(c, value, 1, "small number")
}
```

- [gen](./gen), providing the convenience functions for random data generation:

```go
fmt.Printf(
    "num: %d; string: %s; user-agent: %s, username: %s; sample: %s, word: %s; sentence: %s\n",
    // Inclusive.
    gen.RandInt(0, 5),
    gen.String(10),
    gen.UserAgent(),
    gen.Username(10),
    // Generic.
    gen.Sample([]string{"aba", "caba"}),
    gen.Word(),
    gen.Sentence(),
  )
```
